# contracts sdk go

## Deploy Contract

部署合约可以使用通用的合约部署操作，但必须首先实现对应逻辑合约的Client,并且实现"deploy.IDeployLogic"接口：

```go
type IDeployLogic interface {
	DeployLogic() (address common.Address, txId string, err error)
	UpgradeTo(logicAddr common.Address) (err error)
}

```

部署合约的调用参考以下以部署钱包合约为例：

```go
var opts []chain.ChainOption
opts = append(opts, chain.WithChainURL(URL))
opts = append(opts, chain.WithChainHexKey(Key))

logicCli, err := NewWalletClient(common.Address{}, opts...)
if err != nil {
t.Fatal(err)
}

deployCli, err := deploy.NewContractDeploy(logicCli, opts...)
if err != nil {
t.Fatal(err)
}

l, p, err := deployCli.DeployAll()
if err != nil {
t.Fatal(err)
}

fmt.Println("logic contract address :", l.String())
fmt.Println("proxy contract address :", p.String())
```

## Call Contract

应当通过代理合约地址调用合约，调用合约的示例如下：

```go
	var opts []chain.ChainOption
	opts = append(opts, chain.WithChainURL(URL))
	opts = append(opts, chain.WithChainHexKey(Key))

	logicCli, err := NewWalletClient(common.HexToAddress("0x528BD95D1de5bE1Aa110E3126bf3790d8F635A97"), opts...)
	if err != nil {
		t.Fatal(err)
	}

	session, err := logicCli.NewSession()
	if err != nil {
		t.Fatal(err)
	}

	owner, err := session.Owner()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("contract owner :", owner.String())
```

