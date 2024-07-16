package mytoken

import (
	m "deploy_solidity_contract/contracts/mytoken"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *MyTokenClient) NewSession() (*m.MyTokenSession, error) {
	if s.session != nil {
		return s.Session(), nil
	}
	session, err := m.NewMyToken(s.GetMyTokenProxyAddress(), s.ETHClient().Client())
	if err != nil {
		return nil, err
	}
	transactOpts := s.ETHClient().GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	s.session = &m.MyTokenSession{Contract: session, CallOpts: callOpts, TransactOpts: *transactOpts}
	return s.Session(), err
}

func (s *MyTokenClient) DeployLogic() (address common.Address, txId string, err error) {

	address, txId, err = s.ETHClient().Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		address, tx, _, err := m.DeployMyToken(auth, backend)
		return address, tx, err
	})

	return
}

func (s *MyTokenClient) UpgradeTo(logicAddr common.Address) (err error) {
	fmt.Println("UpgradeTo ", s.GetMyTokenProxyAddress().String(), logicAddr.String())

	session, err := s.NewSession()
	if err != nil {
		return
	}

	var a []byte
	tx, err := session.UpgradeToAndCall(logicAddr, a)
	if err != nil {
		return
	}
	return s.WaitTx(tx.Hash().String())
}
