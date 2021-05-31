package services

// CheckinProcess runs the checkin routine for each checkin attached to service
func CheckinProcess(s *Service) {
	for _, c := range s.Checkins {
		if last := c.LastHit(); last != nil {
			c.Start()
		}
	}
}
