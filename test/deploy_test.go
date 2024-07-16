package test

import (
	"deploy_solidity_contract/core/chain"
	"testing"
)

func TestDeploy(t *testing.T) {
	var opts []chain.ChainOption
	opts = append(opts, chain.WithChainURL(URL))
	opts = append(opts, chain.WithChainHexKey(Key))

	s := NewMyTokenLifecycle(opts...)

	if err := s.Deploy(); err != nil {
		t.Fatal(err)
	}

	if err := s.ContractSetting(); err != nil {
		t.Fatal(err)
	}

	s.ShowAddress()
}
