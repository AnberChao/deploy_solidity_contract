package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
)

func TestUpgradeTO(t *testing.T) {
	cli := NewMyTokenClient(t, Key)

	var a []byte
	total, err := cli.Session().UpgradeToAndCall(common.HexToAddress("0xe765da5e9db0ed9f5266d85e43d0e59e82eca2db"), a)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(total.Hash().String())
}

func TestMint(t *testing.T) {
	cli := NewMyTokenClient(t, Key)

	total, err := cli.Session().Mint(common.HexToAddress("0x0"), new(big.Int).SetInt64(int64(1000)))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(total.Hash().String())
}

func TestTransfer(t *testing.T) {
	cli := NewMyTokenClient(t, Key)

	total, err := cli.Session().Transfer(common.HexToAddress("0x0"), new(big.Int).SetInt64(int64(1000)))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(total.Hash().String())
}

func TestBalanceOf(t *testing.T) {
	cli := NewMyTokenClient(t, Key)

	total, err := cli.Session().BalanceOf(common.HexToAddress("0x0"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(total)
}

func TestOwner(t *testing.T) {
	cli := NewMyTokenClient(t, Key)

	total, err := cli.Session().Owner()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(total.String())
}
