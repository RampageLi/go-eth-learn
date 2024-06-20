package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// var infuraUrl = "https://mainnet.infura.io/v3/e5bdd363a56d4727bbfb0528211a78b3"

var ganacheUrl = "http://127.0.0.1:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheUrl)
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}
	fmt.Println("The block number: ", block.Number())

	addr := "0x847f793D003cA54d7C62640833B3e67e0995d5d5"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get balance: %v", err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	balanceEth := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance: ", balanceEth)
}
