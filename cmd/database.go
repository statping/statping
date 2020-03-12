package main

import (
	"github.com/statping/statping/notifiers"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
)

var (
	// DbSession stores the Statping database session
	DbModels []interface{}
)

func init() {
	DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifiers.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}
}
