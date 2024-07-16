package test

import (
	"deploy_solidity_contract/core/chain"
	"deploy_solidity_contract/core/mytoken"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func NewMyTokenClient(t *testing.T, key string) *mytoken.MyTokenClient {
	var opts []chain.ChainOption
	opts = append(opts, chain.WithChainURL(URL))
	opts = append(opts, chain.WithChainHexKey(key))
	opts = append(opts, chain.WithMyTokenProxyAddress(common.HexToAddress(_MyTokenAddress)))

	logicCli, err := mytoken.NewMyTokenClient(opts...)
	if err != nil {
		t.Fatal(err)
	}
	_, err = logicCli.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	return logicCli
}
