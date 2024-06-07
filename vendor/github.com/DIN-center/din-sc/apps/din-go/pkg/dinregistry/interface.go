package dinregistry

import (
	"math/big"

	web3 "github.com/chenzhijie/go-web3"
)

type IDinRegistryHandler interface {
	GetAllMethodsByEndpointCollection(endpointCollection string) ([]Method, error)
	ListAllMethodsByEndpointCollection(endpointCollection string) ([]string, error)
	GetAllEndpointCollections() ([]EndpointCollection, error)
	GetEndpointCollectionCapabilities(endpointCollection string) (*big.Int, error)
	GetAllProviders() ([]ProviderHandler, error)
	GetProvidersByEndpointCollection(endpointCollection string) ([]ProviderHandler, error)
}

type IServiceHandler interface {
	// function listMethods() public view returns (string[] memory methods)
	ListMethods() ([]string, error)

	// uint256 public capabilities;
	GetCapabilities() (*big.Int, error)

	// function isMethodSupported(uint8 bit) public view returns (bool supported)
	IsMethodSupported(bit uint8) (bool, error)
}

type IEndpointCollectionHandler interface {
	// function getMethodId(string memory name) public view returns (uint8 bit)
	GetMethodId(name string) (uint8, error)

	// function getMethodName(uint8 bit) public view returns (string memory name)
	GetMethodName(bit uint8) (string, error)

	// function capabilities() public view returns (uint256)
	GetCapabilities() (*big.Int, error)

	// function allMethods() public view returns (Method[] memory)
	AllMethods() ([]Method, error)

	// function isMethodSupported(uint8 bit) public view returns (bool supported)
	IsMethodSupported(bit uint8) (bool, error)
}

type IProviderHandler interface {
	AllServices() ([]ServiceHandler, error)
	Name() (string, error)
	Owner() (string, error)
}

type IContractHandler interface {
	Call(method string, args ...interface{}) (interface{}, error)
	GetWeb3Object() *web3.Web3
}
