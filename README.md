# MadGuard

<img src="./pb_public/images/madguard.png" width="512" height="512" />

## TODO

- Learn Gin again: https://github.com/gin-gonic/gin/blob/master/docs/doc.md
- Figure out how to automate release, basic process is like: (change numbers based on semantic versioning)
  ```
  git tag -a v0.1.0 -m "First release"
  git push origin v0.1.0
  ```
  - It would be helpful to automate the version incrementing part, but for now it will be manual
- We'll also need to figure out how to automate getting the correct binaries to the VPS
  - We could just do something like `wget https://github.com/Madness-Gaming-Network/madguard/releases/download/v0.1.0/madguard_Linux_x86_64.tar.gz`
  - Auth is generally a pain on remote servers, so we'll have to research or something
  - For now we can just manually `scp` the tarball/binary over
- Set up basic Caddy config: https://github.com/caddyserver/caddy
- Figure out basic unit/integration test setup
- Document systemd setup for auto restarts (instead of something like pm2 for Node)
  - https://www.linux.org/docs/man1/systemd-run.html
  - Or maybe use https://github.com/cloudflare/tableflip ?
- Figure out how to "publish" or distribute bot?
  - https://discord.com/developers/docs/getting-started
  - https://discord.com/developers/applications

## Decisions

- Golang, because:
  - Wireguard, Caddy, Pocketbase, and more are built with it
  - Way more documentation, resources, and examples
  - Simpler syntax
    - Supposedly better for ChatGPT (GitHub Copilot) prompting too?
- Simplest possible infrastructure
  - Focus on getting MVP working first, then automating
  - Minimize dependencies
  - No Kubernetes until absolutely neccessary, regular VPS deploys are fine
  - When in doubt, "shift left", "reduce scope when possible", and "the best code is no code at all"

## Potential Commands

- Create/read/update/delete (CRUD) Wireguard network
  - Only for trusted users, otherwise they have to request networksk manually (automating too much early on could lead to easier abuse)
  - Inputs: Network name, users, desired expiration, ???
- CRUD user from existing Wireguard network
  - Inputs: Email, IP range, games, ???
- CRUD server/session in emulator
  - Inputs: ???


### Other Feature Ideas / Questions

- Pull Discord roles and use those for permissions
  - Mirror some anonymous Discord user data in Pocketbase, like Discord ID and permissions
- Host on same VPS as Wireguard server?
  - Probably not, maybe separate server?
- Discord donation/sponsorship/boosting gives more features?
  - https://discord.com/developers/docs/monetization/overview
  - Longer network expiration?
  - More players
  - Etc?
- Consider serverless?
  - Only as last resort to minimize cost


## Useful Commands

```
# Download the release (requires public repo, as mentioned in TODO section)
wget https://github.com/Madness-Gaming-Network/madguard/releases/download/v0.1.0/discord_bot_Linux_x86_64.tar.gz

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
