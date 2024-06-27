package ethers

import "math/big"

func (e *EthersHelper) AddNetwork(name string, chainId *big.Int) {
	network := Network{
		Name:    name,
		ChainId: chainId,
	}
	e.Networks = append(e.Networks, network)
}
