package config

import (
	"github.com/juju/errors"
)

var current *ChainConfig

type ChainConfig struct {
	Name      string `json:"name"`
	CoreAsset string `json:"core_asset"`
	Prefix    string `json:"prefix"`
	ID        string `json:"id"`
}

const (
	ChainIDUnknown = "0000000000000000000000000000000000000000000000000000000000000000"
	ChainIDMainnet = "6057d856c398875cac2650fe33caef3d5f6b403d184c5154abbff526ec1143c4"
	ChainIDTestnet = "1ae3653a3105800f5722c5bda2b55530d0e9e8654314e2f3dc6d2b010da641c5"
	ChainIDLocal   = "179db3c6a2e08d610f718f05e9cc2aabad62aff80305b9621b162b8b6f2fd79f"
)

var (
	knownNetworks = []ChainConfig{
		ChainConfig{
			Name:      "Unknown",
			CoreAsset: "n/a",
			Prefix:    "n/a",
			ID:        ChainIDUnknown,
		},
		ChainConfig{
			Name:      "BCXMainnet",
			CoreAsset: "COCOS",
			Prefix:    "COCOS",
			ID:        ChainIDMainnet,
		},
		ChainConfig{
			Name:      "BCXTestnet",
			CoreAsset: "COCOS",
			Prefix:    "COCOS",
			ID:        ChainIDTestnet,
		},
		ChainConfig{
			Name:      "BCXLocal",
			CoreAsset: "COCOS",
			Prefix:    "COCOS",
			ID:        ChainIDLocal,
		},
	}
)

func Current() *ChainConfig {
	return current
}

func Add(cnf ChainConfig) error {
	if FindByID(cnf.ID) != nil {
		return errors.Errorf("ChainConfig for ID %q already available", cnf.ID)
	}

	knownNetworks = append(knownNetworks, cnf)
	return nil
}

func FindByID(chainID string) *ChainConfig {
	for _, cnf := range knownNetworks {
		if cnf.ID == chainID {
			return &cnf
		}
	}

	return nil
}

func SetCurrent(chainID string) error {
	current = FindByID(chainID)
	if current != nil {
		return nil
	}

	return errors.Errorf("ChainConfig for ID %q not found", chainID)
}
