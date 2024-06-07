package dinregistry

import (
	web3 "github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
	"github.com/pkg/errors"
)

// ContractHandler holds common fields for contract interaction.
type ContractHandler struct {
	Contract *eth.Contract
	Web3     *web3.Web3
}

// NewContractHandler creates a new ContractHandler struct
// It is used as a common fields for contract interaction.
func NewContractHandler(w3 *web3.Web3, contractAddress string, abiStr string) (*ContractHandler, error) {
	contract, err := w3.Eth.NewContract(abiStr, contractAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewContractHandler NewContract")
	}
	return &ContractHandler{Contract: contract, Web3: w3}, nil
}

func (c *ContractHandler) GetWeb3Object() *web3.Web3 {
	return c.Web3
}

func (c *ContractHandler) Call(method string, args ...interface{}) (interface{}, error) {
	return c.Contract.Call(method, args...)
}
