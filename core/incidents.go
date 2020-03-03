package core

import (
	"github.com/hunterlong/statping/types"
)

type Incident struct {
	*types.Incident
}

type IncidentUpdate struct {
	*types.IncidentUpdate
}

// ReturnIncident returns *core.Incident based off a *types.Incident
func ReturnIncident(u *types.Incident) *Incident {
	return &Incident{u}
}
