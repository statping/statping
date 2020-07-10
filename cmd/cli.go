package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/statping/statping/source"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/configs"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func assetsCli() error {
	dir := utils.Directory
	if err := utils.InitLogs(); err != nil {
		return err
	}
	if err := source.Assets(); err != nil {
		return err
	}
	if err := source.CreateAllAssets(dir); err != nil {
		return err
	}
	return nil
}

func exportCli(args []string) error {
	filename := fmt.Sprintf("%s/statping-%s.json", utils.Directory, time.Now().Format("01-02-2006-1504"))
	if len(args) == 1 {
		filename = fmt.Sprintf("%s/%s", utils.Directory, args)
	}
	var data []byte
	if err := utils.InitLogs(); err != nil {
		return err
	}
	if err := source.Assets(); err != nil {
		return err
	}
	config, err := configs.LoadConfigs(configFile)
	if err != nil {
		return err
	}
	if err = configs.ConnectConfigs(config, false); err != nil {
		return err
	}
	if _, err := services.SelectAllServices(false); err != nil {
		return err
	}
	if data, err = ExportSettings(); err != nil {
		return fmt.Errorf("could not export settings: %v", err.Error())
	}
	if err = utils.SaveFile(filename, data); err != nil {
		return fmt.Errorf("could not write file statping-export.json: %v", err.Error())
	}
	log.Infoln("Statping export file saved to ", filename)
	return nil
}

func sassCli() error {
	if err := utils.InitLogs(); err != nil {
		return err
	}
	if err := source.Assets(); err != nil {
		return err
	}
	if err := source.CompileSASS(source.DefaultScss...); err != nil {
		return err
	}
	return nil
}

func resetCli() error {
	d := utils.Directory
	fmt.Println("Statping directory: ", d)
	assets := d + "/assets"
	if utils.FolderExists(assets) {
		fmt.Printf("Deleting %s folder.\n", assets)
		if err := utils.DeleteDirectory(assets); err != nil {
			return err
		}
	} else {
		fmt.Printf("Assets folder does not exist %s\n", assets)
	}

	logDir := d + "/logs"
	if utils.FolderExists(logDir) {
		fmt.Printf("Deleting %s directory.\n", logDir)
		if err := utils.DeleteDirectory(logDir); err != nil {
			return err
		}
	} else {
		fmt.Printf("Logs folder does not exist %s\n", logDir)
	}

	c := d + "/config.yml"
	if utils.FileExists(c) {
		fmt.Printf("Deleting %s file.\n", c)
		if err := utils.DeleteFile(c); err != nil {
			return err
		}
	} else {
		fmt.Printf("Config file does not exist %s\n", c)
	}

	dbFile := d + "/statping.db"
	if utils.FileExists(dbFile) {
		fmt.Printf("Backuping up %s file.\n", dbFile)
		if err := utils.RenameDirectory(dbFile, d+"/statping.db.backup"); err != nil {
			return err
		}
	} else {
		fmt.Printf("Statping SQL Database file does not exist %s\n", dbFile)
	}

	fmt.Println("Statping has been reset")
	return nil
}

func envCli() error {
	fmt.Println("Statping Configuration")
	for k, v := range utils.Params.AllSettings() {
		fmt.Printf("%s=%v\n", strings.ToUpper(k), v)
	}
	return nil
}

func onceCli() error {
	if err := utils.InitLogs(); err != nil {
		return err
	}
	if err := source.Assets(); err != nil {
		return err
	}
	log.Infoln("Running 1 time and saving to database...")
	if err := runOnce(); err != nil {
		return err
	}
	//core.CloseDB()
	fmt.Println("Check is complete.")
	return nil
}

func importCli(args []string) error {
	var err error
	var data []byte
	if len(args) < 1 {
		return errors.New("invalid command arguments")
	}
	if data, err = ioutil.ReadFile(args[0]); err != nil {
		return err
	}
	var exportData ExportData
	if err = json.Unmarshal(data, &exportData); err != nil {
		return err
	}
	log.Printf("=== %s ===\n", exportData.Core.Name)
	if exportData.Config != nil {
		log.Printf("Configs:     %s\n", exportData.Config.DbConn)
		if exportData.Config.DbUser != "" {
			log.Printf("   - Host:   %s\n", exportData.Config.DbHost)
			log.Printf("   - User:   %s\n", exportData.Config.DbUser)
		}
	}
	if len(exportData.Services) > 0 {
		log.Printf("Services:   %d\n", len(exportData.Services))
	}
	if len(exportData.Checkins) > 0 {
		log.Printf("Checkins:   %d\n", len(exportData.Checkins))
	}
	if len(exportData.Groups) > 0 {
		log.Printf("Groups:     %d\n", len(exportData.Groups))
	}
	if len(exportData.Messages) > 0 {
		log.Printf("Messages:   %d\n", len(exportData.Messages))
	}
	if len(exportData.Users) > 0 {
		log.Printf("Users:      %d\n", len(exportData.Users))
	}

	if exportData.Config != nil {
		if ask("Create config.yml file from Configs?") {
			log.Printf("Database Host:   	%s\n", exportData.Config.DbHost)
			log.Printf("Database Port:   	%d\n", exportData.Config.DbPort)
			log.Printf("Database User:   	%s\n", exportData.Config.DbUser)
			log.Printf("Database Password:   %s\n", exportData.Config.DbPass)
			if err := exportData.Config.Save(utils.Directory); err != nil {
				return err
			}
		}
	}

	config, err := configs.LoadConfigs(configFile)
	if err != nil {
		return err
	}
	if err = configs.ConnectConfigs(config, false); err != nil {
		return err
	}
	if ask("Create database rows and sample data?") {
		if err := config.ResetCore(); err != nil {
			return err
		}
	}

	if ask("Import Core settings?") {
		c := exportData.Core
		if err := c.Update(); err != nil {
			return err
		}
	}
	for _, s := range exportData.Groups {
		if ask(fmt.Sprintf("Import Group '%s'?", s.Name)) {
			s.Id = 0
			if err := s.Create(); err != nil {
				return err
			}
		}
	}
	for _, s := range exportData.Services {
		if ask(fmt.Sprintf("Import Service '%s'?", s.Name)) {
			s.Id = 0
			if err := s.Create(); err != nil {
				return err
			}
		}
	}
	for _, s := range exportData.Checkins {
		if ask(fmt.Sprintf("Import Checkin '%s'?", s.Name)) {
			s.Id = 0
			if err := s.Create(); err != nil {
				return err
			}
		}
	}
	for _, s := range exportData.Messages {
		if ask(fmt.Sprintf("Import Message '%s'?", s.Title)) {
			s.Id = 0
			if err := s.Create(); err != nil {
				return err
			}
		}
	}
	for _, s := range exportData.Users {
		if ask(fmt.Sprintf("Import User '%s'?", s.Username)) {
			s.Id = 0
			if err := s.Create(); err != nil {
				log.Errorln(err)
			}
		}
	}
	log.Infof("Import complete")
	return nil
}

func ask(format string) bool {
	fmt.Printf(fmt.Sprintf(format + " [y/N]: "))
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return strings.ToLower(text) == "y"
}

func updateDisplay() error {
	gitCurrent, err := checkGithubUpdates()
	if err != nil {
		return errors.Wrap(err, "Issue connecting to https://github.com/statping/statping")
	}
	if gitCurrent.TagName == "" {
		return nil
	}
	if len(gitCurrent.TagName) < 2 {
		return nil
	}
	if VERSION != gitCurrent.TagName[1:] {
		fmt.Printf("New Update %v Available!\n", gitCurrent.TagName[1:])
		fmt.Printf("Update Command:\n")
		fmt.Printf("curl -o- -L https://statping.com/install.sh | bash\n\n")
	}
	return nil
}

// runOnce will initialize the Statping application and check each service 1 time, will not run HTTP server
func runOnce() error {
	config, err := configs.LoadConfigs(configFile)
	if err != nil {
		return errors.Wrap(err, "config.yml file not found")
	}
	err = configs.ConnectConfigs(config, false)
	if err != nil {
		return errors.Wrap(err, "issue connecting to database")
	}
	c, err := core.Select()
	if err != nil {
		return errors.Wrap(err, "core database was not found or setup")
	}

	core.App = c

	_, err = services.SelectAllServices(true)
	if err != nil {
		return errors.Wrap(err, "could not select all services")
	}
	for _, srv := range services.Services() {
		srv.CheckService(true)
	}
	return nil
}

func checkGithubUpdates() (githubResponse, error) {
	url := "https://api.github.com/repos/statping/statping/releases/latest"
	contents, _, err := utils.HttpRequest(url, "GET", nil, nil, nil, time.Duration(2*time.Second), true, nil)
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

// ExportChartsJs renders the charts for the index page

type ExportData struct {
	Config    *configs.DbConfig   `json:"config"`
	Core      *core.Core          `json:"core"`
	Services  []services.Service  `json:"services"`
	Messages  []*messages.Message `json:"messages"`
	Checkins  []*checkins.Checkin `json:"checkins"`
	Users     []*users.User       `json:"users"`
	Groups    []*groups.Group     `json:"groups"`
	Notifiers []core.AllNotifiers `json:"notifiers"`
}

// ExportSettings will export a JSON file containing all of the settings below:
// - Core
// - Notifiers
// - Checkins
// - Users
// - Services
// - Groups
// - Messages
func ExportSettings() ([]byte, error) {
	c, err := core.Select()
	if err != nil {
		return nil, err
	}
	var srvs []services.Service
	for _, s := range services.AllInOrder() {
		s.Failures = nil
		srvs = append(srvs, s)
	}

	cfg, err := configs.LoadConfigs(configFile)
	if err != nil {
		return nil, err
	}

	data := ExportData{
		Config:    cfg,
		Core:      c,
		Notifiers: core.App.Notifications,
		Checkins:  checkins.All(),
		Users:     users.All(),
		Services:  srvs,
		Groups:    groups.All(),
		Messages:  messages.All(),
	}
	export, err := json.Marshal(data)
	return export, err
}

// ExportIndexHTML returns the HTML of the index page as a string
//func ExportIndexHTML() []byte {
//	source.Assets()
//	core.CoreApp.Connect(core.CoreApp., utils.Directory)
//	core.SelectAllServices(false)
//	core.CoreApp.UseCdn = types.NewNullBool(true)
//	for _, srv := range core.Services() {
//		core.CheckService(srv, true)
//	}
//	w := httptest.NewRecorder()
//	r := httptest.NewRequest("GET", "/", nil)
//	handlers.ExecuteResponse(w, r, "index.gohtml", nil, nil)
//	return w.Body.Bytes()
//}
