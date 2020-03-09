package main

import (
	"github.com/hunterlong/statping/types/checkins"
	"github.com/hunterlong/statping/types/failures"
	"github.com/hunterlong/statping/types/groups"
	"github.com/hunterlong/statping/types/hits"
	"github.com/hunterlong/statping/types/incidents"
	"github.com/hunterlong/statping/types/messages"
	"github.com/hunterlong/statping/types/notifications"
	"github.com/hunterlong/statping/types/services"
	"github.com/hunterlong/statping/types/users"
)

var (
	// DbSession stores the Statping database session
	DbModels []interface{}
)

func init() {
	DbModels = []interface{}{&services.Service{}, &users.User{}, &hits.Hit{}, &failures.Failure{}, &messages.Message{}, &groups.Group{}, &checkins.Checkin{}, &checkins.CheckinHit{}, &notifications.Notification{}, &incidents.Incident{}, &incidents.IncidentUpdate{}}
}
