package contract_factory

import (
	servertypes "server/v0/types/server"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// ContractFactory creates contract instances.
type ContractFactory struct {
	client bind.ContractBackend
}

// NewContractFactory creates a new ContractFactory.
func NewContractFactory(client bind.ContractBackend) *ContractFactory {
	return &ContractFactory{client: client}
}

// CreateContract creates a contract instance based on the provided type.
func CreateContract[T any](address common.Address, client bind.ContractBackend, constructor func(common.Address, bind.ContractBackend) (T, error)) (*servertypes.Contract[T], error) {
	instance, err := constructor(address, client)
	if err != nil {
		return nil, err
	}

	return &servertypes.Contract[T]{
		Instance: instance,
	}, nil
}
