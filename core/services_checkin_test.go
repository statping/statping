// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	testCheckin *Checkin
)

func TestCreateCheckin(t *testing.T) {
	service := SelectService(1)
	checkin := &types.Checkin{
		ServiceId:   service.Id,
		Interval:    10,
		GracePeriod: 5,
	}
	id, err := database.Create(checkin)
	assert.Nil(t, err)
	assert.NotZero(t, id)
	assert.NotEmpty(t, testCheckin.ApiKey)
	assert.Equal(t, int64(10), testCheckin.Interval)
	assert.Equal(t, int64(5), testCheckin.GracePeriod)
	assert.True(t, testCheckin.Expected().Minutes() < 0)
}

func TestSelectCheckin(t *testing.T) {
	service := SelectService(1)
	checkins := service.Checkins()
	assert.NotNil(t, checkins)
	assert.Equal(t, 1, len(checkins))
	c := checkins[0]
	assert.Equal(t, int64(10), c.Interval)
	assert.Equal(t, int64(5), c.GracePeriod)
	assert.Equal(t, 7, len(c.ApiKey))
}

func TestUpdateCheckin(t *testing.T) {
	testCheckin.Interval = 60
	testCheckin.GracePeriod = 15
	err := database.Update(testCheckin)
	assert.Nil(t, err)
	assert.NotZero(t, testCheckin.Id)
	assert.NotEmpty(t, testCheckin.ApiKey)
	service := SelectService(1)
	checkin := service.Checkins()[0]
	assert.Equal(t, int64(60), checkin.Interval)
	assert.Equal(t, int64(15), checkin.GracePeriod)
	t.Log(testCheckin.Expected())
	assert.True(t, testCheckin.Expected().Minutes() < 0)
}

func TestCreateCheckinHits(t *testing.T) {
	service := SelectService(1)
	checkins := service.Checkins()
	assert.Equal(t, 1, len(checkins))
	created := time.Now().UTC().Add(-60 * time.Second)
	hit := &types.CheckinHit{
		Checkin:   testCheckin.Id,
		From:      "192.168.1.1",
		CreatedAt: created,
	}
	_, err := database.Create(hit)
	require.Nil(t, err)

	checks := service.Checkins()
	assert.Equal(t, 1, len(checks))
}

func TestSelectCheckinMethods(t *testing.T) {
	time.Sleep(5 * time.Second)
	service := SelectService(1)
	checkins := service.Checkins()
	assert.NotNil(t, checkins)
	assert.Equal(t, float64(60), testCheckin.Period().Seconds())
	assert.Equal(t, float64(15), testCheckin.Grace().Seconds())
	t.Log(testCheckin.Expected())

	lastHit := checkins[0]
	assert.True(t, testCheckin.Expected().Seconds() < -5)
	assert.False(t, lastHit.CreatedAt.IsZero())
}
