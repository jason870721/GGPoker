package main

import (
	"fmt"

	"github.com/jason870721/GGPoker/p2p"
)

func main() {
	// ================================================
	// following code mock a remote gamer's PC
	// like other player start a client on his/her computer.
	remoteUserCfg := p2p.ServerConfig{
		Name:       "Stevie",
		Version:    "GGPOKER V0.1-alpha",
		ListenAddr: ":3000",
	}
	remoteUserClient := p2p.NewServer(remoteUserCfg)
	remoteUserClient.Start()
	// ================================================

	// ================================================
	// start my PC client and try to connect to other player who listen on 3000 port.
	myCfg := p2p.ServerConfig{
		Name:       "Kai",
		Version:    "GGPOKER V0.1-alpha",
		ListenAddr: ":4000",
	}
	myClient := p2p.NewServer(myCfg)
	myClient.Start()

	if err := myClient.Connect(":3000"); err != nil {
		fmt.Printf("connect failed: %v", err)
	}

	for {
	}

	//fmt.Println(deck.New())
}
