//package awesomeEthereumClient
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		 log.Fatal("Woops something went wrong")
	}
	ctx := context.Background()
	tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0xb54ca074e3c1390d0a075dd6acd847abd6659c42eadcebaa84dd90dc90f34a8a"))
	if (!pending) {
		fmt.Println(tx)
	}
}
