package config

import (
	"github.com/anaxexp/anaxexp-cli/pkg/api"
	"github.com/anaxexp/anaxexp-cli/pkg/types"
	"github.com/pkg/errors"
	"fmt"
	"strings"
)

type Config struct {
	UUID          string               `json,mapstructure:"instance"`
	DataContainer string               `json,mapstructure:"data,omitempty"`
	Context       string               `json,mapstructure:"context"`
	API           *api.Config          `json,mapstructure:"api"`
	BuildConfig   *types.BuildConfig   `json,mapstructure:"config"`
	Metadata      *types.BuildMetadata `json,mapstructure:"metadata"`
}

func (config *Config) FindService(serviceName string) (types.Service, error) {
	for _, service := range config.BuildConfig.Services {
		if service.Name == serviceName {
			return service, nil
		}
	}

	return types.Service{}, errors.New("Service not found")
}

func (config *Config) FindServicesByPrefix(prefix string) ([]types.Service, error) {
	var services []types.Service

	for _, service := range config.BuildConfig.Services {
		if strings.HasPrefix(service.Name, prefix) {
			services = append(services, service)
		}
	}

	if len(services) == 0 {
		return services, errors.New(fmt.Sprintf("No matching services found with prefix %s", prefix))
	}

	return services, nil
}
