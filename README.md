# go-mirror-archlinux
A go project to quickly setup an Archlinux mirror with automatic updating.

## Requirements

- rsync
    - `pacman -S rsync` (Arch)
    - `apt install rsync` (Ubuntu)

## Installation

Run `install_or_update.sh` as root. It will create a user and a service (go-mirror-archlinux). The whole thing can take sometime to get online since the initial clone is a bit long.

Edit `/etc/go-mirror-archlinux/config.json` to change the following parameters if the default configuration is suitable for you.

|||
|:-----|-----|
| PORT | The port the mirror will be running on. |
| REPO_DIRECTORY | The directory in which the 50GB-ish of content will be. |
| PRIMARY_SERVER | The server you will be syncing from. A exhaustive list can be found [here]()|
| BACKUP_SERVER | In case the primary one fails. |
| BANDWITDH_LIMIT_KB | The bandwith limit, in KB |
| SYNC_INTERVAL_HOURS | How often to sync from the Tier 1 servers. Should be at least once a day. |
| SYNC_ISO | Sync the "iso" folder |
| SYNC_OTHER | Sync the "other" folder |
| SYNC_SOURCES | Sync the "sources" folder |
|||