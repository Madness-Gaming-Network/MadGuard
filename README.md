# Discord Bot (name pending)

tl;dr A Discord bot that will manage the lifecycles of adhoc Wireguard networks for small to large groups of gamers. When in doubt, "shift left" and "minimize".

TODO:

- Set up some way of distributing this to server that doesn't require manually building
  - https://goreleaser.com/ci/actions/
  - https://github.com/marketplace/actions/go-release-binaries

Decisions:

- Golang, because:
  - Wireguard is built with it
  - Way more documentation, resources, and examples
  - Simpler syntax
    - Supposedly better for ChatGPT (GitHub Copilot) prompting too?
- Simplest possible infrastructure
  - Focus on getting MVP working first, then automating
  - Minimize dependencies
  - No Kubernetes until absolutely neccessary

Potential commands:

- Create/read/update/delete (CRUD) Wireguard network
  - Inputs: Network name, users, desired expiration, ???
- CRUD user from existing Wireguard network
  - Inputs: Email, IP range, games, ???
- CRUD server/session in emulator
  - Inputs: ???

Other features:

- Pull Discord roles and use those for permissions

Ideas/questions:

- Host on same VPS as Wireguard server?
  - Probably not, maybe Aaron's server?
- Discord donation/sponsorship/boosting gives more features?
  - Longer network expiration?
  - More players
  - Etc?

Links:

- https://discord.com/developers/docs/topics/community-resources
  - https://github.com/bwmarrin/discordgo
  - https://github.com/Amatsagu/Tempest