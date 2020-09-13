package ufo

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xddb7d9fc
	TestnetMagic wire.BitcoinNet = 0xdbb8c0fb
	RegtestMagic wire.BitcoinNet = 0x1c55211b
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{27}
	MainNetParams.ScriptHashAddrID = []byte{68}
	MainNetParams.Bech32HRPSegwit = "uf"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{130}
	TestNetParams.Bech32HRPSegwit = "ut"
}

// UFOParser handle
type UFOParser struct {
	*btc.BitcoinParser
}

// NewUFOParser returns new UFOParser instance
func NewUFOParser(params *chaincfg.Params, c *btc.Configuration) *UFOParser {
	return &UFOParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main UFO network,
// and the test UFO network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to UFO parameters
	// UFO has dual standard of addresses and we want to be able to
	// parse both standards
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
