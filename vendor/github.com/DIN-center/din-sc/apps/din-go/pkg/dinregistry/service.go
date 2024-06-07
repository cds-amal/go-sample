package dinregistry

import (
	"embed"
	"io/fs"
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/pkg/errors"
)

type ServiceHandler struct {
	ContractHandler IContractHandler
}

//go:embed abi/service.abi
var abiServiceFS embed.FS

// NewService creates a new ServiceHandler service struct
func NewServiceHandler(w3 *web3.Web3, contractAddress string) (*ServiceHandler, error) {
	abiBytes, err := fs.ReadFile(abiServiceFS, ABIServicePath)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewService ReadFile")
	}

	contractHandler, err := NewContractHandler(w3, contractAddress, string(abiBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewService")
	}
	return &ServiceHandler{contractHandler}, nil
}

// ListMethods returns a list of method strings supported by the service
func (c *ServiceHandler) ListMethods() ([]string, error) {
	methods, err := c.ContractHandler.Call(ListMethods)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to ServiceHandler ListMethods")
	}
	return methods.([]string), nil
}

// GetCapabilities returns the capabilities of the service
func (c *ServiceHandler) GetCapabilities() (*big.Int, error) {
	caps, err := c.ContractHandler.Call(Capabilities)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to ServiceHandler GetCapabilities")
	}
	return caps.(*big.Int), nil
}

// IsMethodSupported returns whether a method is supported by the service
func (c *ServiceHandler) IsMethodSupported(bit uint8) (bool, error) {
	supported, err := c.ContractHandler.Call(IsMethodSupported, bit)
	if err != nil {
		return false, errors.Wrap(err, "failed call to ServiceHandler IsMethodSupported")
	}
	return supported.(bool), nil
}
