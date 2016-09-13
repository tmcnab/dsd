package main

import (
	"net"
	"time"
)

// A _Peer represents
type _Peer struct {
	addr net.Addr
	name string
	time time.Time
}

// A Cluster represents a group of fault-tolerant replicating nodes.
type Cluster struct {
	peers     _Peer
	operating bool
}

// Start the cluster synchronizing with peer nodes.
func (cluster *Cluster) Start() (err error) {
	cluster.operating = true
	go cluster.loop()
	return nil
}

// Stop the cluster from listening.
func (cluster *Cluster) Stop() {
	cluster.operating = false
}

func (cluster *Cluster) loop() {
	for {
		// ask a peer for data

		if !cluster.operating {
			break
		}
	}
}
