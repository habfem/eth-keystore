package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	/* key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Address) */
	password := "xxxxxx"
	b, err := ioutil.ReadFile("./wallet/UTC--2023-10-25T14-58-58.703590800Z--17cb1242a96370937944e9e9c21fb92e986cf718")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv", hexutil.Encode(pData))

	pDate := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub", hexutil.Encode(pDate))

	fmt.Println("Add", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
