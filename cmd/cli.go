package main

import (
	"encoding/json"
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/joho/godotenv"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"upper.io/db.v3/sqlite"
)

const (
	BRAKER = "=============================================================================="
	POINT  = "                     "
)

func CatchCLI(args []string) {
	switch args[1] {
	case "version":
		fmt.Printf("Statup v%v\n", VERSION)
	case "assets":
		core.RenderBoxes()
		core.CreateAllAssets(".")
	case "sass":
		core.CompileSASS(".")
	case "update":
		gitCurrent, err := CheckGithubUpdates()
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Printf("Statup Version: v%v\nLatest Version: %v\n", VERSION, gitCurrent.TagName)
		if VERSION != gitCurrent.TagName[1:] {
			fmt.Printf("You don't have the latest version v%v!\nDownload the latest release at: https://github.com/hunterlong/statup\n", gitCurrent.TagName[1:])
		} else {
			fmt.Printf("You have the latest version of Statup!\n")
		}
	case "test":
		cmd := args[2]
		switch cmd {
		case "plugins":
			LoadPlugins(true)
		}
	case "export":
		var err error
		fmt.Printf("Statup v%v Exporting Static 'index.html' page...\n", VERSION)
		core.RenderBoxes()
		core.Configs, err = core.LoadConfig()
		if err != nil {
			utils.Log(4, "config.yml file not found")
		}
		RunOnce()
		indexSource := core.ExportIndexHTML()
		err = core.SaveFile("./index.html", []byte(indexSource))
		if err != nil {
			utils.Log(4, err)
		}
		utils.Log(1, "Exported Statup index page: 'index.html'")
	case "help":
		HelpEcho()
	case "run":
		utils.Log(1, "Running 1 time and saving to database...")
		RunOnce()
		fmt.Println("Check is complete.")
	case "env":
		fmt.Println("Statup Environment Variables")
		envs, err := godotenv.Read(".env")
		if err != nil {
			utils.Log(4, "No .env file found in current directory.")
		}
		for k, e := range envs {
			fmt.Printf("%v=%v\n", k, e)
		}
	default:
		utils.Log(3, "Statup does not have the command you entered.")
	}
}

func CheckUpdates() {

}

func RunOnce() {
	var err error
	core.Configs, err = core.LoadConfig()
	if err != nil {
		utils.Log(4, "config.yml file not found")
	}
	err = core.DbConnection(core.Configs.Connection, false, ".")
	if err != nil {
		utils.Log(4, err)
	}
	core.CoreApp, err = core.SelectCore()
	if err != nil {
		fmt.Println("Core database was not found, Statup is not setup yet.")
	}
	core.CoreApp.Services, err = core.SelectAllServices()
	if err != nil {
		utils.Log(4, err)
	}
	for _, s := range core.CoreApp.Services {
		out := core.ServiceCheck(s.ToService())
		fmt.Printf("    Service %v | URL: %v | Latency: %0.0fms | Online: %v\n", out.Name, out.Domain, (out.Latency * 1000), out.Online)
	}
}

func HelpEcho() {
	fmt.Printf("Statup v%v - Statup.io\n", VERSION)
	fmt.Printf("A simple Application Status Monitor that is opensource and lightweight.\n")
	fmt.Printf("Commands:\n")
	fmt.Println("     statup                    - Main command to run Statup server")
	fmt.Println("     statup version            - Returns the current version of Statup")
	fmt.Println("     statup run                - Check all services 1 time and then quit")
	fmt.Println("     statup test plugins       - Test all plugins for required information")
	fmt.Println("     statup assets             - Dump all assets used locally to be edited.")
	fmt.Println("     statup env                - Show all environment variables being used for Statup")
	fmt.Println("     statup export             - Exports the index page as a static HTML for pushing")
	fmt.Println("     statup update             - Attempts to update to the latest version")
	fmt.Println("     statup help               - Shows the user basic information about Statup")
	fmt.Println("Give Statup a Star at https://github.com/hunterlong/statup")
}

func TestPlugin(plug types.PluginActions) {
	defer utils.DeleteFile("./.plugin_test.db")
	core.RenderBoxes()

	info := plug.GetInfo()
	fmt.Printf("\n" + BRAKER + "\n")
	fmt.Printf("    Plugin Name:          %v\n", info.Name)
	fmt.Printf("    Plugin Description:   %v\n", info.Description)
	fmt.Printf("    Plugin Routes:        %v\n", len(plug.Routes()))
	for k, r := range plug.Routes() {
		fmt.Printf("      - Route %v      - (%v) /%v \n", k+1, r.Method, r.URL)
	}

	// Function to create a new Core with example services, hits, failures, users, and default communications
	FakeSeed(plug)

	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnLoad(sqlbuilder.Database)'")
	core.OnLoad(core.DbSession)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnSuccess(Service)'")
	core.OnSuccess(core.SelectService(1).ToService())
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnFailure(Service, FailureData)'")
	fakeFailD := core.FailureData{
		Issue: "No issue, just testing this plugin. This would include HTTP failure information though",
	}
	core.OnFailure(core.SelectService(1).ToService(), fakeFailD)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnSettingsSaved(Core)'")
	fmt.Println(BRAKER)
	core.OnSettingsSaved(core.CoreApp.ToCore())
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnNewService(Service)'")
	core.OnNewService(core.SelectService(2).ToService())
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnNewUser(User)'")
	user, _ := core.SelectUser(1)
	core.OnNewUser(user)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnUpdateService(Service)'")
	srv := core.SelectService(2).ToService()
	srv.Type = "http"
	srv.Domain = "https://yahoo.com"
	core.OnUpdateService(srv)
	fmt.Println("\n" + BRAKER)
	fmt.Println(POINT + "Sending 'OnDeletedService(Service)'")
	core.OnDeletedService(core.SelectService(1).ToService())
	fmt.Println("\n" + BRAKER)
}

func FakeSeed(plug types.PluginActions) {
	var err error
	core.CoreApp = core.NewCore()

	core.CoreApp.AllPlugins = []types.PluginActions{plug}

	fmt.Printf("\n" + BRAKER)

	fmt.Println("\nCreating a SQLite database for testing, will be deleted automatically...")
	sqlFake := sqlite.ConnectionURL{
		Database: "./.plugin_test.db",
	}
	core.DbSession, err = sqlite.Open(sqlFake)
	if err != nil {
		utils.Log(3, err)
	}
	up, _ := core.SqlBox.String("sqlite_up.sql")
	requests := strings.Split(up, ";")
	for _, request := range requests {
		_, err := core.DbSession.Exec(request)
		if err != nil {
			utils.Log(2, err)
		}
	}

	fmt.Println("Finished creating Test SQLite database")
	fmt.Println("Inserting example services into test database...")

	core.CoreApp.Name = "Plugin Test"
	core.CoreApp.Description = "This is a fake Core for testing your plugin"
	core.CoreApp.Domain = "http://localhost:8080"
	core.CoreApp.ApiSecret = "0x0x0x0x0"
	core.CoreApp.ApiKey = "abcdefg12345"

	fakeSrv := &types.Service{
		Name:   "Test Plugin Service",
		Domain: "https://google.com",
		Method: "GET",
	}
	core.CreateService(fakeSrv)

	fakeSrv2 := &types.Service{
		Name:   "Awesome Plugin Service",
		Domain: "https://netflix.com",
		Method: "GET",
	}
	core.CreateService(fakeSrv2)

	fakeUser := &types.User{
		Id:        6334,
		Username:  "Bulbasaur",
		Password:  "$2a$14$NzT/fLdE3f9iB1Eux2C84O6ZoPhI4NfY0Ke32qllCFo8pMTkUPZzy",
		Email:     "info@testdomain.com",
		Admin:     true,
		CreatedAt: time.Now(),
	}
	core.CreateUser(fakeUser)

	fakeUser = &types.User{
		Id:        6335,
		Username:  "Billy",
		Password:  "$2a$14$NzT/fLdE3f9iB1Eux2C84O6ZoPhI4NfY0Ke32qllCFo8pMTkUPZzy",
		Email:     "info@awesome.com",
		CreatedAt: time.Now(),
	}
	core.CreateUser(fakeUser)

	for i := 0; i <= 50; i++ {
		dd := core.HitData{
			Latency: rand.Float64(),
		}
		core.CreateServiceHit(fakeSrv, dd)

		dd = core.HitData{
			Latency: rand.Float64(),
		}
		core.CreateServiceHit(fakeSrv2, dd)

		fail := core.FailureData{
			Issue: "This is not an issue, but it would container HTTP response errors.",
		}
		core.CreateServiceFailure(fakeSrv, fail)

		fail = core.FailureData{
			Issue: "HTTP Status Code 521 did not match 200",
		}
		core.CreateServiceFailure(fakeSrv, fail)
	}

	fmt.Println("Seeding example data is complete, running Plugin Tests")

}

func CheckGithubUpdates() (GithubResponse, error) {
	var gitResp GithubResponse
	response, err := http.Get("https://api.github.com/repos/hunterlong/statup/releases/latest")
	if err != nil {
		return GithubResponse{}, err
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return GithubResponse{}, err
		}
		err = json.Unmarshal(contents, &gitResp)
		return gitResp, err
	}
	return gitResp, err
}

type GithubResponse struct {
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
	Author          GitAuthor   `json:"author"`
	Prerelease      bool        `json:"prerelease"`
	CreatedAt       time.Time   `json:"created_at"`
	PublishedAt     time.Time   `json:"published_at"`
	Assets          []GitAssets `json:"assets"`
	TarballURL      string      `json:"tarball_url"`
	ZipballURL      string      `json:"zipball_url"`
	Body            string      `json:"body"`
}

type GitAuthor struct {
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

type GitAssets struct {
	URL                string      `json:"url"`
	ID                 int         `json:"id"`
	NodeID             string      `json:"node_id"`
	Name               string      `json:"name"`
	Label              string      `json:"label"`
	Uploader           GitUploader `json:"uploader"`
	ContentType        string      `json:"content_type"`
	State              string      `json:"state"`
	Size               int         `json:"size"`
	DownloadCount      int         `json:"download_count"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	BrowserDownloadURL string      `json:"browser_download_url"`
}

type GitUploader struct {
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
