package router

import (
	"context"
	"github.com/kent-h/stateful-router/protos/peer"
	"google.golang.org/grpc"
	"time"
)

type DeviceLoader interface {
	Load(ctx context.Context, deviceId string) error
	Unload(deviceId string)
}

func New(server *grpc.Server, ordinal uint32, peerDNSFormat string, loader DeviceLoader, readyCallback func()) *Router {
	ctx, ctxCancelFunc := context.WithCancel(context.Background())
	router := &Router{
		ordinal:       ordinal,
		peerDNSFormat: peerDNSFormat,
		ctx:           ctx,
		ctxCancelFunc: ctxCancelFunc,
		peers:         make(map[uint32]*node),
		devices:       make(map[string]*deviceData),
		deviceCountEventData: deviceCountEventData{
			updateComplete: make(chan struct{}),
			updatingPeers:  make(map[uint32]uint32),
		},
		rebalanceEventData: rebalanceEventData{
			newPeers: make(map[peer.PeerClient]struct{}),
		},
		deviceCountEventHandlerDone: make(chan struct{}),
		rebalanceEventHandlerDone:   make(chan struct{}),
		loader:                      loader,
	}

	// connect to all nodes with smaller ordinal than ours
	for i := uint32(0); i < router.ordinal; i++ {
		router.connect(i)
	}

	go router.startStatsNotifier()

	// node will only start accepting requests after rebalancer is started (default readiness allows no requests)
	time.AfterFunc(waitReadyTime, func() {
		if readyCallback != nil {
			readyCallback()
		}
		router.startRebalancer()
	})

	peer.RegisterPeerServer(server, peerApi{router})
	return router
}

func (router *Router) Stop() {
	router.ctxCancelFunc()
	<-router.rebalanceEventHandlerDone
	<-router.deviceCountEventHandlerDone
}

// this should be called when a device has been, or should be,
// unloaded **from the local node only** due to external circumstance
// (device lock lost, device deleted, inactivity timeout, etc.)
func (router *Router) UnloadDevice(deviceId string) {
	router.UnloadDeviceWithUnloadFunc(deviceId, router.loader.Unload)
}

func (router *Router) UnloadDeviceWithUnloadFunc(deviceId string, unloadFunc func(deviceId string)) {
	router.deviceMutex.Lock()
	device, have := router.devices[deviceId]
	delete(router.devices, deviceId)
	router.deviceMutex.Unlock()

	if have {
		router.deviceCountChanged()
		router.unloadDevice(deviceId, device, unloadFunc)
	}
}

// Locate returns a processor for the given device,
// to either handle the request locally,
// or forward it on to the appropriate peer
func (router *Router) Locate(deviceId string) (interface{ RUnlock() }, *grpc.ClientConn, bool, error) {
	return router.locate(deviceId, router.deviceCountChanged, router.loader.Load)
}

func (router *Router) LocateWithLoadFunc(deviceId string, loadFunc func(ctx context.Context, deviceId string) error) (interface{ RUnlock() }, *grpc.ClientConn, bool, error) {
	return router.locate(deviceId, router.deviceCountChanged, loadFunc)
}
