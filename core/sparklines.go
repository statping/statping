package core

import (
	"github.com/hunterlong/statping/utils"
	"strings"
	"time"
)

// SparklineDayFailures returns a string array of daily service failures
func (s *Service) SparklineDayFailures(days int) string {
	var arr []string
	ago := time.Now().UTC().Add((time.Duration(days) * -24) * time.Hour)
	for day := 1; day <= days; day++ {
		ago = ago.Add(24 * time.Hour)
		failures, _ := s.TotalFailuresOnDate(ago)
		arr = append(arr, utils.ToString(failures))
	}
	return "[" + strings.Join(arr, ",") + "]"
}

// SparklineHourResponse returns a string array for the average response or ping time for a service
func (s *Service) SparklineHourResponse(hours int, method string) string {
	var arr []string
	end := time.Now().UTC()
	start := end.Add(time.Duration(-hours) * time.Hour)
	obj := GraphDataRaw(s, start, end, "hour", method)
	for _, v := range obj.Array {
		arr = append(arr, utils.ToString(v.Value))
	}
	return "[" + strings.Join(arr, ",") + "]"
}
