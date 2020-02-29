package core

import "github.com/hunterlong/statping/core/integrations"

// AddIntegrations will attach all the integrations into the system
func AddIntegrations() error {
	return integrations.AddIntegrations(
		integrations.CsvIntegrator,
		integrations.TraefikIntegrator,
		integrations.DockerIntegrator,
	)
}
