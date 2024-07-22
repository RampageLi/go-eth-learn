package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "passwordd"
	// a, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(a.Address)
	b, err := os.ReadFile("./wallet/UTC--2024-07-22T08-50-32.740334000Z--554dcfd4075863ec5de635c6a8496d1d48909c18")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("priv", hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("pub", hexutil.Encode(pubData))

	fmt.Println("addr", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
