package main

import (
	"time"

	"github.com/tetrex/custom-blockchain/network"
)

// server
// Transport layer ( tcp , udp , sockets )
// block
// txn
// keypairs

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello from remote"))
			time.Sleep(1 * time.Second)
		}

	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}
