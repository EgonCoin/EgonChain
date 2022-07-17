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
	// Ethereum Foundation Go Bootnodes
	"enode://d042cdd576e9d6b9deaf92b4fa5d3c40b2b194f28b2b830e14f68b684a7a4ea15bb4134e6619d9ec89bd50adb527bbb9feae713371169ce9590c214343efc54e@18.195.185.117:30303",
	"enode://f61cfc11069955dc4f91f11bcd0985803f63e6d0d432326f3dda7de2981726a6facbb15f2193dbe549125c2b0a7011472016b86be81904428f714c3fbd32a679@35.170.32.132:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	"enode://c95f6f34a032f15d10125be9a30f2a5b3083a2db5ce8b33f23671bd4f75bb7587f70e91aecb0fc16283c699be22c4f89467be663607fcc582ab741ded29f64b0@3.123.66.255:30303",
	"enode://551e1fd6c879790d9973c3f2e3f4be2dda46533227588a572713739e1eaa19ae952bc320a1923cfe4e6b7b3435c1f06d7a0c47b9ace0c44d831cd8809afda84f@34.199.236.89:30303",
}

var V5Bootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
