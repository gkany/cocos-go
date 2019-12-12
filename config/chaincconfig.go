package config

import "github.com/juju/errors"

var current *ChainConfig

type ChainConfig struct {
	Name      string `json:"name"`
	CoreAsset string `json:"core_asset"`
	Prefix    string `json:"prefix"`
	ID        string `json:"id"`
}

const (
	ChainIDUnknown = "0000000000000000000000000000000000000000000000000000000000000000"
	ChainIDBTS     = "4018d7844c78f6a6c41c6a552b898022310fc5dec06da467ee7905a8dad512c8"
	ChainIDMuse    = "45ad2d3f9ef92a49b55c2227eb06123f613bb35dd08bd876f2aea21925a67a67"
	ChainIDTest    = "39f5e2ede1f8bc1a3a54a7914414e3779e33193f1f5693510e73cb7a87617447"
	ChainIDObelisk = "1cfde7c388b9e8ac06462d68aadbd966b58f88797637d9af805b4560b0e9661e"
	ChainIDGPH     = "b8d1603965b3eb1acba27e62ff59f74efa3154d43a4188d381088ac7cdf35539"
	ChainIDKarma   = "c85b4a30545e09c01aaa7943be89e9785481c1e7bd5ee7d176cb2b3d8dd71a70"
	ChainIDBCX     = "6057d856c398875cac2650fe33caef3d5f6b403d184c5154abbff526ec1143c4"
	ChainIDBCXTest = "1ae3653a3105800f5722c5bda2b55530d0e9e8654314e2f3dc6d2b010da641c5"
	ChainIDBCXDev  = "179db3c6a2e08d610f718f05e9cc2aabad62aff80305b9621b162b8b6f2fd79f"
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
			Name:      "Graphene",
			CoreAsset: "CORE",
			Prefix:    "GPH",
			ID:        ChainIDGPH,
		},
		ChainConfig{
			Name:      "BitShares",
			CoreAsset: "BTS",
			Prefix:    "BTS",
			ID:        ChainIDBTS,
		},
		ChainConfig{
			Name:      "Muse",
			CoreAsset: "MUSE",
			Prefix:    "MUSE",
			ID:        ChainIDMuse,
		},
		ChainConfig{
			Name:      "Test",
			CoreAsset: "TEST",
			Prefix:    "TEST",
			ID:        ChainIDTest,
		},
		ChainConfig{
			Name:      "Obelisk",
			CoreAsset: "GOV",
			Prefix:    "FEW",
			ID:        ChainIDObelisk,
		},
		ChainConfig{
			Name:      "Karma",
			CoreAsset: "KRM",
			Prefix:    "KRM",
			ID:        ChainIDKarma,
		},
		ChainConfig{
			Name:      "cocosBCX",
			CoreAsset: "COCOS",
			Prefix:    "COCOS",
			ID:        ChainIDBCX,
		},
		ChainConfig{
			Name:      "cocosBCXTest",
			CoreAsset: "COCOS",
			Prefix:    "COCOS",
			ID:        ChainIDBCXTest,
		},
		ChainConfig{
			Name:      "cocosBCXDev",
			CoreAsset: "COCOS",
			Prefix:    "COCOS",
			ID:        ChainIDBCXDev,
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
