// Copyright 2015 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package utils

import "github.com/uther/go-uther/p2p/discover"

// FrontierBootNodes are the enode URLs of the P2P bootstrap nodes running on
// the Frontier network.
var FrontierBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
		discover.MustParseNode("enode://50b54003aed5ede7716831769224bc4cbd892ac92eec277e115f4c0664c553b7f0260286e9b24c3d1405f226fad0867e1837a01530bac729ffb821155fca3ccf@128.199.232.240:41303"), 		
		discover.MustParseNode("enode://fc0cd3cd26ac6e58b7350ce1cf0d7116c0153caa758d1cebc93c5afb8c4d1b9fc1d7af7b233e5927f2ae7e365a3d1de6565224633c6e47bf3d95af167fbd2ece@158.69.195.137:41303"),	
		discover.MustParseNode("enode://1f9ae6e8dfb9baf0980bea8bbd0eaf0a19a568e25f70843bc70eee6e8adc79e673af88392660b770e0e28c721e98e34aca81553212acdc123f42a71f15d88294@103.230.120.198:41303"),					
		discover.MustParseNode("enode://96af71904742ab9cf73af953968aafe16852d2586f399165c0d4e82e6a751182901f0a3ba9be8aa64a0b51247b4ba98ef2b7bf4c98be5ab8b497288e6ef3e857@37.59.24.15:41303"),		

	// ETH/DEV Cpp Bootnodes
	//discover.MustParseNode("enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303"),
}

// TestNetBootNodes are the enode URLs of the P2P bootstrap nodes running on the
// Morden test network.
var TestNetBootNodes = []*discover.Node{
	// ETH/DEV Go Bootnodes
		discover.MustParseNode("enode://1f9ae6e8dfb9baf0980bea8bbd0eaf0a19a568e25f70843bc70eee6e8adc79e673af88392660b770e0e28c721e98e34aca81553212acdc123f42a71f15d88294@103.230.120.198:41303"),

	// ETH/DEV Cpp Bootnodes
}
