package database

import "github.com/hunterlong/statping/types"

type CheckinObj struct {
	*types.Checkin
	failures
}

func (o *Object) AsCheckin() HitsFailures {
	return &CheckinObj{
		Checkin: o.model.(*types.Checkin),
	}
}

func Checkin(id int64) (HitsFailures, error) {
	var checkin types.Checkin
	query := database.Model(&types.Checkin{}).Where("id = ?", id).Find(&checkin)
	return &CheckinObj{Checkin: &checkin}, query.Error()
}

func (c *CheckinObj) Hits() *hits {
	return &hits{
		database.Model(&types.Checkin{}).Where("checkin = ?", c.Id),
	}
}

func (c *CheckinObj) Failures() *failures {
	return &failures{
		database.Model(&types.Failure{}).
			Where("method = 'checkin' AND service = ?", c.Id).Order("id desc"),
	}
}
