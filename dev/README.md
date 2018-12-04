

# core
`import "github.com/hunterlong/statping/core"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package core contains the main functionality of Statping. This includes everything for
Services, Hits, Failures, Users, service checking mechanisms, databases, and notifiers
in the notifier package

More info on: <a href="https://github.com/hunterlong/statping">https://github.com/hunterlong/statping</a>




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func CheckHash(password, hash string) bool](#CheckHash)
* [func CloseDB()](#CloseDB)
* [func CountFailures() uint64](#CountFailures)
* [func CountUsers() int64](#CountUsers)
* [func DatabaseMaintence()](#DatabaseMaintence)
* [func Dbtimestamp(group string, column string) string](#Dbtimestamp)
* [func DefaultPort(db string) int64](#DefaultPort)
* [func DeleteAllSince(table string, date time.Time)](#DeleteAllSince)
* [func DeleteConfig() error](#DeleteConfig)
* [func ExportChartsJs() string](#ExportChartsJs)
* [func ExportSettings() ([]byte, error)](#ExportSettings)
* [func GetLocalIP() string](#GetLocalIP)
* [func InitApp()](#InitApp)
* [func InsertLargeSampleData() error](#InsertLargeSampleData)
* [func InsertNotifierDB() error](#InsertNotifierDB)
* [func InsertSampleData() error](#InsertSampleData)
* [func InsertSampleHits() error](#InsertSampleHits)
* [func SampleData() error](#SampleData)
* [func SelectServicer(id int64) types.ServiceInterface](#SelectServicer)
* [func Services() []types.ServiceInterface](#Services)
* [type Checkin](#Checkin)
  * [func AllCheckins() []*Checkin](#AllCheckins)
  * [func ReturnCheckin(c *types.Checkin) *Checkin](#ReturnCheckin)
  * [func SelectCheckin(api string) *Checkin](#SelectCheckin)
  * [func SelectCheckinId(id int64) *Checkin](#SelectCheckinId)
  * [func (c *Checkin) AfterFind() (err error)](#Checkin.AfterFind)
  * [func (c *Checkin) AllFailures() []*types.Failure](#Checkin.AllFailures)
  * [func (c *Checkin) AllHits() []*types.CheckinHit](#Checkin.AllHits)
  * [func (c *Checkin) BeforeCreate() (err error)](#Checkin.BeforeCreate)
  * [func (c *Checkin) Create() (int64, error)](#Checkin.Create)
  * [func (c *Checkin) CreateFailure() (int64, error)](#Checkin.CreateFailure)
  * [func (c *Checkin) Delete() error](#Checkin.Delete)
  * [func (c *Checkin) Expected() time.Duration](#Checkin.Expected)
  * [func (c *Checkin) Grace() time.Duration](#Checkin.Grace)
  * [func (c *Checkin) Last() *CheckinHit](#Checkin.Last)
  * [func (c *Checkin) Link() string](#Checkin.Link)
  * [func (c *Checkin) Period() time.Duration](#Checkin.Period)
  * [func (c *Checkin) RecheckCheckinFailure(guard chan struct{})](#Checkin.RecheckCheckinFailure)
  * [func (c *Checkin) Routine()](#Checkin.Routine)
  * [func (c *Checkin) Service() *Service](#Checkin.Service)
  * [func (c *Checkin) String() string](#Checkin.String)
  * [func (c *Checkin) Update() (int64, error)](#Checkin.Update)
* [type CheckinHit](#CheckinHit)
  * [func ReturnCheckinHit(c *types.CheckinHit) *CheckinHit](#ReturnCheckinHit)
  * [func (c *CheckinHit) AfterFind() (err error)](#CheckinHit.AfterFind)
  * [func (c *CheckinHit) Ago() string](#CheckinHit.Ago)
  * [func (c *CheckinHit) BeforeCreate() (err error)](#CheckinHit.BeforeCreate)
  * [func (c *CheckinHit) Create() (int64, error)](#CheckinHit.Create)
* [type Core](#Core)
  * [func NewCore() *Core](#NewCore)
  * [func SelectCore() (*Core, error)](#SelectCore)
  * [func UpdateCore(c *Core) (*Core, error)](#UpdateCore)
  * [func (c *Core) AfterFind() (err error)](#Core.AfterFind)
  * [func (c Core) AllOnline() bool](#Core.AllOnline)
  * [func (c Core) BaseSASS() string](#Core.BaseSASS)
  * [func (c *Core) Count24HFailures() uint64](#Core.Count24HFailures)
  * [func (c *Core) CountOnline() int](#Core.CountOnline)
  * [func (c Core) CurrentTime() string](#Core.CurrentTime)
  * [func (c Core) Messages() []*Message](#Core.Messages)
  * [func (c Core) MobileSASS() string](#Core.MobileSASS)
  * [func (c Core) SassVars() string](#Core.SassVars)
  * [func (c *Core) SelectAllServices(start bool) ([]*Service, error)](#Core.SelectAllServices)
  * [func (c *Core) ServicesCount() int](#Core.ServicesCount)
  * [func (c *Core) ToCore() *types.Core](#Core.ToCore)
  * [func (c Core) UsingAssets() bool](#Core.UsingAssets)
* [type DateScan](#DateScan)
* [type DateScanObj](#DateScanObj)
  * [func GraphDataRaw(service types.ServiceInterface, start, end time.Time, group string, column string) *DateScanObj](#GraphDataRaw)
  * [func (d *DateScanObj) ToString() string](#DateScanObj.ToString)
* [type DbConfig](#DbConfig)
  * [func EnvToConfig() *DbConfig](#EnvToConfig)
  * [func LoadConfigFile(directory string) (*DbConfig, error)](#LoadConfigFile)
  * [func LoadUsingEnv() (*DbConfig, error)](#LoadUsingEnv)
  * [func (db *DbConfig) Connect(retry bool, location string) error](#DbConfig.Connect)
  * [func (c *DbConfig) CreateCore() *Core](#DbConfig.CreateCore)
  * [func (db *DbConfig) CreateDatabase() error](#DbConfig.CreateDatabase)
  * [func (db *DbConfig) DropDatabase() error](#DbConfig.DropDatabase)
  * [func (db *DbConfig) InsertCore() (*Core, error)](#DbConfig.InsertCore)
  * [func (db *DbConfig) MigrateDatabase() error](#DbConfig.MigrateDatabase)
  * [func (db *DbConfig) Save() (*DbConfig, error)](#DbConfig.Save)
  * [func (db *DbConfig) Update() error](#DbConfig.Update)
* [type ErrorResponse](#ErrorResponse)
* [type ExportData](#ExportData)
* [type Hit](#Hit)
  * [func (h *Hit) AfterFind() (err error)](#Hit.AfterFind)
  * [func (h *Hit) BeforeCreate() (err error)](#Hit.BeforeCreate)
* [type Message](#Message)
  * [func ReturnMessage(m *types.Message) *Message](#ReturnMessage)
  * [func SelectMessage(id int64) (*Message, error)](#SelectMessage)
  * [func SelectMessages() ([]*Message, error)](#SelectMessages)
  * [func SelectServiceMessages(id int64) []*Message](#SelectServiceMessages)
  * [func (u *Message) AfterFind() (err error)](#Message.AfterFind)
  * [func (u *Message) BeforeCreate() (err error)](#Message.BeforeCreate)
  * [func (m *Message) Create() (int64, error)](#Message.Create)
  * [func (m *Message) Delete() error](#Message.Delete)
  * [func (m *Message) Service() *Service](#Message.Service)
  * [func (m *Message) Update() (*Message, error)](#Message.Update)
* [type PluginJSON](#PluginJSON)
* [type PluginRepos](#PluginRepos)
* [type Service](#Service)
  * [func ReturnService(s *types.Service) *Service](#ReturnService)
  * [func SelectService(id int64) *Service](#SelectService)
  * [func (s *Service) ActiveMessages() []*Message](#Service.ActiveMessages)
  * [func (s *Service) AfterFind() (err error)](#Service.AfterFind)
  * [func (s *Service) AllFailures() []*failure](#Service.AllFailures)
  * [func (s *Service) AvgTime() float64](#Service.AvgTime)
  * [func (s *Service) AvgUptime(ago time.Time) string](#Service.AvgUptime)
  * [func (s *Service) AvgUptime24() string](#Service.AvgUptime24)
  * [func (s *Service) BeforeCreate() (err error)](#Service.BeforeCreate)
  * [func (s *Service) Check(record bool)](#Service.Check)
  * [func (s *Service) CheckQueue(record bool)](#Service.CheckQueue)
  * [func (s *Service) CheckinProcess()](#Service.CheckinProcess)
  * [func (s *Service) Checkins() []*Checkin](#Service.Checkins)
  * [func (s *Service) CountHits() (int64, error)](#Service.CountHits)
  * [func (s *Service) Create(check bool) (int64, error)](#Service.Create)
  * [func (s *Service) CreateFailure(fail types.FailureInterface) (int64, error)](#Service.CreateFailure)
  * [func (s *Service) CreateHit(h *types.Hit) (int64, error)](#Service.CreateHit)
  * [func (s *Service) Delete() error](#Service.Delete)
  * [func (s *Service) DeleteFailures()](#Service.DeleteFailures)
  * [func (s *Service) Downtime() time.Duration](#Service.Downtime)
  * [func (s *Service) DowntimeText() string](#Service.DowntimeText)
  * [func (s *Service) Hits() ([]*types.Hit, error)](#Service.Hits)
  * [func (s *Service) HitsBetween(t1, t2 time.Time, group string, column string) *gorm.DB](#Service.HitsBetween)
  * [func (s *Service) LimitedCheckinFailures(amount int64) []*failure](#Service.LimitedCheckinFailures)
  * [func (s *Service) LimitedCheckins() []*Checkin](#Service.LimitedCheckins)
  * [func (s *Service) LimitedFailures(amount int64) []*failure](#Service.LimitedFailures)
  * [func (s *Service) LimitedHits() ([]*types.Hit, error)](#Service.LimitedHits)
  * [func (s *Service) Messages() []*Message](#Service.Messages)
  * [func (s *Service) Online24() float32](#Service.Online24)
  * [func (s *Service) OnlineSince(ago time.Time) float32](#Service.OnlineSince)
  * [func (s *Service) Select() *types.Service](#Service.Select)
  * [func (s *Service) SmallText() string](#Service.SmallText)
  * [func (s *Service) Sum() (float64, error)](#Service.Sum)
  * [func (s *Service) ToJSON() string](#Service.ToJSON)
  * [func (s *Service) TotalFailures() (uint64, error)](#Service.TotalFailures)
  * [func (s *Service) TotalFailures24() (uint64, error)](#Service.TotalFailures24)
  * [func (s *Service) TotalFailuresSince(ago time.Time) (uint64, error)](#Service.TotalFailuresSince)
  * [func (s *Service) TotalHits() (uint64, error)](#Service.TotalHits)
  * [func (s *Service) TotalHitsSince(ago time.Time) (uint64, error)](#Service.TotalHitsSince)
  * [func (s *Service) TotalUptime() string](#Service.TotalUptime)
  * [func (s *Service) Update(restart bool) error](#Service.Update)
  * [func (s *Service) UpdateSingle(attr ...interface{}) error](#Service.UpdateSingle)
* [type ServiceOrder](#ServiceOrder)
  * [func (c ServiceOrder) Len() int](#ServiceOrder.Len)
  * [func (c ServiceOrder) Less(i, j int) bool](#ServiceOrder.Less)
  * [func (c ServiceOrder) Swap(i, j int)](#ServiceOrder.Swap)
* [type User](#User)
  * [func AuthUser(username, password string) (*User, bool)](#AuthUser)
  * [func ReturnUser(u *types.User) *User](#ReturnUser)
  * [func SelectAllUsers() ([]*User, error)](#SelectAllUsers)
  * [func SelectUser(id int64) (*User, error)](#SelectUser)
  * [func SelectUsername(username string) (*User, error)](#SelectUsername)
  * [func (u *User) AfterFind() (err error)](#User.AfterFind)
  * [func (u *User) BeforeCreate() (err error)](#User.BeforeCreate)
  * [func (u *User) Create() (int64, error)](#User.Create)
  * [func (u *User) Delete() error](#User.Delete)
  * [func (u *User) Update() error](#User.Update)


#### <a name="pkg-files">Package files</a>
[checker.go](https://github.com/hunterlong/statping/tree/master/core/checker.go) [checkin.go](https://github.com/hunterlong/statping/tree/master/core/checkin.go) [configs.go](https://github.com/hunterlong/statping/tree/master/core/configs.go) [core.go](https://github.com/hunterlong/statping/tree/master/core/core.go) [database.go](https://github.com/hunterlong/statping/tree/master/core/database.go) [doc.go](https://github.com/hunterlong/statping/tree/master/core/doc.go) [export.go](https://github.com/hunterlong/statping/tree/master/core/export.go) [failures.go](https://github.com/hunterlong/statping/tree/master/core/failures.go) [hits.go](https://github.com/hunterlong/statping/tree/master/core/hits.go) [messages.go](https://github.com/hunterlong/statping/tree/master/core/messages.go) [sample.go](https://github.com/hunterlong/statping/tree/master/core/sample.go) [services.go](https://github.com/hunterlong/statping/tree/master/core/services.go) [users.go](https://github.com/hunterlong/statping/tree/master/core/users.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    Configs   *DbConfig // Configs holds all of the config.yml and database info
    CoreApp   *Core     // CoreApp is a global variable that contains many elements
    SetupMode bool      // SetupMode will be true if Statping does not have a database connection
    VERSION   string    // VERSION is set on build automatically by setting a -ldflag
)
```
``` go
var (
    // DbSession stores the Statping database session
    DbSession *gorm.DB
    DbModels  []interface{}
)
```


## <a name="CheckHash">func</a> [CheckHash](https://github.com/hunterlong/statping/tree/master/core/users.go?s=3032:3074#L112)
``` go
func CheckHash(password, hash string) bool
```
CheckHash returns true if the password matches with a hashed bcrypt password



## <a name="CloseDB">func</a> [CloseDB](https://github.com/hunterlong/statping/tree/master/core/database.go?s=3206:3220#L97)
``` go
func CloseDB()
```
CloseDB will close the database connection if available



## <a name="CountFailures">func</a> [CountFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=3409:3436#L117)
``` go
func CountFailures() uint64
```
CountFailures returns the total count of failures for all services



## <a name="CountUsers">func</a> [CountUsers](https://github.com/hunterlong/statping/tree/master/core/users.go?s=1020:1043#L36)
``` go
func CountUsers() int64
```
CountUsers returns the amount of users



## <a name="DatabaseMaintence">func</a> [DatabaseMaintence](https://github.com/hunterlong/statping/tree/master/core/database.go?s=8692:8716#L283)
``` go
func DatabaseMaintence()
```
DatabaseMaintence will automatically delete old records from 'failures' and 'hits'
this function is currently set to delete records 7+ days old every 60 minutes



## <a name="Dbtimestamp">func</a> [Dbtimestamp](https://github.com/hunterlong/statping/tree/master/core/services.go?s=6195:6247#L219)
``` go
func Dbtimestamp(group string, column string) string
```
Dbtimestamp will return a SQL query for grouping by date



## <a name="DefaultPort">func</a> [DefaultPort](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=3380:3413#L115)
``` go
func DefaultPort(db string) int64
```
DefaultPort accepts a database type and returns its default port



## <a name="DeleteAllSince">func</a> [DeleteAllSince](https://github.com/hunterlong/statping/tree/master/core/database.go?s=9025:9074#L293)
``` go
func DeleteAllSince(table string, date time.Time)
```
DeleteAllSince will delete a specific table's records based on a time.



## <a name="DeleteConfig">func</a> [DeleteConfig](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=4848:4873#L184)
``` go
func DeleteConfig() error
```
DeleteConfig will delete the 'config.yml' file



## <a name="ExportChartsJs">func</a> [ExportChartsJs](https://github.com/hunterlong/statping/tree/master/core/export.go?s=922:950#L28)
``` go
func ExportChartsJs() string
```
ExportChartsJs renders the charts for the index page



## <a name="ExportSettings">func</a> [ExportSettings](https://github.com/hunterlong/statping/tree/master/core/export.go?s=1690:1727#L57)
``` go
func ExportSettings() ([]byte, error)
```


## <a name="GetLocalIP">func</a> [GetLocalIP](https://github.com/hunterlong/statping/tree/master/core/core.go?s=4549:4573#L164)
``` go
func GetLocalIP() string
```
GetLocalIP returns the non loopback local IP of the host



## <a name="InitApp">func</a> [InitApp](https://github.com/hunterlong/statping/tree/master/core/core.go?s=1705:1719#L62)
``` go
func InitApp()
```
InitApp will initialize Statping



## <a name="InsertLargeSampleData">func</a> [InsertLargeSampleData](https://github.com/hunterlong/statping/tree/master/core/sample.go?s=5703:5737#L212)
``` go
func InsertLargeSampleData() error
```
InsertLargeSampleData will create the example/dummy services for testing the Statping server



## <a name="InsertNotifierDB">func</a> [InsertNotifierDB](https://github.com/hunterlong/statping/tree/master/core/core.go?s=1956:1985#L72)
``` go
func InsertNotifierDB() error
```
InsertNotifierDB inject the Statping database instance to the Notifier package



## <a name="InsertSampleData">func</a> [InsertSampleData](https://github.com/hunterlong/statping/tree/master/core/sample.go?s=907:936#L27)
``` go
func InsertSampleData() error
```
InsertSampleData will create the example/dummy services for a brand new Statping installation



## <a name="InsertSampleHits">func</a> [InsertSampleHits](https://github.com/hunterlong/statping/tree/master/core/sample.go?s=3388:3417#L126)
``` go
func InsertSampleHits() error
```
InsertSampleHits will create a couple new hits for the sample services



## <a name="SampleData">func</a> [SampleData](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=4637:4660#L173)
``` go
func SampleData() error
```
SampleData runs all the sample data for a new Statping installation



## <a name="SelectServicer">func</a> [SelectServicer](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1546:1598#L59)
``` go
func SelectServicer(id int64) types.ServiceInterface
```
SelectServicer returns a types.ServiceInterface from in memory



## <a name="Services">func</a> [Services](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1209:1249#L44)
``` go
func Services() []types.ServiceInterface
```



## <a name="Checkin">type</a> [Checkin](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=826:865#L26)
``` go
type Checkin struct {
    *types.Checkin
}

```






### <a name="AllCheckins">func</a> [AllCheckins](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=2636:2665#L94)
``` go
func AllCheckins() []*Checkin
```
AllCheckins returns all checkin in system


### <a name="ReturnCheckin">func</a> [ReturnCheckin](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=1803:1848#L66)
``` go
func ReturnCheckin(c *types.Checkin) *Checkin
```
ReturnCheckin converts *types.Checking to *core.Checkin


### <a name="SelectCheckin">func</a> [SelectCheckin](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=2805:2844#L101)
``` go
func SelectCheckin(api string) *Checkin
```
SelectCheckin will find a Checkin based on the API supplied


### <a name="SelectCheckinId">func</a> [SelectCheckinId](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=3006:3045#L108)
``` go
func SelectCheckinId(id int64) *Checkin
```
SelectCheckin will find a Checkin based on the API supplied





### <a name="Checkin.AfterFind">func</a> (\*Checkin) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=4310:4351#L137)
``` go
func (c *Checkin) AfterFind() (err error)
```
AfterFind for Checkin will set the timezone




### <a name="Checkin.AllFailures">func</a> (\*Checkin) [AllFailures](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=4418:4466#L154)
``` go
func (c *Checkin) AllFailures() []*types.Failure
```
Hits returns all of the CheckinHits for a given Checkin




### <a name="Checkin.AllHits">func</a> (\*Checkin) [AllHits](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=4178:4225#L147)
``` go
func (c *Checkin) AllHits() []*types.CheckinHit
```
AllHits returns all of the CheckinHits for a given Checkin




### <a name="Checkin.BeforeCreate">func</a> (\*Checkin) [BeforeCreate](https://github.com/hunterlong/statping/tree/master/core/database.go?s=6023:6067#L202)
``` go
func (c *Checkin) BeforeCreate() (err error)
```
BeforeCreate for Checkin will set CreatedAt to UTC




### <a name="Checkin.Create">func</a> (\*Checkin) [Create](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=4804:4845#L169)
``` go
func (c *Checkin) Create() (int64, error)
```
Create will create a new Checkin




### <a name="Checkin.CreateFailure">func</a> (\*Checkin) [CreateFailure](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=2142:2190#L80)
``` go
func (c *Checkin) CreateFailure() (int64, error)
```



### <a name="Checkin.Delete">func</a> (\*Checkin) [Delete](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=4670:4702#L162)
``` go
func (c *Checkin) Delete() error
```
Create will create a new Checkin




### <a name="Checkin.Expected">func</a> (\*Checkin) [Expected](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=3659:3701#L127)
``` go
func (c *Checkin) Expected() time.Duration
```
Expected returns the duration of when the serviec should receive a Checkin




### <a name="Checkin.Grace">func</a> (\*Checkin) [Grace](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=3449:3488#L121)
``` go
func (c *Checkin) Grace() time.Duration
```
Grace will return the duration of the Checkin Grace Period (after service hasn't responded, wait a bit for a response)




### <a name="Checkin.Last">func</a> (\*Checkin) [Last](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=3886:3922#L136)
``` go
func (c *Checkin) Last() *CheckinHit
```
Last returns the last checkinHit for a Checkin




### <a name="Checkin.Link">func</a> (\*Checkin) [Link](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=4016:4047#L142)
``` go
func (c *Checkin) Link() string
```



### <a name="Checkin.Period">func</a> (\*Checkin) [Period](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=3197:3237#L115)
``` go
func (c *Checkin) Period() time.Duration
```
Period will return the duration of the Checkin interval




### <a name="Checkin.RecheckCheckinFailure">func</a> (\*Checkin) [RecheckCheckinFailure](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=5812:5872#L211)
``` go
func (c *Checkin) RecheckCheckinFailure(guard chan struct{})
```
RecheckCheckinFailure will check if a Service Checkin has been reported yet




### <a name="Checkin.Routine">func</a> (\*Checkin) [Routine](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=982:1009#L35)
``` go
func (c *Checkin) Routine()
```
Routine for checking if the last Checkin was within its interval




### <a name="Checkin.Service">func</a> (\*Checkin) [Service](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=2045:2081#L75)
``` go
func (c *Checkin) Service() *Service
```



### <a name="Checkin.String">func</a> (\*Checkin) [String](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=1688:1721#L61)
``` go
func (c *Checkin) String() string
```
String will return a Checkin API string




### <a name="Checkin.Update">func</a> (\*Checkin) [Update](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=5073:5114#L182)
``` go
func (c *Checkin) Update() (int64, error)
```
Update will update a Checkin




## <a name="CheckinHit">type</a> [CheckinHit](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=867:912#L30)
``` go
type CheckinHit struct {
    *types.CheckinHit
}

```






### <a name="ReturnCheckinHit">func</a> [ReturnCheckinHit](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=1950:2004#L71)
``` go
func ReturnCheckinHit(c *types.CheckinHit) *CheckinHit
```
ReturnCheckinHit converts *types.checkinHit to *core.checkinHit





### <a name="CheckinHit.AfterFind">func</a> (\*CheckinHit) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=4539:4583#L144)
``` go
func (c *CheckinHit) AfterFind() (err error)
```
AfterFind for checkinHit will set the timezone




### <a name="CheckinHit.Ago">func</a> (\*CheckinHit) [Ago](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=5622:5655#L205)
``` go
func (c *CheckinHit) Ago() string
```
Ago returns the duration of time between now and the last successful checkinHit




### <a name="CheckinHit.BeforeCreate">func</a> (\*CheckinHit) [BeforeCreate](https://github.com/hunterlong/statping/tree/master/core/database.go?s=6234:6281#L211)
``` go
func (c *CheckinHit) BeforeCreate() (err error)
```
BeforeCreate for checkinHit will set CreatedAt to UTC




### <a name="CheckinHit.Create">func</a> (\*CheckinHit) [Create](https://github.com/hunterlong/statping/tree/master/core/checkin.go?s=5299:5343#L192)
``` go
func (c *CheckinHit) Create() (int64, error)
```
Create will create a new successful checkinHit




## <a name="Core">type</a> [Core](https://github.com/hunterlong/statping/tree/master/core/core.go?s=978:1011#L33)
``` go
type Core struct {
    *types.Core
}

```






### <a name="NewCore">func</a> [NewCore](https://github.com/hunterlong/statping/tree/master/core/core.go?s=1439:1459#L49)
``` go
func NewCore() *Core
```
NewCore return a new *core.Core struct


### <a name="SelectCore">func</a> [SelectCore](https://github.com/hunterlong/statping/tree/master/core/core.go?s=3999:4031#L145)
``` go
func SelectCore() (*Core, error)
```
SelectCore will return the CoreApp global variable and the settings/configs for Statping


### <a name="UpdateCore">func</a> [UpdateCore](https://github.com/hunterlong/statping/tree/master/core/core.go?s=2296:2335#L84)
``` go
func UpdateCore(c *Core) (*Core, error)
```
UpdateCore will update the CoreApp variable inside of the 'core' table in database





### <a name="Core.AfterFind">func</a> (\*Core) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=3321:3359#L104)
``` go
func (c *Core) AfterFind() (err error)
```
AfterFind for Core will set the timezone




### <a name="Core.AllOnline">func</a> (Core) [AllOnline](https://github.com/hunterlong/statping/tree/master/core/core.go?s=3771:3801#L135)
``` go
func (c Core) AllOnline() bool
```
AllOnline will be true if all services are online




### <a name="Core.BaseSASS">func</a> (Core) [BaseSASS](https://github.com/hunterlong/statping/tree/master/core/core.go?s=3274:3305#L118)
``` go
func (c Core) BaseSASS() string
```
BaseSASS is the base design , this opens the file /assets/scss/base.scss to be edited in Theme




### <a name="Core.Count24HFailures">func</a> (\*Core) [Count24HFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=3137:3177#L106)
``` go
func (c *Core) Count24HFailures() uint64
```
Count24HFailures returns the amount of failures for a service within the last 24 hours




### <a name="Core.CountOnline">func</a> (\*Core) [CountOnline](https://github.com/hunterlong/statping/tree/master/core/services.go?s=12620:12652#L455)
``` go
func (c *Core) CountOnline() int
```
CountOnline




### <a name="Core.CurrentTime">func</a> (Core) [CurrentTime](https://github.com/hunterlong/statping/tree/master/core/core.go?s=2438:2472#L90)
``` go
func (c Core) CurrentTime() string
```
CurrentTime will return the current local time




### <a name="Core.Messages">func</a> (Core) [Messages](https://github.com/hunterlong/statping/tree/master/core/core.go?s=2652:2687#L98)
``` go
func (c Core) Messages() []*Message
```
Messages will return the current local time




### <a name="Core.MobileSASS">func</a> (Core) [MobileSASS](https://github.com/hunterlong/statping/tree/master/core/core.go?s=3559:3592#L127)
``` go
func (c Core) MobileSASS() string
```
MobileSASS is the -webkit responsive custom css designs. This opens the
file /assets/scss/mobile.scss to be edited in Theme




### <a name="Core.SassVars">func</a> (Core) [SassVars](https://github.com/hunterlong/statping/tree/master/core/core.go?s=3016:3047#L110)
``` go
func (c Core) SassVars() string
```
SassVars opens the file /assets/scss/variables.scss to be edited in Theme




### <a name="Core.SelectAllServices">func</a> (\*Core) [SelectAllServices](https://github.com/hunterlong/statping/tree/master/core/services.go?s=2448:2512#L92)
``` go
func (c *Core) SelectAllServices(start bool) ([]*Service, error)
```
SelectAllServices returns a slice of *core.Service to be store on []*core.Services, should only be called once on startup.




### <a name="Core.ServicesCount">func</a> (\*Core) [ServicesCount](https://github.com/hunterlong/statping/tree/master/core/services.go?s=12541:12575#L450)
``` go
func (c *Core) ServicesCount() int
```
ServicesCount returns the amount of services inside the []*core.Services slice




### <a name="Core.ToCore">func</a> (\*Core) [ToCore](https://github.com/hunterlong/statping/tree/master/core/core.go?s=1613:1648#L57)
``` go
func (c *Core) ToCore() *types.Core
```
ToCore will convert *core.Core to *types.Core




### <a name="Core.UsingAssets">func</a> (Core) [UsingAssets](https://github.com/hunterlong/statping/tree/master/core/core.go?s=2857:2889#L105)
``` go
func (c Core) UsingAssets() bool
```
UsingAssets will return true if /assets folder is present




## <a name="DateScan">type</a> [DateScan](https://github.com/hunterlong/statping/tree/master/core/services.go?s=4539:4631#L168)
``` go
type DateScan struct {
    CreatedAt string `json:"x,omitempty"`
    Value     int64  `json:"y"`
}

```
DateScan struct is for creating the charts.js graph JSON array










## <a name="DateScanObj">type</a> [DateScanObj](https://github.com/hunterlong/statping/tree/master/core/services.go?s=4702:4761#L174)
``` go
type DateScanObj struct {
    Array []DateScan `json:"data"`
}

```
DateScanObj struct is for creating the charts.js graph JSON array







### <a name="GraphDataRaw">func</a> [GraphDataRaw](https://github.com/hunterlong/statping/tree/master/core/services.go?s=7369:7482#L262)
``` go
func GraphDataRaw(service types.ServiceInterface, start, end time.Time, group string, column string) *DateScanObj
```
GraphDataRaw will return all the hits between 2 times for a Service





### <a name="DateScanObj.ToString">func</a> (\*DateScanObj) [ToString](https://github.com/hunterlong/statping/tree/master/core/services.go?s=8567:8606#L299)
``` go
func (d *DateScanObj) ToString() string
```
ToString will convert the DateScanObj into a JSON string for the charts to render




## <a name="DbConfig">type</a> [DbConfig](https://github.com/hunterlong/statping/tree/master/core/database.go?s=1401:1429#L43)
``` go
type DbConfig types.DbConfig
```
DbConfig stores the config.yml file for the statup configuration







### <a name="EnvToConfig">func</a> [EnvToConfig](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=3615:3643#L129)
``` go
func EnvToConfig() *DbConfig
```
EnvToConfig converts environment variables to a DbConfig variable


### <a name="LoadConfigFile">func</a> [LoadConfigFile](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=1032:1088#L34)
``` go
func LoadConfigFile(directory string) (*DbConfig, error)
```
LoadConfigFile will attempt to load the 'config.yml' file in a specific directory


### <a name="LoadUsingEnv">func</a> [LoadUsingEnv](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=1696:1734#L53)
``` go
func LoadUsingEnv() (*DbConfig, error)
```
LoadUsingEnv will attempt to load database configs based on environment variables. If DB_CONN is set if will force this function.





### <a name="DbConfig.Connect">func</a> (\*DbConfig) [Connect](https://github.com/hunterlong/statping/tree/master/core/database.go?s=6920:6982#L235)
``` go
func (db *DbConfig) Connect(retry bool, location string) error
```
Connect will attempt to connect to the sqlite, postgres, or mysql database




### <a name="DbConfig.CreateCore">func</a> (\*DbConfig) [CreateCore](https://github.com/hunterlong/statping/tree/master/core/database.go?s=10187:10224#L340)
``` go
func (c *DbConfig) CreateCore() *Core
```
CreateCore will initialize the global variable 'CoreApp". This global variable contains most of Statping app.




### <a name="DbConfig.CreateDatabase">func</a> (\*DbConfig) [CreateDatabase](https://github.com/hunterlong/statping/tree/master/core/database.go?s=11288:11330#L377)
``` go
func (db *DbConfig) CreateDatabase() error
```
CreateDatabase will CREATE TABLES for each of the Statping elements




### <a name="DbConfig.DropDatabase">func</a> (\*DbConfig) [DropDatabase](https://github.com/hunterlong/statping/tree/master/core/database.go?s=10686:10726#L362)
``` go
func (db *DbConfig) DropDatabase() error
```
DropDatabase will DROP each table Statping created




### <a name="DbConfig.InsertCore">func</a> (\*DbConfig) [InsertCore](https://github.com/hunterlong/statping/tree/master/core/database.go?s=6428:6475#L219)
``` go
func (db *DbConfig) InsertCore() (*Core, error)
```
InsertCore create the single row for the Core settings in Statping




### <a name="DbConfig.MigrateDatabase">func</a> (\*DbConfig) [MigrateDatabase](https://github.com/hunterlong/statping/tree/master/core/database.go?s=11914:11957#L395)
``` go
func (db *DbConfig) MigrateDatabase() error
```
MigrateDatabase will migrate the database structure to current version.
This function will NOT remove previous records, tables or columns from the database.
If this function has an issue, it will ROLLBACK to the previous state.




### <a name="DbConfig.Save">func</a> (\*DbConfig) [Save](https://github.com/hunterlong/statping/tree/master/core/database.go?s=9657:9702#L320)
``` go
func (db *DbConfig) Save() (*DbConfig, error)
```
Save will initially create the config.yml file




### <a name="DbConfig.Update">func</a> (\*DbConfig) [Update](https://github.com/hunterlong/statping/tree/master/core/database.go?s=9294:9328#L302)
``` go
func (db *DbConfig) Update() error
```
Update will save the config.yml file




## <a name="ErrorResponse">type</a> [ErrorResponse](https://github.com/hunterlong/statping/tree/master/core/configs.go?s=902:945#L29)
``` go
type ErrorResponse struct {
    Error string
}

```
ErrorResponse is used for HTTP errors to show to User










## <a name="ExportData">type</a> [ExportData](https://github.com/hunterlong/statping/tree/master/core/export.go?s=1344:1688#L48)
``` go
type ExportData struct {
    Core      *types.Core              `json:"core"`
    Services  []types.ServiceInterface `json:"services"`
    Messages  []*Message               `json:"messages"`
    Checkins  []*Checkin               `json:"checkins"`
    Users     []*User                  `json:"users"`
    Notifiers []types.AllNotifiers     `json:"notifiers"`
}

```









## <a name="Hit">type</a> [Hit](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=790:821#L24)
``` go
type Hit struct {
    *types.Hit
}

```









### <a name="Hit.AfterFind">func</a> (\*Hit) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=3766:3803#L118)
``` go
func (h *Hit) AfterFind() (err error)
```
AfterFind for Hit will set the timezone




### <a name="Hit.BeforeCreate">func</a> (\*Hit) [BeforeCreate](https://github.com/hunterlong/statping/tree/master/core/database.go?s=5059:5099#L159)
``` go
func (h *Hit) BeforeCreate() (err error)
```
BeforeCreate for Hit will set CreatedAt to UTC




## <a name="Message">type</a> [Message](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=797:836#L25)
``` go
type Message struct {
    *types.Message
}

```






### <a name="ReturnMessage">func</a> [ReturnMessage](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=1117:1162#L37)
``` go
func ReturnMessage(m *types.Message) *Message
```
ReturnMessage will convert *types.Message to *core.Message


### <a name="SelectMessage">func</a> [SelectMessage](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=1437:1483#L49)
``` go
func SelectMessage(id int64) (*Message, error)
```
SelectMessage returns a Message based on the ID passed


### <a name="SelectMessages">func</a> [SelectMessages](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=1227:1268#L42)
``` go
func SelectMessages() ([]*Message, error)
```
SelectMessages returns all messages


### <a name="SelectServiceMessages">func</a> [SelectServiceMessages](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=898:945#L30)
``` go
func SelectServiceMessages(id int64) []*Message
```
SelectServiceMessages returns all messages for a service





### <a name="Message.AfterFind">func</a> (\*Message) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=4706:4747#L150)
``` go
func (u *Message) AfterFind() (err error)
```
AfterFind for Message will set the timezone




### <a name="Message.BeforeCreate">func</a> (\*Message) [BeforeCreate](https://github.com/hunterlong/statping/tree/master/core/database.go?s=5607:5651#L184)
``` go
func (u *Message) BeforeCreate() (err error)
```
BeforeCreate for Message will set CreatedAt to UTC




### <a name="Message.Create">func</a> (\*Message) [Create](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=1772:1813#L63)
``` go
func (m *Message) Create() (int64, error)
```
Create will create a Message and insert it into the database




### <a name="Message.Delete">func</a> (\*Message) [Delete](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=2083:2115#L74)
``` go
func (m *Message) Delete() error
```
Delete will delete a Message from database




### <a name="Message.Service">func</a> (\*Message) [Service](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=1592:1628#L55)
``` go
func (m *Message) Service() *Service
```



### <a name="Message.Update">func</a> (\*Message) [Update](https://github.com/hunterlong/statping/tree/master/core/messages.go?s=2216:2260#L80)
``` go
func (m *Message) Update() (*Message, error)
```
Update will update a Message in the database




## <a name="PluginJSON">type</a> [PluginJSON](https://github.com/hunterlong/statping/tree/master/core/core.go?s=909:941#L30)
``` go
type PluginJSON types.PluginJSON
```









## <a name="PluginRepos">type</a> [PluginRepos](https://github.com/hunterlong/statping/tree/master/core/core.go?s=942:976#L31)
``` go
type PluginRepos types.PluginRepos
```









## <a name="Service">type</a> [Service](https://github.com/hunterlong/statping/tree/master/core/services.go?s=910:949#L30)
``` go
type Service struct {
    *types.Service
}

```






### <a name="ReturnService">func</a> [ReturnService](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1138:1183#L40)
``` go
func ReturnService(s *types.Service) *Service
```
ReturnService will convert *types.Service to *core.Service


### <a name="SelectService">func</a> [SelectService](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1336:1373#L49)
``` go
func SelectService(id int64) *Service
```
SelectService returns a *core.Service from in memory





### <a name="Service.ActiveMessages">func</a> (\*Service) [ActiveMessages](https://github.com/hunterlong/statping/tree/master/core/services.go?s=12215:12260#L438)
``` go
func (s *Service) ActiveMessages() []*Message
```
ActiveMessages returns all Messages for a Service




### <a name="Service.AfterFind">func</a> (\*Service) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=3544:3585#L111)
``` go
func (s *Service) AfterFind() (err error)
```
AfterFind for Service will set the timezone




### <a name="Service.AllFailures">func</a> (\*Service) [AllFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=1443:1485#L54)
``` go
func (s *Service) AllFailures() []*failure
```
AllFailures will return all failures attached to a service




### <a name="Service.AvgTime">func</a> (\*Service) [AvgTime](https://github.com/hunterlong/statping/tree/master/core/services.go?s=3460:3495#L127)
``` go
func (s *Service) AvgTime() float64
```
AvgTime will return the average amount of time for a service to response back successfully




### <a name="Service.AvgUptime">func</a> (\*Service) [AvgUptime](https://github.com/hunterlong/statping/tree/master/core/services.go?s=8968:9017#L315)
``` go
func (s *Service) AvgUptime(ago time.Time) string
```
AvgUptime returns average online status for last 24 hours




### <a name="Service.AvgUptime24">func</a> (\*Service) [AvgUptime24](https://github.com/hunterlong/statping/tree/master/core/services.go?s=8798:8836#L309)
``` go
func (s *Service) AvgUptime24() string
```
AvgUptime24 returns a service's average online status for last 24 hours




### <a name="Service.BeforeCreate">func</a> (\*Service) [BeforeCreate](https://github.com/hunterlong/statping/tree/master/core/database.go?s=5815:5859#L193)
``` go
func (s *Service) BeforeCreate() (err error)
```
BeforeCreate for Service will set CreatedAt to UTC




### <a name="Service.Check">func</a> (\*Service) [Check](https://github.com/hunterlong/statping/tree/master/core/checker.go?s=5440:5476#L215)
``` go
func (s *Service) Check(record bool)
```
Check will run checkHttp for HTTP services and checkTcp for TCP services




### <a name="Service.CheckQueue">func</a> (\*Service) [CheckQueue](https://github.com/hunterlong/statping/tree/master/core/checker.go?s=1253:1294#L42)
``` go
func (s *Service) CheckQueue(record bool)
```
CheckQueue is the main go routine for checking a service




### <a name="Service.CheckinProcess">func</a> (\*Service) [CheckinProcess](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1774:1808#L69)
``` go
func (s *Service) CheckinProcess()
```
CheckinProcess runs the checkin routine for each checkin attached to service




### <a name="Service.Checkins">func</a> (\*Service) [Checkins](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1960:1999#L78)
``` go
func (s *Service) Checkins() []*Checkin
```
Checkins will return a slice of Checkins for a Service




### <a name="Service.CountHits">func</a> (\*Service) [CountHits](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=1154:1198#L39)
``` go
func (s *Service) CountHits() (int64, error)
```
CountHits returns a int64 for all hits for a service




### <a name="Service.Create">func</a> (\*Service) [Create](https://github.com/hunterlong/statping/tree/master/core/services.go?s=11608:11659#L416)
``` go
func (s *Service) Create(check bool) (int64, error)
```
Create will create a service and insert it into the database




### <a name="Service.CreateFailure">func</a> (\*Service) [CreateFailure](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=983:1058#L37)
``` go
func (s *Service) CreateFailure(fail types.FailureInterface) (int64, error)
```
CreateFailure will create a new failure record for a service




### <a name="Service.CreateHit">func</a> (\*Service) [CreateHit](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=915:971#L29)
``` go
func (s *Service) CreateHit(h *types.Hit) (int64, error)
```
CreateHit will create a new 'hit' record in the database for a successful/online service




### <a name="Service.Delete">func</a> (\*Service) [Delete](https://github.com/hunterlong/statping/tree/master/core/services.go?s=10255:10287#L369)
``` go
func (s *Service) Delete() error
```
Delete will remove a service from the database, it will also end the service checking go routine




### <a name="Service.DeleteFailures">func</a> (\*Service) [DeleteFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=1825:1859#L66)
``` go
func (s *Service) DeleteFailures()
```
DeleteFailures will delete all failures for a service




### <a name="Service.Downtime">func</a> (\*Service) [Downtime](https://github.com/hunterlong/statping/tree/master/core/services.go?s=7010:7052#L248)
``` go
func (s *Service) Downtime() time.Duration
```
Downtime returns the amount of time of a offline service




### <a name="Service.DowntimeText">func</a> (\*Service) [DowntimeText](https://github.com/hunterlong/statping/tree/master/core/services.go?s=5994:6033#L214)
``` go
func (s *Service) DowntimeText() string
```
DowntimeText will return the amount of downtime for a service based on the duration


	service.DowntimeText()
	// Service has been offline for 15 minutes




### <a name="Service.Hits">func</a> (\*Service) [Hits](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=1363:1409#L47)
``` go
func (s *Service) Hits() ([]*types.Hit, error)
```
Hits returns all successful hits for a service




### <a name="Service.HitsBetween">func</a> (\*Service) [HitsBetween](https://github.com/hunterlong/statping/tree/master/core/database.go?s=2533:2618#L86)
``` go
func (s *Service) HitsBetween(t1, t2 time.Time, group string, column string) *gorm.DB
```
HitsBetween returns the gorm database query for a collection of service hits between a time range




### <a name="Service.LimitedCheckinFailures">func</a> (\*Service) [LimitedCheckinFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=2415:2480#L82)
``` go
func (s *Service) LimitedCheckinFailures(amount int64) []*failure
```
LimitedFailures will return the last amount of failures from a service




### <a name="Service.LimitedCheckins">func</a> (\*Service) [LimitedCheckins](https://github.com/hunterlong/statping/tree/master/core/services.go?s=2165:2211#L85)
``` go
func (s *Service) LimitedCheckins() []*Checkin
```
LimitedCheckins will return a slice of Checkins for a Service




### <a name="Service.LimitedFailures">func</a> (\*Service) [LimitedFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=2124:2182#L75)
``` go
func (s *Service) LimitedFailures(amount int64) []*failure
```
LimitedFailures will return the last amount of failures from a service




### <a name="Service.LimitedHits">func</a> (\*Service) [LimitedHits](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=1630:1683#L55)
``` go
func (s *Service) LimitedHits() ([]*types.Hit, error)
```
LimitedHits returns the last 1024 successful/online 'hit' records for a service




### <a name="Service.Messages">func</a> (\*Service) [Messages](https://github.com/hunterlong/statping/tree/master/core/services.go?s=12059:12098#L432)
``` go
func (s *Service) Messages() []*Message
```
Messages returns all Messages for a Service




### <a name="Service.Online24">func</a> (\*Service) [Online24](https://github.com/hunterlong/statping/tree/master/core/services.go?s=3786:3822#L140)
``` go
func (s *Service) Online24() float32
```
Online24 returns the service's uptime percent within last 24 hours




### <a name="Service.OnlineSince">func</a> (\*Service) [OnlineSince](https://github.com/hunterlong/statping/tree/master/core/services.go?s=3986:4038#L146)
``` go
func (s *Service) OnlineSince(ago time.Time) float32
```
OnlineSince accepts a time since parameter to return the percent of a service's uptime.




### <a name="Service.Select">func</a> (\*Service) [Select](https://github.com/hunterlong/statping/tree/master/core/services.go?s=1011:1052#L35)
``` go
func (s *Service) Select() *types.Service
```
Select will return the *types.Service struct for Service




### <a name="Service.SmallText">func</a> (\*Service) [SmallText](https://github.com/hunterlong/statping/tree/master/core/services.go?s=5121:5157#L191)
``` go
func (s *Service) SmallText() string
```
SmallText returns a short description about a services status


	service.SmallText()
	// Online since Monday 3:04:05PM, Jan _2 2006




### <a name="Service.Sum">func</a> (\*Service) [Sum](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=2682:2722#L87)
``` go
func (s *Service) Sum() (float64, error)
```
Sum returns the added value Latency for all of the services successful hits.




### <a name="Service.ToJSON">func</a> (\*Service) [ToJSON](https://github.com/hunterlong/statping/tree/master/core/services.go?s=3278:3311#L121)
``` go
func (s *Service) ToJSON() string
```
ToJSON will convert a service to a JSON string




### <a name="Service.TotalFailures">func</a> (\*Service) [TotalFailures](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=3860:3909#L134)
``` go
func (s *Service) TotalFailures() (uint64, error)
```
TotalFailures returns the total amount of failures for a service




### <a name="Service.TotalFailures24">func</a> (\*Service) [TotalFailures24](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=3661:3712#L128)
``` go
func (s *Service) TotalFailures24() (uint64, error)
```
TotalFailures24 returns the amount of failures for a service within the last 24 hours




### <a name="Service.TotalFailuresSince">func</a> (\*Service) [TotalFailuresSince](https://github.com/hunterlong/statping/tree/master/core/failures.go?s=4134:4201#L142)
``` go
func (s *Service) TotalFailuresSince(ago time.Time) (uint64, error)
```
TotalFailuresSince returns the total amount of failures for a service since a specific time/date




### <a name="Service.TotalHits">func</a> (\*Service) [TotalHits](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=2113:2158#L71)
``` go
func (s *Service) TotalHits() (uint64, error)
```
TotalHits returns the total amount of successful hits a service has




### <a name="Service.TotalHitsSince">func</a> (\*Service) [TotalHitsSince](https://github.com/hunterlong/statping/tree/master/core/hits.go?s=2358:2421#L79)
``` go
func (s *Service) TotalHitsSince(ago time.Time) (uint64, error)
```
TotalHitsSince returns the total amount of hits based on a specific time/date




### <a name="Service.TotalUptime">func</a> (\*Service) [TotalUptime](https://github.com/hunterlong/statping/tree/master/core/services.go?s=9443:9481#L337)
``` go
func (s *Service) TotalUptime() string
```
TotalUptime returns the total uptime percent of a service




### <a name="Service.Update">func</a> (\*Service) [Update](https://github.com/hunterlong/statping/tree/master/core/services.go?s=10917:10961#L390)
``` go
func (s *Service) Update(restart bool) error
```
Update will update a service in the database, the service's checking routine can be restarted by passing true




### <a name="Service.UpdateSingle">func</a> (\*Service) [UpdateSingle](https://github.com/hunterlong/statping/tree/master/core/services.go?s=10692:10749#L385)
``` go
func (s *Service) UpdateSingle(attr ...interface{}) error
```
UpdateSingle will update a single column for a service




## <a name="ServiceOrder">type</a> [ServiceOrder](https://github.com/hunterlong/statping/tree/master/core/core.go?s=5032:5074#L181)
``` go
type ServiceOrder []types.ServiceInterface
```
ServiceOrder will reorder the services based on 'order_id' (Order)










### <a name="ServiceOrder.Len">func</a> (ServiceOrder) [Len](https://github.com/hunterlong/statping/tree/master/core/core.go?s=5130:5161#L184)
``` go
func (c ServiceOrder) Len() int
```
Sort interface for resroting the Services in order




### <a name="ServiceOrder.Less">func</a> (ServiceOrder) [Less](https://github.com/hunterlong/statping/tree/master/core/core.go?s=5260:5301#L186)
``` go
func (c ServiceOrder) Less(i, j int) bool
```



### <a name="ServiceOrder.Swap">func</a> (ServiceOrder) [Swap](https://github.com/hunterlong/statping/tree/master/core/core.go?s=5190:5226#L185)
``` go
func (c ServiceOrder) Swap(i, j int)
```



## <a name="User">type</a> [User](https://github.com/hunterlong/statping/tree/master/core/users.go?s=827:860#L26)
``` go
type User struct {
    *types.User
}

```






### <a name="AuthUser">func</a> [AuthUser](https://github.com/hunterlong/statping/tree/master/core/users.go?s=2710:2764#L99)
``` go
func AuthUser(username, password string) (*User, bool)
```
AuthUser will return the User and a boolean if authentication was correct.
AuthUser accepts username, and password as a string


### <a name="ReturnUser">func</a> [ReturnUser](https://github.com/hunterlong/statping/tree/master/core/users.go?s=919:955#L31)
``` go
func ReturnUser(u *types.User) *User
```
ReturnUser returns *core.User based off a *types.User


### <a name="SelectAllUsers">func</a> [SelectAllUsers](https://github.com/hunterlong/statping/tree/master/core/users.go?s=2342:2380#L87)
``` go
func SelectAllUsers() ([]*User, error)
```
SelectAllUsers returns all users


### <a name="SelectUser">func</a> [SelectUser](https://github.com/hunterlong/statping/tree/master/core/users.go?s=1163:1203#L43)
``` go
func SelectUser(id int64) (*User, error)
```
SelectUser returns the User based on the User's ID.


### <a name="SelectUsername">func</a> [SelectUsername](https://github.com/hunterlong/statping/tree/master/core/users.go?s=1364:1415#L50)
``` go
func SelectUsername(username string) (*User, error)
```
SelectUsername returns the User based on the User's username





### <a name="User.AfterFind">func</a> (\*User) [AfterFind](https://github.com/hunterlong/statping/tree/master/core/database.go?s=4087:4125#L130)
``` go
func (u *User) AfterFind() (err error)
```
AfterFind for USer will set the timezone




### <a name="User.BeforeCreate">func</a> (\*User) [BeforeCreate](https://github.com/hunterlong/statping/tree/master/core/database.go?s=5402:5443#L175)
``` go
func (u *User) BeforeCreate() (err error)
```
BeforeCreate for User will set CreatedAt to UTC




### <a name="User.Create">func</a> (\*User) [Create](https://github.com/hunterlong/statping/tree/master/core/users.go?s=1903:1941#L70)
``` go
func (u *User) Create() (int64, error)
```
Create will insert a new User into the database




### <a name="User.Delete">func</a> (\*User) [Delete](https://github.com/hunterlong/statping/tree/master/core/users.go?s=1592:1621#L58)
``` go
func (u *User) Delete() error
```
Delete will remove the User record from the database




### <a name="User.Update">func</a> (\*User) [Update](https://github.com/hunterlong/statping/tree/master/core/users.go?s=1713:1742#L63)
``` go
func (u *User) Update() error
```
Update will update the User's record in database










# handlers
`import "github.com/hunterlong/statping/handlers"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package handlers contains the HTTP server along with the requests and routes. All HTTP related
functions are in this package.

More info on: <a href="https://github.com/hunterlong/statping">https://github.com/hunterlong/statping</a>




## <a name="pkg-index">Index</a>
* [func DesktopInit(ip string, port int)](#DesktopInit)
* [func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}, redirect interface{})](#ExecuteResponse)
* [func IsAuthenticated(r *http.Request) bool](#IsAuthenticated)
* [func Router() *mux.Router](#Router)
* [func RunHTTPServer(ip string, port int) error](#RunHTTPServer)
* [type Cacher](#Cacher)
* [type Item](#Item)
  * [func (item Item) Expired() bool](#Item.Expired)
* [type PluginSelect](#PluginSelect)
* [type Storage](#Storage)
  * [func NewStorage() *Storage](#NewStorage)
  * [func (s Storage) Delete(key string)](#Storage.Delete)
  * [func (s Storage) Get(key string) []byte](#Storage.Get)
  * [func (s Storage) Set(key string, content []byte, duration time.Duration)](#Storage.Set)


#### <a name="pkg-files">Package files</a>
[api.go](https://github.com/hunterlong/statping/tree/master/handlers/api.go) [cache.go](https://github.com/hunterlong/statping/tree/master/handlers/cache.go) [checkin.go](https://github.com/hunterlong/statping/tree/master/handlers/checkin.go) [dashboard.go](https://github.com/hunterlong/statping/tree/master/handlers/dashboard.go) [doc.go](https://github.com/hunterlong/statping/tree/master/handlers/doc.go) [handlers.go](https://github.com/hunterlong/statping/tree/master/handlers/handlers.go) [index.go](https://github.com/hunterlong/statping/tree/master/handlers/index.go) [messages.go](https://github.com/hunterlong/statping/tree/master/handlers/messages.go) [notifications.go](https://github.com/hunterlong/statping/tree/master/handlers/notifications.go) [plugins.go](https://github.com/hunterlong/statping/tree/master/handlers/plugins.go) [prometheus.go](https://github.com/hunterlong/statping/tree/master/handlers/prometheus.go) [routes.go](https://github.com/hunterlong/statping/tree/master/handlers/routes.go) [services.go](https://github.com/hunterlong/statping/tree/master/handlers/services.go) [settings.go](https://github.com/hunterlong/statping/tree/master/handlers/settings.go) [setup.go](https://github.com/hunterlong/statping/tree/master/handlers/setup.go) [users.go](https://github.com/hunterlong/statping/tree/master/handlers/users.go) 





## <a name="DesktopInit">func</a> [DesktopInit](https://github.com/hunterlong/statping/tree/master/handlers/index.go?s=1538:1575#L48)
``` go
func DesktopInit(ip string, port int)
```
DesktopInit will run the Statping server on a specific IP and port using SQLite database



## <a name="ExecuteResponse">func</a> [ExecuteResponse](https://github.com/hunterlong/statping/tree/master/handlers/handlers.go?s=5257:5370#L194)
``` go
func ExecuteResponse(w http.ResponseWriter, r *http.Request, file string, data interface{}, redirect interface{})
```
ExecuteResponse will render a HTTP response for the front end user



## <a name="IsAuthenticated">func</a> [IsAuthenticated](https://github.com/hunterlong/statping/tree/master/handlers/handlers.go?s=2076:2118#L69)
``` go
func IsAuthenticated(r *http.Request) bool
```
IsAuthenticated returns true if the HTTP request is authenticated. You can set the environment variable GO_ENV=test
to bypass the admin authenticate to the dashboard features.



## <a name="Router">func</a> [Router](https://github.com/hunterlong/statping/tree/master/handlers/routes.go?s=992:1017#L34)
``` go
func Router() *mux.Router
```
Router returns all of the routes used in Statping



## <a name="RunHTTPServer">func</a> [RunHTTPServer](https://github.com/hunterlong/statping/tree/master/handlers/handlers.go?s=1166:1211#L44)
``` go
func RunHTTPServer(ip string, port int) error
```
RunHTTPServer will start a HTTP server on a specific IP and port




## <a name="Cacher">type</a> [Cacher](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=143:269#L13)
``` go
type Cacher interface {
    Get(key string) []byte
    Delete(key string)
    Set(key string, content []byte, duration time.Duration)
}
```

``` go
var CacheStorage Cacher
```









## <a name="Item">type</a> [Item](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=301:358#L20)
``` go
type Item struct {
    Content    []byte
    Expiration int64
}

```
Item is a cached reference










### <a name="Item.Expired">func</a> (Item) [Expired](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=409:440#L26)
``` go
func (item Item) Expired() bool
```
Expired returns true if the item has expired.




## <a name="PluginSelect">type</a> [PluginSelect](https://github.com/hunterlong/statping/tree/master/handlers/plugins.go?s=729:818#L23)
``` go
type PluginSelect struct {
    Plugin string
    Form   string
    Params map[string]interface{}
}

```









## <a name="Storage">type</a> [Storage](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=588:655#L34)
``` go
type Storage struct {
    // contains filtered or unexported fields
}

```
Storage mecanism for caching strings in memory







### <a name="NewStorage">func</a> [NewStorage](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=707:733#L40)
``` go
func NewStorage() *Storage
```
NewStorage creates a new in memory CacheStorage





### <a name="Storage.Delete">func</a> (Storage) [Delete](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=1033:1068#L60)
``` go
func (s Storage) Delete(key string)
```



### <a name="Storage.Get">func</a> (Storage) [Get](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=848:887#L48)
``` go
func (s Storage) Get(key string) []byte
```
Get a cached content by key




### <a name="Storage.Set">func</a> (Storage) [Set](https://github.com/hunterlong/statping/tree/master/handlers/cache.go?s=1162:1234#L67)
``` go
func (s Storage) Set(key string, content []byte, duration time.Duration)
```
Set a cached content by key










# notifiers
`import "github.com/hunterlong/statping/notifiers"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package notifiers holds all the notifiers for Statping, which also includes
user created notifiers that have been accepted in a Push Request. Read the wiki
to see a full example of a notifier with all events, visit Statping's
notifier example code: <a href="https://github.com/hunterlong/statping/wiki/Notifier-Example">https://github.com/hunterlong/statping/wiki/Notifier-Example</a>

This package shouldn't contain any exports, to see how notifiers work
visit the core/notifier package at: <a href="https://godoc.org/github.com/hunterlong/statping/core/notifier">https://godoc.org/github.com/hunterlong/statping/core/notifier</a>
and learn how to create your own custom notifier.




## <a name="pkg-index">Index</a>


#### <a name="pkg-files">Package files</a>
[command.go](https://github.com/hunterlong/statping/tree/master/notifiers/command.go) [discord.go](https://github.com/hunterlong/statping/tree/master/notifiers/discord.go) [doc.go](https://github.com/hunterlong/statping/tree/master/notifiers/doc.go) [email.go](https://github.com/hunterlong/statping/tree/master/notifiers/email.go) [line_notify.go](https://github.com/hunterlong/statping/tree/master/notifiers/line_notify.go) [mobile.go](https://github.com/hunterlong/statping/tree/master/notifiers/mobile.go) [slack.go](https://github.com/hunterlong/statping/tree/master/notifiers/slack.go) [twilio.go](https://github.com/hunterlong/statping/tree/master/notifiers/twilio.go) [webhook.go](https://github.com/hunterlong/statping/tree/master/notifiers/webhook.go) 












# plugin
`import "github.com/hunterlong/statping/plugin"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package plugin contains the interfaces to build your own Golang Plugin that will receive triggers on Statping events.




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func LoadPlugin(file string) error](#LoadPlugin)
* [func LoadPlugins()](#LoadPlugins)


#### <a name="pkg-files">Package files</a>
[doc.go](https://github.com/hunterlong/statping/tree/master/plugin/doc.go) [plugin.go](https://github.com/hunterlong/statping/tree/master/plugin/plugin.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    AllPlugins []*types.PluginObject
)
```


## <a name="LoadPlugin">func</a> [LoadPlugin](https://github.com/hunterlong/statping/tree/master/plugin/plugin.go?s=1188:1222#L51)
``` go
func LoadPlugin(file string) error
```


## <a name="LoadPlugins">func</a> [LoadPlugins](https://github.com/hunterlong/statping/tree/master/plugin/plugin.go?s=2685:2703#L96)
``` go
func LoadPlugins()
```









# source
`import "github.com/hunterlong/statping/source"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package source holds all the assets for Statping. This includes
CSS, JS, SCSS, HTML and other website related content.
This package uses Rice to compile all assets into a single 'rice-box.go' file.

### Required Dependencies
- rice -> <a href="https://github.com/GeertJohan/go.rice">https://github.com/GeertJohan/go.rice</a>
- sass -> <a href="https://sass-lang.com/install">https://sass-lang.com/install</a>

### Compile Assets
To compile all the HTML, JS, SCSS, CSS and image assets you'll need to have rice and sass installed on your local system.


	sass source/scss/base.scss source/css/base.css
	cd source && rice embed-go

More info on: <a href="https://github.com/hunterlong/statping">https://github.com/hunterlong/statping</a>




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func Assets()](#Assets)
* [func CompileSASS(folder string) error](#CompileSASS)
* [func CopyAllToPublic(box *rice.Box, folder string) error](#CopyAllToPublic)
* [func CopyToPublic(box *rice.Box, folder, file string) error](#CopyToPublic)
* [func CreateAllAssets(folder string) error](#CreateAllAssets)
* [func DeleteAllAssets(folder string) error](#DeleteAllAssets)
* [func HelpMarkdown() string](#HelpMarkdown)
* [func MakePublicFolder(folder string) error](#MakePublicFolder)
* [func OpenAsset(folder, file string) string](#OpenAsset)
* [func SaveAsset(data []byte, folder, file string) error](#SaveAsset)
* [func UsingAssets(folder string) bool](#UsingAssets)

#### <a name="pkg-examples">Examples</a>
* [OpenAsset](#example_OpenAsset)
* [SaveAsset](#example_SaveAsset)

#### <a name="pkg-files">Package files</a>
[doc.go](https://github.com/hunterlong/statping/tree/master/source/doc.go) [rice-box.go](https://github.com/hunterlong/statping/tree/master/source/rice-box.go) [source.go](https://github.com/hunterlong/statping/tree/master/source/source.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    CssBox  *rice.Box // CSS files from the 'source/css' directory, this will be loaded into '/assets/css'
    ScssBox *rice.Box // SCSS files from the 'source/scss' directory, this will be loaded into '/assets/scss'
    JsBox   *rice.Box // JS files from the 'source/js' directory, this will be loaded into '/assets/js'
    TmplBox *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
    FontBox *rice.Box // HTML and other small files from the 'source/tmpl' directory, this will be loaded into '/assets'
)
```


## <a name="Assets">func</a> [Assets](https://github.com/hunterlong/statping/tree/master/source/source.go?s=1487:1500#L37)
``` go
func Assets()
```
Assets will load the Rice boxes containing the CSS, SCSS, JS, and HTML files.



## <a name="CompileSASS">func</a> [CompileSASS](https://github.com/hunterlong/statping/tree/master/source/source.go?s=2015:2052#L57)
``` go
func CompileSASS(folder string) error
```
CompileSASS will attempt to compile the SASS files into CSS



## <a name="CopyAllToPublic">func</a> [CopyAllToPublic](https://github.com/hunterlong/statping/tree/master/source/source.go?s=5779:5835#L167)
``` go
func CopyAllToPublic(box *rice.Box, folder string) error
```
CopyAllToPublic will copy all the files in a rice box into a local folder



## <a name="CopyToPublic">func</a> [CopyToPublic](https://github.com/hunterlong/statping/tree/master/source/source.go?s=6377:6436#L189)
``` go
func CopyToPublic(box *rice.Box, folder, file string) error
```
CopyToPublic will create a file from a rice Box to the '/assets' directory



## <a name="CreateAllAssets">func</a> [CreateAllAssets](https://github.com/hunterlong/statping/tree/master/source/source.go?s=4239:4280#L129)
``` go
func CreateAllAssets(folder string) error
```
CreateAllAssets will dump HTML, CSS, SCSS, and JS assets into the '/assets' directory



## <a name="DeleteAllAssets">func</a> [DeleteAllAssets](https://github.com/hunterlong/statping/tree/master/source/source.go?s=5434:5475#L156)
``` go
func DeleteAllAssets(folder string) error
```
DeleteAllAssets will delete the '/assets' folder



## <a name="HelpMarkdown">func</a> [HelpMarkdown](https://github.com/hunterlong/statping/tree/master/source/source.go?s=1742:1768#L46)
``` go
func HelpMarkdown() string
```
HelpMarkdown will return the Markdown of help.md into HTML



## <a name="MakePublicFolder">func</a> [MakePublicFolder](https://github.com/hunterlong/statping/tree/master/source/source.go?s=6944:6986#L206)
``` go
func MakePublicFolder(folder string) error
```
MakePublicFolder will create a new folder



## <a name="OpenAsset">func</a> [OpenAsset](https://github.com/hunterlong/statping/tree/master/source/source.go?s=3929:3971#L119)
``` go
func OpenAsset(folder, file string) string
```
OpenAsset returns a file's contents as a string



#### <a name="example_OpenAsset">Example</a>

Code:
``` go
OpenAsset("js", "main.js")
```

## <a name="SaveAsset">func</a> [SaveAsset](https://github.com/hunterlong/statping/tree/master/source/source.go?s=3556:3610#L108)
``` go
func SaveAsset(data []byte, folder, file string) error
```
SaveAsset will save an asset to the '/assets/' folder.



#### <a name="example_SaveAsset">Example</a>

Code:
``` go
data := []byte("alert('helloooo')")
SaveAsset(data, "js", "test.js")
```

## <a name="UsingAssets">func</a> [UsingAssets](https://github.com/hunterlong/statping/tree/master/source/source.go?s=3011:3047#L88)
``` go
func UsingAssets(folder string) bool
```
UsingAssets returns true if the '/assets' folder is found in the directory










# types
`import "github.com/hunterlong/statping/types"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package types contains all of the structs for objects in Statping including services, hits, failures, Core, and others.

More info on: <a href="https://github.com/hunterlong/statping">https://github.com/hunterlong/statping</a>




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [type AllNotifiers](#AllNotifiers)
* [type Asseter](#Asseter)
* [type Checkin](#Checkin)
  * [func (s *Checkin) Close()](#Checkin.Close)
  * [func (s *Checkin) IsRunning() bool](#Checkin.IsRunning)
  * [func (s *Checkin) Start()](#Checkin.Start)
* [type CheckinHit](#CheckinHit)
* [type Core](#Core)
* [type Databaser](#Databaser)
* [type DbConfig](#DbConfig)
* [type FailSort](#FailSort)
  * [func (s FailSort) Len() int](#FailSort.Len)
  * [func (s FailSort) Less(i, j int) bool](#FailSort.Less)
  * [func (s FailSort) Swap(i, j int)](#FailSort.Swap)
* [type Failure](#Failure)
* [type FailureInterface](#FailureInterface)
* [type Hit](#Hit)
* [type Info](#Info)
* [type Message](#Message)
* [type NullBool](#NullBool)
  * [func NewNullBool(s bool) NullBool](#NewNullBool)
  * [func (nb *NullBool) MarshalJSON() ([]byte, error)](#NullBool.MarshalJSON)
  * [func (nf *NullBool) UnmarshalJSON(b []byte) error](#NullBool.UnmarshalJSON)
* [type NullFloat64](#NullFloat64)
  * [func NewNullFloat64(s float64) NullFloat64](#NewNullFloat64)
  * [func (ni *NullFloat64) MarshalJSON() ([]byte, error)](#NullFloat64.MarshalJSON)
  * [func (nf *NullFloat64) UnmarshalJSON(b []byte) error](#NullFloat64.UnmarshalJSON)
* [type NullInt64](#NullInt64)
  * [func NewNullInt64(s int64) NullInt64](#NewNullInt64)
  * [func (ni *NullInt64) MarshalJSON() ([]byte, error)](#NullInt64.MarshalJSON)
  * [func (nf *NullInt64) UnmarshalJSON(b []byte) error](#NullInt64.UnmarshalJSON)
* [type NullString](#NullString)
  * [func NewNullString(s string) NullString](#NewNullString)
  * [func (ns *NullString) MarshalJSON() ([]byte, error)](#NullString.MarshalJSON)
  * [func (nf *NullString) UnmarshalJSON(b []byte) error](#NullString.UnmarshalJSON)
* [type Plugin](#Plugin)
* [type PluginActions](#PluginActions)
* [type PluginInfo](#PluginInfo)
* [type PluginJSON](#PluginJSON)
* [type PluginObject](#PluginObject)
* [type PluginRepos](#PluginRepos)
* [type PluginRoute](#PluginRoute)
* [type PluginRouting](#PluginRouting)
* [type Pluginer](#Pluginer)
* [type Router](#Router)
* [type Service](#Service)
  * [func (s *Service) Close()](#Service.Close)
  * [func (s *Service) IsRunning() bool](#Service.IsRunning)
  * [func (s *Service) Start()](#Service.Start)
* [type ServiceInterface](#ServiceInterface)
* [type User](#User)
* [type UserInterface](#UserInterface)


#### <a name="pkg-files">Package files</a>
[checkin.go](https://github.com/hunterlong/statping/tree/master/types/checkin.go) [core.go](https://github.com/hunterlong/statping/tree/master/types/core.go) [doc.go](https://github.com/hunterlong/statping/tree/master/types/doc.go) [failure.go](https://github.com/hunterlong/statping/tree/master/types/failure.go) [message.go](https://github.com/hunterlong/statping/tree/master/types/message.go) [null.go](https://github.com/hunterlong/statping/tree/master/types/null.go) [plugin.go](https://github.com/hunterlong/statping/tree/master/types/plugin.go) [service.go](https://github.com/hunterlong/statping/tree/master/types/service.go) [time.go](https://github.com/hunterlong/statping/tree/master/types/time.go) [types.go](https://github.com/hunterlong/statping/tree/master/types/types.go) [user.go](https://github.com/hunterlong/statping/tree/master/types/user.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    TIME_NANO     = "2006-01-02T15:04:05Z"
    TIME          = "2006-01-02 15:04:05"
    POSTGRES_TIME = "2006-01-02 15:04"
    CHART_TIME    = "2006-01-02T15:04:05.999999-07:00"
    TIME_DAY      = "2006-01-02"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    NOW = func() time.Time { return time.Now() }()
)
```



## <a name="AllNotifiers">type</a> [AllNotifiers](https://github.com/hunterlong/statping/tree/master/types/core.go?s=761:790#L23)
``` go
type AllNotifiers interface{}
```
AllNotifiers contains all the Notifiers loaded










## <a name="Asseter">type</a> [Asseter](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=966:1023#L64)
``` go
type Asseter interface {
    Asset(string) ([]byte, error)
}
```









## <a name="Checkin">type</a> [Checkin](https://github.com/hunterlong/statping/tree/master/types/checkin.go?s=815:1548#L23)
``` go
type Checkin struct {
    Id          int64         `gorm:"primary_key;column:id" json:"id"`
    ServiceId   int64         `gorm:"index;column:service" json:"service_id"`
    Name        string        `gorm:"column:name"  json:"name"`
    Interval    int64         `gorm:"column:check_interval" json:"interval"`
    GracePeriod int64         `gorm:"column:grace_period"  json:"grace"`
    ApiKey      string        `gorm:"column:api_key"  json:"api_key"`
    CreatedAt   time.Time     `gorm:"column:created_at" json:"created_at"`
    UpdatedAt   time.Time     `gorm:"column:updated_at" json:"updated_at"`
    Running     chan bool     `gorm:"-" json:"-"`
    Hits        []*CheckinHit `gorm:"-" json:"hits"`
    Failures    []*Failure    `gorm:"-" json:"failures"`
}

```
Checkin struct will allow an application to send a recurring HTTP GET to confirm a service is online










### <a name="Checkin.Close">func</a> (\*Checkin) [Close](https://github.com/hunterlong/statping/tree/master/types/checkin.go?s=2055:2080#L51)
``` go
func (s *Checkin) Close()
```
Close will stop the checkin routine




### <a name="Checkin.IsRunning">func</a> (\*Checkin) [IsRunning](https://github.com/hunterlong/statping/tree/master/types/checkin.go?s=2191:2225#L58)
``` go
func (s *Checkin) IsRunning() bool
```
IsRunning returns true if the checkin go routine is running




### <a name="Checkin.Start">func</a> (\*Checkin) [Start](https://github.com/hunterlong/statping/tree/master/types/checkin.go?s=1956:1981#L46)
``` go
func (s *Checkin) Start()
```
Start will create a channel for the checkin checking go routine




## <a name="CheckinHit">type</a> [CheckinHit](https://github.com/hunterlong/statping/tree/master/types/checkin.go?s=1604:1887#L38)
``` go
type CheckinHit struct {
    Id        int64     `gorm:"primary_key;column:id" json:"id"`
    Checkin   int64     `gorm:"index;column:checkin" json:"checkin"`
    From      string    `gorm:"column:from_location" json:"from"`
    CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

```
CheckinHit is a successful response from a Checkin










## <a name="Core">type</a> [Core](https://github.com/hunterlong/statping/tree/master/types/core.go?s=1046:2599#L28)
``` go
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
Core struct contains all the required fields for Statping. All application settings
will be saved into 1 row in the 'core' table. You can use the core.CoreApp
global variable to interact with the attributes to the application, such as services.










## <a name="Databaser">type</a> [Databaser](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=806:862#L55)
``` go
type Databaser interface {
    StatpingDatabase(*gorm.DB)
}
```









## <a name="DbConfig">type</a> [DbConfig](https://github.com/hunterlong/statping/tree/master/types/types.go?s=1132:1737#L32)
``` go
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
DbConfig struct is used for the database connection and creates the 'config.yml' file










## <a name="FailSort">type</a> [FailSort](https://github.com/hunterlong/statping/tree/master/types/failure.go?s=1544:1576#L40)
``` go
type FailSort []FailureInterface
```









### <a name="FailSort.Len">func</a> (FailSort) [Len](https://github.com/hunterlong/statping/tree/master/types/failure.go?s=1578:1605#L42)
``` go
func (s FailSort) Len() int
```



### <a name="FailSort.Less">func</a> (FailSort) [Less](https://github.com/hunterlong/statping/tree/master/types/failure.go?s=1687:1724#L48)
``` go
func (s FailSort) Less(i, j int) bool
```



### <a name="FailSort.Swap">func</a> (FailSort) [Swap](https://github.com/hunterlong/statping/tree/master/types/failure.go?s=1625:1657#L45)
``` go
func (s FailSort) Swap(i, j int)
```



## <a name="Failure">type</a> [Failure](https://github.com/hunterlong/statping/tree/master/types/failure.go?s=866:1335#L24)
``` go
type Failure struct {
    Id        int64     `gorm:"primary_key;column:id" json:"id"`
    Issue     string    `gorm:"column:issue" json:"issue"`
    Method    string    `gorm:"column:method" json:"method,omitempty"`
    MethodId  int64     `gorm:"column:method_id" json:"method_id,omitempty"`
    Service   int64     `gorm:"index;column:service" json:"-"`
    PingTime  float64   `gorm:"column:ping_time"  json:"ping"`
    CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

```
Failure is a failed attempt to check a service. Any a service does not meet the expected requirements,
a new Failure will be inserted into database.










## <a name="FailureInterface">type</a> [FailureInterface](https://github.com/hunterlong/statping/tree/master/types/failure.go?s=1337:1542#L34)
``` go
type FailureInterface interface {
    Select() *Failure
    Ago() string        // Ago returns a human readable timestamp
    ParseError() string // ParseError returns a human readable error for a service failure
}
```









## <a name="Hit">type</a> [Hit](https://github.com/hunterlong/statping/tree/master/types/types.go?s=785:1041#L23)
``` go
type Hit struct {
    Id        int64     `gorm:"primary_key;column:id"`
    Service   int64     `gorm:"column:service"`
    Latency   float64   `gorm:"column:latency"`
    PingTime  float64   `gorm:"column:ping_time"`
    CreatedAt time.Time `gorm:"column:created_at"`
}

```
Hit struct is a 'successful' ping or web response entry for a service.










## <a name="Info">type</a> [Info](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=499:579#L34)
``` go
type Info struct {
    Name        string
    Description string
    Form        string
}

```









## <a name="Message">type</a> [Message](https://github.com/hunterlong/statping/tree/master/types/message.go?s=797:1754#L23)
``` go
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
Message is for creating Announcements, Alerts and other messages for the end users










## <a name="NullBool">type</a> [NullBool](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1460:1498#L49)
``` go
type NullBool struct {
    sql.NullBool
}

```
NullBool is an alias for sql.NullBool data type







### <a name="NewNullBool">func</a> [NewNullBool](https://github.com/hunterlong/statping/tree/master/types/null.go?s=939:972#L29)
``` go
func NewNullBool(s bool) NullBool
```
NewNullBool returns a sql.NullBool for JSON parsing





### <a name="NullBool.MarshalJSON">func</a> (\*NullBool) [MarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=2064:2113#L80)
``` go
func (nb *NullBool) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullBool




### <a name="NullBool.UnmarshalJSON">func</a> (\*NullBool) [UnmarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=2716:2765#L110)
``` go
func (nf *NullBool) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullBool




## <a name="NullFloat64">type</a> [NullFloat64](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1656:1700#L59)
``` go
type NullFloat64 struct {
    sql.NullFloat64
}

```
NullFloat64 is an alias for sql.NullFloat64 data type







### <a name="NewNullFloat64">func</a> [NewNullFloat64](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1220:1262#L39)
``` go
func NewNullFloat64(s float64) NullFloat64
```
NewNullFloat64 returns a sql.NullFloat64 for JSON parsing





### <a name="NullFloat64.MarshalJSON">func</a> (\*NullFloat64) [MarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1897:1949#L72)
``` go
func (ni *NullFloat64) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullFloat64




### <a name="NullFloat64.UnmarshalJSON">func</a> (\*NullFloat64) [UnmarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=2554:2606#L103)
``` go
func (nf *NullFloat64) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullFloat64




## <a name="NullInt64">type</a> [NullInt64](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1367:1407#L44)
``` go
type NullInt64 struct {
    sql.NullInt64
}

```
NullInt64 is an alias for sql.NullInt64 data type







### <a name="NewNullInt64">func</a> [NewNullInt64](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1075:1111#L34)
``` go
func NewNullInt64(s int64) NullInt64
```
NewNullInt64 returns a sql.NullInt64 for JSON parsing





### <a name="NullInt64.MarshalJSON">func</a> (\*NullInt64) [MarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1731:1781#L64)
``` go
func (ni *NullInt64) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullInt64




### <a name="NullInt64.UnmarshalJSON">func</a> (\*NullInt64) [UnmarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=2393:2443#L96)
``` go
func (nf *NullInt64) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullInt64




## <a name="NullString">type</a> [NullString](https://github.com/hunterlong/statping/tree/master/types/null.go?s=1555:1597#L54)
``` go
type NullString struct {
    sql.NullString
}

```
NullString is an alias for sql.NullString data type







### <a name="NewNullString">func</a> [NewNullString](https://github.com/hunterlong/statping/tree/master/types/null.go?s=795:834#L24)
``` go
func NewNullString(s string) NullString
```
NewNullString returns a sql.NullString for JSON parsing





### <a name="NullString.MarshalJSON">func</a> (\*NullString) [MarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=2227:2278#L88)
``` go
func (ns *NullString) MarshalJSON() ([]byte, error)
```
MarshalJSON for NullString




### <a name="NullString.UnmarshalJSON">func</a> (\*NullString) [UnmarshalJSON](https://github.com/hunterlong/statping/tree/master/types/null.go?s=2874:2925#L117)
``` go
func (nf *NullString) UnmarshalJSON(b []byte) error
```
Unmarshaler for NullString




## <a name="Plugin">type</a> [Plugin](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=65:127#L8)
``` go
type Plugin struct {
    Name        string
    Description string
}

```









## <a name="PluginActions">type</a> [PluginActions](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=169:234#L17)
``` go
type PluginActions interface {
    GetInfo() *Info
    OnLoad() error
}
```









## <a name="PluginInfo">type</a> [PluginInfo](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=581:644#L40)
``` go
type PluginInfo struct {
    Info   *Info
    Routes []*PluginRoute
}

```









## <a name="PluginJSON">type</a> [PluginJSON](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=287:497#L26)
``` go
type PluginJSON struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Repo        string `json:"repo"`
    Author      string `json:"author"`
    Namespace   string `json:"namespace"`
}

```









## <a name="PluginObject">type</a> [PluginObject](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=129:167#L13)
``` go
type PluginObject struct {
    Pluginer
}

```









## <a name="PluginRepos">type</a> [PluginRepos](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=236:285#L22)
``` go
type PluginRepos struct {
    Plugins []PluginJSON
}

```









## <a name="PluginRoute">type</a> [PluginRoute](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=1025:1107#L68)
``` go
type PluginRoute struct {
    Url    string
    Method string
    Func   http.HandlerFunc
}

```









## <a name="PluginRouting">type</a> [PluginRouting](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=646:757#L45)
``` go
type PluginRouting struct {
    URL     string
    Method  string
    Handler func(http.ResponseWriter, *http.Request)
}

```









## <a name="Pluginer">type</a> [Pluginer](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=759:804#L51)
``` go
type Pluginer interface {
    Select() *Plugin
}
```









## <a name="Router">type</a> [Router](https://github.com/hunterlong/statping/tree/master/types/plugin.go?s=864:964#L59)
``` go
type Router interface {
    Routes() []*PluginRoute
    AddRoute(string, string, http.HandlerFunc) error
}
```









## <a name="Service">type</a> [Service](https://github.com/hunterlong/statping/tree/master/types/service.go?s=754:2863#L23)
``` go
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
    AllowNotifications NullBool           `gorm:"default:false;column:allow_notifications" json:"allow_notifications"`
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
    LastOnline         time.Time          `gorm:"-" json:"last_online"`
    Failures           []FailureInterface `gorm:"-" json:"failures,omitempty"`
}

```
Service is the main struct for Services










### <a name="Service.Close">func</a> (\*Service) [Close](https://github.com/hunterlong/statping/tree/master/types/service.go?s=3223:3248#L68)
``` go
func (s *Service) Close()
```
Close will stop the go routine that is checking if service is online or not




### <a name="Service.IsRunning">func</a> (\*Service) [IsRunning](https://github.com/hunterlong/statping/tree/master/types/service.go?s=3359:3393#L75)
``` go
func (s *Service) IsRunning() bool
```
IsRunning returns true if the service go routine is running




### <a name="Service.Start">func</a> (\*Service) [Start](https://github.com/hunterlong/statping/tree/master/types/service.go?s=3084:3109#L63)
``` go
func (s *Service) Start()
```
Start will create a channel for the service checking go routine




## <a name="ServiceInterface">type</a> [ServiceInterface](https://github.com/hunterlong/statping/tree/master/types/service.go?s=2865:3015#L53)
``` go
type ServiceInterface interface {
    Select() *Service
    CheckQueue(bool)
    Check(bool)
    Create(bool) (int64, error)
    Update(bool) error
    Delete() error
}
```









## <a name="User">type</a> [User](https://github.com/hunterlong/statping/tree/master/types/user.go?s=748:1514#L23)
``` go
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










## <a name="UserInterface">type</a> [UserInterface](https://github.com/hunterlong/statping/tree/master/types/user.go?s=1567:1656#L37)
``` go
type UserInterface interface {
    Create() (int64, error)
    Update() error
    Delete() error
}
```
UserInterface interfaces the database functions
















# utils
`import "github.com/hunterlong/statping/utils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
Package utils contains common methods used in most packages in Statping.
This package contains multiple function like:
Logging, encryption, type conversions, setting utils.Directory as the current directory,
running local CMD commands, and creating/deleting files/folder.

You can overwrite the utils.Directory global variable by including
STATPING_DIR environment variable to be an absolute path.

More info on: <a href="https://github.com/hunterlong/statping">https://github.com/hunterlong/statping</a>




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func Command(cmd string) (string, string, error)](#Command)
* [func DeleteDirectory(directory string) error](#DeleteDirectory)
* [func DeleteFile(file string) error](#DeleteFile)
* [func DurationReadable(d time.Duration) string](#DurationReadable)
* [func FileExists(name string) bool](#FileExists)
* [func FormatDuration(d time.Duration) string](#FormatDuration)
* [func HashPassword(password string) string](#HashPassword)
* [func Http(r *http.Request) string](#Http)
* [func HttpRequest(url, method string, content interface{}, headers []string, body io.Reader, timeout time.Duration) ([]byte, *http.Response, error)](#HttpRequest)
* [func InitLogs() error](#InitLogs)
* [func Log(level int, err interface{}) error](#Log)
* [func NewSHA1Hash(n ...int) string](#NewSHA1Hash)
* [func RandomString(n int) string](#RandomString)
* [func SaveFile(filename string, data []byte) error](#SaveFile)
* [func Timezoner(t time.Time, zone float32) time.Time](#Timezoner)
* [func ToInt(s interface{}) int64](#ToInt)
* [func ToString(s interface{}) string](#ToString)
* [func UnderScoreString(str string) string](#UnderScoreString)
* [type LogRow](#LogRow)
  * [func GetLastLine() *LogRow](#GetLastLine)
  * [func (o *LogRow) FormatForHtml() string](#LogRow.FormatForHtml)
* [type Timestamp](#Timestamp)
  * [func (t Timestamp) Ago() string](#Timestamp.Ago)
* [type Timestamper](#Timestamper)

#### <a name="pkg-examples">Examples</a>
* [DurationReadable](#example_DurationReadable)
* [ToString](#example_ToString)

#### <a name="pkg-files">Package files</a>
[doc.go](https://github.com/hunterlong/statping/tree/master/utils/doc.go) [encryption.go](https://github.com/hunterlong/statping/tree/master/utils/encryption.go) [log.go](https://github.com/hunterlong/statping/tree/master/utils/log.go) [time.go](https://github.com/hunterlong/statping/tree/master/utils/time.go) [utils.go](https://github.com/hunterlong/statping/tree/master/utils/utils.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    FlatpickrTime     = "2006-01-02 15:04"
    FlatpickrDay      = "2006-01-02"
    FlatpickrReadable = "Mon, 02 Jan 2006"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    LastLines []*LogRow
    LockLines sync.Mutex
)
```
``` go
var (
    // Directory returns the current path or the STATPING_DIR environment variable
    Directory string
)
```


## <a name="Command">func</a> [Command](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=3714:3762#L151)
``` go
func Command(cmd string) (string, string, error)
```
Command will run a terminal command with 'sh -c COMMAND' and return stdout and errOut as strings


	in, out, err := Command("sass assets/scss assets/css/base.css")



## <a name="DeleteDirectory">func</a> [DeleteDirectory](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=3464:3508#L145)
``` go
func DeleteDirectory(directory string) error
```
DeleteDirectory will attempt to delete a directory and all contents inside


	DeleteDirectory("assets")



## <a name="DeleteFile">func</a> [DeleteFile](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=3215:3249#L134)
``` go
func DeleteFile(file string) error
```
DeleteFile will attempt to delete a file


	DeleteFile("newfile.json")



## <a name="DurationReadable">func</a> [DurationReadable](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=5058:5103#L209)
``` go
func DurationReadable(d time.Duration) string
```
DurationReadable will return a time.Duration into a human readable string


	t := time.Duration(5 * time.Minute)
	DurationReadable(t)
	// 5 minutes



#### <a name="example_DurationReadable">Example</a>

Code:
``` go
dur, _ := time.ParseDuration("25m")
readable := DurationReadable(dur)
fmt.Print(readable)
```
Output:

    25 minutes
    


## <a name="FileExists">func</a> [FileExists](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=2997:3030#L123)
``` go
func FileExists(name string) bool
```
FileExists returns true if a file exists


	exists := FileExists("assets/css/base.css")



## <a name="FormatDuration">func</a> [FormatDuration](https://github.com/hunterlong/statping/tree/master/utils/time.go?s=1127:1170#L38)
``` go
func FormatDuration(d time.Duration) string
```
FormatDuration converts a time.Duration into a string



## <a name="HashPassword">func</a> [HashPassword](https://github.com/hunterlong/statping/tree/master/utils/encryption.go?s=837:878#L27)
``` go
func HashPassword(password string) string
```
HashPassword returns the bcrypt hash of a password string



## <a name="Http">func</a> [Http](https://github.com/hunterlong/statping/tree/master/utils/log.go?s=3074:3107#L126)
``` go
func Http(r *http.Request) string
```
Http returns a log for a HTTP request



## <a name="HttpRequest">func</a> [HttpRequest](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=5648:5794#L228)
``` go
func HttpRequest(url, method string, content interface{}, headers []string, body io.Reader, timeout time.Duration) ([]byte, *http.Response, error)
```
HttpRequest is a global function to send a HTTP request



## <a name="InitLogs">func</a> [InitLogs](https://github.com/hunterlong/statping/tree/master/utils/log.go?s=1419:1440#L58)
``` go
func InitLogs() error
```
InitLogs will create the '/logs' directory and creates a file '/logs/statup.log' for application logging



## <a name="Log">func</a> [Log](https://github.com/hunterlong/statping/tree/master/utils/log.go?s=2195:2237#L93)
``` go
func Log(level int, err interface{}) error
```
Log creates a new entry in the Logger. Log has 1-5 levels depending on how critical the log/error is



## <a name="NewSHA1Hash">func</a> [NewSHA1Hash](https://github.com/hunterlong/statping/tree/master/utils/encryption.go?s=1038:1071#L33)
``` go
func NewSHA1Hash(n ...int) string
```
NewSHA1Hash returns a random SHA1 hash based on a specific length



## <a name="RandomString">func</a> [RandomString](https://github.com/hunterlong/statping/tree/master/utils/encryption.go?s=1451:1482#L48)
``` go
func RandomString(n int) string
```
RandomString generates a random string of n length



## <a name="SaveFile">func</a> [SaveFile](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=5475:5524#L222)
``` go
func SaveFile(filename string, data []byte) error
```
SaveFile will create a new file with data inside it


	SaveFile("newfile.json", []byte('{"data": "success"}')



## <a name="Timezoner">func</a> [Timezoner](https://github.com/hunterlong/statping/tree/master/utils/time.go?s=905:956#L30)
``` go
func Timezoner(t time.Time, zone float32) time.Time
```
Timezoner returns the time.Time with the user set timezone



## <a name="ToInt">func</a> [ToInt](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=1222:1253#L49)
``` go
func ToInt(s interface{}) int64
```
ToInt converts a int to a string



## <a name="ToString">func</a> [ToString](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=1408:1443#L60)
``` go
func ToString(s interface{}) string
```
ToString converts a int to a string



#### <a name="example_ToString">Example</a>

Code:
``` go
amount := 42
fmt.Print(ToString(amount))
```
Output:

    42
    


## <a name="UnderScoreString">func</a> [UnderScoreString](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=2264:2304#L98)
``` go
func UnderScoreString(str string) string
```
UnderScoreString will return a string that replaces spaces and other characters to underscores


	UnderScoreString("Example String")
	// example_string




## <a name="LogRow">type</a> [LogRow](https://github.com/hunterlong/statping/tree/master/utils/log.go?s=3709:3765#L153)
``` go
type LogRow struct {
    Date time.Time
    Line interface{}
}

```






### <a name="GetLastLine">func</a> [GetLastLine](https://github.com/hunterlong/statping/tree/master/utils/log.go?s=3556:3582#L144)
``` go
func GetLastLine() *LogRow
```
GetLastLine returns 1 line for a recent log entry





### <a name="LogRow.FormatForHtml">func</a> (\*LogRow) [FormatForHtml](https://github.com/hunterlong/statping/tree/master/utils/log.go?s=4075:4114#L177)
``` go
func (o *LogRow) FormatForHtml() string
```



## <a name="Timestamp">type</a> [Timestamp](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=1837:1861#L84)
``` go
type Timestamp time.Time
```









### <a name="Timestamp.Ago">func</a> (Timestamp) [Ago](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=1995:2026#L90)
``` go
func (t Timestamp) Ago() string
```
Ago returns a human readable timestamp based on the Timestamp (time.Time) interface




## <a name="Timestamper">type</a> [Timestamper](https://github.com/hunterlong/statping/tree/master/utils/utils.go?s=1862:1906#L85)
``` go
type Timestamper interface {
    Ago() string
}
```













