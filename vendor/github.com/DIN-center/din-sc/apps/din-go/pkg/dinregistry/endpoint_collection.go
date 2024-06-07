package dinregistry

import (
	"embed"
	"io/fs"
	"math/big"

	"github.com/chenzhijie/go-web3"
	"github.com/pkg/errors"
)

type EndpointCollectionHandler struct {
	ContractHandler IContractHandler
}

//go:embed abi/endpoint_collection.abi
var abiEndpointCollectionFS embed.FS

// NewEndpointCollection creates a new Endpoint Collection
func NewEndpointCollection(w3 *web3.Web3, contractAddress string) (*EndpointCollectionHandler, error) {
	abiBytes, err := fs.ReadFile(abiEndpointCollectionFS, ABIEndpointCollectionPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to EndpointCollectionHandler ReadFile")
	}
	abiStr := string(abiBytes)
	contractHandler, err := NewContractHandler(w3, contractAddress, abiStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to EndpointCollectionHandler")
	}
	return &EndpointCollectionHandler{ContractHandler: contractHandler}, nil
}

// GetMethodId returns the method id for a given method name associated with the endpoint collection
func (e *EndpointCollectionHandler) GetMethodId(name string) (uint8, error) {
	bit, err := e.ContractHandler.Call(GetMethodId, name)
	if err != nil {
		return 0, errors.Wrap(err, "failed call to EndpointCollectionHandler GetMethodId")
	}
	return bit.(uint8), nil
}

// GetMethodName returns the method name for a given method id associated with the endpoint collection
func (e *EndpointCollectionHandler) GetMethodName(bit uint8) (string, error) {
	name, err := e.ContractHandler.Call(GetMethodName, bit)
	if err != nil {
		return "", errors.Wrap(err, "failed call to EndpointCollectionHandler GetMethodName")
	}
	return name.(string), nil
}

// GetCapabilities returns the capabilities of the endpoint collection
func (e *EndpointCollectionHandler) GetCapabilities() (*big.Int, error) {
	caps, err := e.ContractHandler.Call(Capabilities)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to EndpointCollectionHandler GetCapabilities")
	}
	return caps.(*big.Int), nil
}

// AllMethods returns a list of methods supported by the endpoint collection
func (e *EndpointCollectionHandler) AllMethods() ([]Method, error) {
	allMethods, err := e.ContractHandler.Call(AllMethods)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to EndpointCollectionHandler AllMethods")
	}

	var methods []Method
	// Need to convert the raw data into an anonymous struct to access the fields.
	for _, method := range allMethods.([]struct {
		Name        string `json:"name"`
		Bit         uint8  `json:"bit"`
		Deactivated bool   `json:"deactivated"`
	}) {
		methods = append(methods, Method{
			Name:        method.Name,
			Bit:         method.Bit,
			Deactivated: method.Deactivated,
		})
	}
	return methods, nil
}

// IsMethodSupported returns whether a method is supported by the endpoint collection
func (e *EndpointCollectionHandler) IsMethodSupported(bit uint8) (bool, error) {
	supported, err := e.ContractHandler.Call(IsMethodSupported, bit)
	if err != nil {
		return false, errors.Wrap(err, "failed call to EndpointCollectionHandler IsMethodSupported")
	}
	return supported.(bool), nil
}
