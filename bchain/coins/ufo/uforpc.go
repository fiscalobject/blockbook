package ufo

import (
	"encoding/json"

	"github.com/golang/glog"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// UFORPC is an interface to JSON-RPC bitcoind service.
type UFORPC struct {
	*btc.BitcoinRPC
}

// NewUFORPC returns new UFORPC instance.
func NewUFORPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &UFORPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV2{}
	s.ChainConfig.SupportsEstimateFee = false

	return s, nil
}

// Initialize initializes UFORPC instance.
func (b *UFORPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewUFOParser(params, b.ChainConfig)

	// parameters for getInfo request
	if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	} else {
		b.Testnet = true
		b.Network = "testnet"
	}

	glog.Info("rpc: block chain ", params.Name)

	return nil
}
