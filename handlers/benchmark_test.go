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

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleIndex(b *testing.B) {
	b.ReportAllocs()
	r := request(b, "/")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		indexHandler(rw, r)
	}
}

func BenchmarkServicesHandlerIndex(b *testing.B) {
	r := request(b, "/")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		servicesHandler(rw, r)
	}
}

func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
