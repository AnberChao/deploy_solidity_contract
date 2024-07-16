package event

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	_MyTokenProxyAddress = "0x1d9215D7dD5228b3897114468276f5D193826b50"
)

type ContractEvent struct {
	contractAddress common.Address
	contractName    string
	contractAbi     *abi.ABI

	events map[string]EventOption
}

func (c *ContractEvent) ContractName() string {
	return c.contractName
}

func (c *ContractEvent) ContractAddress() common.Address {
	return c.contractAddress
}

func (c *ContractEvent) ShowLog(log *types.Log) {
	id := log.Topics[0]
	event, err := c.contractAbi.EventByID(id)
	if err != nil {
		return
	}
	eop, ok := c.events[event.Name]
	if ok {
		data := eop.DataStruct()
		err := c.UnPackEvent(log, data, event)
		if err != nil {
			return
		}
		msg := fmt.Sprintf("%s(%s):%s(%s)", c.contractName, c.contractAddress, event.Name, eop.ShowData())
		fmt.Println(msg)
	} else {
		msg := fmt.Sprintf("%s(%s):%s(%s)", c.contractName, c.contractAddress, event.Name, "none")
		fmt.Println(msg)
	}
}

func (c *ContractEvent) UnPackEvent(log *types.Log, data interface{}, event *abi.Event) error {
	if len(log.Data) > 0 {
		err := c.contractAbi.UnpackIntoInterface(data, event.Name, log.Data)
		if err != nil {
			return err
		}
	}

	var indexed abi.Arguments
	for _, arg := range event.Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(data, indexed, log.Topics[1:])
}
