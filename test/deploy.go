package test

import (
	"deploy_solidity_contract/core/chain"
	"deploy_solidity_contract/core/deploy"
	"deploy_solidity_contract/core/mytoken"
	"fmt"
)

type MyTokenLifecycle struct {
	ContractConfig Config
	opts           []chain.ChainOption

	mytokenCli *mytoken.MyTokenClient
}

func NewMyTokenLifecycle(opts ...chain.ChainOption) *MyTokenLifecycle {

	return &MyTokenLifecycle{
		opts: opts,
	}
}

func (s *MyTokenLifecycle) Deploy() error {
	if err := s.deployMyTokenContract(); err != nil {
		return err
	}

	return nil
}

func (s *MyTokenLifecycle) deployMyTokenContract() error {
	logicCli, err := mytoken.NewMyTokenClient(s.opts...)
	if err != nil {
		return err
	}

	deployCli, err := deploy.NewContractDeploy(logicCli, s.opts...)
	if err != nil {
		return err
	}

	l, p, err := deployCli.DeployAll()
	if err != nil {
		return err
	}

	s.ContractConfig.MyTokenLogicAddress = l
	s.ContractConfig.MyTokenProxyAddress = p

	return nil
}

func (s *MyTokenLifecycle) ChainOpts() []chain.ChainOption {
	var opts []chain.ChainOption
	opts = append(opts, s.opts...)
	opts = append(opts, chain.WithMyTokenProxyAddress(s.ContractConfig.MyTokenProxyAddress))
	return opts
}

func (s *MyTokenLifecycle) connSession() (err error) {
	s.mytokenCli, err = mytoken.NewMyTokenClient(s.ChainOpts()...)
	if err != nil {
		return err
	}
	if _, err := s.mytokenCli.NewSession(); err != nil {
		return err
	}

	return nil
}

func (s *MyTokenLifecycle) ContractSetting() error {

	err := s.connSession()
	if err != nil {
		return err
	}

	return nil
}

func (s *MyTokenLifecycle) ShowAddress() {

	fmt.Println("Stablecoin Logic Address", s.ContractConfig.MyTokenLogicAddress)
	fmt.Println("Stablecoin Proxy Address", s.ContractConfig.MyTokenProxyAddress)
}
