package services

import (
	"bufio"
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"
	"os"
)

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

func ServicesFromEnvFile() error {
	servicesEnv := utils.Getenv("SERVICES_FILE", "").(string)
	if servicesEnv == "" {
		return nil
	}

	file, err := os.Open(servicesEnv)
	if err != nil {
		return errors.Wrapf(err, "error opening 'SERVICES_FILE' at: %s", servicesEnv)
	}
	defer file.Close()

	var serviceLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		serviceLines = append(serviceLines, scanner.Text())
	}

	if len(serviceLines) == 0 {
		return nil
	}

	for k, service := range serviceLines {

		svr, err := ValidateService(service)
		if err != nil {
			return errors.Wrapf(err, "invalid service at index %d in SERVICES_FILE environment variable", k)
		}
		if findServiceByHash(svr.Hash()) == nil {
			if err := svr.Create(); err != nil {
				return errors.Wrapf(err, "could not create service %s", svr.Name)
			}
			log.Infof("Created new service '%s'", svr.Name)
		}

	}

	return nil
}
