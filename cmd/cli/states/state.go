package states

import "github.com/ethereum/go-ethereum/common"

type BlockStates struct {
	balanceCounter map[common.Address]int
	nonceCounter   map[common.Address]int
	codesCounter   map[common.Address]int
	storageCounter map[common.Address]map[common.Hash]int
}

func NewBlockStates() *BlockStates {
	return &BlockStates{
		balanceCounter: make(map[common.Address]int),
		nonceCounter:   make(map[common.Address]int),
		codesCounter:   make(map[common.Address]int),
		storageCounter: make(map[common.Address]map[common.Hash]int),
	}
}

func (bs *BlockStates) AddBalance(address common.Address) {
	if _, ok := bs.balanceCounter[address]; !ok {
		bs.balanceCounter[address] = 0
	}
	bs.balanceCounter[address]++
}

func (bs *BlockStates) AddNonce(address common.Address) {
	if _, ok := bs.nonceCounter[address]; !ok {
		bs.nonceCounter[address] = 0
	}
	bs.nonceCounter[address]++
}

func (bs *BlockStates) AddCode(address common.Address) {
	if _, ok := bs.codesCounter[address]; !ok {
		bs.codesCounter[address] = 0
	}
	bs.codesCounter[address]++
}

func (bs *BlockStates) AddStorage(address common.Address, key common.Hash) {
	if _, ok := bs.storageCounter[address]; !ok {
		bs.storageCounter[address] = make(map[common.Hash]int)
	}
	if _, ok := bs.storageCounter[address][key]; !ok {
		bs.storageCounter[address][key] = 0
	}
	bs.storageCounter[address][key]++
}

func (bs *BlockStates) AddAccountState(addr common.Address, acc *Account) {
	if len(acc.Code) != 0 {
		bs.AddCode(addr)
	}
	if acc.Balance != nil {
		bs.AddBalance(addr)
	}
	if acc.Nonce != nil {
		bs.AddNonce(addr)
	}
	for key, _ := range acc.Storage {
		bs.AddStorage(addr, key)
	}
}
