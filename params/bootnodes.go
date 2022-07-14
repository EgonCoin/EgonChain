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
	"enode://02f9ea3ea8f1c29fcf546fb579751243f93e914c63c135929f9bcf52ccde5f714630fe189108d0cfe4b470d5f67705af1adce4bcf6359df55b0e1373288a24ad@18.195.185.117:30303",  // egon-mainnet-node-boot-01
	"enode://612df9ad10f607bb0eaf6d4770eed8a7188ee98a294763e5ce8c4c0a9ac7478eaf1ecb2eaa14d15f91fae51fbf2a4c594d34a1e24e18dcf46f3940a373bbe3cc@35.170.32.132:30303",  //egon-mainnet-node-boot-02
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://45083c078484a87ffcd24e839b1b2db8d38014a9530f1b26d0d99d8feff8ead2bfaddbb4c05164aacfbe2ff52f3f33d6a36eb8737fb71737a6f218ca7ddfadca@3.123.66.255:30303", // egon-testnet-node-boot-01
	"enode://003663afbc1d6b6a3660ab78af9b1fdf9421ee5b4a7438824dadfa22c08efdb80c2ff1a083c0dbda55ea0cb655680e35f85468cadfdd86c5614d0c5c7d3d3d4e@34.199.236.89:30303",  // egon-testnet-node-boot-02
}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
