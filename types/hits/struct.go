package hits

import "time"

// Hit struct is a 'successful' ping or web response entry for a service.
type Hit struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Service   int64     `gorm:"index;column:service" json:"-"`
	Latency   int64     `gorm:"column:latency" json:"latency"`
	PingTime  int64     `gorm:"column:ping_time" json:"ping_time"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// BeforeCreate for Hit will set CreatedAt to UTC
func (h *Hit) BeforeCreate() (err error) {
	if h.CreatedAt.IsZero() {
		h.CreatedAt = time.Now().UTC()
	}
	return
}
