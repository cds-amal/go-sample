package dinregistry

import (
	"embed"
	"io/fs"

	"github.com/chenzhijie/go-web3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type ProviderHandler struct {
	ContractHandler IContractHandler
}

//go:embed abi/provider.abi
var abiProvideFS embed.FS

func NewProviderHandler(w3 *web3.Web3, contractAddress string) (*ProviderHandler, error) {
	abiBytes, err := fs.ReadFile(abiProvideFS, ABIProviderPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewProviderHandler ReadFile")
	}
	abiStr := string(abiBytes)
	contractHandler, err := NewContractHandler(w3, contractAddress, abiStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to NewProviderHandler NewContractWrapper")
	}
	return &ProviderHandler{contractHandler}, nil
}

func (p *ProviderHandler) Name() (string, error) {
	name, err := p.ContractHandler.Call(Name)
	if err != nil {
		return "", errors.Wrap(err, "failed call to Provider Name")
	}
	return name.(string), nil
}

func (p *ProviderHandler) Owner() (string, error) {
	owner, err := p.ContractHandler.Call(Owner)
	if err != nil {
		return "", errors.Wrap(err, "failed call to ProviderHandler Owner")
	}
	return owner.(common.Address).String(), nil
}

func (p *ProviderHandler) AllServices() ([]ServiceHandler, error) {
	rawData, err := p.ContractHandler.Call(AllServices)
	if err != nil {
		return nil, errors.Wrap(err, "failed call to ProviderHandler AllServices Call")
	}

	var services []ServiceHandler
	for _, address := range rawData.([]common.Address) {
		service, err := NewServiceHandler(p.ContractHandler.GetWeb3Object(), address.String())
		if err != nil {
			return nil, errors.Wrap(err, "failed call to ProviderHandler AllServices")
		}
		services = append(services, *service)
	}
	return services, nil
}
