package event

import (
	m "deploy_solidity_contract/contracts/mytoken"
	"github.com/ethereum/go-ethereum/common"
)

func NewMyTokenContractEvent(address string) *ContractEvent {
	e := &ContractEvent{
		contractAddress: common.HexToAddress(address),
		contractName:    "MyToken",
	}

	abi, _ := m.MyTokenMetaData.GetAbi()
	e.contractAbi = abi
	e.events = make(map[string]EventOption)

	return e
}
