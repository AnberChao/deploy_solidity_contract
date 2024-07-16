package chain

import (
	"context"
	"deploy_solidity_contract/client"
	"deploy_solidity_contract/core/event"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

const (
	URL        = ""
	Key        = ""
	KeyAddress = ""
)

func NewChainClient(opts ...ChainOption) (*ChainClient, error) {
	conf := NewDefConfig()
	for _, option := range opts {
		err := option(conf)
		if err != nil {
			return nil, err
		}
	}

	ethCli, err := ethclient.Dial(conf.url)
	if err != nil {
		return nil, err
	}

	cli, err := client.NewETHClient(ethCli, conf.key)
	if err != nil {
		return nil, err
	}

	c := &ChainClient{
		cli:                 cli,
		timeout:             defTimeout,
		myTokenProxyAddress: conf.myTokenProxyAddress,
	}
	c.events = make(map[string]*event.ContractEvent)
	c.events[c.myTokenProxyAddress.String()] = event.NewMyTokenContractEvent(c.myTokenProxyAddress.String())

	return c, nil

}

type ChainClient struct {
	cli     *client.ETHClient
	timeout time.Duration

	events              map[string]*event.ContractEvent
	myTokenProxyAddress common.Address
}

func (c *ChainClient) GetMyTokenProxyAddress() common.Address {
	return c.myTokenProxyAddress
}

func (s *ChainClient) ETHClient() *client.ETHClient {
	return s.cli
}
func (s *ChainClient) WaitTx(txId string) error {
	return s.waitTx(txId, false)
}
func (s *ChainClient) WaitTxAndShowEvent(txId string) error {
	return s.waitTx(txId, true)
}

func (s *ChainClient) waitTx(txId string, showEvent bool) error {
	tick := time.NewTicker(s.timeout)
	for {
		select {
		case <-tick.C:
			return errors.New("wait time out")
		default:
			_, p, err := s.cli.Client().TransactionByHash(context.Background(), common.HexToHash(txId))
			if err != nil {
				return err
			}
			if !p {
				receipt, err := s.cli.Client().TransactionReceipt(context.Background(), common.HexToHash(txId))
				if err != nil {
					return err
				}
				fmt.Println("tx status:", receipt.Status)
				if receipt.Status != 1 {
					return errors.New("tx exec failure ")
				} else {
					fmt.Println("tx success！")
					if showEvent {
						s.ShowLogs(receipt.Logs)
					}
					return nil
				}

			}
			time.Sleep(time.Second)
		}
	}
}

func (s *ChainClient) ShowEvent(txId string) error {
	receipt, err := s.cli.Client().TransactionReceipt(context.Background(), common.HexToHash(txId))
	if err != nil {
		return err
	}

	return s.ShowLogs(receipt.Logs)
}

func (s *ChainClient) ShowLogs(logs []*types.Log) error {
	if len(logs) == 0 {
		return nil
	}
	for _, log := range logs {
		ce, ok := s.events[log.Address.String()]
		if ok {
			ce.ShowLog(log)
		}
	}

	return nil
}
