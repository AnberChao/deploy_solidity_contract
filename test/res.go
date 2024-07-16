package test

import (
	"context"
	"deploy_solidity_contract/client"
	"deploy_solidity_contract/contracts/msgSmart"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
	"testing"
)

// 在没有根本abi生成合约类的时候，解析合约事件（log.data）
func res(t *testing.T) {
	client, err := ethclient.Dial(URL)
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0xabbfcd4920363c4df65e481e7380dd25fe2d923e9feacfb9379693bf53ee454d")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	data := receipt.Logs[0].Data

	const abiJson = `[
	{
		"inputs": [],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "address",
				"name": "previousAdmin",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "newAdmin",
				"type": "address"
			}
		],
		"name": "AdminChanged",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "beacon",
				"type": "address"
			}
		],
		"name": "BeaconUpgraded",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "OwnershipTransferred",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "caller",
				"type": "address"
			},
			{
				"components": [
					{
						"internalType": "string",
						"name": "serialNo",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromBankID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toBankID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromCoin",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toCoin",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "amount",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_PayerWallet",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "tb_PayeeWallet",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "rate",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "remarks",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_txID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_txDate",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_payeeWallet",
						"type": "string"
					},
					{
						"internalType": "int8",
						"name": "fb_status",
						"type": "int8"
					},
					{
						"internalType": "int8",
						"name": "tb_status",
						"type": "int8"
					},
					{
						"internalType": "string",
						"name": "fb_remarks",
						"type": "string"
					}
				],
				"indexed": false,
				"internalType": "struct IMsgSmart.TXMsg",
				"name": "txMsg",
				"type": "tuple"
			}
		],
		"name": "SendSwapTx",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "caller",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "account",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "bool",
				"name": "approval",
				"type": "bool"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "remarks",
				"type": "string"
			}
		],
		"name": "SetOperator",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "caller",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "serialNo",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "toBankID",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tb_txID",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tb_txDate",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tb_payerWallet",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tb_payeeWallet",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "amount",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "int8",
				"name": "tb_status",
				"type": "int8"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tb_remarks",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "swapNo",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "swapDate",
				"type": "uint256"
			}
		],
		"name": "SetSwapTx",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "serialNo",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "int8",
				"name": "fb_status",
				"type": "int8"
			},
			{
				"indexed": false,
				"internalType": "int8",
				"name": "tb_status",
				"type": "int8"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "rate",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "tb_amount",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "refund",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "reDate",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "remark",
				"type": "string"
			}
		],
		"name": "UpdateSwapTx",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "implementation",
				"type": "address"
			}
		],
		"name": "Upgraded",
		"type": "event"
	},
	{
		"inputs": [],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "renounceOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"components": [
					{
						"internalType": "string",
						"name": "serialNo",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromBankID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toBankID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromCoin",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toCoin",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "amount",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_PayerWallet",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "tb_PayeeWallet",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "rate",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "remarks",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_txID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_txDate",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fb_payeeWallet",
						"type": "string"
					},
					{
						"internalType": "int8",
						"name": "fb_status",
						"type": "int8"
					},
					{
						"internalType": "int8",
						"name": "tb_status",
						"type": "int8"
					},
					{
						"internalType": "string",
						"name": "fb_remarks",
						"type": "string"
					}
				],
				"internalType": "struct IMsgSmart.TXMsg",
				"name": "txMsg",
				"type": "tuple"
			}
		],
		"name": "sendSwapTx",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "account",
				"type": "address"
			},
			{
				"internalType": "bool",
				"name": "approval",
				"type": "bool"
			},
			{
				"internalType": "string",
				"name": "remarks",
				"type": "string"
			}
		],
		"name": "setOperator",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "serialNo",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "toBankID",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "tb_txID",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "tb_txDate",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "tb_payerWallet",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "tb_payeeWallet",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "amount",
				"type": "string"
			},
			{
				"internalType": "int8",
				"name": "tb_status",
				"type": "int8"
			},
			{
				"internalType": "string",
				"name": "tb_remarks",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "swapNo",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "swapDate",
				"type": "uint256"
			}
		],
		"name": "setSwapTx",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "serialNo",
				"type": "string"
			}
		],
		"name": "swapTxOf",
		"outputs": [
			{
				"components": [
					{
						"internalType": "string",
						"name": "serialNo",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromBankID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toBankID",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromCoin",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toCoin",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "amount",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "fromBank_PayerWallet",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "toBank_PayeeWallet",
						"type": "string"
					},
					{
						"internalType": "string",
						"name": "rate",
						"type": "string"
					},
					{
						"internalType": "uint256",
						"name": "date",
						"type": "uint256"
					},
					{
						"components": [
							{
								"internalType": "string",
								"name": "bankID",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "txID",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "date",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "coin",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "payerWallet",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "payeeWallet",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "amount",
								"type": "string"
							},
							{
								"internalType": "int8",
								"name": "status",
								"type": "int8"
							},
							{
								"internalType": "string",
								"name": "remarks",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "refundNo",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "refundDate",
								"type": "string"
							}
						],
						"internalType": "struct IMsgSmart.BankTxInfo",
						"name": "fromBankTx",
						"type": "tuple"
					},
					{
						"components": [
							{
								"internalType": "string",
								"name": "bankID",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "txID",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "date",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "coin",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "payerWallet",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "payeeWallet",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "amount",
								"type": "string"
							},
							{
								"internalType": "int8",
								"name": "status",
								"type": "int8"
							},
							{
								"internalType": "string",
								"name": "remarks",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "refundNo",
								"type": "string"
							},
							{
								"internalType": "string",
								"name": "refundDate",
								"type": "string"
							}
						],
						"internalType": "struct IMsgSmart.BankTxInfo",
						"name": "toBankTx",
						"type": "tuple"
					},
					{
						"internalType": "string",
						"name": "swapNo",
						"type": "string"
					},
					{
						"internalType": "uint256",
						"name": "swapDate",
						"type": "uint256"
					},
					{
						"internalType": "string",
						"name": "remarks",
						"type": "string"
					}
				],
				"internalType": "struct IMsgSmart.TXInfo",
				"name": "",
				"type": "tuple"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			}
		],
		"name": "transferOwnership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "serialNo",
				"type": "string"
			},
			{
				"internalType": "int8",
				"name": "fb_status",
				"type": "int8"
			},
			{
				"internalType": "int8",
				"name": "tb_status",
				"type": "int8"
			},
			{
				"internalType": "string",
				"name": "rate",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "tb_amount",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "refund",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "reDate",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "remark",
				"type": "string"
			}
		],
		"name": "updateSwapTx",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newImplementation",
				"type": "address"
			}
		],
		"name": "upgradeTo",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "newImplementation",
				"type": "address"
			},
			{
				"internalType": "bytes",
				"name": "data",
				"type": "bytes"
			}
		],
		"name": "upgradeToAndCall",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	}
]`

	//msgSmart.MsgSmartMetaData.ABI
	edabi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		fmt.Println("Invalid abi:", err)
	}

	var SetSwapTxEvent struct {
		Caller        common.Address
		SerialNo      string
		ToBankID      string
		TbTxID        string
		TbTxDate      string
		TbPayerWallet string
		TbPayeeWallet string
		Amount        string
		TbStatus      int8
		TbRemarks     string
		SwapNo        string
		SwapDate      *big.Int
	}
	err = edabi.UnpackIntoInterface(&SetSwapTxEvent, "SetSwapTx", data)
	//a, err := edabi.Unpack("SetSwapTxEvent", data)
	if err != nil {
		fmt.Println("Failed to unpack:", err)
	}
	//fmt.Println("a:", a)
	SetSwapTxEvent.Caller = common.BytesToAddress(receipt.Logs[0].Topics[1].Bytes())
	fmt.Println("Caller:", SetSwapTxEvent.Caller.Hex())
	fmt.Println("SerialNo:", SetSwapTxEvent.SerialNo)
	fmt.Println("ToBankID:", SetSwapTxEvent.ToBankID)
	fmt.Println("Tb_txID:", SetSwapTxEvent.TbTxID)
	fmt.Println("Tb_txDate:", SetSwapTxEvent.TbTxDate)
	fmt.Println("Tb_payerWallet:", SetSwapTxEvent.TbPayerWallet)
	fmt.Println("Tb_payeeWallet:", SetSwapTxEvent.TbPayeeWallet)
	fmt.Println("Amount:", SetSwapTxEvent.Amount)
	fmt.Println("Tb_status:", SetSwapTxEvent.TbStatus)
	fmt.Println("Tb_remarks:", SetSwapTxEvent.TbRemarks)
	fmt.Println("SwapNo:", SetSwapTxEvent.SwapNo)
	fmt.Println("SwapDate:", SetSwapTxEvent.SwapDate)
}

//func addFNEvent(blockHash common.Hash) {
//	cli, err := client.NewETHCliByURL("ws://", Key)
//	if err != nil {
//	}
//
//	contractABi, err := abi.JSON(strings.NewReader(msgSmart.MsgSmartABI))
//	if err != nil {
//	}
//
//	query := ethereum.FilterQuery{
//		BlockHash: &blockHash,
//		Addresses: []common.Address{},
//		Topics:    [][]common.Hash{{contractABi.Events["SetSwapTx"].ID}},
//	}
//
//	logs, err := cli.Client().FilterLogs(context.Background(), query)
//	if err != nil {
//	}
//
//	session,err:=msg.
//}

func NewMsgSmartSession(cli *client.ETHClient, addr common.Address) (*msgSmart.MsgSmartSession, error) {
	session, err := msgSmart.NewMsgSmart(addr, cli.Client())
	if err != nil {
		return nil, err
	}
	transactOpts := cli.GetTrandOpts()
	callOpts := bind.CallOpts{From: transactOpts.From}
	return &msgSmart.MsgSmartSession{session, callOpts, *transactOpts}, err
}
