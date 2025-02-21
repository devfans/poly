/*
 * Copyright (C) 2021 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"fmt"
	"github.com/polynetwork/poly/common"
	"github.com/polynetwork/poly/common/config"
)

type BtcNetType int

const (
	TyTestnet3 BtcNetType = iota
	TyRegtest
	TySimnet
	TyMainnet
)

var (
	BYTE_FALSE = []byte{0}
	BYTE_TRUE  = []byte{1}

	HeaderSyncContractAddress, _        = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
	CrossChainManagerContractAddress, _ = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03})
	SideChainManagerContractAddress, _  = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04})
	NodeManagerContractAddress, _       = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05})
	RelayerManagerContractAddress, _    = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06})
	Neo3StateManagerContractAddress, _  = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07})
	ReplenishContractAddress, _         = common.AddressParseFromBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09})

	VOTE_ROUTER             = uint64(0)
	BTC_ROUTER              = uint64(1)
	ETH_ROUTER              = uint64(2)
	ONT_ROUTER              = uint64(3)
	NEO_ROUTER              = uint64(4)
	COSMOS_ROUTER           = uint64(5)
	BSC_ROUTER              = uint64(6)
	HECO_ROUTER             = uint64(7)
	QUORUM_ROUTER           = uint64(8)
	ZILLIQA_LEGACY_ROUTER   = uint64(9)
	MSC_ROUTER              = uint64(10)
	NEO3_LEGACY_ROUTER      = uint64(11)
	OKEX_ROUTER             = uint64(12)
	NEO3_ROUTER             = uint64(14)
	POLYGON_HEIMDALL_ROUTER = uint64(15)
	POLYGON_BOR_ROUTER      = uint64(16)
	ZILLIQA_ROUTER          = uint64(17)
	STARCOIN_ROUTER         = uint64(18)
	PIXIECHAIN_ROUTER       = uint64(19)
	HSC_ROUTER              = uint64(20)
	HARMONY_ROUTER          = uint64(21)
	BYTOM_ROUTER            = uint64(22)
)

//Check router StartBlock to prevent hard forks
func CheckRouterStartBlock(router uint64, block uint32) (err error) {
	var startBLock uint32
	switch router {
	case HARMONY_ROUTER, HSC_ROUTER, BYTOM_ROUTER:
		if config.DefConfig.P2PNode.NetworkId == config.NETWORK_ID_MAIN_NET {
			startBLock = 18823000
		}
	}
	if startBLock > 0 && block < startBLock {
		return fmt.Errorf("not a supported router:%d", router)
	}
	return
}