package ethers

import (
	ERC20 "fraktal/go-mev-bot/contracts/ERC20"

	"github.com/ethereum/go-ethereum/common"
)

type Erc20 struct {
	Address  common.Address
	Contract ERC20.ERC20
}

type Erc721 common.Address

type Etc1155 common.Address
