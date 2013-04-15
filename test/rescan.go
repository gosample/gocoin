package main

import (
	"fmt"
	"encoding/hex"
	"github.com/piotrnar/gocoin/btc"
	"github.com/piotrnar/gocoin/btc/leveldb"
	"flag"
)

var testnet *bool = flag.Bool("t", false, "use testnet")
var rescan *bool = flag.Bool("r", false, "rescan")
var listunspent *bool = flag.Bool("l", false, "list unspent")

var GenesisBlock *btc.Uint256

func main() {
	flag.Parse()
	var addr string

	if *testnet { // testnet3
		GenesisBlock = btc.NewUint256FromString("000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943")
		leveldb.Testnet = true
		addr = "mwZSC78JGfS6NY7R57aFeJQp4HgRCadHze"
	} else {
		GenesisBlock = btc.NewUint256FromString("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
		//addr = "19vPUYV7JE45ZP9z11RZCFcBHU1KXpUcNv"
		//addr = "1MBfj713pjbtFHeegKwxrB8oZwYPHgC9mL"
		addr = "13tzBMErCWcdBDvM69fxbrF9nWGdi99cMY"
	}

	//btc.TestRollback = true
	chain := btc.NewChain(GenesisBlock, *rescan)
	
	if *listunspent {
		println(chain.Stats(), addr)

		a, e := btc.NewAddrFromString(addr)
		if e != nil {
			println(e.Error())
			return
		}
		fmt.Println(hex.EncodeToString(a.OutScript()[:]))
		unsp := chain.GetUnspentFromPkScr(a.OutScript())
		var sum uint64
		for i := range unsp {
			fmt.Println(unsp[i].Output.String())
			sum += unsp[i].Value
		}
		fmt.Printf("Total %.8f unspent BTC at address %s\n", float64(sum)/1e8, a.Enc58str);
	}
}

