package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}

	pData := crypto.FromECDSA(pKey)
	fmt.Println(hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&pKey.PublicKey)
	fmt.Println(hexutil.Encode(pubData))

	addr := crypto.PubkeyToAddress(pKey.PublicKey).Hex()
	fmt.Println(addr)
}
