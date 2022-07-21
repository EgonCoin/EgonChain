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
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes//TODO
	"enode://c74dfe5fcdf36fc771c4dfd7e0fdcbbb647a54d55404b0cc475ced7d3991b1f08db965f601505ec069221971ada3cee7a7aaa602efdd7e673806dfb2020c05c8@18.195.185.117:30303",
	"enode://4d1bd814c28671100a5821f1de9f50a3386437ea7d347e0da2b807faa85d83cd9e0446a505f57052a3ebde4c6c5707ead8be0155ff4a2ca560621b569c62cba5@35.170.32.132:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://b2bea79e224b92d4e274b266cc8a7a382372e2fd41e4a7b3ca0ac5983ad8cf182be7cbbb60653cdba80fb6c99920d7ea01bef0fdf7089a684f7dd12830ce6ffb@3.123.66.255:30303",
	"enode://4d2805d28fca021f5cec537d4a36c0fdc4108c65f62040d6d16410cf39d189ceeea6509277a847714406d91e3388371f349d5892fdcc6d357aba349b5abfeda1@34.199.236.89:30303",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
