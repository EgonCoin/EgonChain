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
	"enode://7503b589d2052e396ab3e5cf53c4f65475ae5ab36f17fa1385800fe8ea8fae8338fe953c893ab16165dda616f9f1a98a72cb7c3db5e0b7cc5aaced9eaee9fdf9@18.195.185.117:30303",  // egon-mainnet-node-boot-01
	"enode://b9954bc65bb04337e18e6dd788c6e20cdd7fb2336feeb397e5607d5b989ef4271f091c387d4913fbf3a65ac09a2ca491c69bda9d0bb0f67a4797fd44b8bc4458@35.170.32.132:30303",  //egon-mainnet-node-boot-02
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://d68153cd7b790598bcfe6305b23d2d26eefea17f01c4a87c2bfb08c75a0c48a198acc9689cfcfa023a7b2cf809e064a27f58574a93560362bb6ae22853d0244c@3.123.66.255:30303", // egon-testnet-node-boot-01
	"enode://4413034f877eca3fb2c2dcd40354d02b418062a232f3f82fc268115c9bcec2f9c54d8de898fea9838447033fe7ca8b57f77df4cf049f8d8fef6602f12d1b6c87@34.199.236.89:30303",  // egon-testnet-node-boot-02
}

var V5Bootnodes []string

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
