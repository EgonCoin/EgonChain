// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main EGON network.
var MainnetBootnodes = []string{
	"enode://bba58887951c76e93ec548707c01b5bc81159d914195132d52d962fac98f942b53cbc8bf50b63070641c7bb918598adf90405ddeab7e432c98df02bbfb81a7a0@18.195.185.117:30303",  // egon-mainnet-node-boot-01
	"enode://5ab4fa0a4a9966e1ae152c93defa1064cbad3b266b1af5a4688f5350adbc935a7c3ea85251a7a13986a305a9cec50e95c8a5ade453724e6b52cb0e60ba4364c8@35.170.32.132:30303",  //egon-mainnet-node-boot-02
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://c6b65588a55f73ecc5d02d65b852961fd18347a2be19efc2c6a44088e44059a0b5578ac4beac2997c69dd02b44c0e6d3ebfe069399aba66c43257d19c5dd3ef7@3.123.66.255:30303", // egon-testnet-node-boot-01
	"enode://441c74f4ea5d1fa791d2b49966a9c621117cbe9bf49c22348a01af6478db52bcbba6bcd3c04a53c63945f5d6e89f134b9fc81c724d57510c1269b8ab38776ba3@34.199.236.89:30303",  // egon-testnet-node-boot-02
}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
