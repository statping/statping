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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/handlers"
	"github.com/hunterlong/statping/plugin"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http/httptest"
	"time"
)

// catchCLI will run functions based on the commands sent to Statping
func catchCLI(args []string) error {
	dir := utils.Directory
	runLogs := utils.InitLogs
	runAssets := source.Assets
	loadDotEnvs()

	switch args[0] {
	case "version":
		if COMMIT != "" {
			fmt.Printf("%v (%v)\n", VERSION, COMMIT)
		} else {
			fmt.Printf("%v\n", VERSION)
		}
		return errors.New("end")
	case "assets":
		var err error
		if err = runLogs(); err != nil {
			return err
		}
		if err = runAssets(); err != nil {
			return err
		}
		if err = source.CreateAllAssets(dir); err != nil {
			return err
		}
		return errors.New("end")
	case "sass":
		if err := runLogs(); err != nil {
			return err
		}
		if err := runAssets(); err != nil {
			return err
		}
		if err := source.CompileSASS(dir); err != nil {
			return err
		}
		return errors.New("end")
	case "update":
		updateDisplay()
		return errors.New("end")
	case "test":
		cmd := args[1]
		switch cmd {
		case "plugins":
			plugin.LoadPlugins()
		}
		return errors.New("end")
	case "static":
		var err error
		if err = runLogs(); err != nil {
			return err
		}
		if err = runAssets(); err != nil {
			return err
		}
		fmt.Printf("Statping v%v Exporting Static 'index.html' page...\n", VERSION)
		if core.CoreApp.Config, err = core.LoadConfigFile(dir); err != nil {
			log.Errorln("config.yml file not found")
			return err
		}
		indexSource := ExportIndexHTML()
		//core.CloseDB()
		if err = utils.SaveFile(dir+"/index.html", indexSource); err != nil {
			log.Errorln(err)
			return err
		}
		log.Infoln("Exported Statping index page: 'index.html'")
	case "help":
		HelpEcho()
		return errors.New("end")
	case "export":
		var err error
		var data []byte
		if err = runLogs(); err != nil {
			return err
		}
		if err = runAssets(); err != nil {
			return err
		}
		if core.CoreApp.Config, err = core.LoadConfigFile(dir); err != nil {
			return err
		}
		if err = core.CoreApp.Connect(false, dir); err != nil {
			return err
		}
		if data, err = core.ExportSettings(); err != nil {
			return fmt.Errorf("could not export settings: %v", err.Error())
		}
		//core.CloseDB()
		if err = utils.SaveFile(dir+"/statping-export.json", data); err != nil {
			return fmt.Errorf("could not write file statping-export.json: %v", err.Error())
		}
		return errors.New("end")
	case "import":
		var err error
		var data []byte
		if len(args) != 2 {
			return fmt.Errorf("did not include a JSON file to import\nstatping import filename.json")
		}
		filename := args[1]
		if data, err = ioutil.ReadFile(filename); err != nil {
			return err
		}
		var exportData core.ExportData
		if err = json.Unmarshal(data, &exportData); err != nil {
			return err
		}
		return errors.New("end")
	case "run":
		if err := runLogs(); err != nil {
			return err
		}
		if err := runAssets(); err != nil {
			return err
		}
		log.Infoln("Running 1 time and saving to database...")
		runOnce()
		//core.CloseDB()
		fmt.Println("Check is complete.")
		return errors.New("end")
	case "env":
		fmt.Println("Statping Environment Variable")
		if err := runLogs(); err != nil {
			return err
		}
		if err := runAssets(); err != nil {
			return err
		}
		envs, err := godotenv.Read(".env")
		if err != nil {
			log.Errorln("No .env file found in current directory.")
			return err
		}
		for k, e := range envs {
			fmt.Printf("%v=%v\n", k, e)
		}
	default:
		return nil
	}
	return errors.New("end")
}

// ExportIndexHTML returns the HTML of the index page as a string
func ExportIndexHTML() []byte {
	source.Assets()
	core.CoreApp.Connect(false, utils.Directory)
	core.CoreApp.SelectAllServices(false)
	core.CoreApp.UseCdn = types.NewNullBool(true)
	for _, srv := range core.CoreApp.Services {
		service := srv.(*core.Service)
		service.Check(true)
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handlers.ExecuteResponse(w, r, "index.gohtml", nil, nil)
	return w.Body.Bytes()
}

func updateDisplay() error {
	var err error
	var gitCurrent githubResponse
	if gitCurrent, err = checkGithubUpdates(); err != nil {
		fmt.Printf("Issue connecting to https://github.com/hunterlong/statping\n%v\n", err)
		return err
	}
	if VERSION != gitCurrent.TagName[1:] {
		fmt.Printf("\nNew Update %v Available!\n", gitCurrent.TagName[1:])
		fmt.Printf("Update Command:\n")
		fmt.Printf("curl -o- -L https://statping.com/install.sh | bash\n\n")
	}
	return nil
}

// runOnce will initialize the Statping application and check each service 1 time, will not run HTTP server
func runOnce() {
	var err error
	core.CoreApp.Config, err = core.LoadConfigFile(utils.Directory)
	if err != nil {
		log.Errorln("config.yml file not found")
	}
	err = core.CoreApp.Connect(false, utils.Directory)
	if err != nil {
		log.Errorln(err)
	}
	core.CoreApp, err = core.SelectCore()
	if err != nil {
		fmt.Println("Core database was not found, Statping is not setup yet.")
	}
	_, err = core.CoreApp.SelectAllServices(true)
	if err != nil {
		log.Errorln(err)
	}
	for _, out := range core.CoreApp.Services {
		out.Check(true)
	}
}

// HelpEcho prints out available commands and flags for Statping
func HelpEcho() {
	fmt.Printf("Statping v%v - Statping.com\n", VERSION)
	fmt.Printf("A simple Application Status Monitor that is opensource and lightweight.\n")
	fmt.Printf("Commands:\n")
	fmt.Println("     statping                    - Main command to run Statping server")
	fmt.Println("     statping version            - Returns the current version of Statping")
	fmt.Println("     statping run                - Check all services 1 time and then quit")
	fmt.Println("     statping assets             - Dump all assets used locally to be edited.")
	fmt.Println("     statping static             - Creates a static HTML file of the index page")
	fmt.Println("     statping sass               - Compile .scss files into the css directory")
	fmt.Println("     statping test plugins       - Test all plugins for required information")
	fmt.Println("     statping env                - Show all environment variables being used for Statping")
	fmt.Println("     statping update             - Attempts to update to the latest version")
	fmt.Println("     statping export             - Exports your Statping settings to a 'statping-export.json' file.")
	fmt.Println("     statping import <file>      - Imports settings from a previously saved JSON file.")
	fmt.Println("     statping help               - Shows the user basic information about Statping")
	fmt.Printf("Flags:\n")
	fmt.Println("     -ip 127.0.0.1             - Run HTTP server on specific IP address (default: localhost)")
	fmt.Println("     -port 8080                - Run HTTP server on Port (default: 8080)")
	fmt.Println("     -verbose 1                - Verbose mode levels 1 - 4 (default: 1)")
	fmt.Println("     -env path/debug.env       - Optional .env file to set as environment variables while running server")
	fmt.Printf("Environment Variables:\n")
	fmt.Println("     PORT                      - Set the outgoing port for the HTTP server (or use -port)")
	fmt.Println("     IP                        - Bind a specific IP address to the HTTP server (or use -ip)")
	fmt.Println("     VERBOSE                   - Display more logs in verbose mode. (1 - 4)")
	fmt.Println("     STATPING_DIR              - Set a absolute path for the root path of Statping server (logs, assets, SQL db)")
	fmt.Println("     DISABLE_LOGS              - Disable viewing and writing to the log file (default is false)")
	fmt.Println("     DB_CONN                   - Automatic Database connection (sqlite, postgres, mysql)")
	fmt.Println("     DB_HOST                   - Database hostname or IP address")
	fmt.Println("     DB_USER                   - Database username")
	fmt.Println("     DB_PASS                   - Database password")
	fmt.Println("     DB_PORT                   - Database port (5432, 3306, ...)")
	fmt.Println("     DB_DATABASE               - Database connection's database name")
	fmt.Println("     POSTGRES_SSLMODE          - Enable Postgres SSL Mode 'ssl_mode=VALUE' (enable/disable/verify-full/verify-ca)")
	fmt.Println("     GO_ENV                    - Run Statping in testmode, will bypass HTTP authentication (if set as 'test')")
	fmt.Println("     NAME                      - Set a name for the Statping status page")
	fmt.Println("     DESCRIPTION               - Set a description for the Statping status page")
	fmt.Println("     DOMAIN                    - Set a URL for the Statping status page")
	fmt.Println("     ADMIN_USER                - Username for administrator account (default: admin)")
	fmt.Println("     ADMIN_PASS                - Password for administrator account (default: admin)")
	fmt.Println("     SASS                      - Set the absolute path to the sass binary location")
	fmt.Println("     HTTP_PROXY                - Use a HTTP Proxy for HTTP Requests")
	fmt.Println("     AUTH_USERNAME             - HTTP Basic Authentication username")
	fmt.Println("     AUTH_PASSWORD             - HTTP Basic Authentication password")
	fmt.Println("     BASE_PATH                 - Set the base URL prefix (set to 'monitor' if URL is domain.com/monitor)")
	fmt.Println("   * You can insert environment variables into a '.env' file in root directory.")
	fmt.Println("Give Statping a Star at https://github.com/hunterlong/statping")
}

func checkGithubUpdates() (githubResponse, error) {
	url := "https://api.github.com/repos/hunterlong/statping/releases/latest"
	contents, _, err := utils.HttpRequest(url, "GET", nil, nil, nil, time.Duration(2*time.Second), true)
	if err != nil {
		return githubResponse{}, err
	}
	var gitResp githubResponse
	err = json.Unmarshal(contents, &gitResp)
	return gitResp, err
}

type githubResponse struct {
	URL             string      `json:"url"`
	AssetsURL       string      `json:"assets_url"`
	UploadURL       string      `json:"upload_url"`
	HTMLURL         string      `json:"html_url"`
	ID              int         `json:"id"`
	NodeID          string      `json:"node_id"`
	TagName         string      `json:"tag_name"`
	TargetCommitish string      `json:"target_commitish"`
	Name            string      `json:"name"`
	Draft           bool        `json:"draft"`
	Author          gitAuthor   `json:"author"`
	Prerelease      bool        `json:"prerelease"`
	CreatedAt       time.Time   `json:"created_at"`
	PublishedAt     time.Time   `json:"published_at"`
	Assets          []gitAssets `json:"assets"`
	TarballURL      string      `json:"tarball_url"`
	ZipballURL      string      `json:"zipball_url"`
	Body            string      `json:"body"`
}

type gitAuthor struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type gitAssets struct {
	URL                string      `json:"url"`
	ID                 int         `json:"id"`
	NodeID             string      `json:"node_id"`
	Name               string      `json:"name"`
	Label              string      `json:"label"`
	Uploader           gitUploader `json:"uploader"`
	ContentType        string      `json:"content_type"`
	State              string      `json:"state"`
	Size               int         `json:"size"`
	DownloadCount      int         `json:"download_count"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	BrowserDownloadURL string      `json:"browser_download_url"`
}

type gitUploader struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
