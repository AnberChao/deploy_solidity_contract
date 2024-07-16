package mytoken

import (
	m "deploy_solidity_contract/contracts/mytoken"
	"deploy_solidity_contract/core/chain"
)

type MyTokenClient struct {
	*chain.ChainClient

	session *m.MyTokenSession
}

func (s *MyTokenClient) Session() *m.MyTokenSession {
	return s.session
}

func NewMyTokenClient(options ...chain.ChainOption) (*MyTokenClient, error) {

	chainCli, err := chain.NewChainClient(options...)
	if err != nil {
		return nil, err
	}

	return &MyTokenClient{
		ChainClient: chainCli,
	}, nil
}
