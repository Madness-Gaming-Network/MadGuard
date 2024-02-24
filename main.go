package main

import (
	"sync"

	"github.com/Madness-Gaming-Network/madguard/packages/api"
	"github.com/Madness-Gaming-Network/madguard/packages/chat"
	"github.com/Madness-Gaming-Network/madguard/packages/core"
	// "github.com/Madness-Gaming-Network/madguard/packages/proxy"
	"github.com/Madness-Gaming-Network/madguard/packages/vpn"
)

func main() {
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
		vpn.Create()
	}()
}
