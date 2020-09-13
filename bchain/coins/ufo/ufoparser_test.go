// +build unittest

package ufo

import (
	"encoding/hex"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func Test_GetAddrDescFromAddress_Testnet(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "P2PKH",
			args:    args{address: "mxfTDuMX9VTp7q21dE6wvZGMjxA53Eogyz"},
			want:    "76a914bc159d4058b6ff99664df850c45c6c936b7b0efc88ac",
			wantErr: false,
		},
		{
			name:    "P2SH",
			args:    args{address: "uQEfJmrW8vd5oPy1xopoKT9QX4xFDb94DL"},
			want:    "a9143a1ba41901d0a2044998ba37c4d5a33b7aff1be587",
			wantErr: false,
		},
	}
	parser := NewUFOParser(GetChainParams("test"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.GetAddrDescFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddrDescFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("GetAddrDescFromAddress() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_GetAddrDescFromAddress_Mainnet(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "P2PKH",
			args:    args{address: "Bxdk1m73MZpBLY9QQUkxME3ycVqcQHGC6Y"},
			want:    "76a9143c6ac0bb6adb7de28a307e572c0ab755b931b91588ac",
			wantErr: false,
		},
		{
			name:    "P2SH",
			args:    args{address: "UdVVV2GaJZBUa1Ahv6K5FRX79TacceJLaF"},
			want:    "a914aa1cedd74795b100ab2aa827667aed982221a1ee87",
			wantErr: false,
		},
		{
			name:    "witness_v0_keyhash",
			args:    args{address: "uf1qynpx7rueyj24t7up0nufm92lscwh55r8aeqf94"},
			want:    "001424c26f0f99249555fb817cf89d955f861d7a5067",
			wantErr: false,
		},
	}
	parser := NewUFOParser(GetChainParams("main"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.GetAddrDescFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddrDescFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("GetAddrDescFromAddress() = %v, want %v", h, tt.want)
			}
		})
	}
}
