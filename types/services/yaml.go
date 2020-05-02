package services

import (
	"github.com/pkg/errors"
	"github.com/statping/statping/utils"
	"gopkg.in/yaml.v2"
)

type yamlFile struct {
	Services []*Service `yaml:"services,flow"`
}

// LoadServicesYaml will attempt to load the 'services.yml' file for Service Auto Creation on startup.
func LoadServicesYaml() (*yamlFile, error) {
	f, err := utils.OpenFile(utils.Directory + "/services.yml")
	if err != nil {
		return nil, err
	}

	var svrs *yamlFile
	if err := yaml.Unmarshal([]byte(f), &svrs); err != nil {
		log.Errorln("Unable to parse the services.yml file", err)
		return nil, err
	}

	log.Infof("Found %d services inside services.yml file", len(svrs.Services))

	for _, svr := range svrs.Services {
		serviceByHash := findServiceByHash(svr.Hash())
		if serviceByHash == nil {
			if err := svr.Create(); err != nil {
				return nil, errors.Wrapf(err, "could not create service %s", svr.Name)
			}
			log.Infof("Automatically creating service '%s' checking %s", svr.Name, svr.Domain)

			go ServiceCheckQueue(svr, true)
		} else {
			log.Infof("Service %s #%d, already inserted", svr.Name, serviceByHash.Id)
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
