package main

import (
	"github.com/bcext/gcash/chaincfg"
	"github.com/bcext/gcash/chaincfg/chainhash"
)

var (
	param = &chaincfg.TestNet3Params
)

const (
	defaultSignatureSize = 107
	defaultSequence      = 0xffffffff
)

type utxo struct {
	hash   *chainhash.Hash
	value  int64
	script []byte
}

func main() {

}
