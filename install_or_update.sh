#!/bin/bash

# This script will install or update a previous installation of go-mirror-archlinux
if [[ $EUID -ne 0 ]]; then
   echo "This script must be run as root" 
   exit 1
fi

if [ -e /etc/go-mirror-archlinux/config.json ]
then
    echo "Pre-existing configuration found. Updating binaries only..."
else
    echo "No pre-existing configuration found. Creating /etc/go-mirror-archlinux/config.json..."
    mkdir /etc/go-mirror-archlinux/
    cp ./go-mirror-archlinux/config.json /etc/go-mirror-archlinux/
fi

echo "Creating go-mirror-archlinux user..."
useradd -m -s /bin/false go-mirror-archlinux

echo "Stopping service if running..."
systemctl stop go-mirror-archlinux

echo "Updating binaries..."
rm /usr/bin/go-mirror-archlinux
cp ./go-mirror-archlinux/go-mirror-archlinux /usr/bin/go-mirror-archlinux

echo "Updating systemctl unit..."
echo "
[Unit]
Description=An ArchLinux mirroring service!
After=network.target
[Service]
Type=simple
User=go-mirror-archlinux
Group=go-mirror-archlinux
ExecStart=/usr/bin/go-mirror-archlinux -config=/etc/go-mirror-archlinux/config.json
[Install]
WantedBy=multi-user.target" > ./go-mirror-archlinux.service

mv ./go-mirror-archlinux.service /etc/systemd/system/

systemctl daemon-reload
systemctl start go-mirror-archlinux
