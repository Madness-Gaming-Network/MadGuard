package main

import (
	"sync"

	"github.com/Madness-Gaming-Network/discord_bot/packages/discord"
	"github.com/Madness-Gaming-Network/discord_bot/packages/pocketbase"
	"github.com/Madness-Gaming-Network/discord_bot/packages/server"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		pocketbase.RunPocketbase()
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		server.RunServer()
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		discord.RunBot()
	}()

	waitGroup.Wait()
}
