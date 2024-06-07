package dinregistry

import (
	"embed"
	"io/fs"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	web3 "github.com/chenzhijie/go-web3"
	"github.com/pkg/errors"
)

// Compile time check to ensure that the DinWrapper implements the Din interface
var _ IDinRegistryHandler = &DinRegistryHandler{}
var _ IProviderHandler = &ProviderHandler{}
var _ IServiceHandler = &ServiceHandler{}
var _ IEndpointCollectionHandler = &EndpointCollectionHandler{}

type DinRegistryHandler struct {
	ContractHandler IContractHandler
}

//go:embed abi/din_registry.abi
var abiDinRegistryFS embed.FS

func NewDinRegistryHandler(w3 *web3.Web3, contractAddress string) (*DinRegistryHandler, error) {
	abiBytes, err := fs.ReadFile(abiDinRegistryFS, ABIDinRegistryPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewDinRegistryHandler ReadFile")
	}

	contractHandler, err := NewContractHandler(w3, contractAddress, string(abiBytes))
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewDinRegistryHandler")
	}
	return &DinRegistryHandler{ContractHandler: contractHandler}, nil
}

func (d *DinRegistryHandler) GetAllMethodsByEndpointCollection(endpointCollection string) ([]Method, error) {
	allMethods, err := d.ContractHandler.Call(AllMethods, endpointCollection)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to DinRegistryHandler GetAllMethodsByEndpointCollection")
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

func (d *DinRegistryHandler) ListAllMethodsByEndpointCollection(endpointCollection string) ([]string, error) {
	allMethods, err := d.ContractHandler.Call(ListMethods, endpointCollection)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to DinRegistryHandler ListAllMethodsByEndpointCollection")
	}

	return allMethods.([]string), nil
}

func (d *DinRegistryHandler) GetAllEndpointCollections() ([]EndpointCollection, error) {
	rawData, err := d.ContractHandler.Call(GetEndpointCollections)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to DinRegistryHandler GetEndpointCollections")
	}

	entries, ok := rawData.([]struct {
		Id   *big.Int `json:"id"`
		Meta struct {
			Name        string         `json:"name"`
			Description string         `json:"description"`
			Collection  common.Address `json:"collection"`
		} `json:"meta"`
	})
	if !ok {
		return nil, errors.New("unexpected data structure returned in GetEndpointCollections")
	}

	var result []EndpointCollection
	for _, entry := range entries {
		convertedEntry := EndpointCollection{
			Name:        entry.Meta.Name,
			Description: entry.Meta.Description,

			//TODO: instantiate the contract for collection.
			Collection: entry.Meta.Collection,
		}
		result = append(result, convertedEntry)
	}

	return result, nil
}

func (d *DinRegistryHandler) GetEndpointCollectionCapabilities(endpointCollection string) (*big.Int, error) {
	capabilities, err := d.ContractHandler.Call(GetCollectionCapabilities, endpointCollection)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to DinRegistryHandler GetEndpointCollectionCapabilities")
	}

	return capabilities.(*big.Int), nil
}

func (d *DinRegistryHandler) GetAllProviders() ([]ProviderHandler, error) {
	rawData, err := d.ContractHandler.Call(GetAllProviders)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to DinRegistryHandler GetAllProviders Call")
	}

	var providers []ProviderHandler
	for _, address := range rawData.([]common.Address) {
		provider, err := NewProviderHandler(d.ContractHandler.GetWeb3Object(), address.String())
		if err != nil {
			return nil, errors.Wrap(err, "failed call to DinRegistryHandler GetAllProviders NewProvider")
		}
		providers = append(providers, *provider)
	}
	return providers, nil
}

func (d *DinRegistryHandler) GetProvidersByEndpointCollection(endpointCollection string) ([]ProviderHandler, error) {
	rawData, err := d.ContractHandler.Call(GetProviders, endpointCollection)
	if err != nil {
		return nil, err
	}

	var providers []ProviderHandler
	for _, address := range rawData.([]common.Address) {
		provider, err := NewProviderHandler(d.ContractHandler.GetWeb3Object(), address.String())
		if err != nil {
			return nil, err
		}
		providers = append(providers, *provider)
	}
	return providers, nil
}

func GetDINRegistry() string {
  return "DIN Registry"
}
