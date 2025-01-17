package zilliqa

import (
	"fmt"
	"math/big"

)

type BouncerWallet struct{}

func main() {
	// initialize the Zillean
	zil := zillean.NewZillean("https://api.zilliqa.com")

	// generate a private key
	privKey := zil.GeneratePrivateKey()
	fmt.Printf("private key: %s\n", privKey) // private key: b6c00064b10d33c4a9fadb5b473d834b1995f132acdbe4b831ab5343702c174e

	// get a public key
	pubKey, _ := zil.GetPublicKeyFromPrivateKey(privKey)
	fmt.Printf("public key: %s\n", pubKey) // public key: 03dcb21aaaa918f91a708858dc271343b4bee059e53202ce0358b68effa7e64378

	// get a address
	address, _ := zil.GetAddressFromPrivateKey(privKey)
	fmt.Printf("address: %s\n", address) // address: 5f0e26adf701bb6a4535f0485fe3400e6e90c9ae

	// sign the transaction
	rawTx := zillean.RawTransaction{
		Version:  0,
		Nonce:    2,
		To:       "to address",
		Amount:   "1000000000000",
		PubKey:   pubKey,
		GasPrice: big.NewInt(1000000000),
		GasLimit: 1,
	}
	signature, _ := zil.SignTransaction(rawTx, privKey)
	txID, _ := zil.RPC.CreateTransaction(rawTx, signature)
	fmt.Printf("txID: %s\n", txID)
}
