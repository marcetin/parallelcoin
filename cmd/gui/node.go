package gui

import "github.com/p9c/pod/cmd/node/rpc"

func (r *rcvar) GetNetworkLastBlock() int32 {
	for _, g := range r.cx.RPCServer.Cfg.ConnMgr.ConnectedPeers() {
		l := g.ToPeer().StatsSnapshot().LastBlock
		if l > r.status.NetworkLastBlock {
			r.status.NetworkLastBlock = l
		}
	}
	return r.status.NetworkLastBlock
}

func (r *rcvar) GetBlockCount() int64 {
	getBlockCount, err := rpc.HandleGetBlockCount(r.cx.RPCServer, nil, nil)
	if err != nil {
		r.alert.PushDuOSalert("Error", err.Error(), "error")
	}
	r.status.BlockCount = getBlockCount.(int64)
	return r.status.BlockCount
}

func (r *rcvar) GetConnectionCount() int32 {
	r.status.ConnectionCount = r.cx.RPCServer.Cfg.ConnMgr.ConnectedCount()
	return r.status.ConnectionCount
}
