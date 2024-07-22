package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://sepolia.infura.io/v3/e5bdd363a56d4727bbfb0528211a78b3"
)

func main() {
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 38954ab95516c2e4befe5a62c70f50598e7fb8cb
	// d2d5f2b082037ea0bcdb208cca5b7819b0ec6547
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	addr1 := common.HexToAddress("38954ab95516c2e4befe5a62c70f50598e7fb8cb")
	addr2 := common.HexToAddress("d2d5f2b082037ea0bcdb208cca5b7819b0ec6547")

	// fmt.Println("Addr1", addr1)
	// fmt.Println("Addr2", addr2)

	b1, err := client.BalanceAt(context.Background(), addr1, nil)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := client.BalanceAt(context.Background(), addr2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1", b1)
	fmt.Println("Balance 2", b2)
}
