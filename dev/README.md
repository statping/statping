This readme is automatically generated from the Golang documentation. [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/hunterlong/statping)

# statping
--
    import "github.com/hunterlong/statping"

Package statping is a server monitoring application that includs a status page
server. Visit the Statping repo at https://github.com/hunterlong/statping to get
a full understanding of what this application can do.


### Install Statping

Statping is available for Mac, Linux and Windows 64x. You can download the
tar.gz file or use a couple other methods. Download the latest release at
https://github.com/hunterlong/statping/releases/latest or view below. If you're
on windows, download the zip file from the latest releases link.

    // MacOS using homebrew
    brew tap hunterlong/statping
    brew install statping

    // Linux installation
    bash <(curl -s https://assets.statping.com/install.sh)
    statping version


### Docker

Statping can be built in many way, the best way is to use Docker!

    docker run -it -p 8080:8080 hunterlong/statping

Enjoy Statping and tell me any issues you might be having on Github.
https://github.com/hunterlong

## Usage
# cmd
--
Package main for building the Statping CLI binary application. This package
connects to all the other packages to make a runnable binary for multiple
operating system.


### Compile Assets

Before building, you must compile the Statping Assets with Rice, to install rice
run the command below:

    go get github.com/GeertJohan/go.rice
    go get github.com/GeertJohan/go.rice/rice

Once you have rice install, you can run the following command to build all
assets inside the source directory.

    cd source && rice embed-go


### Build Statping Binary

To build the statup binary for your local environment, run the command below:

    go build -o statup ./cmd


Build All Binary Arch's

To build Statping for Mac, Windows, Linux, and ARM devices, you can run xgo to
build for all. xgo is an awesome golang package that requires Docker.
https://github.com/karalabe/xgo

    docker pull karalabe/xgo-latest
    build-all

More info on: https://github.com/hunterlong/statping
# core
--
    import "github.com/hunterlong/statping/core"

Package core contains the main functionality of Statping. This includes
everything for Services, Hits, Failures, Users, service checking mechanisms,
databases, and notifiers in the notifier package

More info on: https://github.com/hunterlong/statping

## Usage

```go
var (
	Configs   *DbConfig // Configs holds all of the config.yml and database info
	CoreApp   *Core     // CoreApp is a global variable that contains many elements
	SetupMode bool      // SetupMode will be true if Statping does not have a database connection
	VERSION   string    // VERSION is set on build automatically by setting a -ldflag
)
```

```go
var (
	// DbSession stores the Statping database session
	DbSession *gorm.DB
	DbModels  []interface{}
)
```

#### func  CheckHash

```go
func CheckHash(password, hash string) bool
```
CheckHash returns true if the password matches with a hashed bcrypt password

#### func  CloseDB

```go
func CloseDB()
```
CloseDB will close the database connection if available

#### func  CountFailures

```go
func CountFailures() uint64
```
CountFailures returns the total count of failures for all services

#### func  CountUsers

```go
func CountUsers() int64
```
CountUsers returns the amount of users

#### func  DatabaseMaintence

```go
func DatabaseMaintence()
```
DatabaseMaintence will automatically delete old records from 'failures' and
'hits' this function is currently set to delete records 7+ days old every 60
minutes

#### func  Dbtimestamp

```go
func Dbtimestamp(group string, column string) string
```
Dbtimestamp will return a SQL query for grouping by date

#### func  DefaultPort

```go
func DefaultPort(db string) int64
```
DefaultPort accepts a database type and returns its default port

#### func  DeleteAllSince

```go
func DeleteAllSince(table string, date time.Time)
```
DeleteAllSince will delete a specific table's records based on a time.

#### func  DeleteConfig

```go
func DeleteConfig() error
```
DeleteConfig will delete the 'config.yml' file

#### func  ExportChartsJs

```go
func ExportChartsJs() string
```
ExportChartsJs renders the charts for the index page

#### func  ExportSettings

```go
func ExportSettings() ([]byte, error)
```
ExportSettings will export a JSON file containing all of the settings below: -
Core - Notifiers - Checkins - Users - Services - Groups - Messages

#### func  GetLocalIP

```go
func GetLocalIP() string
```
GetLocalIP returns the non loopback local IP of the host

#### func  InitApp

```go
func InitApp()
```
InitApp will initialize Statping

#### func  InsertLargeSampleData

```go
func InsertLargeSampleData() error
```
InsertLargeSampleData will create the example/dummy services for testing the
Statping server

#### func  InsertNotifierDB

```go
func InsertNotifierDB() error
```
InsertNotifierDB inject the Statping database instance to the Notifier package

#### func  InsertSampleData

```go
func InsertSampleData() error
```
InsertSampleData will create the example/dummy services for a brand new Statping
installation

#### func  InsertSampleHits

```go
func InsertSampleHits() error
```
InsertSampleHits will create a couple new hits for the sample services

#### func  SampleData

```go
func SampleData() error
```
SampleData runs all the sample data for a new Statping installation

#### func  Services

```go
func Services() []types.ServiceInterface
```

#### type Checkin

```go
type Checkin struct {
	*types.Checkin
}
```


#### func  AllCheckins

```go
func AllCheckins() []*Checkin
```
AllCheckins returns all checkin in system

#### func  ReturnCheckin

```go
func ReturnCheckin(c *types.Checkin) *Checkin
```
ReturnCheckin converts *types.Checking to *core.Checkin

#### func  SelectCheckin

```go
func SelectCheckin(api string) *Checkin
```
SelectCheckin will find a Checkin based on the API supplied

#### func (*Checkin) AfterFind

```go
func (c *Checkin) AfterFind() (err error)
```
AfterFind for Checkin will set the timezone

#### func (*Checkin) AllFailures

```go
func (c *Checkin) AllFailures() []*types.Failure
```
Hits returns all of the CheckinHits for a given Checkin

#### func (*Checkin) AllHits

```go
func (c *Checkin) AllHits() []*types.CheckinHit
```
AllHits returns all of the CheckinHits for a given Checkin

#### func (*Checkin) Create

```go
func (c *Checkin) Create() (int64, error)
```
Create will create a new Checkin

#### func (*Checkin) CreateFailure

```go
func (c *Checkin) CreateFailure() (int64, error)
```

#### func (*Checkin) Delete

```go
func (c *Checkin) Delete() error
```
Create will create a new Checkin

#### func (*Checkin) Expected

```go
func (c *Checkin) Expected() time.Duration
```
Expected returns the duration of when the serviec should receive a Checkin

#### func (*Checkin) Grace

```go
func (c *Checkin) Grace() time.Duration
```
Grace will return the duration of the Checkin Grace Period (after service hasn't
responded, wait a bit for a response)

#### func (*Checkin) Last

```go
func (c *Checkin) Last() *CheckinHit
```
Last returns the last checkinHit for a Checkin

#### func (*Checkin) LimitedFailures

```go
func (c *Checkin) LimitedFailures(amount int64) []types.FailureInterface
```
Hits returns all of the CheckinHits for a given Checkin

#### func (*Checkin) LimitedHits

```go
func (c *Checkin) LimitedHits(amount int64) []*types.CheckinHit
```
LimitedHits will return the last amount of successful hits from a checkin

#### func (*Checkin) Link

```go
func (c *Checkin) Link() string
```

#### func (*Checkin) Period

```go
func (c *Checkin) Period() time.Duration
```
Period will return the duration of the Checkin interval

#### func (*Checkin) RecheckCheckinFailure

```go
func (c *Checkin) RecheckCheckinFailure(guard chan struct{})
```
RecheckCheckinFailure will check if a Service Checkin has been reported yet

#### func (*Checkin) Routine

```go
func (c *Checkin) Routine()
```
Routine for checking if the last Checkin was within its interval

#### func (*Checkin) Select

```go
func (c *Checkin) Select() *types.Checkin
```
Select returns a *types.Checkin

#### func (*Checkin) Service

```go
func (c *Checkin) Service() *Service
```

#### func (*Checkin) String

```go
func (c *Checkin) String() string
```
String will return a Checkin API string

#### func (*Checkin) Update

```go
func (c *Checkin) Update() (int64, error)
```
Update will update a Checkin

#### type CheckinHit

```go
type CheckinHit struct {
	*types.CheckinHit
}
```


#### func  ReturnCheckinHit

```go
func ReturnCheckinHit(c *types.CheckinHit) *CheckinHit
```
ReturnCheckinHit converts *types.checkinHit to *core.checkinHit

#### func (*CheckinHit) AfterFind

```go
func (c *CheckinHit) AfterFind() (err error)
```
AfterFind for checkinHit will set the timezone

#### func (*CheckinHit) Ago

```go
func (c *CheckinHit) Ago() string
```
Ago returns the duration of time between now and the last successful checkinHit

#### func (*CheckinHit) Create

```go
func (c *CheckinHit) Create() (int64, error)
```
Create will create a new successful checkinHit

#### type Core

```go
type Core struct {
	*types.Core
}
```


#### func  NewCore

```go
func NewCore() *Core
```
NewCore return a new *core.Core struct

#### func  SelectCore

```go
func SelectCore() (*Core, error)
```
SelectCore will return the CoreApp global variable and the settings/configs for
Statping

#### func  UpdateCore

```go
func UpdateCore(c *Core) (*Core, error)
```
UpdateCore will update the CoreApp variable inside of the 'core' table in
database

#### func (*Core) AfterFind

```go
func (c *Core) AfterFind() (err error)
```
AfterFind for Core will set the timezone

#### func (Core) AllOnline

```go
func (c Core) AllOnline() bool
```
AllOnline will be true if all services are online

#### func (Core) BaseSASS

```go
func (c Core) BaseSASS() string
```
BaseSASS is the base design , this opens the file /assets/scss/base.scss to be
edited in Theme

#### func (*Core) Count24HFailures

```go
func (c *Core) Count24HFailures() uint64
```
Count24HFailures returns the amount of failures for a service within the last 24
hours

#### func (*Core) CountOnline

```go
func (c *Core) CountOnline() int
```
CountOnline returns the amount of services online

#### func (Core) CurrentTime

```go
func (c Core) CurrentTime() string
```
CurrentTime will return the current local time

#### func (Core) Messages

```go
func (c Core) Messages() []*Message
```
Messages will return the current local time

#### func (Core) MobileSASS

```go
func (c Core) MobileSASS() string
```
MobileSASS is the -webkit responsive custom css designs. This opens the file
/assets/scss/mobile.scss to be edited in Theme

#### func (Core) SassVars

```go
func (c Core) SassVars() string
```
SassVars opens the file /assets/scss/variables.scss to be edited in Theme

#### func (*Core) SelectAllServices

```go
func (c *Core) SelectAllServices(start bool) ([]*Service, error)
```
SelectAllServices returns a slice of *core.Service to be store on
[]*core.Services, should only be called once on startup.

#### func (*Core) ToCore

```go
func (c *Core) ToCore() *types.Core
```
ToCore will convert *core.Core to *types.Core

#### func (Core) UsingAssets

```go
func (c Core) UsingAssets() bool
```
UsingAssets will return true if /assets folder is present

#### type DateScan

```go
type DateScan struct {
	CreatedAt string `json:"x,omitempty"`
	Value     int64  `json:"y"`
}
```

DateScan struct is for creating the charts.js graph JSON array

#### type DateScanObj

```go
type DateScanObj struct {
	Array []DateScan `json:"data"`
}
```

DateScanObj struct is for creating the charts.js graph JSON array

#### func  GraphDataRaw

```go
func GraphDataRaw(service types.ServiceInterface, start, end time.Time, group string, column string) *DateScanObj
```
GraphDataRaw will return all the hits between 2 times for a Service

#### func (*DateScanObj) ToString

```go
func (d *DateScanObj) ToString() string
```
ToString will convert the DateScanObj into a JSON string for the charts to
render

#### type DbConfig

```go
type DbConfig types.DbConfig
```

DbConfig stores the config.yml file for the statup configuration

#### func  EnvToConfig

```go
func EnvToConfig() *DbConfig
```
EnvToConfig converts environment variables to a DbConfig variable

#### func  LoadConfigFile

```go
func LoadConfigFile(directory string) (*DbConfig, error)
```
LoadConfigFile will attempt to load the 'config.yml' file in a specific
directory

#### func  LoadUsingEnv

```go
func LoadUsingEnv() (*DbConfig, error)
```
LoadUsingEnv will attempt to load database configs based on environment
variables. If DB_CONN is set if will force this function.

#### func (*DbConfig) Connect

```go
func (db *DbConfig) Connect(retry bool, location string) error
```
Connect will attempt to connect to the sqlite, postgres, or mysql database

#### func (*DbConfig) CreateCore

```go
func (c *DbConfig) CreateCore() *Core
```
CreateCore will initialize the global variable 'CoreApp". This global variable
contains most of Statping app.

#### func (*DbConfig) CreateDatabase

```go
func (db *DbConfig) CreateDatabase() error
```
CreateDatabase will CREATE TABLES for each of the Statping elements

#### func (*DbConfig) DropDatabase

```go
func (db *DbConfig) DropDatabase() error
```
DropDatabase will DROP each table Statping created

#### func (*DbConfig) InsertCore

```go
func (db *DbConfig) InsertCore() (*Core, error)
```
InsertCore create the single row for the Core settings in Statping

#### func (*DbConfig) MigrateDatabase

```go
func (db *DbConfig) MigrateDatabase() error
```
MigrateDatabase will migrate the database structure to current version. This
function will NOT remove previous records, tables or columns from the database.
If this function has an issue, it will ROLLBACK to the previous state.

#### func (*DbConfig) Save

```go
func (db *DbConfig) Save() (*DbConfig, error)
```
Save will initially create the config.yml file

#### func (*DbConfig) Update

```go
func (db *DbConfig) Update() error
```
Update will save the config.yml file

#### type ErrorResponse

```go
type ErrorResponse struct {
	Error string
}
```

ErrorResponse is used for HTTP errors to show to User

#### type ExportData

```go
type ExportData struct {
	Core      *types.Core              `json:"core"`
	Services  []types.ServiceInterface `json:"services"`
	Messages  []*Message               `json:"messages"`
	Checkins  []*Checkin               `json:"checkins"`
	Users     []*User                  `json:"users"`
	Groups    []*Group                 `json:"groups"`
	Notifiers []types.AllNotifiers     `json:"notifiers"`
}
```


#### type Failure

```go
type Failure struct {
	*types.Failure
}
```


#### func (*Failure) AfterFind

```go
func (f *Failure) AfterFind() (err error)
```
AfterFind for Failure will set the timezone

#### func (*Failure) Ago

```go
func (f *Failure) Ago() string
```
Ago returns a human readable timestamp for a Failure

#### func (*Failure) Delete

```go
func (f *Failure) Delete() error
```
Delete will remove a Failure record from the database

#### func (*Failure) ParseError

```go
func (f *Failure) ParseError() string
```
ParseError returns a human readable error for a Failure

#### func (*Failure) Select

```go
func (f *Failure) Select() *types.Failure
```
Select returns a *types.Failure

#### type Group

```go
type Group struct {
	*types.Group
}
```


#### func  SelectGroup

```go
func SelectGroup(id int64) *Group
```
SelectGroup returns a *core.Group

#### func  SelectGroups

```go
func SelectGroups(includeAll bool, auth bool) []*Group
```
SelectGroups returns all groups

#### func (*Group) Create

```go
func (g *Group) Create() (int64, error)
```
Create will create a group and insert it into the database

#### func (*Group) Delete

```go
func (g *Group) Delete() error
```
Delete will remove a group

#### func (*Group) Services

```go
func (g *Group) Services() []*Service
```
Services returns all services belonging to a group

#### type Hit

```go
type Hit struct {
	*types.Hit
}
```


#### func (*Hit) AfterFind

```go
func (h *Hit) AfterFind() (err error)
```
AfterFind for Hit will set the timezone

#### type Message

```go
type Message struct {
	*types.Message
}
```


#### func  ReturnMessage

```go
func ReturnMessage(m *types.Message) *Message
```
ReturnMessage will convert *types.Message to *core.Message

#### func  SelectMessage

```go
func SelectMessage(id int64) (*Message, error)
```
SelectMessage returns a Message based on the ID passed

#### func  SelectMessages

```go
func SelectMessages() ([]*Message, error)
```
SelectMessages returns all messages

#### func  SelectServiceMessages

```go
func SelectServiceMessages(id int64) []*Message
```
SelectServiceMessages returns all messages for a service

#### func (*Message) AfterFind

```go
func (u *Message) AfterFind() (err error)
```
AfterFind for Message will set the timezone

#### func (*Message) Create

```go
func (m *Message) Create() (int64, error)
```
Create will create a Message and insert it into the database

#### func (*Message) Delete

```go
func (m *Message) Delete() error
```
Delete will delete a Message from database

#### func (*Message) Service

```go
func (m *Message) Service() *Service
```

#### func (*Message) Update

```go
func (m *Message) Update() (*Message, error)
```
Update will update a Message in the database

#### type PluginJSON

```go
type PluginJSON types.PluginJSON
```


#### type PluginRepos

```go
type PluginRepos types.PluginRepos
```


#### type Service

```go
type Service struct {
	*types.Service
}
```


#### func  ReturnService

```go
func ReturnService(s *types.Service) *Service
```
ReturnService will convert *types.Service to *core.Service

#### func  SelectService

```go
func SelectService(id int64) *Service
```
SelectService returns a *core.Service from in memory

#### func  SelectServiceLink

```go
func SelectServiceLink(permalink string) *Service
```
SelectServiceLink returns a *core.Service from the service permalink

#### func (*Service) ActiveMessages

```go
func (s *Service) ActiveMessages() []*Message
```
ActiveMessages returns all service messages that are available based on the
current time

#### func (*Service) AfterFind

```go
func (s *Service) AfterFind() (err error)
```
AfterFind for Service will set the timezone

#### func (*Service) AllCheckins

```go
func (s *Service) AllCheckins() []*Checkin
```
AllCheckins will return a slice of AllCheckins for a Service

#### func (*Service) AllFailures

```go
func (s *Service) AllFailures() []*Failure
```
AllFailures will return all failures attached to a service

#### func (*Service) AvgTime

```go
func (s *Service) AvgTime() string
```
AvgTime will return the average amount of time for a service to response back
successfully

#### func (*Service) AvgUptime

```go
func (s *Service) AvgUptime(ago time.Time) string
```
AvgUptime returns average online status for last 24 hours

#### func (*Service) AvgUptime24

```go
func (s *Service) AvgUptime24() string
```
AvgUptime24 returns a service's average online status for last 24 hours

#### func (*Service) Check

```go
func (s *Service) Check(record bool)
```
Check will run checkHttp for HTTP services and checkTcp for TCP services

#### func (*Service) CheckQueue

```go
func (s *Service) CheckQueue(record bool)
```
CheckQueue is the main go routine for checking a service

#### func (*Service) CheckinProcess

```go
func (s *Service) CheckinProcess()
```
CheckinProcess runs the checkin routine for each checkin attached to service

#### func (*Service) CountHits

```go
func (s *Service) CountHits() (int64, error)
```
CountHits returns a int64 for all hits for a service

#### func (*Service) Create

```go
func (s *Service) Create(check bool) (int64, error)
```
Create will create a service and insert it into the database

#### func (*Service) CreateFailure

```go
func (s *Service) CreateFailure(fail types.FailureInterface) (int64, error)
```
CreateFailure will create a new Failure record for a service

#### func (*Service) CreateHit

```go
func (s *Service) CreateHit(h *types.Hit) (int64, error)
```
CreateHit will create a new 'hit' record in the database for a successful/online
service

#### func (*Service) Delete

```go
func (s *Service) Delete() error
```
Delete will remove a service from the database, it will also end the service
checking go routine

#### func (*Service) DeleteFailures

```go
func (s *Service) DeleteFailures()
```
DeleteFailures will delete all failures for a service

#### func (*Service) Downtime

```go
func (s *Service) Downtime() time.Duration
```
Downtime returns the amount of time of a offline service

#### func (*Service) DowntimeText

```go
func (s *Service) DowntimeText() string
```
DowntimeText will return the amount of downtime for a service based on the
duration

    service.DowntimeText()
    // Service has been offline for 15 minutes

#### func (*Service) FailuresDaysAgo

```go
func (s *Service) FailuresDaysAgo(days int) uint64
```
FailuresDaysAgo returns the amount of failures since days ago

#### func (*Service) Hits

```go
func (s *Service) Hits() ([]*types.Hit, error)
```
Hits returns all successful hits for a service

#### func (*Service) HitsBetween

```go
func (s *Service) HitsBetween(t1, t2 time.Time, group string, column string) *gorm.DB
```
HitsBetween returns the gorm database query for a collection of service hits
between a time range

#### func (*Service) LimitedCheckinFailures

```go
func (s *Service) LimitedCheckinFailures(amount int64) []*Failure
```
LimitedFailures will return the last amount of failures from a service

#### func (*Service) LimitedFailures

```go
func (s *Service) LimitedFailures(amount int64) []*Failure
```
LimitedFailures will return the last amount of failures from a service

#### func (*Service) LimitedHits

```go
func (s *Service) LimitedHits(amount int64) ([]*types.Hit, error)
```
LimitedHits returns the last 1024 successful/online 'hit' records for a service

#### func (*Service) Messages

```go
func (s *Service) Messages() []*Message
```
Messages returns all Messages for a Service

#### func (*Service) OnlineDaysPercent

```go
func (s *Service) OnlineDaysPercent(days int) float32
```
OnlineDaysPercent returns the service's uptime percent within last 24 hours

#### func (*Service) OnlineSince

```go
func (s *Service) OnlineSince(ago time.Time) float32
```
OnlineSince accepts a time since parameter to return the percent of a service's
uptime.

#### func (*Service) Select

```go
func (s *Service) Select() *types.Service
```
Select will return the *types.Service struct for Service

#### func (*Service) SmallText

```go
func (s *Service) SmallText() string
```
SmallText returns a short description about a services status

    service.SmallText()
    // Online since Monday 3:04:05PM, Jan _2 2006

#### func (*Service) SparklineDayFailures

```go
func (s *Service) SparklineDayFailures(days int) string
```
SparklineDayFailures returns a string array of daily service failures

#### func (*Service) SparklineHourResponse

```go
func (s *Service) SparklineHourResponse(hours int, method string) string
```
SparklineHourResponse returns a string array for the average response or ping
time for a service

#### func (*Service) Sum

```go
func (s *Service) Sum() float64
```
Sum returns the added value Latency for all of the services successful hits.

#### func (*Service) ToJSON

```go
func (s *Service) ToJSON() string
```
ToJSON will convert a service to a JSON string

#### func (*Service) TotalFailures

```go
func (s *Service) TotalFailures() (uint64, error)
```
TotalFailures returns the total amount of failures for a service

#### func (*Service) TotalFailures24

```go
func (s *Service) TotalFailures24() (uint64, error)
```
TotalFailures24 returns the amount of failures for a service within the last 24
hours

#### func (*Service) TotalFailuresOnDate

```go
func (s *Service) TotalFailuresOnDate(ago time.Time) (uint64, error)
```
TotalFailuresOnDate returns the total amount of failures for a service on a
specific time/date

#### func (*Service) TotalFailuresSince

```go
func (s *Service) TotalFailuresSince(ago time.Time) (uint64, error)
```
TotalFailuresSince returns the total amount of failures for a service since a
specific time/date

#### func (*Service) TotalHits

```go
func (s *Service) TotalHits() (uint64, error)
```
TotalHits returns the total amount of successful hits a service has

#### func (*Service) TotalHitsSince

```go
func (s *Service) TotalHitsSince(ago time.Time) (uint64, error)
```
TotalHitsSince returns the total amount of hits based on a specific time/date

#### func (*Service) TotalUptime

```go
func (s *Service) TotalUptime() string
```
TotalUptime returns the total uptime percent of a service

#### func (*Service) Update

```go
func (s *Service) Update(restart bool) error
```
Update will update a service in the database, the service's checking routine can
be restarted by passing true

#### type ServiceOrder

```go
type ServiceOrder []types.ServiceInterface
```

ServiceOrder will reorder the services based on 'order_id' (Order)

#### func (ServiceOrder) Len

```go
func (c ServiceOrder) Len() int
```
Sort interface for resroting the Services in order

#### func (ServiceOrder) Less

```go
func (c ServiceOrder) Less(i, j int) bool
```

#### func (ServiceOrder) Swap

```go
func (c ServiceOrder) Swap(i, j int)
```

#### type User

```go
type User struct {
	*types.User
}
```


#### func  AuthUser

```go
func AuthUser(username, password string) (*User, bool)
```
AuthUser will return the User and a boolean if authentication was correct.
AuthUser accepts username, and password as a string

#### func  ReturnUser

```go
func ReturnUser(u *types.User) *User
```
ReturnUser returns *core.User based off a *types.User

#### func  SelectAllUsers

```go
func SelectAllUsers() ([]*User, error)
```
SelectAllUsers returns all users

#### func  SelectUser

```go
func SelectUser(id int64) (*User, error)
```
SelectUser returns the User based on the User's ID.

#### func  SelectUsername

```go
func SelectUsername(username string) (*User, error)
```
SelectUsername returns the User based on the User's username

#### func (*User) AfterFind

```go
func (u *User) AfterFind() (err error)
```
AfterFind for USer will set the timezone

#### func (*User) Create

```go
func (u *User) Create() (int64, error)
```
Create will insert a new User into the database

#### func (*User) Delete

```go
func (u *User) Delete() error
```
Delete will remove the User record from the database

#### func (*User) Update

```go
func (u *User) Update() error
```
Update will update the User's record in database
# handlers
--
    import "github.com/hunterlong/statping/handlers"

Package handlers contains the HTTP server along with the requests and routes.
All HTTP related functions are in this package.

More info on: https://github.com/hunterlong/statping

## Usage

#### func  ExecuteResponse

```go
func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}, redirect interface{})
```
ExecuteResponse will render a HTTP response for the front end user

#### func  IsAdmin

```go
func IsAdmin(r *http.Request) bool
```
IsAdmin returns true if the user session is an administrator

#### func  IsFullAuthenticated

```go
func IsFullAuthenticated(r *http.Request) bool
```
IsFullAuthenticated returns true if the HTTP request is authenticated. You can
set the environment variable GO_ENV=test to bypass the admin authenticate to the
dashboard features.

#### func  IsReadAuthenticated

```go
func IsReadAuthenticated(r *http.Request) bool
```
IsReadAuthenticated will allow Read Only authentication for some routes

#### func  IsUser

```go
func IsUser(r *http.Request) bool
```
IsUser returns true if the user is registered

#### func  Router

```go
func Router() *mux.Router
```
Router returns all of the routes used in Statping. Server will use static assets
if the 'assets' directory is found in the root directory.

#### func  RunHTTPServer

```go
func RunHTTPServer(ip string, port int) error
```
RunHTTPServer will start a HTTP server on a specific IP and port

#### type Cacher

```go
type Cacher interface {
	Get(key string) []byte
	Delete(key string)
	Set(key string, content []byte, duration time.Duration)
	List() map[string]Item
}
```


```go
var CacheStorage Cacher
```

#### type Item

```go
type Item struct {
	Content    []byte
	Expiration int64
}
```

Item is a cached reference

#### func (Item) Expired

```go
func (item Item) Expired() bool
```
Expired returns true if the item has expired.

#### type PluginSelect

```go
type PluginSelect struct {
	Plugin string
	Form   string
	Params map[string]interface{}
}
```


#### type Storage

```go
type Storage struct {
}
```

Storage mecanism for caching strings in memory

#### func  NewStorage

```go
func NewStorage() *Storage
```
NewStorage creates a new in memory CacheStorage

#### func (Storage) Delete

```go
func (s Storage) Delete(key string)
```

#### func (Storage) Get

```go
func (s Storage) Get(key string) []byte
```
Get a cached content by key

#### func (Storage) List

```go
func (s Storage) List() map[string]Item
```

#### func (Storage) Set

```go
func (s Storage) Set(key string, content []byte, duration time.Duration)
```
Set a cached content by key
# notifiers
--
    import "github.com/hunterlong/statping/notifiers"

Package notifiers holds all the notifiers for Statping, which also includes user
created notifiers that have been accepted in a Push Request. Read the wiki to
see a full example of a notifier with all events, visit Statping's notifier
example code: https://github.com/hunterlong/statping/wiki/Notifier-Example

This package shouldn't contain any exports, to see how notifiers work visit the
core/notifier package at:
https://godoc.org/github.com/hunterlong/statping/core/notifier and learn how to
create your own custom notifier.

## Usage

#### type MobileResponse

```go
type MobileResponse struct {
	Counts  int                   `json:"counts"`
	Logs    []*MobileResponseLogs `json:"logs"`
	Success string                `json:"success"`
}
```


#### type MobileResponseLogs

```go
type MobileResponseLogs struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}
```


#### type PushArray

```go
type PushArray struct {
	Tokens   []string               `json:"tokens"`
	Platform int64                  `json:"platform"`
	Message  string                 `json:"message"`
	Topic    string                 `json:"topic"`
	Title    string                 `json:"title,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty"`
}
```


#### type PushNotification

```go
type PushNotification struct {
	Array []*PushArray `json:"notifications"`
}
```
# plugin
--
    import "github.com/hunterlong/statping/plugin"

Package plugin contains the interfaces to build your own Golang Plugin that will
receive triggers on Statping events.

## Usage

```go
var (
	AllPlugins []*types.PluginObject
)
```

#### func  LoadPlugin

```go
func LoadPlugin(file string) error
```

#### func  LoadPlugins

```go
func LoadPlugins()
```
# source
--
    import "github.com/hunterlong/statping/source"

Package source holds all the assets for Statping. This includes CSS, JS, SCSS,
HTML and other website related content. This package uses Rice to compile all
assets into a single 'rice-box.go' file.


### Required Dependencies

- rice -> https://github.com/GeertJohan/go.rice - sass ->
https://sass-lang.com/install


### Compile Assets

To compile all the HTML, JS, SCSS, CSS and image assets you'll need to have rice
and sass installed on your local system.

    sass source/scss/base.scss source/css/base.css
    cd source && rice embed-go

More info on: https://github.com/hunterlong/statping

Code generated by go generate; DO NOT EDIT. This file was generated by robots at
2019-02-06 12:42:08.202468 -0800 PST m=+0.598756678

This contains the most recently Markdown source for the Statping Wiki.

## Usage

```go
var (
	CssBox  *rice.Box // CSS files from the 'source/css' directory, this will be loaded into '/assets/css'
	ScssBox *rice.Box // SCSS files from the 'source/scss' directory, this will be loaded into '/assets/scss'
	JsBox   *rice.Box // JS files from the 'source/js' directory, this will be loaded into '/assets/js'
	TmplBox *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
	FontBox *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
)
```

```go
var CompiledWiki = []byte("<a class=\"scrollclick\" href=\"#\" data-id=\"page_0\">Types of Monitoring</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_1\">Features</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_2\">Start Statping</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_3\">Linux</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_4\">Mac</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_5\">Windows</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_6\">AWS EC2</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_7\">Docker</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_8\">Mobile App</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_9\">Heroku</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_10\">API</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_11\">Makefile</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_12\">Notifiers</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_13\">Notifier Events</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_14\">Notifier Example</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_15\">Prometheus Exporter</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_16\">SSL</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_17\">Config with .env File</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_18\">Static Export</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_19\">Statping Plugins</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_20\">Statuper</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_21\">Build and Test</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_22\">Contributing</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_23\">PGP Signature</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_24\">Testing</a><br><a class=\"scrollclick\" href=\"#\" data-id=\"page_25\">Deployment</a><br>\n\n<div class=\"mt-5\" id=\"page_0\"><h1>Types of Monitoring</h1></div>\nYou can monitor your application by using a simple HTTP GET to the endpoint to return back a response and status code. Normally you want a 200 status code on an HTTP request. You might want to require a 404 or 500 error as a response code though. With each service you can include a Timeout in seconds to work with your long running services.\n\n# HTTP Endpoints with Custom POST\nFor more advanced monitoring you can add a data as a HTTP POST request. This is useful for automatically submitting JSON, or making sure your signup form is working correctly.\n\n<p align=\"center\">\n<img width=\"100%\" src=\"https://img.cjx.io/statup-httpservice.png\">\n</p>\n\nWith a HTTP service, you can POST a JSON string to your endpoint to retrieve any type of response back. You can then use Regex in the Expected Response field to parse a custom response that exactly matches your status requirements. \n\n# TCP Services\nFor other services that don't use HTTP, you can monitor any type of service by using the PORT of the service. If you're Ethereum Blockchain server is running on 8545, you can use TCP to monitor your server. With a TCP service, you can monitor your Docker containers, or remove service running on a custom port. You don't need to include `http` in the endpoint field, just IP or Hostname.\n\n<p align=\"center\">\n<img width=\"100%\" src=\"https://img.cjx.io/statup-tcpservice.png\">\n</p>\n\n\n<div class=\"mt-5\" id=\"page_1\"><h1>Features</h1></div>\nStatping is a great Status Page that can be deployed with 0 effort.\n\n# 3 Different Databases\nYou can use MySQL, Postgres, or SQLite as a database for your Statping status page. The server will automatically upgrade your database tables depending on which database you have.\n\n# Easy to Startup\nStatping is an extremely easy to setup website monitoring tool without fussing with dependencies or packages. Simply download and install the precompile binary for your operating system. Statping works on Windows, Mac, Linux, Docker, and even the Raspberry Pi.\n\n# Plugins\nStatping is an awesome Status Page generator that allows you to create your own plugins with Golang Plugins! You don't need to request a PR or even tell us about your plugin. Plugin's are compiled and then send as a binary to the Statping `/plugins` folder. Test your plugins using the `statup test plugin` command, checkout the [Plugin Wiki](https://github.com/hunterlong/statping/wiki/Statping-Plugins) to see detailed information about creating plugins.\n\n# No Maintence\nMany other website monitoring applications will collect data until the server fails because of hard drive is 100% full. Statping will automatically delete records to make sure your server will stay UP for years. The EC2 AMI Image is a great way to host your status page without worrying about it crashing one day. Statping will automatically upgrade its software when you reboot your computer.\n\n# Email & Slack Notifications\nReceive email notifications if your website or application goes offline. Statping includes SMTP connections so you can use AWS SES, or any other SMTP emailing service. Go in the Email Settings in Settings to configure these options.\n\n# Prometheus Exporter\nIf you want a deeper view of your applications status, you can use Grafana and Prometheus to graph all types of data about your services. Read more about the [Prometheus Exporter](https://github.com/hunterlong/statping/wiki/Prometheus-Exporter)\n\n<div class=\"mt-5\" id=\"page_2\"><h1>Start Statping</h1></div>\n\n\n<div class=\"mt-5\" id=\"page_3\"><h1>Linux</h1></div>\n# Installing on Linux\nInstalling Statping on Linux is a 1 line command. It's that easy.\n```\nbash <(curl -s https://assets.statup.io/install.sh)\nstatping version\n```\n[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/statping)\n\nIf you are using [snap](https://snapcraft.io/statping), you can simply run this command to install Statping.\n```shell\nsudo snap install statping\n```\n\n## Systemd Service\nSetting up a systemd service is a great way to make sure your Statping server will automatically reboot when needed. You can use the file below for your service. You should have Statping already installed by this step.\n###### /etc/systemd/system/statping.service\n```\n[Unit]\nDescription=Statping Server\nAfter=network.target\nAfter=systemd-user-sessions.service\nAfter=network-online.target\n\n[Service]\nType=simple\nRestart=always\nExecStart=/usr/local/bin/statping\n\n[Install]\nWantedBy=multi-user.target\n```\nThen you can enable and start your systemd service with:\n```\nsystemctl daemon-reload\n\nsystemctl enable statping.service\n\nsystemctl start statping\n```\nYou're Statping server will now automatically restart when your server restarts.\n\n## Raspberry Pi\nYou can even run Statping on your Raspberry Pi by installing the precompiled binary from [Latest Releases](https://github.com/hunterlong/statping/releases/latest). For the Raspberry Pi 3 you'll want to download the `statping-linux-arm7.tar.gz` file. Be sure to change `VERSION` to the latest version in Releases, and include the 'v'.\n\n```\nVERSION=$(curl -s \"https://github.com/hunterlong/statping/releases/latest\" | grep -o 'tag/[v.0-9]*' | awk -F/ '{print $2}')\nwget https://github.com/hunterlong/statping/releases/download/$VERSION/statping-linux-arm7.tar.gz\ntar -xvzf statping-linux-arm7.tar.gz\nchmod +x statping\nmv statping /usr/local/bin/statping\n\nstatping version\n``` \n\n## Alpine Linux\nThe Docker image is using the Statping Alpine binary since it's so incredibly small. You can run it on your own alpine image by downloading `statping-linux-alpine.tar.gz` from [Latest Releases](https://github.com/hunterlong/statping/releases/latest).\n\n<div class=\"mt-5\" id=\"page_4\"><h1>Mac</h1></div>\n# Installing on Mac\nStatping includes an easy to use [Homebrew Formula](https://github.com/hunterlong/homebrew-statping) to quick get your Status Page up and running locally. Statping on brew is automatically generated for each new release to master. Install with the commands below,\n```bash\nbrew tap hunterlong/statping\nbrew install statping\n```\n\n<p align=\"center\">\n<img width=\"80%\" src=\"https://img.cjx.io/statupbrewinstall.gif\">\n</p>\n\nIf you don't have brew, then you can install it with this command below:\n```bash\nbash <(curl -s https://statping.com/install.sh)\n```\n\nOnce you've installed it, checkout which version you have by running `statping version`.\n\n<div class=\"mt-5\" id=\"page_5\"><h1>Windows</h1></div>\n# Installing on Windows\nCurrently, Statping only works on Windows 64-bit computers. Just download the exe file from [Latest Releases](https://github.com/hunterlong/statping/releases/latest) and run it in your command prompt. It will create a HTTP server on port 8080, so you can visit `http://localhost:8080` to see your Statping Status Page.\n\n## Known Issues with Windows\nUnfortunately, Statping only works on Windows 64-bit processors. If you have more than 4gb of ram, there's a good chance you already have a 64-bit processor. Download the [Latest Releases](https://github.com/hunterlong/statping/releases/latest) of Statping, extract the ZIP file, then double click on the `statping.exe` file. You can use a SQLite database for a quick setup, or connect to a local/remote Postgres or MySQL database server.\n\n<div class=\"mt-5\" id=\"page_6\"><h1>AWS EC2</h1></div>\nRunning Statping on the smallest EC2 server is very quick using the AWS AMI Image. The AWS AMI Image will automatically start a Statping Docker container that will automatically update to the latest version. Once the EC2 is booted, you can go to the Public DNS domain to view the Statping installation page. The Statping root folder is located at: `/statping` on the server.\n\n# AMI Image\nChoose the correct AMI Image ID based on your AWS region.\n- us-east-1 `ami-09ccd23d9c7afba61` (Virginia)\n- us-east-2 `ami-0c6c9b714a501cdb3` (Ohio)\n- us-west-1 `ami-02159cc1fc701a77e` (California)\n- us-west-2 `ami-007c6990949f5ccee` (Oregon)\n- eu-central-1 `ami-06e252d6d8b0c2f1f` (Frankfurt)\n\n# Instructions\n\n### 1. Create an EC2 instance from AMI Image\nGo to the main EC2 dashboard and click 'Launch Instance'. Then type `Statping` inside the search field for 'Community AMI'. Once you've found it in your region, click Select!\n\n<img src=\"https://img.cjx.io/statpingaws-ami.png\">\n\n### 2. Get the Public DNS for EC2 Instance\nCopy the 'Public DNS' URL and paste it into your browser.\n\n<img src=\"https://img.cjx.io/statping-aws-ec2.png\">\n\n### 3. Setup Statping\nUse SQLite if you don't want to connect to a remote MySQL or Postgres database.\n\n<img src=\"https://img.cjx.io/statping-setup.png\">\n\n# EC2 Server Features\nRunning your Statping server on a small EC2 instance is perfect for most users. Below you'll find some commands to get up and running in seconds.\n- Super cheap on the t2.nano (~$4.60 monthly)\n- Small usage, 8gb of hard drive\n- Automatic SSL certificate if you require it\n- Automatic reboot when the server needs it\n- Automatic database cleanup, so you'll never be at 100% full.\n- Automatic docker containers/images removal\n\n## Create Security Groups\nUsing the AWS CLI you can copy and paste the commands below to auto create everything for you. The server opens port 80 and 443.\n```bash\naws ec2 create-security-group --group-name StatpingPublicHTTP --description \"Statping HTTP Server on port 80 and 443\"\n# will response back a Group ID. Copy ID and use it for --group-id below.\n```\n```bash\nGROUPS=sg-7e8b830f\naws ec2 authorize-security-group-ingress --group-id $GROUPS --protocol tcp --port 80 --cidr 0.0.0.0/0\naws ec2 authorize-security-group-ingress --group-id $GROUPS --protocol tcp --port 443 --cidr 0.0.0.0/0\n```\n## Create EC2 without SSL\nOnce your server has started, go to the EC2 Public DNS endpoint. You should be redirected to /setup to continue your installation process! The database information is already inputed for you.\n```bash\nGROUPS=sg-7e8b830f\nKEY=MYKEYHERE\nAMI_IMAGE=ami-7be8a103\n\naws ec2 run-instances \\\n    --image-id $AMI_IMAGE \\\n    --count 1 --instance-type t2.nano \\\n    --key-name $KEY \\\n    --security-group-ids $GROUPS\n```\n## Create EC2 with Automatic SSL Certification\nStart a Statping server with an SSL cert that will automatically regenerate when it's near expiration time. You'll need to point your domain's A record (IP address) or CNAME (public DNS endpoint) to use this feature.\n\n```bash\nwget https://raw.githubusercontent.com/hunterlong/statping/master/dev/ec2-ssl.sh\n```\n\n```bash\n# edit the contents inside of ec2-ssl.sh then continue\nLETSENCRYPT_HOST=\"status.MYDOMAIN.com\"\nLETSENCRYPT_EMAIL=\"noreply@MYEMAIL.com\"\n```\nEdit ec2-ssl.sh and insert your domain you want to use, then run command below. Use the Security Group ID that you used above for --security-group-ids\n```\nGROUPS=sg-7e8b830f\nAMI_IMAGE=ami-7be8a103\nKEY=MYKEYHERE\n\naws ec2 run-instances \\\n    --user-data file://ec2-ssl.sh \\\n    --image-id $AMI_IMAGE \\\n    --count 1 --instance-type t2.nano \\\n    --key-name $KEY \\\n    --security-group-ids $GROUPS\n```\n\n### EC2 Server Specs\n- t2.nano ($4.60 monthly)\n- 8gb SSD Memory\n- 0.5gb RAM\n- Docker with Docker Compose installed\n- Running Statping, NGINX, and Postgres\n- boot scripts to automatically clean unused containers.\n\n\n\n<div class=\"mt-5\" id=\"page_7\"><h1>Docker</h1></div>\nStatping is easily ran on Docker with the light weight Alpine linux image. View on [Docker Hub](https://hub.docker.com/r/hunterlong/statping).\n\n[![](https://images.microbadger.com/badges/image/hunterlong/statping.svg)](https://microbadger.com/images/hunterlong/statping) [![Docker Pulls](https://img.shields.io/docker/pulls/hunterlong/statping.svg)](https://hub.docker.com/r/hunterlong/statping/builds/)\n\n# Latest Docker Image\nThe `latest` Docker image uses Alpine Linux to keep it ultra small.\n```bash\ndocker run -d \\\n  -p 8080:8080 \\\n  --restart always \\\n  hunterlong/statping\n```\n\n# Mounting Volume\nYou can mount a volume to the `/app` Statping directory. This folder will contain `logs`, `config.yml`, and static assets if you want to edit the SCSS/CSS. \n```bash\ndocker run -d \\\n  -p 8080:8080 \\\n  -v /mydir/statping:/app \\\n  --restart always \\\n  hunterlong/statping\n```\n\n# Attach a SSL Certificate\nWhen you mount `server.crt` and `server.key` to the `/app` directory, Statping will run a HTTPS server on port 443. Checkout the [SSL Wiki](https://github.com/hunterlong/statping/wiki/SSL) documentation to see more information about this.\n```bash\ndocker run -d \\\n  -p 443:443 \\\n  -v /mydir/domain.crt:/app/server.crt \\\n  -v /mydir/domain.key:/app/server.key \\\n  -v /mydir:/app \\\n  --restart always \\\n  hunterlong/statping\n```\n\n# Development Docker Image\nIf you want to run Statping that was build from the source, use the `dev` Docker image.\n```bash\ndocker run -d -p 8080:8080 hunterlong/statping:dev\n```\n\n# Cypress Testing Docker Image\nThis Docker image will pull the latest version of Statping and test the web interface with [Cypress](https://www.cypress.io/).\n```bash\ndocker run -it -p 8080:8080 hunterlong/statping:cypress\n```\n\n#### Or use Docker Compose\nThis Docker Compose file inlcudes NGINX, Postgres, and Statping.\n\n### Docker Compose with NGINX and Postgres\nOnce you initiate the `docker-compose.yml` file below go to http://localhost and you'll be forwarded to the /setup page. \nDatabase Authentication\n- database: `postgres`\n- port: `5432`\n- username: `statup`\n- password: `password123`\n- database: `statup`\n\n```yaml\nversion: '2.3'\n\nservices:\n\n  nginx:\n    container_name: nginx\n    image: jwilder/nginx-proxy\n    ports:\n      - 0.0.0.0:80:80\n      - 0.0.0.0:443:443\n    networks:\n      - internet\n    restart: always\n    volumes:\n      - /var/run/docker.sock:/tmp/docker.sock:ro\n      - ./statup/nginx/certs:/etc/nginx/certs:ro\n      - ./statup/nginx/vhost:/etc/nginx/vhost.d\n      - ./statup/nginx/html:/usr/share/nginx/html:ro\n      - ./statup/nginx/dhparam:/etc/nginx/dhparam\n    environment:\n      DEFAULT_HOST: localhost\n\n  statup:\n    container_name: statup\n    image: hunterlong/statping:latest\n    restart: always\n    networks:\n      - internet\n      - database\n    depends_on:\n      - postgres\n    volumes:\n      - ./statup/app:/app\n    environment:\n      VIRTUAL_HOST: localhost\n      VIRTUAL_PORT: 8080\n      DB_CONN: postgres\n      DB_HOST: postgres\n      DB_USER: statup\n      DB_PASS: password123\n      DB_DATABASE: statup\n      NAME: EC2 Example\n      DESCRIPTION: This is a Statping Docker Compose instance\n\n  postgres:\n    container_name: postgres\n    image: postgres\n    restart: always\n    networks:\n      - database\n    volumes:\n      - ./statup/postgres:/var/lib/postgresql/data\n    environment:\n      POSTGRES_PASSWORD: password123\n      POSTGRES_USER: statup\n      POSTGRES_DB: statup\n\nnetworks:\n  internet:\n    driver: bridge\n  database:\n    driver: bridge\n```\nOr a simple wget...\n```bash\nwget https://raw.githubusercontent.com/hunterlong/statping/master/servers/docker-compose.yml\ndocker-compose up -d\n```\n\n#### Docker Compose with Automatic SSL\nYou can automatically start a Statping server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statping server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.\n```bash\nwget https://raw.githubusercontent.com/hunterlong/statping/master/servers/docker-compose-ssl.yml\n\nLETSENCRYPT_HOST=mydomain.com \\\n    LETSENCRYPT_EMAIL=info@mydomain.com \\\n    docker-compose -f docker-compose-ssl.yml up -d\n```\n\n### Full docker-compose with Automatic SSL\n\n```yaml\nversion: '2.3'\n\nservices:\n\n  nginx:\n    container_name: nginx\n    image: jwilder/nginx-proxy\n    ports:\n      - 0.0.0.0:80:80\n      - 0.0.0.0:443:443\n    labels:\n      - \"com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy\"\n    networks:\n      - internet\n    restart: always\n    volumes:\n      - /var/run/docker.sock:/tmp/docker.sock:ro\n      - ./statup/nginx/certs:/etc/nginx/certs:ro\n      - ./statup/nginx/vhost:/etc/nginx/vhost.d\n      - ./statup/nginx/html:/usr/share/nginx/html:ro\n      - ./statup/nginx/dhparam:/etc/nginx/dhparam\n    environment:\n      DEFAULT_HOST: ${LETSENCRYPT_HOST}\n\n  letsencrypt:\n    container_name: letsencrypt\n    image: jrcs/letsencrypt-nginx-proxy-companion\n    networks:\n      - internet\n    restart: always\n    volumes:\n      - /var/run/docker.sock:/var/run/docker.sock:ro\n      - ./statup/nginx/certs:/etc/nginx/certs\n      - ./statup/nginx/vhost:/etc/nginx/vhost.d\n      - ./statup/nginx/html:/usr/share/nginx/html\n      - ./statup/nginx/dhparam:/etc/nginx/dhparam\n\n  statup:\n    container_name: statup\n    image: hunterlong/statping:latest\n    restart: always\n    networks:\n      - internet\n      - database\n    depends_on:\n      - postgres\n    volumes:\n      - ./statup/app:/app\n    environment:\n      VIRTUAL_HOST: ${LETSENCRYPT_HOST}\n      VIRTUAL_PORT: 8080\n      LETSENCRYPT_HOST: ${LETSENCRYPT_HOST}\n      LETSENCRYPT_EMAIL: ${LETSENCRYPT_EMAIL}\n      DB_CONN: postgres\n      DB_HOST: postgres\n      DB_USER: statup\n      DB_PASS: password123\n      DB_DATABASE: statup\n      NAME: SSL Example\n      DESCRIPTION: This Status Status Page should be running ${LETSENCRYPT_HOST} with SSL.\n\n  postgres:\n    container_name: postgres\n    image: postgres\n    restart: always\n    networks:\n      - database\n    volumes:\n      - ./statup/postgres:/var/lib/postgresql/data\n    environment:\n      POSTGRES_PASSWORD: password123\n      POSTGRES_USER: statup\n      POSTGRES_DB: statup\n\nnetworks:\n  internet:\n    driver: bridge\n  database:\n    driver: bridge\n```\n\n<div class=\"mt-5\" id=\"page_8\"><h1>Mobile App</h1></div>\nStatping has a free mobile app so you can monitor your websites and applications without the need of a computer. It's currently in **beta** on Google Play and the App Store, try it out!\n\n<p align=\"center\">\n<a href=\"https://play.google.com/store/apps/details?id=com.statping\"><img src=\"https://img.cjx.io/google-play.svg\"></a>\n<a href=\"https://itunes.apple.com/us/app/apple-store/id1445513219\"><img src=\"https://img.cjx.io/app-store-badge.svg\"></a>\n</p>\n\n<p align=\"center\">\n<img src=\"https://img.cjx.io/statping_iphone_bk.png\">\n</p>\n\n\n\n<div class=\"mt-5\" id=\"page_9\"><h1>Heroku</h1></div>\nYou can now instantly deploy your Statping instance on a free Heroku container. Simply click the deploy button below and get up in running within seconds. This Heroku deployment is based on the Statping Docker image so you will have all the great features including SASS and all the notifiers without any setup. \n\n[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/hunterlong/statping/tree/master)\n\nView the live Heroku Statping instance at: [https://statping.herokuapp.com](https://statping.herokuapp.com)\n\n# Database Configuration\nYou will need to deploy a Postgres database to your instance and insert some configuration variables. View the image below to see what environment variable you need to configure. If you insert `DB_CONN`, Statping will attempt to automatically connect to the database without the need for the `config.yml` file. \n\n![](https://img.cjx.io/herokustatping.png)\n\n\n<div class=\"mt-5\" id=\"page_10\"><h1>API</h1></div>\nStatping includes a RESTFUL API so you can view, update, and edit your services with easy to use routes. You can currently view, update and delete services, view, create, update users, and get detailed information about the Statping instance. To make life easy, try out a Postman or Swagger JSON file and use it on your Statping Server.\n\n<p align=\"center\">\n<a href=\"https://documenter.getpostman.com/view/1898229/RzfiJUd6\">Postman</a> | <a href=\"https://github.com/hunterlong/statping/blob/master/source/tmpl/postman.json\">Postman JSON Export</a> | <a href=\"https://github.com/hunterlong/statping/blob/master/dev/swagger.json\">Swagger Export</a>\n</p>\n\n## Authentication\nAuthentication uses the Statping API Secret to accept remote requests. You can find the API Secret in the Settings page of your Statping server. To send requests to your Statping API, include a Authorization Header when you send the request. The API will accept any one of the headers below.\n\n- HTTP Header: `Authorization: API SECRET HERE`\n- HTTP Header: `Authorization: Bearer API SECRET HERE`\n\n## Main Route `/api`\nThe main API route will show you all services and failures along with them.\n\n## Services\nThe services API endpoint will show you detailed information about services and will allow you to edit/delete services with POST/DELETE http methods.\n\n### Viewing All Services\n- Endpoint: `/api/services`\n- Method: `GET`\n- Response: Array of [Services](https://github.com/hunterlong/statping/wiki/API#service-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\n### Viewing Service\n- Endpoint: `/api/services/{id}`\n- Method: `GET`\n- Response: [Service](https://github.com/hunterlong/statping/wiki/API#service-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\n### Updating Service\n- Endpoint: `/api/services/{id}`\n- Method: `POST`\n- Response: [Service](https://github.com/hunterlong/statping/wiki/API#service-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\nPOST Data:\n```json\n{\n    \"name\": \"Updated Service\",\n    \"domain\": \"https://google.com\",\n    \"expected\": \"\",\n    \"expected_status\": 200,\n    \"check_interval\": 15,\n    \"type\": \"http\",\n    \"method\": \"GET\",\n    \"post_data\": \"\",\n    \"port\": 0,\n    \"timeout\": 10,\n    \"order_id\": 0\n}\n```\n\n### Deleting Service\n- Endpoint: `/api/services/{id}`\n- Method: `DELETE`\n- Response: [Object Response](https://github.com/hunterlong/statping/wiki/API#object-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\nResponse:\n```json\n{\n    \"status\": \"success\",\n    \"id\": 4,\n    \"type\": \"service\",\n    \"method\": \"delete\"\n}\n```\n\n## Users\nThe users API endpoint will show you users that are registered inside your Statping instance.\n\n### View All Users\n- Endpoint: `/api/users`\n- Method: `GET`\n- Response: Array of [Users](https://github.com/hunterlong/statping/wiki/API#user-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\n### Viewing User\n- Endpoint: `/api/users/{id}`\n- Method: `GET`\n- Response: [User](https://github.com/hunterlong/statping/wiki/API#user-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\n### Creating New User\n- Endpoint: `/api/users`\n- Method: `POST`\n- Response: [User](https://github.com/hunterlong/statping/wiki/API#user-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\nPOST Data:\n```json\n{\n    \"username\": \"newadmin\",\n    \"email\": \"info@email.com\",\n    \"password\": \"password123\",\n    \"admin\": true\n}\n```\n\n### Updating User\n- Endpoint: `/api/users/{id}`\n- Method: `POST`\n- Response: [User](https://github.com/hunterlong/statping/wiki/API#user-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\nPOST Data:\n```json\n{\n    \"username\": \"updatedadmin\",\n    \"email\": \"info@email.com\",\n    \"password\": \"password123\",\n    \"admin\": true\n}\n```\n\n### Deleting User\n- Endpoint: `/api/services/{id}`\n- Method: `DELETE`\n- Response: [Object Response](https://github.com/hunterlong/statping/wiki/API#object-response)\n- Response Type: `application/json`\n- Request Type: `application/json`\n\nResponse:\n```json\n{\n    \"status\": \"success\",\n    \"id\": 3,\n    \"type\": \"user\",\n    \"method\": \"delete\"\n}\n```\n\n# Service Response\n```json\n{\n    \"id\": 8,\n    \"name\": \"Test Service 0\",\n    \"domain\": \"https://status.coinapp.io\",\n    \"expected\": \"\",\n    \"expected_status\": 200,\n    \"check_interval\": 1,\n    \"type\": \"http\",\n    \"method\": \"GET\",\n    \"post_data\": \"\",\n    \"port\": 0,\n    \"timeout\": 30,\n    \"order_id\": 0,\n    \"created_at\": \"2018-09-12T09:07:03.045832088-07:00\",\n    \"updated_at\": \"2018-09-12T09:07:03.046114305-07:00\",\n    \"online\": false,\n    \"latency\": 0.031411064,\n    \"24_hours_online\": 0,\n    \"avg_response\": \"\",\n    \"status_code\": 502,\n    \"last_online\": \"0001-01-01T00:00:00Z\",\n    \"dns_lookup_time\": 0.001727175,\n    \"failures\": [\n        {\n            \"id\": 5187,\n            \"issue\": \"HTTP Status Code 502 did not match 200\",\n            \"created_at\": \"2018-09-12T10:41:46.292277471-07:00\"\n        },\n        {\n            \"id\": 5188,\n            \"issue\": \"HTTP Status Code 502 did not match 200\",\n            \"created_at\": \"2018-09-12T10:41:47.337659862-07:00\"\n        }\n    ]\n}\n```\n\n# User Response\n```json\n{\n    \"id\": 1,\n    \"username\": \"admin\",\n    \"api_key\": \"02f324450a631980121e8fd6ea7dfe4a7c685a2f\",\n    \"admin\": true,\n    \"created_at\": \"2018-09-12T09:06:53.906398511-07:00\",\n    \"updated_at\": \"2018-09-12T09:06:54.972440207-07:00\"\n}\n```\n\n# Object Response\n```json\n{\n    \"type\": \"service\",\n    \"id\": 19,\n    \"method\": \"delete\",\n    \"status\": \"success\"\n}\n```\n\n# Main API Response\n```json\n{\n    \"name\": \"Awesome Status\",\n    \"description\": \"An awesome status page by Statping\",\n    \"footer\": \"This is my custom footer\",\n    \"domain\": \"https://demo.statping.com\",\n    \"version\": \"v0.56\",\n    \"migration_id\": 1536768413,\n    \"created_at\": \"2018-09-12T09:06:53.905374829-07:00\",\n    \"updated_at\": \"2018-09-12T09:07:01.654201225-07:00\",\n    \"database\": \"sqlite\",\n    \"started_on\": \"2018-09-12T10:43:07.760729349-07:00\",\n    \"services\": [\n        {\n            \"id\": 1,\n            \"name\": \"Google\",\n            \"domain\": \"https://google.com\",\n            \"expected\": \"\",\n            \"expected_status\": 200,\n            \"check_interval\": 10,\n            \"type\": \"http\",\n            \"method\": \"GET\",\n            \"post_data\": \"\",\n            \"port\": 0,\n            \"timeout\": 10,\n            \"order_id\": 0,\n            \"created_at\": \"2018-09-12T09:06:54.97549122-07:00\",\n            \"updated_at\": \"2018-09-12T09:06:54.975624103-07:00\",\n            \"online\": true,\n            \"latency\": 0.09080986,\n            \"24_hours_online\": 0,\n            \"avg_response\": \"\",\n            \"status_code\": 200,\n            \"last_online\": \"2018-09-12T10:44:07.931990439-07:00\",\n            \"dns_lookup_time\": 0.005543935\n        }\n    ]\n}\n```\n\n\n<div class=\"mt-5\" id=\"page_11\"><h1>Makefile</h1></div>\nHere's a simple list of Makefile commands you can run using `make`. The [Makefile](https://github.com/hunterlong/statping/blob/master/Makefile) may change often, so i'll try to keep this Wiki up-to-date.\n\n- Ubuntu `apt-get install build-essential`\n- MacOSX `sudo xcode-select -switch /Applications/Xcode.app/Contents/Developer`\n- Windows [Install Guide for GNU make utility](http://gnuwin32.sourceforge.net/packages/make.htm)\n- CentOS/RedHat `yum groupinstall \"Development Tools\"`\n\n### Commands\n```bash\nmake build                         # build the binary\nmake install\nmake run\nmake test\nmake coverage\nmake docs\n# Building Statping\nmake build-all\nmake build-alpine\nmake docker\nmake docker-run\nmake docker-dev\nmake docker-run-dev\nmake databases\nmake dep\nmake dev-deps\nmake clean\nmake compress\nmake cypress-install\nmake cypress-test\n```\n\n<div class=\"mt-5\" id=\"page_12\"><h1>Notifiers</h1></div>\n<p align=\"center\">\n<img width=\"80%\" src=\"https://s3-us-west-2.amazonaws.com/gitimgs/statupnotifiers.png\">\n</p>\n\nStatping includes multiple Notifiers to alert you when your services are offline. You can also create your own notifier and send a Push Request to this repo! Creating a custom notifier is pretty easy as long as you follow the requirements. A notifier will automatically be installed into the users Statping database, and form values will save without any hassles. 💃\n\n<p align=\"center\">\n<a href=\"https://github.com/hunterlong/statping/wiki/Notifier-Example\">Example Code</a> | <a href=\"https://github.com/hunterlong/statping/wiki/Notifier-Events\">Events</a> | <a href=\"https://github.com/hunterlong/statping/tree/master/notifiers\">View Notifiers</a><br>\n<a href=\"https://godoc.org/github.com/hunterlong/statping/core/notifier\"><img src=\"https://godoc.org/github.com/golang/gddo?status.svg\"></a>\n</p>\n\n## Notifier Requirements\n- Must have a unique `METHOD` name\n- Struct must have `*notifier.Notification` pointer in it. \n- Must create and add your notifier variable in `init()`\n- Should have a form for user to input their variables/keys. `Form: []notifier.NotificationForm`\n\n## Notifier Interface (required)\nStatping has the `Notifier` interface which you'll need to include in your notifier. Statping includes many other events/triggers for your notifier, checkout <a href=\"https://github.com/hunterlong/statping/wiki/Notifier-Events\">Notifier Events</a> to see all of them.\n```go\n// Notifier interface is required to create a new Notifier\ntype Notifier interface {\n\tOnSave() error          // OnSave is triggered when the notifier is saved\n\tSend(interface{}) error // OnSave is triggered when the notifier is saved\n\tSelect() *Notification  // Select returns the *Notification for a notifier\n}\n```\n\n### Basic Interface (required)\nInclude `OnSuccess` and `OnFailure` to receive events when a service is online or offline.\n```go\n// BasicEvents includes the most minimal events, failing and successful service triggers\ntype BasicEvents interface {\n\t// OnSuccess is triggered when a service is successful\n\tOnSuccess(*types.Service)\n\t// OnFailure is triggered when a service is failing\n\tOnFailure(*types.Service, *types.Failure)\n}\n```\n\n### Test Interface\nThe OnTest method will give the front end user the ability to test your notifier without saving, the OnTest method for your notifier run the functionality to test the user's submitted parameters and respond an error if notifier is not correctly setup.\n```go\n// Tester interface will include a function to Test users settings before saving\ntype Tester interface {\n\tOnTest() error\n}\n```\nIf your notifier includes this interface, the Test button will appear.\n\n## Notifier Struct\n```go\nvar example = &Example{&notifier.Notification{\n\tMethod: \"example\",                               // unique method name\n\tHost:   \"http://exmaplehost.com\",                // default 'host' field\n\tForm: []notifier.NotificationForm{{\n\t\tType:        \"text\",                     // text, password, number, or email\n\t\tTitle:       \"Host\",                     // The title of value in form\n\t\tPlaceholder: \"Insert your Host here.\",   // Optional placeholder in input\n\t\tDbField:     \"host\",                     // An accepted DbField value (read below)\n\t}},\n}\n```\n\n## Notifier Form\nInclude a form with your notifier so other users can save API keys, username, passwords, and other values. \n```go\n// NotificationForm contains the HTML fields for each variable/input you want the notifier to accept.\ntype NotificationForm struct {\n\tType        string `json:\"type\"`        // the html input type (text, password, email)\n\tTitle       string `json:\"title\"`       // include a title for ease of use\n\tPlaceholder string `json:\"placeholder\"` // add a placeholder for the input\n\tDbField     string `json:\"field\"`       // true variable key for input\n\tSmallText   string `json:\"small_text\"`  // insert small text under a html input\n\tRequired    bool   `json:\"required\"`    // require this input on the html form\n\tIsHidden    bool   `json:\"hidden\"`      // hide this form element from end user\n\tIsList      bool   `json:\"list\"`        // make this form element a comma separated list\n\tIsSwitch    bool   `json:\"switch\"`      // make the notifier a boolean true/false switch\n}\n```\n\n### Example Notifier Form\nThis is the Slack Notifier `Form` fields.\n```go\nForm: []notifier.NotificationForm{{\n\t\tType:        \"text\",\n\t\tTitle:       \"Incoming webhooker Url\",\n\t\tPlaceholder: \"Insert your slack webhook URL here.\",\n\t\tSmallText:   \"Incoming webhooker URL from <a href=\\\"https://api.slack.com/apps\\\" target=\\\"_blank\\\">slack Apps</a>\",\n\t\tDbField:     \"Host\",\n\t\tRequired:    true,\n\t}}\n}\n```\n\n### Accepted DbField Values\nThe `notifier.NotificationForm` has a field called `DbField` which is the column to save the value into the database. Below are the acceptable DbField string names to include in your form. \n- `host` used for a URL or API endpoint\n- `username` used for a username\n- `password` used for a password\n- `port` used for a integer port number\n- `api_key` used for some kind of API key\n- `api_secret` used for some API secret\n- `var1` used for any type of string\n- `var2` used for any type of string (extra)\n\n### Form Elements\nYou can completely custom your notifications to include a detailed form. \n- `Type` is a HTML input type for your field\n- `Title` give your input element a title\n- `Placeholder` optional field if you want a placeholder in input\n- `DbField` required field to save variable into database (read above)\n- `Placeholder` optional field for inserting small hint under the input\n\n<div class=\"mt-5\" id=\"page_13\"><h1>Notifier Events</h1></div>\nEvents are handled by added interfaces for the elements you want to monitor.\n\n## Required Notifier Interface\n```go\n// Notifier interface is required to create a new Notifier\ntype Notifier interface {\n\t// Run will trigger inside of the notifier when enabled\n\tRun() error\n\t// OnSave is triggered when the notifier is saved\n\tOnSave() error\n\t// Test will run a function inside the notifier to Test if it works\n\tTest() error\n\t// Select returns the *Notification for a notifier\n\tSelect() *Notification\n}\n```\n\n## Basic Success/Failure Interface\n```go\n// BasicEvents includes the most minimal events, failing and successful service triggers\ntype BasicEvents interface {\n\t// OnSuccess is triggered when a service is successful\n\tOnSuccess(*types.Service)\n\t// OnFailure is triggered when a service is failing\n\tOnFailure(*types.Service, *types.Failure)\n}\n```\n\n\n## Service Events\n```go\n// ServiceEvents are events for Services\ntype ServiceEvents interface {\n\tOnNewService(*types.Service)\n\tOnUpdatedService(*types.Service)\n\tOnDeletedService(*types.Service)\n}\n```\n\n## User Events\n```go\n// UserEvents are events for Users\ntype UserEvents interface {\n\tOnNewUser(*types.User)\n\tOnUpdatedUser(*types.User)\n\tOnDeletedUser(*types.User)\n}\n```\n\n## Core Events\n```go\n// CoreEvents are events for the main Core app\ntype CoreEvents interface {\n\tOnUpdatedCore(*types.Core)\n}\n```\n\n## Notifier Events\n```go\n// NotifierEvents are events for other Notifiers\ntype NotifierEvents interface {\n\tOnNewNotifier(*Notification)\n\tOnUpdatedNotifier(*Notification)\n}\n```\n\n<div class=\"mt-5\" id=\"page_14\"><h1>Notifier Example</h1></div>\nBelow is a full example of a Statping notifier which will give you a good example of how to create your own. Insert your new notifier inside the `/notifiers` folder once your ready!\n\n[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/hunterlong/statping/core/notifier)\n\n```go\npackage notifiers\n\nimport (\n\t\"errors\"\n\t\"fmt\"\n\t\"github.com/hunterlong/statping/types\"\n        \"github.com/hunterlong/statping/core/notifier\"\n\t\"time\"\n)\n\ntype Example struct {\n\t*notifier.Notification\n}\n\nvar example = &Example{&notifier.Notification{\n\tMethod:      METHOD,\n\tTitle:       \"Example\",\n\tDescription: \"Example Notifier\",\n\tAuthor:      \"Hunter Long\",\n\tAuthorUrl:   \"https://github.com/hunterlong\",\n\tDelay:       time.Duration(5 * time.Second),\n\tForm: []notifier.NotificationForm{{\n\t\tType:        \"text\",\n\t\tTitle:       \"Host\",\n\t\tPlaceholder: \"Insert your Host here.\",\n\t\tDbField:     \"host\",\n\t\tSmallText:   \"this is where you would put the host\",\n\t}, {\n\t\tType:        \"text\",\n\t\tTitle:       \"Username\",\n\t\tPlaceholder: \"Insert your Username here.\",\n\t\tDbField:     \"username\",\n\t}, {\n\t\tType:        \"password\",\n\t\tTitle:       \"Password\",\n\t\tPlaceholder: \"Insert your Password here.\",\n\t\tDbField:     \"password\",\n\t}, {\n\t\tType:        \"number\",\n\t\tTitle:       \"Port\",\n\t\tPlaceholder: \"Insert your Port here.\",\n\t\tDbField:     \"port\",\n\t}, {\n\t\tType:        \"text\",\n\t\tTitle:       \"API Key\",\n\t\tPlaceholder: \"Insert your API Key here\",\n\t\tDbField:     \"api_key\",\n\t}, {\n\t\tType:        \"text\",\n\t\tTitle:       \"API Secret\",\n\t\tPlaceholder: \"Insert your API Secret here\",\n\t\tDbField:     \"api_secret\",\n\t}, {\n\t\tType:        \"text\",\n\t\tTitle:       \"Var 1\",\n\t\tPlaceholder: \"Insert your Var1 here\",\n\t\tDbField:     \"var1\",\n\t}, {\n\t\tType:        \"text\",\n\t\tTitle:       \"Var2\",\n\t\tPlaceholder: \"Var2 goes here\",\n\t\tDbField:     \"var2\",\n\t}},\n}}\n\n// REQUIRED init() will install/load the notifier\nfunc init() {\n\tnotifier.AddNotifier(example)\n}\n\n// REQUIRED - Send is where you would put the action's of your notifier\nfunc (n *Example) Send(msg interface{}) error {\n\tmessage := msg.(string)\n\tfmt.Printf(\"i received this string: %v\\n\", message)\n\treturn nil\n}\n\n// REQUIRED\nfunc (n *Example) Select() *notifier.Notification {\n\treturn n.Notification\n}\n\n// REQUIRED\nfunc (n *Example) OnSave() error {\n\tmsg := fmt.Sprintf(\"received on save trigger\")\n\tn.AddQueue(msg)\n\treturn nil\n}\n\n// REQUIRED\nfunc (n *Example) Test() error {\n\tmsg := fmt.Sprintf(\"received a test trigger\\n\")\n\tn.AddQueue(msg)\n\treturn nil\n}\n\n// REQUIRED - BASIC EVENT\nfunc (n *Example) OnSuccess(s *types.Service) {\n\tmsg := fmt.Sprintf(\"received a count trigger for service: %v\\n\", s.Name)\n\tn.AddQueue(msg)\n}\n\n// REQUIRED - BASIC EVENT\nfunc (n *Example) OnFailure(s *types.Service, f *types.Failure) {\n\tmsg := fmt.Sprintf(\"received a failure trigger for service: %v\\n\", s.Name)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnNewService(s *types.Service) {\n\tmsg := fmt.Sprintf(\"received a new service trigger for service: %v\\n\", s.Name)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnUpdatedService(s *types.Service) {\n\tmsg := fmt.Sprintf(\"received a update service trigger for service: %v\\n\", s.Name)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnDeletedService(s *types.Service) {\n\tmsg := fmt.Sprintf(\"received a delete service trigger for service: %v\\n\", s.Name)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnNewUser(s *types.User) {\n\tmsg := fmt.Sprintf(\"received a new user trigger for user: %v\\n\", s.Username)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnUpdatedUser(s *types.User) {\n\tmsg := fmt.Sprintf(\"received a updated user trigger for user: %v\\n\", s.Username)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnDeletedUser(s *types.User) {\n\tmsg := fmt.Sprintf(\"received a deleted user trigger for user: %v\\n\", s.Username)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnUpdatedCore(s *types.Core) {\n\tmsg := fmt.Sprintf(\"received a updated core trigger for core: %v\\n\", s.Name)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnNewNotifier(s *Notification) {\n\tmsg := fmt.Sprintf(\"received a new notifier trigger for notifier: %v\\n\", s.Method)\n\tn.AddQueue(msg)\n}\n\n// OPTIONAL\nfunc (n *Example) OnUpdatedNotifier(s *Notification) {\n\tmsg := fmt.Sprintf(\"received a update notifier trigger for notifier: %v\\n\", s.Method)\n\tn.AddQueue(msg)\n}\n```\n\n<div class=\"mt-5\" id=\"page_15\"><h1>Prometheus Exporter</h1></div>\nStatping includes a prometheus exporter so you can have even more monitoring power with your services. The prometheus exporter can be seen on `/metrics`, simply create another exporter in your prometheus config. Use your Statping API Secret for the Authorization Bearer header, the `/metrics` URL is dedicated for Prometheus and requires the correct API Secret has `Authorization` header.\n\n# Grafana Dashboard\nStatping has a [Grafana Dashboard](https://grafana.com/dashboards/6950) that you can quickly implement if you've added your Statping service to Prometheus. Import Dashboard ID: `6950` into your Grafana dashboard and watch the metrics come in!\n\n<p align=\"center\"><img width=\"80%\" src=\"https://img.cjx.io/statupgrafana.png\"></p>\n\n## Basic Prometheus Exporter\nIf you have Statping and the Prometheus server in the same Docker network, you can use the yaml config below.\n```yaml\nscrape_configs:\n  - job_name: 'statup'\n    scrape_interval: 30s\n    bearer_token: 'SECRET API KEY HERE'\n    static_configs:\n      - targets: ['statup:8080']\n```\n\n## Remote URL Prometheus Exporter\nThis exporter yaml below has `scheme: https`, which you can remove if you arn't using HTTPS.\n```yaml\nscrape_configs:\n  - job_name: 'statup'\n    scheme: https\n    scrape_interval: 30s\n    bearer_token: 'SECRET API KEY HERE'\n    static_configs:\n      - targets: ['status.mydomain.com']\n```\n\n### `/metrics` Output\n```\nstatup_total_failures 206\nstatup_total_services 4\nstatup_service_failures{id=\"1\" name=\"Google\"} 0\nstatup_service_latency{id=\"1\" name=\"Google\"} 12\nstatup_service_online{id=\"1\" name=\"Google\"} 1\nstatup_service_status_code{id=\"1\" name=\"Google\"} 200\nstatup_service_response_length{id=\"1\" name=\"Google\"} 10777\nstatup_service_failures{id=\"2\" name=\"Statping.com\"} 0\nstatup_service_latency{id=\"2\" name=\"Statping.com\"} 3\nstatup_service_online{id=\"2\" name=\"Statping.com\"} 1\nstatup_service_status_code{id=\"2\" name=\"Statping.com\"} 200\nstatup_service_response_length{id=\"2\" name=\"Statping.com\"} 2\n```\n\n<div class=\"mt-5\" id=\"page_16\"><h1>SSL</h1></div>\nYou can run Statping with a valid certificate by including 2 files in the root directory. Although, I personally recommend using NGINX or Apache to serve the SSL and then have the webserver direct traffic to the Statping instance. This guide will show you how to implement SSL onto your Statping server with multiple options.\n\n## SSL Certificate with Statping\nTo run the Statping HTTP server in SSL mode, you must include 2 files in the root directory of your Statping application. The 2 files you must include are:\n- `server.crt` SSL Certificate File\n- `server.key` SSL Certificate Key File\n\nThe filenames and extensions must match the exact naming above. If these 2 files are found, Statping will automatically start the HTTP server in SSL mode using your certificates. You can also generate your own SSL certificates, but you will receive a \"ERR_CERT_AUTHORITY_INVALID\" error. To generate your own, follow the commands below:\n\n```shell\nopenssl req -new -sha256 -key server.key -out server.csr\nopenssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650\n```\nThis will generate a self signed certificate that you can use for your Statup instance. I recommend using a web server to do SSL termination for your server though.\n\n## Choose a Web Server or Environment\n\n**Choose the environment running the Statping instance.**\n- [Docker](#docker)\n- [NGINX](#nginx)\n- [Apache](#apache)\n\n## Docker\nDocker might be the easiest way to get up and running with a SSL certificate. Below is a `docker-compose.yml` file that will run NGINX, LetEncrypt, and Statping.\n\n1. Point your domain or subdomain to the IP address of the Docker server. This would be done on CloudFlare, Route53, or some other DNS provider.\n\n2. Replace the `docker-compose.yml` contents:\n- `MY.DOMAIN.COM` with the domain you want to use\n- `MY@EMAIL.COM` with your email address\n\n3. Run the docker container by running command `docker-compose up -d`. Give a little bit of time for LetEncrypt to automatically generate your SSL certificate.\n\n###### `docker-compose.yml`\n```yaml\nversion: '2.3'\nservices:\n  nginx:\n    container_name: nginx\n    image: jwilder/nginx-proxy\n    ports:\n      - 0.0.0.0:80:80\n      - 0.0.0.0:443:443\n    labels:\n      - \"com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy\"\n    networks:\n      - internet\n    restart: always\n    volumes:\n      - /var/run/docker.sock:/tmp/docker.sock:ro\n      - ./statping/nginx/certs:/etc/nginx/certs:ro\n      - ./statping/nginx/vhost:/etc/nginx/vhost.d\n      - ./statping/nginx/html:/usr/share/nginx/html:ro\n      - ./statping/nginx/dhparam:/etc/nginx/dhparam\n    environment:\n      DEFAULT_HOST: MY.DOMAIN.COM\n\n  letsencrypt:\n    container_name: letsencrypt\n    image: jrcs/letsencrypt-nginx-proxy-companion\n    networks:\n      - internet\n    restart: always\n    volumes:\n      - /var/run/docker.sock:/var/run/docker.sock:ro\n      - ./statping/nginx/certs:/etc/nginx/certs\n      - ./statping/nginx/vhost:/etc/nginx/vhost.d\n      - ./statping/nginx/html:/usr/share/nginx/html\n      - ./statping/nginx/dhparam:/etc/nginx/dhparam\n\n  statping:\n    container_name: statping\n    image: hunterlong/statping:latest\n    restart: always\n    networks:\n      - internet\n    depends_on:\n      - nginx\n    volumes:\n      - ./statping/app:/app\n    environment:\n      VIRTUAL_HOST: MY.DOMAIN.COM\n      VIRTUAL_PORT: 8080\n      LETSENCRYPT_HOST: MY.DOMAIN.COM\n      LETSENCRYPT_EMAIL: MY@EMAIL.COM\n\nnetworks:\n  internet:\n    driver: bridge\n```\n\n## NGINX\nIf you already have a NGINX web server running, you just have to add a proxy pass and your SSL certs to the nginx config or as a vhost. By default Statping runs on port 8080, you can change this port by starting server with `statping -ip 127.0.0.1 -port 9595`.\n\n- Replace `/my/absolute/directory/for/cert/server.crt` with SSL certificate file.\n- Replace `/my/absolute/directory/for/key/server.key` with SSL key file.\n- Run `service nginx restart` and try out https on your domain.\n\n##### Tutorials\n- [NGINX Guide](https://docs.nginx.com/nginx/admin-guide/security-controls/terminating-ssl-http/)\n- [How To Set Up Nginx Load Balancing with SSL Termination](https://www.digitalocean.com/community/tutorials/how-to-set-up-nginx-load-balancing-with-ssl-termination)\n\n###### `/etc/nginx/nginx.conf`\n```\n#user  nobody;\nworker_processes  1;\nevents {\n    worker_connections  1024;\n}\nhttp {\n    include            mime.types;\n    default_type       application/octet-stream;\n    send_timeout       1800;\n    sendfile           on;\n    keepalive_timeout  6500;\n    server {\n        listen       80;\n        server_name  localhost;\n        location / {\n          proxy_pass          http://localhost:8080;\n          proxy_set_header    Host             $host;\n          proxy_set_header    X-Real-IP        $remote_addr;\n          proxy_set_header    X-Forwarded-For  $proxy_add_x_forwarded_for;\n          proxy_set_header    X-Client-Verify  SUCCESS;\n          proxy_set_header    X-Client-DN      $ssl_client_s_dn;\n          proxy_set_header    X-SSL-Subject    $ssl_client_s_dn;\n          proxy_set_header    X-SSL-Issuer     $ssl_client_i_dn;\n          proxy_read_timeout 1800;\n          proxy_connect_timeout 1800;\n        }\n    }\n    # HTTPS server\n    \n    server {\n        listen       443;\n        server_name  localhost;\n    \n        ssl                  on;\n        ssl_certificate      /my/absolute/directory/for/cert/server.crt;\n        ssl_certificate_key  /my/absolute/directory/for/key/server.key;\n        ssl_session_timeout  5m;\n    \n        ssl_protocols  SSLv2 SSLv3 TLSv1;\n        ssl_ciphers  ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;\n        ssl_prefer_server_ciphers   on;\n    \n        location / {\n          proxy_pass          http://localhost:8080;\n          proxy_set_header    Host             $host;\n          proxy_set_header    X-Real-IP        $remote_addr;\n          proxy_set_header    X-Forwarded-For  $proxy_add_x_forwarded_for;\n          proxy_set_header    X-Client-Verify  SUCCESS;\n          proxy_set_header    X-Client-DN      $ssl_client_s_dn;\n          proxy_set_header    X-SSL-Subject    $ssl_client_s_dn;\n          proxy_set_header    X-SSL-Issuer     $ssl_client_i_dn;\n          proxy_read_timeout 1800;\n          proxy_connect_timeout 1800;\n        }\n    }\n}\n```\n\n## Apache\n\n<div class=\"mt-5\" id=\"page_17\"><h1>Config with .env File</h1></div>\nIt may be useful to load your environment using a `.env` file in the root directory of your Statping server. The .env file will be automatically loaded on startup and will overwrite all values you have in config.yml.\n\nIf you have the `DB_CONN` environment variable set Statping will bypass all values in config.yml and will require you to have the other DB_* variables in place. You can pass in these environment variables without requiring a .env file.\n\n## `.env` File\n```bash\nDB_CONN=postgres\nDB_HOST=0.0.0.0\nDB_PORT=5432\nDB_USER=root\nDB_PASS=password123\nDB_DATABASE=root\n\nNAME=Demo\nDESCRIPTION=This is an awesome page\nDOMAIN=https://domain.com\nADMIN_USER=admin\nADMIN_PASS=admin\nADMIN_EMAIL=info@admin.com\nUSE_CDN=true\nPOSTGRES_SSL=false # enable ssl_mode for postgres (true/false)\n\nIS_DOCKER=false\nIS_AWS=false\nSASS=/usr/local/bin/sass\nCMD_FILE=/bin/bash\n```\nThis .env file will include additional variables in the future, subscribe to this repo to keep up-to-date with changes and updates. \n\n<div class=\"mt-5\" id=\"page_18\"><h1>Static Export</h1></div>\nIf you want to use Statping as a CLI application without running a server, you can export your status page to a static HTML.\nThis export tool is very useful for people who want to export their HTML and upload/commit it to Github Pages or an FTP server.\n```dash\nstatup export\n```\n###### Creates `index.html` in the current directory with CDN asset URL's. 💃 \n\n<div class=\"mt-5\" id=\"page_19\"><h1>Statping Plugins</h1></div>\nSince Statping is built in Go Language we can use the [Go Plugin](https://golang.org/pkg/plugin/) feature to create dynamic plugins that run on load. Statping has an event anytime anything happens, you can create your own plugins and do any type of function. To implement your own ideas into Statping, use the plugin using the [statup/plugin](https://github.com/hunterlong/statping/blob/master/plugin/main.go) package.\n```\ngo get github.com/hunterlong/statping/plugin\n```\n\n## Example Plugin\nStart off with the [Example Statping Plugin](https://github.com/hunterlong/statping_plugin) that includes all the interfaces and some custom options for you to expand on. You can include any type of function in your own plugin!\n\n<p align=\"center\">\n<img width=\"95%\" src=\"https://img.cjx.io/statuppluginrun.gif\">\n</p>\n\n## Building Plugins\nPlugins don't need a push request and they can be private! You'll need to compile your plugin to the Golang `.so` binary format. Once you've built your plugin, insert it into the `plugins` folder in your Statping directory and reboot the application. Clone the [Example Statping Plugin](https://github.com/hunterlong/statping_plugin) repo and try to build it yourself!\n\n#### Build Requirements\n- You must have `main.go`\n- You must create the Plugin variable on `init()`\n\n```bash\ngit clone https://github.com/hunterlong/statping_plugin\ncd statup-plugin\ngo build -buildmode=plugin -o example.so\n```\n###### Insert `example.so` into the `plugins` directory and reload Statping\n\n## Testing Statping Plugins\nStatping includes a couple tools to help you on your Plugin journey, you can use `statup test plugins` command to test all plugins in your `/plugins` folder. This test will attempt to parse your plugin details, and then it will send events for your plugin to be fired.\n```\nstatup test plugins\n```\n<p align=\"center\">\n<img width=\"85%\" src=\"https://img.cjx.io/statupplugintester.gif\">\n</p>\n\nYour plugin should be able to parse and receive events before distributing it. The test tools creates a temporary database (SQLite) that your plugin can interact with. Statping uses [upper.io/db.v3](https://upper.io/db.v3) for database interactions. The database is passed to your plugin `OnLoad(db sqlbuilder.Database)`, so you can use the `db` variable passed here.\n\n## Statping Plugin Interface\nPlease remember Golang plugin's are very new and Statping plugin package may change and 'could' brake your plugin. Checkout the [statup/plugin package](https://github.com/hunterlong/statping/blob/master/plugin/main.go) to see the most current interfaces.\n```go\ntype PluginActions interface {\n\tGetInfo() Info\n\tGetForm() string\n\tSetInfo(map[string]interface{}) Info\n\tRoutes() []Routing\n\tOnSave(map[string]interface{})\n\tOnFailure(map[string]interface{})\n\tOnSuccess(map[string]interface{})\n\tOnSettingsSaved(map[string]interface{})\n\tOnNewUser(map[string]interface{})\n\tOnNewService(map[string]interface{})\n\tOnUpdatedService(map[string]interface{})\n\tOnDeletedService(map[string]interface{})\n\tOnInstall(map[string]interface{})\n\tOnUninstall(map[string]interface{})\n\tOnBeforeRequest(map[string]interface{})\n\tOnAfterRequest(map[string]interface{})\n\tOnShutdown()\n\tOnLoad(sqlbuilder.Database)\n}\n```\n\n## Event Parameters\nAll event interfaces for the Statping Plugin will return a `map[string]interface{}` type, this is because the plugin package will most likely update and change in the future, but using this type will allow your plugin to continue even after updates.\n\n## Example of an Event\nKnowing what happens during an event is important for your plugin. For example, lets have an event that echo something when a service has a Failure status being issued. Checkout some example below to see how this golang plugin action works. \n\n```go\nfunc (p pkg) OnSuccess(data map[string]interface{}) {\n    fmt.Println(\"Statping Example Plugin received a successful service hit! \")\n    fmt.Println(\"Name:    \", data[\"Name\"])\n    fmt.Println(\"Domain:  \", data[\"Domain\"])\n    fmt.Println(\"Method:  \", data[\"Method\"])\n    fmt.Println(\"Latency: \", data[\"Latency\"])\n}\n```\n###### OnSuccess is fired every time a service has check it be online\n\n```go\nfunc OnFailure(service map[string]interface{}) {\n    fmt.Println(\"oh no! an event is failing right now! do something!\")\n    fmt.Println(service)\n}\n```\n###### OnFailure is fired every time a service is failing\n\n```go\nfunc (p pkg) OnLoad(db sqlbuilder.Database) {\n    fmt.Println(\"=============================================================\")\n    fmt.Printf(\"  Statping Example Plugin Loaded using %v database\\n\", db.Name())\n    fmt.Println(\"=============================================================\")\n}\n```\n###### OnLoad is fired after plugin is loaded into the environment\n\n\n## Interacting with Database\nThe Example Statping Plugin includes a variable `Database` that will allow you to interact with the Statping database. Checkout [database.go](https://github.com/hunterlong/statping_plugin/blob/master/database.go) to see a full example of Create, Read, Update and then Deleting a custom Communication entry into the database.\n```go\n// Insert a new communication into database\n// once inserted, return the Communication\nfunc (c *Communication) Create() *Communication {\n\tuuid, err := CommunicationTable().Insert(c)\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\tc.Id = uuid.(int64)\n\treturn c\n}\n```\n\n## Custom HTTP Routes\nPlugin's can include their own HTTP route to accept GET/POST requests. Route are loaded after Statping loads all of it's Routes. Checkout [routes.go](https://github.com/hunterlong/statping_plugin/blob/master/routes.go) on the example plugin to see a full example of how to use it.\n```go\n// You must have a Routes() method in your plugin\nfunc (p *pkg) Routes() []plugin.Routing {\n\treturn []plugin.Routing{{\n\t\tURL:     \"hello\",\n\t\tMethod:  \"GET\",\n\t\tHandler: CustomInfoHandler,\n\t}}\n}\n\n// This is the HTTP handler for the '/hello' URL created above\nfunc CustomInfoHandler(w http.ResponseWriter, r *http.Request) {\n\tw.WriteHeader(http.StatusOK)\n\tfmt.Fprintln(w, \"Oh Wow!!! This is cool...\")\n}\n```\n\n\n## Plugin To-Do List\n- [ ] Ability to includes assets like jpg, json, etc\n\n<div class=\"mt-5\" id=\"page_20\"><h1>Statuper</h1></div>\nStatping includes a simple to use installation shell script that will help you install locally, Docker, and even onto a AWS EC2 instance.\n\n<p align=\"center\">\n<img width=\"90%\" src=\"https://img.cjx.io/statuper1.png\">\n</p>\n\n## Installation\n```bash\ncurl -O https://assets.statup.io/statuper && chmod +x statuper\n```\n\n## Usage\n- `statuper`\n\n<div class=\"mt-5\" id=\"page_21\"><h1>Build and Test</h1></div>\nBuilding from the Go Language source code is pretty easy if you already have Go installed. Clone this repo and `cd` into it. \n\n### Git n' Go Get\n```bash\ngit clone https://github.com/hunterlong/statping.git\ncd statup\ngo get -v\n```\n\n### Install go.rice\nStatping uses go.rice to compile HTML, JS, and CSS files into it's single binary output.\n```\ngo get github.com/GeertJohan/go.rice\ngo get github.com/GeertJohan/go.rice/rice\n```\n\n### Build Statping Binary\nStatping uses go.rice to compile HTML, JS, and CSS files into it's single binary output.\n```\nrice embed-go\ngo build -o statup .\n./statup version\n```\n\n### Test Coverage\nYou can also test Statio on your localhost, but it does require a MySQL, and Postgres server to be accessible since testing does create/drop tables for multiple databases. \n```\ngo test -v\n```\n\n<div class=\"mt-5\" id=\"page_22\"><h1>Contributing</h1></div>\nHave a feature you want to implement into Statping!? Awesome! Follow this guide to see how you can test, compile and build Statping for production use. I recommend you use `make` with this process, it will save you time and it will auto include many customized parameters to get everything working correctly.\n\n# Dependencies\nStatping has a couple of required dependencies when testing and compiling the binary. The [Makefile](https://github.com/hunterlong/statping/blob/master/Makefile) will make these tasks a lot easier. Take a look at the Makefile to see what commands are ran. Run the command below to get setup right away.\n```bash\nmake dev-deps\n```\nList of requirements for compiling assets, building binary, and testing.\n- [Go Language](https://golang.org/) (currently `1.10.3`)\n- [Docker](https://docs.docker.com/)\n- [SASS](https://sass-lang.com/install)\n- [Cypress](https://www.cypress.io/) (only used for UI testing, `make cypress-install`)\n\n# Compiling Assets\nThis Golang project uses [rice](https://github.com/GeertJohan/go.rice) to compile static assets into a single file. The file `source/rice-box.go` is never committed to the Github repo, it is automatically created on build. Statping also requires `sass` to be installed on your local OS. To compile all the static assets run the command below:\n\n```bash\nmake compile\n```\nAfter this is complete, you'll notice the `source/rice-box.go` file has been generated. You can now continue to build, and test.\n\n# Testing\nStatping includes multiple ways to Test the application, you can run the `make` command, or the normal `go test` command. To see the full experience of your updates, you can even run Cypress tests which is in the `.dev/test` folder.\n\nStatping will run all tests in `cmd` folder on MySQL, Postgres, and SQLite databases. You can run `make databases` to automatically create MySQL and Postgres with Docker.\n\n###### Go Unit Testing:\n```bash\nmake test\n```\n\n###### Cypress UI Testing:\n```bash\nmake cypress-test\n```\n\n###### Test Everything:\n```bash\nmake test-all\n```\n\n# Build\nStatping will build on all operating systems except Windows 32-bit. I personally use [xgo](https://github.com/karalabe/xgo) to cross-compile on multiple systems using Docker. Follow the commands below to build on your local system.\n\n###### Build for local operating system:\n```bash\nmake build\n```\n\n# Compile for Production\nOnce you've tested and built locally, you can compile Statping for all available operating systems using the command below. This command will require you to have Docker.\n\n```bash\nmake build-all\n```\n\n# What Now\nEverything tested, compiled and worked out!? Awesome! 💃 You can now commit your changes, and submit a Pull Request with the features/bugs you added or removed.\n\n\n\n\n\n<div class=\"mt-5\" id=\"page_23\"><h1>PGP Signature</h1></div>\nYou can check if the Statping binary you downloaded is authentic by running a few commands.\n\n### Steps to Authenticate\n1. Download the Statping `tar.gz` file from [Latest Releases](https://github.com/hunterlong/statping/releases/latest) and extract the `statping` binary and the `statup.asc` file.\n2. Run command: `gpg --verify statping.asc`\n3. You should see `Good signature from \"Hunter Long <info@statping.com>\" [ultimate]`.\n\n# Statping Public Key\n- [https://statping.com/statping.gpg](https://statping.com/statping.gpg)\n\nYou can also download the key with the command below:\n```\nwget https://statping.com/statping.gpg\n```\n\n###### `statping.gpg`\n```\n-----BEGIN PGP PUBLIC KEY BLOCK-----\n\nmQINBFwGUYIBEADNsDY4aUOx8EoZuTRFPtjuadJzFRyKtHhw/tLlAnoFACanZPIT\nNZoRYvRR5v6lMDXdxsteDbJEOhZ1WDiKIr4OyMahPsyyH6ULzSBKgePUswa0sDef\nUnXYzPFQCzqQyQQFbp9AYfDP7dW6dTL9I6qU2NqlJvjxJiiZTAq87SmsLqHiASnI\n+ottnQuu6vJQBJz2PFIuaS1c3js/+HBbth9GK5B9YN1BIIyZoFmWKVU9HnJf+aM3\nUs6OLjjwYwWzQH38ZV84IjVXyiP9PQVhlCXeHK7XdhPZvnSP1m5Wszj/jowwY6Mz\nLgLotfL540X7yOJ7hJTFYLFBOtJdJr/3Ov8SH4HXdPFPVG+UqxsmtmPqUQ9iAxAE\njRFfkAxBvH5Szf2WZdaLnlrrOcOKJIIjZgHqalquBTAhlh5ul0lUVSSPxetwIBlW\n60L41k94NJFGDt8xOJ+122mLcywmQ1CzhDfeIKlxl6JDiVHjoRqlQQrqIoNZMV85\nrzGfrmbuwv1MXGBJoiNy3330ujOBmhQ9dQVwKpxhBKdjnAgIGM9szbUYxIkGgM1O\nU4b1WF3AF/9JOpKJ0LewslpM3BFFYnemGsHXAv3TBPqKidNdwMAiBOtNykGoXF6i\n0D6jOW/IB1da0gUA+kr5JdAOwIG7JXKhur2MO7Ncid59DL2N8RePRWj+jwARAQAB\ntB9IdW50ZXIgTG9uZyA8aW5mb0BzdGF0cGluZy5jb20+iQJOBBMBCAA4FiEEt21h\n+qbbdZRm6D2ZZLnGquLVUngFAlwGUYICGwMFCwkIBwIGFQoJCAsCBBYCAwECHgEC\nF4AACgkQZLnGquLVUnizwA//c7vmwTMq/8LYlbo37WM2kDE9AKIrz6VSMq4RhGbC\nLikH0X0epa+if79n9BZrVU/Af3aKTn7vu2J4XrvzcdCXtcsR0YmCWML2Y6OSFmhX\nw3o6woiFcp+SUWdcM/kithRun+j9sKV4akdgkdBQUdh/RMVln+radz1c6G59iTdh\nS+Ip3ObO7Gn5VnrLwxix+W9Jhg8YhDgDGEDt8e1yvjuMRY+WhjHFlwEMoE0kvQL8\nQvQH2dGD3dExWAuIL7+0xC0ZGU0PR8vRrq1ukdIsWlDY+42vvhcyPZKFFDTM/QLF\nFcCNiPSGhiK/NQq67xnRMFdh0fnqbydWj2atMpacIrheEkOt8db2/UMyDOwlIxgy\nKOG8x+yNKiG9LyvW4axRLctN608/+TbvtFo5TVOFJYxJQp4b5uz7LgJAJw7PBvfC\nbqx64BH8WGzgyGcAl9unQEtpDuxXoKvP2kbsS7hjvhK0gJgW9llpV4sRJJGApTBc\nWtbcS9DBGs3k1aZdA72bxnayD32syVz7czl4+tkRsbQZ4VgJh1yrHIDsdWQXFnYu\nEQJfCgX5HvvC13MpDUth0NWCFtWQirY3EFbIgSuhB/D5iXA+Dt1Dq5c1u7wQlUVi\nLQCU++oMGrlU3gZrnov5lnBGCEjn0O9bKQm8zmLdEcENFxUZvfPjOIY64YprZxD9\nBv65Ag0EXAZRggEQAMmjHmnvH8SvNJhku/oI96dFKen3bg9xdaFUD1vAuNglCalH\nwgXcCZd0RdobYNG46cXTzTQadtHS4hi/UBJ+oy5ZUpIRglW12eTYtqM2G11VbcQi\nj6rLITP9NIP+G1xBICSYK4UwmH55BolMEQ/1ZX0a9rESM9stDNglheCCudbMGR/1\nZYnufdEsh0yPwyC/1upZeu8LPWK62pt9mE/gccx77QTeDi5OJcRf1fPbUTCm3vSS\nwPPV2AGANodIhostjDymt5vh0tGwc7oUZZLnVdErfuctv7yMgZdiCpYu0jFy1NYf\nJgOpZasrcK7/1ozGzsfAo/sSU4kIkMwuWGgqfx5kGRK2CgU4T0i7oI6DMpOX9ZS8\ns3+oCWu83X0ofvm5R2CbjiUj2gR6JOhBQbJpCeTkLe+SFcUpnyrr7lG8B8QZHm5N\nnBi05V/s63RE3g/6WpR/fWuh+uswe01uqlSx9deW7jT49BL/MdSxwjfwLBLz/hLM\n0ld385XAd9bqMjUtp0XhZX2YORx3f/aKY7PYA62baGibb5RdPRw6viEAWU20eb+8\nX9Pa7hGmwUeal5lka4SD/TGl7wdY+g4oYP+jtKinH/ZftWA5wHTe3jWT5bdWrT2d\ne+0qA0SBkmKIDLpktvtTa19w2nfwBIwJ6fN36ZjYpOn/stxR7aRtnhSqvzxbABEB\nAAGJAjYEGAEIACAWIQS3bWH6ptt1lGboPZlkucaq4tVSeAUCXAZRggIbDAAKCRBk\nucaq4tVSeGWmD/9Pg1x6s98zdZCQa2apmUnuoQAQA9Gf2RBBuglCDGsY67wbvdHZ\n9wdFRs2QEhl2O3oFmidxthBOBRl9z62nXliLwNn1Lcy/yDfaB8wH6gMm4jn2N/z9\nvQXnyIzg8m4PItZ1p5mnY3qH5lpGF8r9Gb7tzK10rqulM2XTDojZOevlEGI6LGw8\nFjccXtNquqGZwxzytmKF3T7UBmpmt2qock8N5iJn987m6WeYmbFNc0ii0guHfdtO\nzQcItz2ngCdyvfgQPwCAoAv72ysSGhz5KZgAXRrEdcqj6Jw3ivoEUKq1aUrYncXQ\n3zC3ED6AjWOGRzjvTZzj22IVacUZ0gqx0x/oldXLOhMB9u6nFXHKj1n9nc0XHMNi\nLp9EuvQgcNLjFZGE9sxh25u9V+OhItfT/aarYEu/Xq0IkUUcdz4GehXth1/Cq1wH\nlSUie4nCs7I7OWhqMNClqP7ywElDXsQ66MCgvf01Dh64YUVjJNnyyK0QiYlCx/JQ\nZ85hNLtVXZfYqC5BRZlVFp8I8Rs2Qos9YEgn2M22+Rj+RIeD74LZFB7Q4myRvTMB\n/P466dFI83KYhwvjBYOP3jPTrV7Ky8poEGifQp2mM294CFIPS7z0z7a8+yMzcsRP\nOluFxewsEO0QNDrfFb+0gnjYlnGqOFcZjUMXbDdY5oLSPtXohynuTK1qyQ==\n=Xn0G\n-----END PGP PUBLIC KEY BLOCK-----\n```\n\n<div class=\"mt-5\" id=\"page_24\"><h1>Testing</h1></div>\nIf you want to test your updates with the current golang testing units, you can follow the guide below to run a full test process. Each test for Statping will run in MySQL, Postgres, and SQlite to make sure all database types work correctly.\n\n## Create Docker Databases\nThe easiest way to run the tests on all 3 databases is by starting temporary databases servers with Docker. Docker is available for Linux, Mac and Windows. You can download/install it by going to the [Docker Installation](https://docs.docker.com/install/) site.\n\n```go\ndocker run -it -d \\\n   -p 3306:3306 \\\n   -env MYSQL_ROOT_PASSWORD=password123 \\\n   -env MYSQL_DATABASE=root mysql\n```\n\n```go\ndocker run -it -d \\\n   -p 5432:5432 \\\n   -env POSTGRES_PASSWORD=password123 \\\n   -env POSTGRES_USER=root \\\n   -env POSTGRES_DB=root postgres\n```\n\nOnce you have MySQL and Postgres running, you can begin the testing. SQLite database will automatically create a `statup.db` file and will delete after testing.\n\n## Run Tests\nInsert the database environment variables to auto connect the the databases and run the normal test command: `go test -v`. You'll see a verbose output of each test. If all tests pass, make a push request! 💃\n```go\nDB_DATABASE=root \\\n   DB_USER=root \\\n   DB_PASS=password123 \\\n   DB_HOST=localhost \\\n   go test -v\n```\n\n<div class=\"mt-5\" id=\"page_25\"><h1>Deployment</h1></div>\nStatping is a pretty cool server for monitoring your services. The way we deploy might be a little cooler though. Statping is using the most bleeding edge technology to release updates and distribute binary files automatically.\n\n1. Source code commits get pushed to Github\n2. [Rice](https://github.com/GeertJohan/go.rice) will compile all the static assets into 1 file (rice-box.go in source)\n3. SASS will generate  a compiled version of the CSS. \n4. Statping Help page is generated by cloning the Wiki repo using `go generate`.\n5. Travis-CI tests the Golang application.\n6. Travis-CI tests the Statping API using [Postman](https://github.com/hunterlong/statping/blob/master/source/tmpl/postman.json).\n7. If all tests are successful, Travis-CI will compile the binaries using [xgo](https://github.com/karalabe/xgo).\n8. Binaries are code signed using the official [PGP key](https://github.com/hunterlong/statping/wiki/PGP-Signature) and compressed.\n9. [Docker](https://cloud.docker.com/repository/docker/hunterlong/statping/builds) receives a trigger to build for the `latest` tag.\n10. Travis-CI uploads the [latest release](https://github.com/hunterlong/statping/releases) as a tagged version on Github.\n11. Travis-CI updates the [homebrew-statping](https://github.com/hunterlong/homebrew-statping) repo with the latest version.\n\nAnd that's it! Statping is ready to be shipped and installed.\n\n")
```
CompiledWiki contains all of the Statping Wiki pages from the Github Wiki repo.

#### func  Assets

```go
func Assets()
```
Assets will load the Rice boxes containing the CSS, SCSS, JS, and HTML files.

#### func  CompileSASS

```go
func CompileSASS(folder string) error
```
CompileSASS will attempt to compile the SASS files into CSS

#### func  CopyAllToPublic

```go
func CopyAllToPublic(box *rice.Box, folder string) error
```
CopyAllToPublic will copy all the files in a rice box into a local folder

#### func  CopyToPublic

```go
func CopyToPublic(box *rice.Box, folder, file string) error
```
CopyToPublic will create a file from a rice Box to the '/assets' directory

#### func  CreateAllAssets

```go
func CreateAllAssets(folder string) error
```
CreateAllAssets will dump HTML, CSS, SCSS, and JS assets into the '/assets'
directory

#### func  DeleteAllAssets

```go
func DeleteAllAssets(folder string) error
```
DeleteAllAssets will delete the '/assets' folder

#### func  HelpMarkdown

```go
func HelpMarkdown() string
```
HelpMarkdown will return the Markdown of help.md into HTML

#### func  MakePublicFolder

```go
func MakePublicFolder(folder string) error
```
MakePublicFolder will create a new folder

#### func  OpenAsset

```go
func OpenAsset(folder, file string) string
```
OpenAsset returns a file's contents as a string

#### func  SaveAsset

```go
func SaveAsset(data []byte, folder, file string) error
```
SaveAsset will save an asset to the '/assets/' folder.

#### func  UsingAssets

```go
func UsingAssets(folder string) bool
```
UsingAssets returns true if the '/assets' folder is found in the directory
# types
--
    import "github.com/hunterlong/statping/types"

Package types contains all of the structs for objects in Statping including
services, hits, failures, Core, and others.

More info on: https://github.com/hunterlong/statping

## Usage

```go
const (
	TIME_NANO     = "2006-01-02T15:04:05Z"
	TIME          = "2006-01-02 15:04:05"
	POSTGRES_TIME = "2006-01-02 15:04"
	CHART_TIME    = "2006-01-02T15:04:05.999999-07:00"
	TIME_DAY      = "2006-01-02"
)
```

```go
var (
	NOW = func() time.Time { return time.Now() }()
)
```

#### type AllNotifiers

```go
type AllNotifiers interface{}
```

AllNotifiers contains all the Notifiers loaded

#### type Asseter

```go
type Asseter interface {
	Asset(string) ([]byte, error)
}
```


#### type Checkin

```go
type Checkin struct {
	Id          int64              `gorm:"primary_key;column:id" json:"id"`
	ServiceId   int64              `gorm:"index;column:service" json:"service_id"`
	Name        string             `gorm:"column:name" json:"name"`
	Interval    int64              `gorm:"column:check_interval" json:"interval"`
	GracePeriod int64              `gorm:"column:grace_period"  json:"grace"`
	ApiKey      string             `gorm:"column:api_key"  json:"api_key"`
	CreatedAt   time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Running     chan bool          `gorm:"-" json:"-"`
	Failing     bool               `gorm:"-" json:"failing"`
	LastHit     time.Time          `gorm:"-" json:"last_hit"`
	Hits        []*CheckinHit      `gorm:"-" json:"hits"`
	Failures    []FailureInterface `gorm:"-" json:"failures"`
}
```

Checkin struct will allow an application to send a recurring HTTP GET to confirm
a service is online

#### func (*Checkin) BeforeCreate

```go
func (c *Checkin) BeforeCreate() (err error)
```
BeforeCreate for Checkin will set CreatedAt to UTC

#### func (*Checkin) Close

```go
func (s *Checkin) Close()
```
Close will stop the checkin routine

#### func (*Checkin) IsRunning

```go
func (s *Checkin) IsRunning() bool
```
IsRunning returns true if the checkin go routine is running

#### func (*Checkin) Start

```go
func (s *Checkin) Start()
```
Start will create a channel for the checkin checking go routine

#### type CheckinHit

```go
type CheckinHit struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Checkin   int64     `gorm:"index;column:checkin" json:"-"`
	From      string    `gorm:"column:from_location" json:"from"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
```

CheckinHit is a successful response from a Checkin

#### func (*CheckinHit) BeforeCreate

```go
func (c *CheckinHit) BeforeCreate() (err error)
```
BeforeCreate for checkinHit will set CreatedAt to UTC

#### type CheckinInterface

```go
type CheckinInterface interface {
	Select() *Checkin
}
```


#### type Core

```go
type Core struct {
	Name          string             `gorm:"not null;column:name" json:"name"`
	Description   string             `gorm:"not null;column:description" json:"description,omitempty"`
	Config        string             `gorm:"column:config" json:"-"`
	ApiKey        string             `gorm:"column:api_key" json:"-"`
	ApiSecret     string             `gorm:"column:api_secret" json:"-"`
	Style         string             `gorm:"not null;column:style" json:"style,omitempty"`
	Footer        NullString         `gorm:"column:footer" json:"footer"`
	Domain        string             `gorm:"not null;column:domain" json:"domain"`
	Version       string             `gorm:"column:version" json:"version"`
	MigrationId   int64              `gorm:"column:migration_id" json:"migration_id,omitempty"`
	UseCdn        NullBool           `gorm:"column:use_cdn;default:false" json:"using_cdn,omitempty"`
	Timezone      float32            `gorm:"column:timezone;default:-8.0" json:"timezone,omitempty"`
	CreatedAt     time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time          `gorm:"column:updated_at" json:"updated_at"`
	DbConnection  string             `gorm:"-" json:"database"`
	Started       time.Time          `gorm:"-" json:"started_on"`
	Services      []ServiceInterface `gorm:"-" json:"-"`
	Plugins       []*Info            `gorm:"-" json:"-"`
	Repos         []PluginJSON       `gorm:"-" json:"-"`
	AllPlugins    []PluginActions    `gorm:"-" json:"-"`
	Notifications []AllNotifiers     `gorm:"-" json:"-"`
}
```

Core struct contains all the required fields for Statping. All application
settings will be saved into 1 row in the 'core' table. You can use the
core.CoreApp global variable to interact with the attributes to the application,
such as services.

#### type Databaser

```go
type Databaser interface {
	StatpingDatabase(*gorm.DB)
}
```


#### type DbConfig

```go
type DbConfig struct {
	DbConn      string `yaml:"connection"`
	DbHost      string `yaml:"host"`
	DbUser      string `yaml:"user"`
	DbPass      string `yaml:"password"`
	DbData      string `yaml:"database"`
	DbPort      int64  `yaml:"port"`
	ApiKey      string `yaml:"api_key"`
	ApiSecret   string `yaml:"api_secret"`
	Project     string `yaml:"-"`
	Description string `yaml:"-"`
	Domain      string `yaml:"-"`
	Username    string `yaml:"-"`
	Password    string `yaml:"-"`
	Email       string `yaml:"-"`
	Error       error  `yaml:"-"`
	Location    string `yaml:"location"`
	LocalIP     string `yaml:"-"`
}
```

DbConfig struct is used for the database connection and creates the 'config.yml'
file

#### type FailSort

```go
type FailSort []FailureInterface
```


#### func (FailSort) Len

```go
func (s FailSort) Len() int
```

#### func (FailSort) Less

```go
func (s FailSort) Less(i, j int) bool
```

#### func (FailSort) Swap

```go
func (s FailSort) Swap(i, j int)
```

#### type Failure

```go
type Failure struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Issue     string    `gorm:"column:issue" json:"issue"`
	Method    string    `gorm:"column:method" json:"method,omitempty"`
	MethodId  int64     `gorm:"column:method_id" json:"method_id,omitempty"`
	ErrorCode int       `gorm:"column:error_code" json:"error_code"`
	Service   int64     `gorm:"index;column:service" json:"-"`
	Checkin   int64     `gorm:"index;column:checkin" json:"-"`
	PingTime  float64   `gorm:"column:ping_time"  json:"ping"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
```

Failure is a failed attempt to check a service. Any a service does not meet the
expected requirements, a new Failure will be inserted into database.

#### func (*Failure) BeforeCreate

```go
func (f *Failure) BeforeCreate() (err error)
```
BeforeCreate for Failure will set CreatedAt to UTC

#### type FailureInterface

```go
type FailureInterface interface {
	Select() *Failure
	Ago() string        // Ago returns a human readable timestamp
	ParseError() string // ParseError returns a human readable error for a service failure
}
```


#### type Group

```go
type Group struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Public    NullBool  `gorm:"default:true;column:public" json:"public"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
```

Group is the main struct for Groups

#### type Hit

```go
type Hit struct {
	Id        int64     `gorm:"primary_key;column:id" json:"id"`
	Service   int64     `gorm:"column:service" json:"-"`
	Latency   float64   `gorm:"column:latency" json:"latency"`
	PingTime  float64   `gorm:"column:ping_time" json:"ping_time"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
```

Hit struct is a 'successful' ping or web response entry for a service.

#### func (*Hit) BeforeCreate

```go
func (h *Hit) BeforeCreate() (err error)
```
BeforeCreate for Hit will set CreatedAt to UTC

#### type Info

```go
type Info struct {
	Name        string
	Description string
	Form        string
}
```


#### type Message

```go
type Message struct {
	Id                int64     `gorm:"primary_key;column:id" json:"id"`
	Title             string    `gorm:"column:title" json:"title"`
	Description       string    `gorm:"column:description" json:"description"`
	StartOn           time.Time `gorm:"column:start_on" json:"start_on"`
	EndOn             time.Time `gorm:"column:end_on" json:"end_on"`
	ServiceId         int64     `gorm:"index;column:service" json:"service"`
	NotifyUsers       NullBool  `gorm:"column:notify_users" json:"notify_users"`
	NotifyMethod      string    `gorm:"column:notify_method" json:"notify_method"`
	NotifyBefore      NullInt64 `gorm:"column:notify_before" json:"notify_before"`
	NotifyBeforeScale string    `gorm:"column:notify_before_scale" json:"notify_before_scale"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"created_at" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at" json:"updated_at" json:"updated_at"`
}
```

Message is for creating Announcements, Alerts and other messages for the end
users

#### func (*Message) BeforeCreate

```go
func (u *Message) BeforeCreate() (err error)
```
BeforeCreate for Message will set CreatedAt to UTC

#### type NullBool

```go
type NullBool struct {
	sql.NullBool
}
```

NullBool is an alias for sql.NullBool data type

#### func  NewNullBool

```go
func NewNullBool(s bool) NullBool
```
NewNullBool returns a sql.NullBool for JSON parsing

#### func (*NullBool) MarshalJSON

```go
func (nb *NullBool) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullBool

#### func (*NullBool) UnmarshalJSON

```go
func (nf *NullBool) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullBool

#### type NullFloat64

```go
type NullFloat64 struct {
	sql.NullFloat64
}
```

NullFloat64 is an alias for sql.NullFloat64 data type

#### func  NewNullFloat64

```go
func NewNullFloat64(s float64) NullFloat64
```
NewNullFloat64 returns a sql.NullFloat64 for JSON parsing

#### func (*NullFloat64) MarshalJSON

```go
func (ni *NullFloat64) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullFloat64

#### func (*NullFloat64) UnmarshalJSON

```go
func (nf *NullFloat64) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullFloat64

#### type NullInt64

```go
type NullInt64 struct {
	sql.NullInt64
}
```

NullInt64 is an alias for sql.NullInt64 data type

#### func  NewNullInt64

```go
func NewNullInt64(s int64) NullInt64
```
NewNullInt64 returns a sql.NullInt64 for JSON parsing

#### func (*NullInt64) MarshalJSON

```go
func (ni *NullInt64) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullInt64

#### func (*NullInt64) UnmarshalJSON

```go
func (nf *NullInt64) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullInt64

#### type NullString

```go
type NullString struct {
	sql.NullString
}
```

NullString is an alias for sql.NullString data type

#### func  NewNullString

```go
func NewNullString(s string) NullString
```
NewNullString returns a sql.NullString for JSON parsing

#### func (*NullString) MarshalJSON

```go
func (ns *NullString) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullString

#### func (*NullString) UnmarshalJSON

```go
func (nf *NullString) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullString

#### type Plugin

```go
type Plugin struct {
	Name        string
	Description string
}
```


#### type PluginActions

```go
type PluginActions interface {
	GetInfo() *Info
	OnLoad() error
}
```


#### type PluginInfo

```go
type PluginInfo struct {
	Info   *Info
	Routes []*PluginRoute
}
```


#### type PluginJSON

```go
type PluginJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Repo        string `json:"repo"`
	Author      string `json:"author"`
	Namespace   string `json:"namespace"`
}
```


#### type PluginObject

```go
type PluginObject struct {
	Pluginer
}
```


#### type PluginRepos

```go
type PluginRepos struct {
	Plugins []PluginJSON
}
```


#### type PluginRoute

```go
type PluginRoute struct {
	Url    string
	Method string
	Func   http.HandlerFunc
}
```


#### type PluginRouting

```go
type PluginRouting struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}
```


#### type Pluginer

```go
type Pluginer interface {
	Select() *Plugin
}
```


#### type Router

```go
type Router interface {
	Routes() []*PluginRoute
	AddRoute(string, string, http.HandlerFunc) error
}
```


#### type Service

```go
type Service struct {
	Id                 int64              `gorm:"primary_key;column:id" json:"id"`
	Name               string             `gorm:"column:name" json:"name"`
	Domain             string             `gorm:"column:domain" json:"domain"`
	Expected           NullString         `gorm:"column:expected" json:"expected"`
	ExpectedStatus     int                `gorm:"default:200;column:expected_status" json:"expected_status"`
	Interval           int                `gorm:"default:30;column:check_interval" json:"check_interval"`
	Type               string             `gorm:"column:check_type" json:"type"`
	Method             string             `gorm:"column:method" json:"method"`
	PostData           NullString         `gorm:"column:post_data" json:"post_data"`
	Port               int                `gorm:"not null;column:port" json:"port"`
	Timeout            int                `gorm:"default:30;column:timeout" json:"timeout"`
	Order              int                `gorm:"default:0;column:order_id" json:"order_id"`
	AllowNotifications NullBool           `gorm:"default:true;column:allow_notifications" json:"allow_notifications"`
	Public             NullBool           `gorm:"default:true;column:public" json:"public"`
	GroupId            int                `gorm:"default:0;column:group_id" json:"group_id"`
	Permalink          NullString         `gorm:"column:permalink" json:"permalink"`
	CreatedAt          time.Time          `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time          `gorm:"column:updated_at" json:"updated_at"`
	Online             bool               `gorm:"-" json:"online"`
	Latency            float64            `gorm:"-" json:"latency"`
	PingTime           float64            `gorm:"-" json:"ping_time"`
	Online24Hours      float32            `gorm:"-" json:"online_24_hours"`
	AvgResponse        string             `gorm:"-" json:"avg_response"`
	Running            chan bool          `gorm:"-" json:"-"`
	Checkpoint         time.Time          `gorm:"-" json:"-"`
	SleepDuration      time.Duration      `gorm:"-" json:"-"`
	LastResponse       string             `gorm:"-" json:"-"`
	LastStatusCode     int                `gorm:"-" json:"status_code"`
	LastOnline         time.Time          `gorm:"-" json:"last_success"`
	Failures           []FailureInterface `gorm:"-" json:"failures,omitempty"`
	Checkins           []CheckinInterface `gorm:"-" json:"checkins,omitempty"`
}
```

Service is the main struct for Services

#### func (*Service) BeforeCreate

```go
func (s *Service) BeforeCreate() (err error)
```
BeforeCreate for Service will set CreatedAt to UTC

#### func (*Service) Close

```go
func (s *Service) Close()
```
Close will stop the go routine that is checking if service is online or not

#### func (*Service) IsRunning

```go
func (s *Service) IsRunning() bool
```
IsRunning returns true if the service go routine is running

#### func (*Service) Start

```go
func (s *Service) Start()
```
Start will create a channel for the service checking go routine

#### type ServiceInterface

```go
type ServiceInterface interface {
	Select() *Service
	CheckQueue(bool)
	Check(bool)
	Create(bool) (int64, error)
	Update(bool) error
	Delete() error
}
```


#### type User

```go
type User struct {
	Id            int64     `gorm:"primary_key;column:id" json:"id"`
	Username      string    `gorm:"type:varchar(100);unique;column:username;" json:"username,omitempty"`
	Password      string    `gorm:"column:password" json:"password,omitempty"`
	Email         string    `gorm:"type:varchar(100);unique;column:email" json:"email,omitempty"`
	ApiKey        string    `gorm:"column:api_key" json:"api_key,omitempty"`
	ApiSecret     string    `gorm:"column:api_secret" json:"api_secret,omitempty"`
	Admin         NullBool  `gorm:"column:administrator" json:"admin,omitempty"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserInterface `gorm:"-" json:"-"`
}
```

User is the main struct for Users

#### func (*User) BeforeCreate

```go
func (u *User) BeforeCreate() (err error)
```
BeforeCreate for User will set CreatedAt to UTC

#### type UserInterface

```go
type UserInterface interface {
	Create() (int64, error)
	Update() error
	Delete() error
}
```

UserInterface interfaces the database functions
# utils
--
    import "github.com/hunterlong/statping/utils"

Package utils contains common methods used in most packages in Statping. This
package contains multiple function like: Logging, encryption, type conversions,
setting utils.Directory as the current directory, running local CMD commands,
and creating/deleting files/folder.

You can overwrite the utils.Directory global variable by including STATPING_DIR
environment variable to be an absolute path.

More info on: https://github.com/hunterlong/statping

## Usage

```go
const (
	FlatpickrTime     = "2006-01-02 15:04"
	FlatpickrDay      = "2006-01-02"
	FlatpickrReadable = "Mon, 02 Jan 2006"
)
```

```go
var (
	LastLines []*LogRow
	LockLines sync.Mutex
)
```

```go
var (
	// Directory returns the current path or the STATPING_DIR environment variable
	Directory string
)
```

#### func  Command

```go
func Command(cmd string) (string, string, error)
```
Command will run a terminal command with 'sh -c COMMAND' and return stdout and
errOut as strings

    in, out, err := Command("sass assets/scss assets/css/base.css")

#### func  DeleteDirectory

```go
func DeleteDirectory(directory string) error
```
DeleteDirectory will attempt to delete a directory and all contents inside

    DeleteDirectory("assets")

#### func  DeleteFile

```go
func DeleteFile(file string) error
```
DeleteFile will attempt to delete a file

    DeleteFile("newfile.json")

#### func  DurationReadable

```go
func DurationReadable(d time.Duration) string
```
DurationReadable will return a time.Duration into a human readable string // t
:= time.Duration(5 * time.Minute) // DurationReadable(t) // returns: 5 minutes

#### func  FileExists

```go
func FileExists(name string) bool
```
FileExists returns true if a file exists

    exists := FileExists("assets/css/base.css")

#### func  FormatDuration

```go
func FormatDuration(d time.Duration) string
```
FormatDuration converts a time.Duration into a string

#### func  HashPassword

```go
func HashPassword(password string) string
```
HashPassword returns the bcrypt hash of a password string

#### func  Http

```go
func Http(r *http.Request) string
```
Http returns a log for a HTTP request

#### func  HttpRequest

```go
func HttpRequest(url, method string, content interface{}, headers []string, body io.Reader, timeout time.Duration) ([]byte, *http.Response, error)
```
HttpRequest is a global function to send a HTTP request // url - The URL for
HTTP request // method - GET, POST, DELETE, PATCH // content - The HTTP request
content type (text/plain, application/json, or nil) // headers - An array of
Headers to be sent (KEY=VALUE) []string{"Authentication=12345", ...} // body -
The body or form data to send with HTTP request // timeout - Specific duration
to timeout on. time.Duration(30 * time.Seconds) // You can use a HTTP Proxy if
you HTTP_PROXY environment variable

#### func  InitLogs

```go
func InitLogs() error
```
InitLogs will create the '/logs' directory and creates a file '/logs/statup.log'
for application logging

#### func  Log

```go
func Log(level int, err interface{}) error
```
Log creates a new entry in the utils.Log. Log has 1-5 levels depending on how
critical the log/error is

#### func  NewSHA1Hash

```go
func NewSHA1Hash(n ...int) string
```
NewSHA1Hash returns a random SHA1 hash based on a specific length

#### func  RandomString

```go
func RandomString(n int) string
```
RandomString generates a random string of n length

#### func  SaveFile

```go
func SaveFile(filename string, data []byte) error
```
SaveFile will create a new file with data inside it

    SaveFile("newfile.json", []byte('{"data": "success"}')

#### func  Timezoner

```go
func Timezoner(t time.Time, zone float32) time.Time
```
Timezoner returns the time.Time with the user set timezone

#### func  ToInt

```go
func ToInt(s interface{}) int64
```
ToInt converts a int to a string

#### func  ToString

```go
func ToString(s interface{}) string
```
ToString converts a int to a string

#### func  UnderScoreString

```go
func UnderScoreString(str string) string
```
UnderScoreString will return a string that replaces spaces and other characters
to underscores

    UnderScoreString("Example String")
    // example_string

#### type LogRow

```go
type LogRow struct {
	Date time.Time
	Line interface{}
}
```


#### func  GetLastLine

```go
func GetLastLine() *LogRow
```
GetLastLine returns 1 line for a recent log entry

#### func (*LogRow) FormatForHtml

```go
func (o *LogRow) FormatForHtml() string
```

#### type Timestamp

```go
type Timestamp time.Time
```


#### func (Timestamp) Ago

```go
func (t Timestamp) Ago() string
```
Ago returns a human readable timestamp based on the Timestamp (time.Time)
interface

#### type Timestamper

```go
type Timestamper interface {
	Ago() string
}
```
