package dinregistry

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	// Contract calls
	// Shared
	IsMethodSupported = "isMethodSupported"
	Capabilities      = "capabilities"
	AllMethods        = "allMethods"
	ListMethods       = "listMethods"
	// Provider
	Name        = "name"
	Owner       = "owner"
	AllServices = "allServices"
	// Endpoint
	GetMethodId   = "getMethodId"
	GetMethodName = "getMethodName"
	// DinRegistry
	GetEndpointCollections    = "getEndpointCollections"
	GetCollectionCapabilities = "getCollectionCapabilities"
	GetAllProviders           = "getAllProviders"
	GetProviders              = "getProviders"

	// ABI Paths
	ABIEndpointCollectionPath = "abi/endpoint_collection.abi"
	ABIProviderPath           = "abi/provider.abi"
	ABIServicePath            = "abi/service.abi"
	ABIDinRegistryPath        = "abi/din_registry.abi"
)

// Method corresponds to the `struct Method` in the ABI.
type Method struct {
	Name        string `json:"name"`
	Bit         uint8  `json:"bit"`
	Deactivated bool   `json:"deactivated"`
}

type EndpointCollection struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Collection  common.Address `json:"collection"`
}
