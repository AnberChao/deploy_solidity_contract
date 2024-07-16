package chain

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"time"
)

const defTimeout = 20 * time.Second

type ChainOption func(config *ChainConfig) error

func NewDefConfig() *ChainConfig {
	return &ChainConfig{
		timeout: defTimeout,
	}

}

type ChainConfig struct {
	url string
	key *ecdsa.PrivateKey

	timeout time.Duration

	myTokenProxyAddress common.Address
}

func WithChainURL(url string) ChainOption {
	return func(config *ChainConfig) error {
		config.url = url
		return nil
	}
}

func WithChainHexKey(key string) ChainOption {
	return func(config *ChainConfig) error {
		hex, err := hexutil.Decode(key)
		if err != nil {
			return errors.WithMessage(err, "decode key has error")
		}
		pk, err := crypto.ToECDSA(hex)
		if err != nil {
			return errors.WithMessage(err, "crypto to ecdsa has error")
		}
		config.key = pk
		return nil
	}
}

func WithMyTokenProxyAddress(addr common.Address) ChainOption {
	return func(config *ChainConfig) error {
		config.myTokenProxyAddress = addr
		return nil
	}
}

func WithChainPrivateKey(key *ecdsa.PrivateKey) ChainOption {
	return func(config *ChainConfig) error {
		config.key = key
		return nil
	}
}

func WithTimeout(timeout time.Duration) ChainOption {
	return func(config *ChainConfig) error {
		config.timeout = timeout
		return nil
	}
}
