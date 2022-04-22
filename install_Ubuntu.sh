#!/bin/bash
echo ""
echo "Dev-dependencies will now be installed for Ubuntu 18.04"
echo ""

sudo apt-get update && sudo apt-get upgrade -y

sudo apt-get install build-essential software-properties-common make curl -y

echo ""
echo "Installing Go"
curl -OL https://go.dev/dl/go1.17.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.17.6.linux-amd64.tar.gz
rm -rf go1.17.6.linux-amd64.tar.gz

echo ""
echo "Build EgonCoin"
make all

echo ""
echo "Installation is now complete"
