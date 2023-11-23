package main

import (
	"log"

	"github.com/datachainlab/ethereum-ibc-relay-chain/pkg/relay/ethereum"
	"github.com/datachainlab/ethereum-ibc-relay-chain/pkg/relay/ethereum/signers/hd"
	ethereumlc "github.com/datachainlab/ethereum-ibc-relay-prover/relay"
	lcp "github.com/datachainlab/lcp-go/relay"
	lcptm "github.com/datachainlab/lcp-go/relay/tendermint"
	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		ethereum.Module{},
		ethereumlc.Module{},
		hd.Module{},
		lcp.Module{},
		lcptm.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
