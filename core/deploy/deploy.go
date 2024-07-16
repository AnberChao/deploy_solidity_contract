package deploy

import (
	"deploy_solidity_contract/contracts/proxy"
	"deploy_solidity_contract/core/chain"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type IDeployLogic interface {
	DeployLogic() (address common.Address, txId string, err error)
	UpgradeTo(logicAddr common.Address) (err error)
}

type ContractDeploy struct {
	*chain.ChainClient
	logic IDeployLogic
}

func NewContractDeploy(logic IDeployLogic, options ...chain.ChainOption) (*ContractDeploy, error) {

	chainCli, err := chain.NewChainClient(options...)
	if err != nil {
		return nil, err
	}

	return &ContractDeploy{
		ChainClient: chainCli,
		logic:       logic,
	}, nil

}

// DeployAll Do deploy logic contract and proxy contract at the same time.
func (s *ContractDeploy) DeployAll() (logicAddr common.Address, proxyAddr common.Address, err error) {

	logicAddr, txId, err := s.logic.DeployLogic()
	if err != nil {
		return
	}

	if err = s.WaitTx(txId); err != nil {
		return
	}

	proxyAddr, txId, err = s.DeployProxy(logicAddr)
	if err = s.WaitTx(txId); err != nil {
		return
	}

	return
}

// Upgrade to your logic contract
func (s *ContractDeploy) Upgrade() (err error) {
	logicAddr, txId, err := s.logic.DeployLogic()
	if err != nil {
		return
	}

	if err = s.WaitTx(txId); err != nil {
		return
	}
	return s.logic.UpgradeTo(logicAddr)
}

// DeployProxy only deploy proxy contract
func (c *ContractDeploy) DeployProxy(logic common.Address) (address common.Address, txId string, err error) {
	initData, _ := hex.DecodeString("8129fc1c")
	address, txId, err = c.ETHClient().Deploy(func(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, error) {
		address, tx, _, err := proxy.DeployProxy(auth, backend, logic, initData)
		return address, tx, err
	})

	return
}
