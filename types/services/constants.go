package services

const FAILURE_THRESHOLD = 3

const (
	CRITICAL = "critical"
	PARTIAL  = "partial"
	DELAYED  = "delayed"
	NO       = "no"
)

const (
	STATUS_UP       = "up"
	STATUS_DOWN     = "down"
	STATUS_DEGRADED = "degraded"
)

func ApplyStatus(current string, apply string, defaultStatus string) string {
	switch current {
	case STATUS_DOWN:
		return STATUS_DOWN
	case STATUS_DEGRADED:
		if apply == STATUS_DOWN {
			return apply
		}
		return STATUS_DEGRADED
	case STATUS_UP:
		return apply
	default:
		return defaultStatus
	}
}

func HandleEmptyStatus(status string) string {
	if status == "" {
		return STATUS_DOWN
	} else {
		return status
	}
}

const INCIDENTS = "Incidents"
