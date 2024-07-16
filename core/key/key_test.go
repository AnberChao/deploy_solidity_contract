package store

import (
	"deploy_solidity_contract/comm"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestGenKey(t *testing.T) {

	key, _ := crypto.GenerateKey()

	address := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println(address)

	keyB := crypto.FromECDSA(key)

	keyHex := hex.EncodeToString(keyB)

	comm.WriteFile([]byte(keyHex), fmt.Sprintf("./account/%s.key", address.String()), true)

}
