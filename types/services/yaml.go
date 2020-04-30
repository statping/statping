package services

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
	"gopkg.in/yaml.v2"
)

type ServicesYaml struct {
	Services []Service `yaml:"services,flow"`
}

// LoadServicesYaml will attempt to load the 'services.yml' file for Service Auto Creation on startup.
func LoadServicesYaml() (*ServicesYaml, error) {
	f, err := utils.OpenFile("services.yml")
	if err != nil {
		return nil, err
	}

	var svrs *ServicesYaml
	if err := yaml.Unmarshal([]byte(f), &svrs); err != nil {
		return nil, err
	}

	for _, svr := range svrs.Services {
		if findServiceByHash(svr.Hash()) == nil {
			if err := svr.Create(); err != nil {
				return nil, errors.Wrapf(err, "could not create service %s", svr.Name)
			}
			log.Infof("Automatically created service '%s' checking %s", svr.Name, svr.Domain)
		}

	}

	return svrs, nil
}

// findServiceByHas will return a service that matches the SHA256 hash of a service
// Service hash example: sha256(name:EXAMPLEdomain:HTTP://DOMAIN.COMport:8080type:HTTPmethod:GET)
func findServiceByHash(hash string) *Service {
	for _, service := range All() {
		if service.Hash() == hash {
			return service
		}
	}
	return nil
}
