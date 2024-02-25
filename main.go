package main

import (
	"log"
	"sync"

	"github.com/Madness-Gaming-Network/MadGuard/packages/api"
	"github.com/Madness-Gaming-Network/MadGuard/packages/chat"
	"github.com/Madness-Gaming-Network/MadGuard/packages/core"

	// "github.com/Madness-Gaming-Network/MadGuard/packages/proxy"
	"github.com/Madness-Gaming-Network/MadGuard/packages/vpn"
)

func main() {
	// TODO: Use https://pkg.go.dev/golang.org/x/sync/errgroup instead?
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		core.Run()
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		api.RunServer()
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		chat.RunBot()
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		client, err := vpn.Create()
		if err != nil {
			log.Fatalf("VPN error: %s\n", err)
		}
		log.Printf("VPN works!: %s\n", client)
	}()
}
