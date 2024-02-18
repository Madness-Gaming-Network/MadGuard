# Discord Bot (name pending)

tl;dr A Discord bot that will manage the lifecycles of adhoc Wireguard networks for small to large groups of gamers. When in doubt, "shift left" and "minimize".

TODO:

- Set up some way of distributing this to server that doesn't require manually building
  - https://github.com/marketplace/actions/go-release-binaries
- https://github.com/gin-gonic/gin/blob/master/docs/doc.md
- Figure out how to automate release, basic process is like:
  ```
  git tag -a v0.1.0 -m "First release"
  git push origin v0.1.0
  ```
  - It would be helpful to automate the version incrementing part, but for now it will be manual
- Since this is a private repo, we'll also need to figure out how to get the correct binaries to the VPS
  - Ideally we could just do something like `wget https://github.com/Madness-Gaming-Network/discord_bot/releases/download/v0.1.0/discord_bot_Linux_x86_64.tar.gz`
  - Auth is generally a pain on remote servers, so we'll have to research or something
  - For now we can just manually `scp` the tarball/binary over

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

Useful commands

```
# Download the release (requires public repo, as mentioned in TODO section)
wget https://github.com/Madness-Gaming-Network/discord_bot/releases/download/v0.1.0/discord_bot_Linux_x86_64.tar.gz

# Extract `discord_bot` executable from GitHub release tarball
tar -xvzf discord_bot_Linux_x86_64.tar.gz

# Set up SSH keys on server (will prompt for root user password on server)
ssh-copy-id -i ~/.ssh/madness-server-ssh-key.pub root@142.171.162.80

# SSH into server with specific SSH key (will ask for SSH key password)
ssh -i ~/.ssh/madness-server-ssh-key root@142.171.162.80

# Copy over `discord_bot` executable
scp -i ~/.ssh/madness-server-ssh-key root@142.171.162.80:/root/discord_bot ./discord_bot

# Run it (on server)
./discord_bot
```