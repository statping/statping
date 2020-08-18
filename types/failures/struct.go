package failures

import "time"

// Failure is a failed attempt to check a service. Any a service does not meet the expected requirements,
// a new Failure will be inserted into Db.
type Failure struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Issue     string    `gorm:"column:issue" json:"issue"`
	Method    string    `gorm:"column:method" json:"method,omitempty"`
	MethodId  int64     `gorm:"column:method_id" json:"method_id,omitempty"`
	ErrorCode int       `gorm:"column:error_code" json:"error_code"`
	Service   int64     `gorm:"index;column:service" json:"-"`
	Checkin   int64     `gorm:"index;column:checkin" json:"-"`
	PingTime  int64     `gorm:"column:ping_time"  json:"ping"`
	Reason    string    `gorm:"column:reason" json:"reason,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type FailSort []Failure

func (s FailSort) Len() int {
	return len(s)
}
func (s FailSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FailSort) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}
