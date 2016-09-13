package main

import (
	"flag"
	"net"
	"strings"
	"time"
)

// A Peer represents
type Peer struct {
	addr net.IP
	name string
	time time.Time
}

// A Cluster represents a group of fault-tolerant replicating nodes.
type Cluster struct {
	engine    *Engine
	operating bool
	peers     []Peer
	port      int
}

// NewCluster creates and initializes the Cluster struct.
func NewCluster(engine *Engine) (cluster *Cluster) {
	ipstr := *flag.String("--peers", "",
		"comma-separated string of IP addresses")
	items := strings.Split(ipstr, ",")
	count := len(items)
	peers := make([]Peer, count)

	for index := 0; index < count; index++ {
		peer := Peer{}
		peer.addr = net.ParseIP(items[index])
		peers[index] = peer
	}

	port := *flag.Int("--port", 13579,
		"integer, port which clients and peers communicate over")

	return &Cluster{engine: engine, operating: false, peers: peers, port: port}
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
	duration := time.Millisecond * time.Duration(*flag.Int64("interval", 500, "integer, milliseconds"))

	for {
		for _, peer := range cluster.peers {
			cluster.poll(&peer)
		}

		if cluster.operating {
			time.Sleep(duration)
		} else {
			break
		}
	}
}

func (cluster *Cluster) poll(peer *Peer) {
	_, err := net.Dial("tcp", (*peer).addr.String())
	if err == nil {
		// Have a chat with the peer. See docs/Replication
	}
}
