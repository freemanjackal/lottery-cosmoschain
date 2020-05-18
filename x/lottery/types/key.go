package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	// ModuleName is the name of the module
	ModuleName = "lottery"
	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName
)

var (
	// GlobalStoreKeyPrefix is a prefix for global primitive state variable
	GlobalStoreKeyPrefix = []byte{0x00}

	// OrdersCountStoreKey is a key that help getting to current Lottery count state variable
	LotteryCountStoreKey = append(GlobalStoreKeyPrefix, []byte("LotteryCount")...)

	BetsCountStoreKey = append(GlobalStoreKeyPrefix, []byte("BetsCount")...)

	// ChannelStoreKeyPrefix is a prefix for storing channel
	ChannelStoreKeyPrefix = []byte{0x01}

	// LotteryStoreKeyPrefix is a prefix for storing lottery
	LotteryStoreKeyPrefix = []byte{0x02}

	// BetStoreKeyPrefix is a prefix for storing bet
	BetStoreKeyPrefix = []byte{0x03}
)

// ChannelStoreKey is a function to generate key for each verified channel in store
func ChannelStoreKey(chainName, channelPort string) []byte {
	buf := append(ChannelStoreKeyPrefix, []byte(chainName)...)
	buf = append(buf, []byte(channelPort)...)
	return buf
}

// LotteryStoreKey is a function to generate key for each Lottery in store
func LotteryStoreKey(LotteryID uint64) []byte {
	return append(LotteryStoreKeyPrefix, uint64ToBytes(LotteryID)...)
}

// BetStoreKey is a function to generate key for each Bet in store
func BetStoreKey(BetID uint64, LotteryId uint64) []byte {
	return append(BetStoreKeyPrefix, uint64ToBytes(BetID)...)
}

func uint64ToBytes(num uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, num)
	return result
}

func GetEscrowAddress() sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte("COLLATERAL")))
}
