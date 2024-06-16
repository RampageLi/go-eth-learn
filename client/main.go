package main

import (
	"context"
	"fmt"
	"log"

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

	fmt.Println(block.Number())
}
