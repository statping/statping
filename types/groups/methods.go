package groups

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/services"
)

func (g *Group) Services() []*services.Service {
	var services []*services.Service
	database.DB().Where("group = ?", g.Id).Find(&services)
	return services
}

// GroupOrder will reorder the groups based on 'order_id' (Order)
type GroupOrder []*Group

// Sort interface for resorting the Groups in order
func (c GroupOrder) Len() int           { return len(c) }
func (c GroupOrder) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c GroupOrder) Less(i, j int) bool { return c[i].Order < c[j].Order }
