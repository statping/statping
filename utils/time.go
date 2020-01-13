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

package utils

import (
	"fmt"
	"time"
)

const (
	FlatpickrTime     = "2006-01-02 15:04"
	FlatpickrDay      = "2006-01-02"
	FlatpickrReadable = "Mon, 02 Jan 2006"
)

// Timezoner returns the time.Time with the user set timezone
func Timezoner(t time.Time, zone float32) time.Time {
	zoneInt := float32(3600) * zone
	loc := time.FixedZone("", int(zoneInt))
	timez := t.In(loc)
	return timez
}

// Now returns the UTC timestamp
func Now() time.Time {
	return time.Now().UTC()
}

// FormatDuration converts a time.Duration into a string
func FormatDuration(d time.Duration) string {
	var out string
	if d.Hours() >= 24 {
		out = fmt.Sprintf("%0.0f day", d.Hours()/24)
		if (d.Hours() / 24) >= 2 {
			out += "s"
		}
		return out
	} else if d.Hours() >= 1 {
		out = fmt.Sprintf("%0.0f hour", d.Hours())
		if d.Hours() >= 2 {
			out += "s"
		}
		return out
	} else if d.Minutes() >= 1 {
		out = fmt.Sprintf("%0.0f minute", d.Minutes())
		if d.Minutes() >= 2 {
			out += "s"
		}
		return out
	} else if d.Seconds() >= 1 {
		out = fmt.Sprintf("%0.0f second", d.Seconds())
		if d.Seconds() >= 2 {
			out += "s"
		}
		return out
	} else if rev(d.Hours()) >= 24 {
		out = fmt.Sprintf("%0.0f day", rev(d.Hours()/24))
		if rev(d.Hours()/24) >= 2 {
			out += "s"
		}
		return out
	} else if rev(d.Hours()) >= 1 {
		out = fmt.Sprintf("%0.0f hour", rev(d.Hours()))
		if rev(d.Hours()) >= 2 {
			out += "s"
		}
		return out
	} else if rev(d.Minutes()) >= 1 {
		out = fmt.Sprintf("%0.0f minute", rev(d.Minutes()))
		if rev(d.Minutes()) >= 2 {
			out += "s"
		}
		return out
	} else {
		out = fmt.Sprintf("%0.0f second", rev(d.Seconds()))
		if rev(d.Seconds()) >= 2 {
			out += "s"
		}
	}
	return out
}

func rev(f float64) float64 {
	return f * -1
}
