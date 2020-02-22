// Statping
// Copyright (C) 2020.  Hunter Long and the project contributors
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

package notifiers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestMatrixNotifier(t *testing.T) {
	urlRegExp := regexp.MustCompile(`^\/_matrix\/client\/r0\/rooms\/(.[^/]+)\/send\/m\.room\.message\/(\d+)$`)

	serverIds := map[string]bool{}

	var token = "super-secret-token"
	var room = "!testroom1234:matrix.org"

	var matrixHandler = http.HandlerFunc(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errResp := map[string]string{}
		defer r.Body.Close()

		auth := r.Header.Get("Authorization")
		if auth == "" {
			errResp["ErrCode"] = "M_MISSING_TOKEN"
			errResp["error"] = "Missing access token"
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		userToken := strings.Split(auth, " ")
		if len(userToken) != 2 {
			t.Errorf("invalid token, expected Bearer <token>, got: %s", auth)
			errResp["ErrCode"] = "M_MISSING_TOKEN"
			errResp["error"] = "Missing access token"
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		if userToken[0] != "Bearer" || userToken[1] != token {
			errResp["ErrCode"] = "M_UNKNOWN_TOKEN"
			errResp["error"] = "Unknown token"
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		contentType := r.Header.Get("content-type")
		if contentType != "application/json" {
			t.Errorf("expect content type: application/json, got: %s", contentType)
		}

		data := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			t.Errorf("parse json: %v", err)
			errResp["ErrCode"] = "M_NOT_JSON"
			errResp["error"] = "Content not json"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		if data["msgtype"] != "m.text" && data["msgtype"] != "m.notice" {
			t.Errorf("invalid matrix event type: %v", data["msgtype"])
			errResp["ErrCode"] = "M_UNKNOWN"
			errResp["error"] = "No body"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		params := urlRegExp.FindStringSubmatch(r.URL.Path)
		if len(params) != 3 {
			errResp["ErrCode"] = "M_UNRECOGNIZED"
			errResp["error"] = "Unrecognized request"
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		if params[1] != room {
			errResp["ErrCode"] = "M_FORBIDDEN"
			errResp["error"] = "Unknown room"
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		if serverIds[params[2]] == true {
			// server actually returns 200 and silently drops message
			t.Errorf("duplicate transaction id")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errResp)
			return
		}

		serverIds[params[2]] = true
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Matrix test request valid", func(t *testing.T) {
		server := httptest.NewServer(matrixHandler)
		defer server.Close()
		Matrix.Host = server.URL
		Matrix.Var1 = room
		Matrix.ApiSecret = token
		Matrix.Var2 = "m.text"

		err := Matrix.Send("a test message")
		if err != nil {
			t.Errorf("send message: %v", err)
		}

		// send second message to verify transaction id is unique
		time.Sleep(time.Millisecond)
		err = Matrix.Send("a test message")
		if err != nil {
			t.Errorf("send second message: %v", err)
		}
	})

	t.Run("Matrix test unauthorized", func(t *testing.T) {
		server := httptest.NewServer(matrixHandler)
		defer server.Close()

		Matrix.Host = server.URL
		Matrix.Var1 = "!testroom1234:matrix.org"
		Matrix.ApiSecret = "invalid-token"
		Matrix.Var2 = "m.text"

		err := Matrix.Send("a test message")
		if err == nil {
			t.Error("send message, expect error on invalid token, got nil")
		} else if err.Error() != "invalid token" {
			t.Errorf("send message, expted 'invalid token', got: %v", err)
		}
	})

	t.Run("Matrix test unknown room", func(t *testing.T) {
		server := httptest.NewServer(matrixHandler)
		defer server.Close()

		Matrix.Host = server.URL
		Matrix.Var1 = "!no-room:matrix.org"
		Matrix.ApiSecret = token
		Matrix.Var2 = "m.text"

		err := Matrix.Send("a test message")
		if err == nil {
			t.Error("send message, expect error on invalid token, got nil")
		} else if err.Error() != "Unknown room" {
			t.Errorf("send message, expted 'Unknown room', got: %v", err)
		}
	})

	t.Run("Matrix test invalid room", func(t *testing.T) {
		server := httptest.NewServer(matrixHandler)
		defer server.Close()

		Matrix.Host = server.URL
		Matrix.Var1 = room + "/invalid"
		Matrix.ApiSecret = token
		Matrix.Var2 = "m.text"

		err := Matrix.Send("a test message")
		if err == nil {
			t.Error("send message, expect error on invalid token, got nil")
		} else if err.Error() != "invalid homeserver url or room name" {
			t.Errorf("send message, expted 'invalid homeserver url or room name', got: %v", err)
		}
	})
}
