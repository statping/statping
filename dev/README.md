

# core
`import "github.com/hunterlong/statup/core"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package core contains the main functionality of Statup. This includes everything for
Services, Hits, Failures, Users, service checking mechanisms, databases, and notifiers
in the notifier package

More info on: <a href="https://github.com/hunterlong/statup">https://github.com/hunterlong/statup</a>




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [func CheckHash(password, hash string) bool](#CheckHash)
* [func CloseDB()](#CloseDB)
* [func CountFailures() uint64](#CountFailures)
* [func DatabaseMaintence()](#DatabaseMaintence)
* [func Dbtimestamp(group string, column string) string](#Dbtimestamp)
* [func DeleteAllSince(table string, date time.Time)](#DeleteAllSince)
* [func DeleteConfig()](#DeleteConfig)
* [func ExportChartsJs() string](#ExportChartsJs)
* [func ExportIndexHTML() string](#ExportIndexHTML)
* [func InitApp()](#InitApp)
* [func InsertLargeSampleData() error](#InsertLargeSampleData)
* [func InsertSampleData() error](#InsertSampleData)
* [func InsertSampleHits() error](#InsertSampleHits)
* [func SaveFile(filename string, data []byte) error](#SaveFile)
* [type Checkin](#Checkin)
  * [func ReturnCheckin(s *types.Checkin) *Checkin](#ReturnCheckin)
  * [func SelectCheckin(api string) *Checkin](#SelectCheckin)
  * [func (s *Checkin) AfterFind() (err error)](#Checkin.AfterFind)
  * [func (u *Checkin) BeforeCreate() (err error)](#Checkin.BeforeCreate)
  * [func (u *Checkin) Create() (int64, error)](#Checkin.Create)
  * [func (u *Checkin) Expected() time.Duration](#Checkin.Expected)
  * [func (u *Checkin) Grace() time.Duration](#Checkin.Grace)
  * [func (u *Checkin) Hits() []CheckinHit](#Checkin.Hits)
  * [func (u *Checkin) Last() CheckinHit](#Checkin.Last)
  * [func (u *Checkin) Period() time.Duration](#Checkin.Period)
  * [func (c *Checkin) RecheckCheckinFailure(guard chan struct{})](#Checkin.RecheckCheckinFailure)
  * [func (c *Checkin) String() string](#Checkin.String)
  * [func (u *Checkin) Update() (int64, error)](#Checkin.Update)
* [type CheckinHit](#CheckinHit)
  * [func ReturnCheckinHit(h *types.CheckinHit) *CheckinHit](#ReturnCheckinHit)
  * [func (s *CheckinHit) AfterFind() (err error)](#CheckinHit.AfterFind)
  * [func (f *CheckinHit) Ago() string](#CheckinHit.Ago)
  * [func (u *CheckinHit) BeforeCreate() (err error)](#CheckinHit.BeforeCreate)
  * [func (u *CheckinHit) Create() (int64, error)](#CheckinHit.Create)
* [type Core](#Core)
  * [func NewCore() *Core](#NewCore)
  * [func SelectCore() (*Core, error)](#SelectCore)
  * [func UpdateCore(c *Core) (*Core, error)](#UpdateCore)
  * [func (c Core) AllOnline() bool](#Core.AllOnline)
  * [func (c Core) BaseSASS() string](#Core.BaseSASS)
  * [func (c *Core) Count24HFailures() uint64](#Core.Count24HFailures)
  * [func (c *Core) CountOnline() int](#Core.CountOnline)
  * [func (c Core) CurrentTime() string](#Core.CurrentTime)
  * [func (c Core) MobileSASS() string](#Core.MobileSASS)
  * [func (c Core) SassVars() string](#Core.SassVars)
  * [func (c *Core) SelectAllServices() ([]*Service, error)](#Core.SelectAllServices)
  * [func (c *Core) ServicesCount() int](#Core.ServicesCount)
  * [func (c *Core) ToCore() *types.Core](#Core.ToCore)
  * [func (c Core) UsingAssets() bool](#Core.UsingAssets)
* [type DateScan](#DateScan)
* [type DateScanObj](#DateScanObj)
  * [func GraphDataRaw(service types.ServiceInterface, start, end time.Time, group string, column string) *DateScanObj](#GraphDataRaw)
  * [func (d *DateScanObj) ToString() string](#DateScanObj.ToString)
* [type DbConfig](#DbConfig)
  * [func LoadConfig(directory string) (*DbConfig, error)](#LoadConfig)
  * [func LoadUsingEnv() (*DbConfig, error)](#LoadUsingEnv)
  * [func (db *DbConfig) Close() error](#DbConfig.Close)
  * [func (db *DbConfig) Connect(retry bool, location string) error](#DbConfig.Connect)
  * [func (c *DbConfig) CreateCore() *Core](#DbConfig.CreateCore)
  * [func (db *DbConfig) CreateDatabase() error](#DbConfig.CreateDatabase)
  * [func (db *DbConfig) DropDatabase() error](#DbConfig.DropDatabase)
  * [func (db *DbConfig) InsertCore() (*Core, error)](#DbConfig.InsertCore)
  * [func (db *DbConfig) MigrateDatabase() error](#DbConfig.MigrateDatabase)
  * [func (c *DbConfig) Save() (*DbConfig, error)](#DbConfig.Save)
  * [func (c *DbConfig) Update() error](#DbConfig.Update)
* [type ErrorResponse](#ErrorResponse)
* [type Failure](#Failure)
  * [func (f *Failure) AfterFind() (err error)](#Failure.AfterFind)
  * [func (f *Failure) Ago() string](#Failure.Ago)
  * [func (u *Failure) BeforeCreate() (err error)](#Failure.BeforeCreate)
  * [func (f *Failure) Delete() error](#Failure.Delete)
  * [func (f *Failure) ParseError() string](#Failure.ParseError)
* [type Hit](#Hit)
  * [func (s *Hit) AfterFind() (err error)](#Hit.AfterFind)
  * [func (u *Hit) BeforeCreate() (err error)](#Hit.BeforeCreate)
* [type PluginJSON](#PluginJSON)
* [type PluginRepos](#PluginRepos)
* [type Service](#Service)
  * [func ReturnService(s *types.Service) *Service](#ReturnService)
  * [func SelectService(id int64) *Service](#SelectService)
  * [func (s *Service) AfterFind() (err error)](#Service.AfterFind)
  * [func (s *Service) AllFailures() []*Failure](#Service.AllFailures)
  * [func (s *Service) AvgTime() float64](#Service.AvgTime)
  * [func (s *Service) AvgUptime(ago time.Time) string](#Service.AvgUptime)
  * [func (s *Service) AvgUptime24() string](#Service.AvgUptime24)
  * [func (u *Service) BeforeCreate() (err error)](#Service.BeforeCreate)
  * [func (s *Service) Check(record bool)](#Service.Check)
  * [func (s *Service) CheckQueue(record bool)](#Service.CheckQueue)
  * [func (s *Service) Checkins() []*Checkin](#Service.Checkins)
  * [func (u *Service) Create(check bool) (int64, error)](#Service.Create)
  * [func (s *Service) CreateFailure(f *types.Failure) (int64, error)](#Service.CreateFailure)
  * [func (s *Service) CreateHit(h *types.Hit) (int64, error)](#Service.CreateHit)
  * [func (u *Service) Delete() error](#Service.Delete)
  * [func (u *Service) DeleteFailures()](#Service.DeleteFailures)
  * [func (s *Service) Downtime() time.Duration](#Service.Downtime)
  * [func (s *Service) DowntimeText() string](#Service.DowntimeText)
  * [func (s *Service) GraphData() string](#Service.GraphData)
  * [func (s *Service) Hits() ([]*types.Hit, error)](#Service.Hits)
  * [func (s *Service) HitsBetween(t1, t2 time.Time, group string, column string) *gorm.DB](#Service.HitsBetween)
  * [func (s *Service) LimitedFailures() []*Failure](#Service.LimitedFailures)
  * [func (s *Service) LimitedHits() ([]*types.Hit, error)](#Service.LimitedHits)
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
  * [func (u *Service) Update(restart bool) error](#Service.Update)
  * [func (u *Service) UpdateSingle(attr ...interface{}) error](#Service.UpdateSingle)
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
[checker.go](/src/github.com/hunterlong/statup/core/checker.go) [checkin.go](/src/github.com/hunterlong/statup/core/checkin.go) [configs.go](/src/github.com/hunterlong/statup/core/configs.go) [core.go](/src/github.com/hunterlong/statup/core/core.go) [database.go](/src/github.com/hunterlong/statup/core/database.go) [doc.go](/src/github.com/hunterlong/statup/core/doc.go) [export.go](/src/github.com/hunterlong/statup/core/export.go) [failures.go](/src/github.com/hunterlong/statup/core/failures.go) [hits.go](/src/github.com/hunterlong/statup/core/hits.go) [sample.go](/src/github.com/hunterlong/statup/core/sample.go) [services.go](/src/github.com/hunterlong/statup/core/services.go) [users.go](/src/github.com/hunterlong/statup/core/users.go) 



## <a name="pkg-variables">Variables</a>
``` go
var (
    Configs   *DbConfig // Configs holds all of the config.yml and database info
    CoreApp   *Core     // CoreApp is a global variable that contains many elements
    SetupMode bool      // SetupMode will be true if Statup does not have a database connection
    VERSION   string    // VERSION is set on build automatically by setting a -ldflag
)
```
``` go
var (
    DbSession *gorm.DB
)
```


## <a name="CheckHash">func</a> [CheckHash](/src/target/users.go?s=2896:2938#L105)
``` go
func CheckHash(password, hash string) bool
```
CheckHash returns true if the password matches with a hashed bcrypt password



## <a name="CloseDB">func</a> [CloseDB](/src/target/database.go?s=2503:2517#L81)
``` go
func CloseDB()
```
CloseDB will close the database connection if available



## <a name="CountFailures">func</a> [CountFailures](/src/target/failures.go?s=2819:2846#L99)
``` go
func CountFailures() uint64
```
CountFailures returns the total count of failures for all services



## <a name="DatabaseMaintence">func</a> [DatabaseMaintence](/src/target/database.go?s=7073:7097#L247)
``` go
func DatabaseMaintence()
```
DatabaseMaintence will automatically delete old records from 'failures' and 'hits'
this function is currently set to delete records 7+ days old every 60 minutes



## <a name="Dbtimestamp">func</a> [Dbtimestamp](/src/target/services.go?s=5171:5223#L180)
``` go
func Dbtimestamp(group string, column string) string
```
Dbtimestamp will return a SQL query for grouping by date



## <a name="DeleteAllSince">func</a> [DeleteAllSince](/src/target/database.go?s=7406:7455#L257)
``` go
func DeleteAllSince(table string, date time.Time)
```
DeleteAllSince will delete a specific table's records based on a time.



## <a name="DeleteConfig">func</a> [DeleteConfig](/src/target/configs.go?s=3785:3804#L133)
``` go
func DeleteConfig()
```
DeleteConfig will delete the 'config.yml' file



## <a name="ExportChartsJs">func</a> [ExportChartsJs](/src/target/export.go?s=2118:2146#L86)
``` go
func ExportChartsJs() string
```


## <a name="ExportIndexHTML">func</a> [ExportIndexHTML](/src/target/export.go?s=889:918#L31)
``` go
func ExportIndexHTML() string
```


## <a name="InitApp">func</a> [InitApp](/src/target/core.go?s=1690:1704#L60)
``` go
func InitApp()
```
InitApp will initialize Statup



## <a name="InsertLargeSampleData">func</a> [InsertLargeSampleData](/src/target/sample.go?s=4671:4705#L179)
``` go
func InsertLargeSampleData() error
```
InsertSampleData will create the example/dummy services for a brand new Statup installation



## <a name="InsertSampleData">func</a> [InsertSampleData](/src/target/sample.go?s=897:926#L27)
``` go
func InsertSampleData() error
```
InsertSampleData will create the example/dummy services for a brand new Statup installation



## <a name="InsertSampleHits">func</a> [InsertSampleHits](/src/target/sample.go?s=3132:3161#L118)
``` go
func InsertSampleHits() error
```
InsertSampleHits will create a couple new hits for the sample services



## <a name="SaveFile">func</a> [SaveFile](/src/target/export.go?s=2540:2589#L106)
``` go
func SaveFile(filename string, data []byte) error
```



## <a name="Checkin">type</a> [Checkin](/src/target/checkin.go?s=818:857#L26)
``` go
type Checkin struct {
    *types.Checkin
}

```






### <a name="ReturnCheckin">func</a> [ReturnCheckin](/src/target/checkin.go?s=1064:1109#L40)
``` go
func ReturnCheckin(s *types.Checkin) *Checkin
```
ReturnCheckin converts *types.Checking to *core.Checkin


### <a name="SelectCheckin">func</a> [SelectCheckin](/src/target/checkin.go?s=1369:1408#L50)
``` go
func SelectCheckin(api string) *Checkin
```
SelectCheckin will find a Checkin based on the API supplied





### <a name="Checkin.AfterFind">func</a> (\*Checkin) [AfterFind](/src/target/database.go?s=3376:3417#L117)
``` go
func (s *Checkin) AfterFind() (err error)
```
AfterFind for Checkin will set the timezone




### <a name="Checkin.BeforeCreate">func</a> (\*Checkin) [BeforeCreate](/src/target/database.go?s=4403:4447#L161)
``` go
func (u *Checkin) BeforeCreate() (err error)
```
BeforeCreate for Checkin will set CreatedAt to UTC




### <a name="Checkin.Create">func</a> (\*Checkin) [Create](/src/target/checkin.go?s=2642:2683#L92)
``` go
func (u *Checkin) Create() (int64, error)
```
Create will create a new Checkin




### <a name="Checkin.Expected">func</a> (\*Checkin) [Expected](/src/target/checkin.go?s=2028:2070#L69)
``` go
func (u *Checkin) Expected() time.Duration
```
Expected returns the duration of when the serviec should receive a checkin




### <a name="Checkin.Grace">func</a> (\*Checkin) [Grace](/src/target/checkin.go?s=1818:1857#L63)
``` go
func (u *Checkin) Grace() time.Duration
```
Grace will return the duration of the Checkin Grace Period (after service hasn't responded, wait a bit for a response)




### <a name="Checkin.Hits">func</a> (\*Checkin) [Hits](/src/target/checkin.go?s=2442:2479#L85)
``` go
func (u *Checkin) Hits() []CheckinHit
```
Hits returns all of the CheckinHits for a given Checkin




### <a name="Checkin.Last">func</a> (\*Checkin) [Last](/src/target/checkin.go?s=2255:2290#L78)
``` go
func (u *Checkin) Last() CheckinHit
```
Last returns the last CheckinHit for a Checkin




### <a name="Checkin.Period">func</a> (\*Checkin) [Period](/src/target/checkin.go?s=1566:1606#L57)
``` go
func (u *Checkin) Period() time.Duration
```
Period will return the duration of the Checkin interval




### <a name="Checkin.RecheckCheckinFailure">func</a> (\*Checkin) [RecheckCheckinFailure](/src/target/checkin.go?s=3622:3682#L132)
``` go
func (c *Checkin) RecheckCheckinFailure(guard chan struct{})
```
RecheckCheckinFailure will check if a Service Checkin has been reported yet




### <a name="Checkin.String">func</a> (\*Checkin) [String](/src/target/checkin.go?s=949:982#L35)
``` go
func (c *Checkin) String() string
```
String will return a Checkin API string




### <a name="Checkin.Update">func</a> (\*Checkin) [Update](/src/target/checkin.go?s=2884:2925#L103)
``` go
func (u *Checkin) Update() (int64, error)
```
Update will update a Checkin




## <a name="CheckinHit">type</a> [CheckinHit](/src/target/checkin.go?s=859:904#L30)
``` go
type CheckinHit struct {
    *types.CheckinHit
}

```






### <a name="ReturnCheckinHit">func</a> [ReturnCheckinHit](/src/target/checkin.go?s=1211:1265#L45)
``` go
func ReturnCheckinHit(h *types.CheckinHit) *CheckinHit
```
ReturnCheckinHit converts *types.CheckinHit to *core.CheckinHit





### <a name="CheckinHit.AfterFind">func</a> (\*CheckinHit) [AfterFind](/src/target/database.go?s=3543:3587#L123)
``` go
func (s *CheckinHit) AfterFind() (err error)
```
AfterFind for CheckinHit will set the timezone




### <a name="CheckinHit.Ago">func</a> (\*CheckinHit) [Ago](/src/target/checkin.go?s=3432:3465#L126)
``` go
func (f *CheckinHit) Ago() string
```
Ago returns the duration of time between now and the last successful CheckinHit




### <a name="CheckinHit.BeforeCreate">func</a> (\*CheckinHit) [BeforeCreate](/src/target/database.go?s=4581:4628#L169)
``` go
func (u *CheckinHit) BeforeCreate() (err error)
```
BeforeCreate for CheckinHit will set CreatedAt to UTC




### <a name="CheckinHit.Create">func</a> (\*CheckinHit) [Create](/src/target/checkin.go?s=3110:3154#L113)
``` go
func (u *CheckinHit) Create() (int64, error)
```
Create will create a new successful CheckinHit




## <a name="Core">type</a> [Core](/src/target/core.go?s=967:1000#L31)
``` go
type Core struct {
    *types.Core
}

```






### <a name="NewCore">func</a> [NewCore](/src/target/core.go?s=1426:1446#L47)
``` go
func NewCore() *Core
```
NewCore return a new *core.Core struct


### <a name="SelectCore">func</a> [SelectCore](/src/target/core.go?s=3778:3810#L136)
``` go
func SelectCore() (*Core, error)
```
SelectCore will return the CoreApp global variable and the settings/configs for Statup


### <a name="UpdateCore">func</a> [UpdateCore](/src/target/core.go?s=2257:2296#L82)
``` go
func UpdateCore(c *Core) (*Core, error)
```
UpdateCore will update the CoreApp variable inside of the 'core' table in database





### <a name="Core.AllOnline">func</a> (Core) [AllOnline](/src/target/core.go?s=3552:3582#L126)
``` go
func (c Core) AllOnline() bool
```
AllOnline will be true if all services are online




### <a name="Core.BaseSASS">func</a> (Core) [BaseSASS](/src/target/core.go?s=3055:3086#L109)
``` go
func (c Core) BaseSASS() string
```
BaseSASS is the base design , this opens the file /assets/scss/base.scss to be edited in Theme




### <a name="Core.Count24HFailures">func</a> (\*Core) [Count24HFailures](/src/target/failures.go?s=2547:2587#L88)
``` go
func (c *Core) Count24HFailures() uint64
```
Count24HFailures returns the amount of failures for a service within the last 24 hours




### <a name="Core.CountOnline">func</a> (\*Core) [CountOnline](/src/target/services.go?s=10815:10847#L382)
``` go
func (c *Core) CountOnline() int
```
CountOnline




### <a name="Core.CurrentTime">func</a> (Core) [CurrentTime](/src/target/core.go?s=2410:2444#L88)
``` go
func (c Core) CurrentTime() string
```
UsingAssets will return true if /assets folder is present




### <a name="Core.MobileSASS">func</a> (Core) [MobileSASS](/src/target/core.go?s=3340:3373#L118)
``` go
func (c Core) MobileSASS() string
```
MobileSASS is the -webkit responsive custom css designs. This opens the
file /assets/scss/mobile.scss to be edited in Theme




### <a name="Core.SassVars">func</a> (Core) [SassVars](/src/target/core.go?s=2797:2828#L101)
``` go
func (c Core) SassVars() string
```
SassVars opens the file /assets/scss/variables.scss to be edited in Theme




### <a name="Core.SelectAllServices">func</a> (\*Core) [SelectAllServices](/src/target/services.go?s=1729:1783#L62)
``` go
func (c *Core) SelectAllServices() ([]*Service, error)
```
SelectAllServices returns a slice of *core.Service to be store on []*core.Services, should only be called once on startup.




### <a name="Core.ServicesCount">func</a> (\*Core) [ServicesCount](/src/target/services.go?s=10736:10770#L377)
``` go
func (c *Core) ServicesCount() int
```
ServicesCount returns the amount of services inside the []*core.Services slice




### <a name="Core.ToCore">func</a> (\*Core) [ToCore](/src/target/core.go?s=1600:1635#L55)
``` go
func (c *Core) ToCore() *types.Core
```
ToCore will convert *core.Core to *types.Core




### <a name="Core.UsingAssets">func</a> (Core) [UsingAssets](/src/target/core.go?s=2638:2670#L96)
``` go
func (c Core) UsingAssets() bool
```
UsingAssets will return true if /assets folder is present




## <a name="DateScan">type</a> [DateScan](/src/target/services.go?s=3675:3757#L133)
``` go
type DateScan struct {
    CreatedAt string `json:"x"`
    Value     int64  `json:"y"`
}

```
DateScan struct is for creating the charts.js graph JSON array










## <a name="DateScanObj">type</a> [DateScanObj](/src/target/services.go?s=3828:3887#L139)
``` go
type DateScanObj struct {
    Array []DateScan `json:"data"`
}

```
DateScanObj struct is for creating the charts.js graph JSON array







### <a name="GraphDataRaw">func</a> [GraphDataRaw](/src/target/services.go?s=6296:6409#L216)
``` go
func GraphDataRaw(service types.ServiceInterface, start, end time.Time, group string, column string) *DateScanObj
```
GraphDataRaw will return all the hits between 2 times for a Service





### <a name="DateScanObj.ToString">func</a> (\*DateScanObj) [ToString](/src/target/services.go?s=7072:7111#L238)
``` go
func (d *DateScanObj) ToString() string
```
ToString will convert the DateScanObj into a JSON string for the charts to render




## <a name="DbConfig">type</a> [DbConfig](/src/target/database.go?s=1099:1127#L37)
``` go
type DbConfig types.DbConfig
```






### <a name="LoadConfig">func</a> [LoadConfig](/src/target/configs.go?s=1020:1072#L34)
``` go
func LoadConfig(directory string) (*DbConfig, error)
```
LoadConfig will attempt to load the 'config.yml' file in a specific directory


### <a name="LoadUsingEnv">func</a> [LoadUsingEnv](/src/target/configs.go?s=1680:1718#L53)
``` go
func LoadUsingEnv() (*DbConfig, error)
```
LoadUsingEnv will attempt to load database configs based on environment variables. If DB_CONN is set if will force this function.





### <a name="DbConfig.Close">func</a> (\*DbConfig) [Close](/src/target/database.go?s=2617:2650#L88)
``` go
func (db *DbConfig) Close() error
```
Close shutsdown the database connection




### <a name="DbConfig.Connect">func</a> (\*DbConfig) [Connect](/src/target/database.go?s=5265:5327#L193)
``` go
func (db *DbConfig) Connect(retry bool, location string) error
```
Connect will attempt to connect to the sqlite, postgres, or mysql database




### <a name="DbConfig.CreateCore">func</a> (\*DbConfig) [CreateCore](/src/target/database.go?s=8558:8595#L304)
``` go
func (c *DbConfig) CreateCore() *Core
```
CreateCore will initialize the global variable 'CoreApp". This global variable contains most of Statup app.




### <a name="DbConfig.CreateDatabase">func</a> (\*DbConfig) [CreateDatabase](/src/target/database.go?s=9609:9651#L340)
``` go
func (db *DbConfig) CreateDatabase() error
```
CreateDatabase will CREATE TABLES for each of the Statup elements




### <a name="DbConfig.DropDatabase">func</a> (\*DbConfig) [DropDatabase](/src/target/database.go?s=9055:9095#L326)
``` go
func (db *DbConfig) DropDatabase() error
```
DropDatabase will DROP each table Statup created




### <a name="DbConfig.InsertCore">func</a> (\*DbConfig) [InsertCore](/src/target/database.go?s=4773:4820#L177)
``` go
func (db *DbConfig) InsertCore() (*Core, error)
```
InsertCore create the single row for the Core settings in Statup




### <a name="DbConfig.MigrateDatabase">func</a> (\*DbConfig) [MigrateDatabase](/src/target/database.go?s=10390:10433#L357)
``` go
func (db *DbConfig) MigrateDatabase() error
```
MigrateDatabase will migrate the database structure to current version.
This function will NOT remove previous records, tables or columns from the database.
If this function has an issue, it will ROLLBACK to the previous state.




### <a name="DbConfig.Save">func</a> (\*DbConfig) [Save](/src/target/database.go?s=8035:8079#L284)
``` go
func (c *DbConfig) Save() (*DbConfig, error)
```
Save will initially create the config.yml file




### <a name="DbConfig.Update">func</a> (\*DbConfig) [Update](/src/target/database.go?s=7674:7707#L266)
``` go
func (c *DbConfig) Update() error
```
Update will save the config.yml file




## <a name="ErrorResponse">type</a> [ErrorResponse](/src/target/configs.go?s=894:937#L29)
``` go
type ErrorResponse struct {
    Error string
}

```
ErrorResponse is used for HTTP errors to show to user










## <a name="Failure">type</a> [Failure](/src/target/failures.go?s=829:868#L27)
``` go
type Failure struct {
    *types.Failure
}

```









### <a name="Failure.AfterFind">func</a> (\*Failure) [AfterFind](/src/target/database.go?s=3054:3095#L105)
``` go
func (f *Failure) AfterFind() (err error)
```
AfterFind for Failure will set the timezone




### <a name="Failure.Ago">func</a> (\*Failure) [Ago](/src/target/failures.go?s=2207:2237#L76)
``` go
func (f *Failure) Ago() string
```
Ago returns a human readable timestamp for a failure




### <a name="Failure.BeforeCreate">func</a> (\*Failure) [BeforeCreate](/src/target/database.go?s=3884:3928#L137)
``` go
func (u *Failure) BeforeCreate() (err error)
```
BeforeCreate for Failure will set CreatedAt to UTC




### <a name="Failure.Delete">func</a> (\*Failure) [Delete](/src/target/failures.go?s=2372:2404#L82)
``` go
func (f *Failure) Delete() error
```
Delete will remove a failure record from the database




### <a name="Failure.ParseError">func</a> (\*Failure) [ParseError](/src/target/failures.go?s=3855:3892#L132)
``` go
func (f *Failure) ParseError() string
```
ParseError returns a human readable error for a failure




## <a name="Hit">type</a> [Hit](/src/target/hits.go?s=782:813#L24)
``` go
type Hit struct {
    *types.Hit
}

```









### <a name="Hit.AfterFind">func</a> (\*Hit) [AfterFind](/src/target/database.go?s=2894:2931#L99)
``` go
func (s *Hit) AfterFind() (err error)
```
AfterFind for Hit will set the timezone




### <a name="Hit.BeforeCreate">func</a> (\*Hit) [BeforeCreate](/src/target/database.go?s=3713:3753#L129)
``` go
func (u *Hit) BeforeCreate() (err error)
```
BeforeCreate for Hit will set CreatedAt to UTC




## <a name="PluginJSON">type</a> [PluginJSON](/src/target/core.go?s=898:930#L28)
``` go
type PluginJSON types.PluginJSON
```









## <a name="PluginRepos">type</a> [PluginRepos](/src/target/core.go?s=931:965#L29)
``` go
type PluginRepos types.PluginRepos
```









## <a name="Service">type</a> [Service](/src/target/services.go?s=900:939#L30)
``` go
type Service struct {
    *types.Service
}

```






### <a name="ReturnService">func</a> [ReturnService](/src/target/services.go?s=1128:1173#L40)
``` go
func ReturnService(s *types.Service) *Service
```
ReturnService will convert *types.Service to *core.Service


### <a name="SelectService">func</a> [SelectService](/src/target/services.go?s=1255:1292#L45)
``` go
func SelectService(id int64) *Service
```
SelectService returns a *core.Service from in memory





### <a name="Service.AfterFind">func</a> (\*Service) [AfterFind](/src/target/database.go?s=2734:2775#L93)
``` go
func (s *Service) AfterFind() (err error)
```
AfterFind for Service will set the timezone




### <a name="Service.AllFailures">func</a> (\*Service) [AllFailures](/src/target/failures.go?s=1249:1291#L44)
``` go
func (s *Service) AllFailures() []*Failure
```
AllFailures will return all failures attached to a service




### <a name="Service.AvgTime">func</a> (\*Service) [AvgTime](/src/target/services.go?s=2596:2631#L92)
``` go
func (s *Service) AvgTime() float64
```
AvgTime will return the average amount of time for a service to response back successfully




### <a name="Service.AvgUptime">func</a> (\*Service) [AvgUptime](/src/target/services.go?s=7817:7866#L267)
``` go
func (s *Service) AvgUptime(ago time.Time) string
```
AvgUptime returns average online status for last 24 hours




### <a name="Service.AvgUptime24">func</a> (\*Service) [AvgUptime24](/src/target/services.go?s=7647:7685#L261)
``` go
func (s *Service) AvgUptime24() string
```
AvgUptime24 returns a service's average online status for last 24 hours




### <a name="Service.BeforeCreate">func</a> (\*Service) [BeforeCreate](/src/target/database.go?s=4228:4272#L153)
``` go
func (u *Service) BeforeCreate() (err error)
```
BeforeCreate for Service will set CreatedAt to UTC




### <a name="Service.Check">func</a> (\*Service) [Check](/src/target/checker.go?s=5538:5574#L222)
``` go
func (s *Service) Check(record bool)
```
Check will run checkHttp for HTTP services and checkTcp for TCP services




### <a name="Service.CheckQueue">func</a> (\*Service) [CheckQueue](/src/target/checker.go?s=1256:1297#L43)
``` go
func (s *Service) CheckQueue(record bool)
```
CheckQueue is the main go routine for checking a service




### <a name="Service.Checkins">func</a> (\*Service) [Checkins](/src/target/services.go?s=1463:1502#L55)
``` go
func (s *Service) Checkins() []*Checkin
```
Checkins will return a slice of Checkins for a Service




### <a name="Service.Create">func</a> (\*Service) [Create](/src/target/services.go?s=10250:10301#L361)
``` go
func (u *Service) Create(check bool) (int64, error)
```
Create will create a service and insert it into the database




### <a name="Service.CreateFailure">func</a> (\*Service) [CreateFailure](/src/target/failures.go?s=934:998#L32)
``` go
func (s *Service) CreateFailure(f *types.Failure) (int64, error)
```
CreateFailure will create a new failure record for a service




### <a name="Service.CreateHit">func</a> (\*Service) [CreateHit](/src/target/hits.go?s=907:963#L29)
``` go
func (s *Service) CreateHit(h *types.Hit) (int64, error)
```
CreateHit will create a new 'hit' record in the database for a successful/online service




### <a name="Service.Delete">func</a> (\*Service) [Delete](/src/target/services.go?s=9104:9136#L321)
``` go
func (u *Service) Delete() error
```
Delete will remove a service from the database, it will also end the service checking go routine




### <a name="Service.DeleteFailures">func</a> (\*Service) [DeleteFailures](/src/target/failures.go?s=1672:1706#L59)
``` go
func (u *Service) DeleteFailures()
```
DeleteFailures will delete all failures for a service




### <a name="Service.Downtime">func</a> (\*Service) [Downtime](/src/target/services.go?s=5906:5948#L202)
``` go
func (s *Service) Downtime() time.Duration
```
Downtime returns the amount of time of a offline service




### <a name="Service.DowntimeText">func</a> (\*Service) [DowntimeText](/src/target/services.go?s=4970:5009#L175)
``` go
func (s *Service) DowntimeText() string
```
DowntimeText will return the amount of downtime for a service based on the duration




### <a name="Service.GraphData">func</a> (\*Service) [GraphData](/src/target/services.go?s=7303:7339#L248)
``` go
func (s *Service) GraphData() string
```
GraphData returns the JSON object used by Charts.js to render the chart




### <a name="Service.Hits">func</a> (\*Service) [Hits](/src/target/hits.go?s=1139:1185#L39)
``` go
func (s *Service) Hits() ([]*types.Hit, error)
```
Hits returns all successful hits for a service




### <a name="Service.HitsBetween">func</a> (\*Service) [HitsBetween](/src/target/database.go?s=2097:2182#L75)
``` go
func (s *Service) HitsBetween(t1, t2 time.Time, group string, column string) *gorm.DB
```
HitsBetween returns the gorm database query for a collection of service hits between a time range




### <a name="Service.LimitedFailures">func</a> (\*Service) [LimitedFailures](/src/target/failures.go?s=1964:2010#L68)
``` go
func (s *Service) LimitedFailures() []*Failure
```
LimitedFailures will return the last 10 failures from a service




### <a name="Service.LimitedHits">func</a> (\*Service) [LimitedHits](/src/target/hits.go?s=1406:1459#L47)
``` go
func (s *Service) LimitedHits() ([]*types.Hit, error)
```
LimitedHits returns the last 1024 successful/online 'hit' records for a service




### <a name="Service.Online24">func</a> (\*Service) [Online24](/src/target/services.go?s=2922:2958#L105)
``` go
func (s *Service) Online24() float32
```
Online24 returns the service's uptime percent within last 24 hours




### <a name="Service.OnlineSince">func</a> (\*Service) [OnlineSince](/src/target/services.go?s=3122:3174#L111)
``` go
func (s *Service) OnlineSince(ago time.Time) float32
```
OnlineSince accepts a time since parameter to return the percent of a service's uptime.




### <a name="Service.Select">func</a> (\*Service) [Select](/src/target/services.go?s=1001:1042#L35)
``` go
func (s *Service) Select() *types.Service
```
Select will return the *types.Service struct for Service




### <a name="Service.SmallText">func</a> (\*Service) [SmallText](/src/target/services.go?s=4172:4208#L154)
``` go
func (s *Service) SmallText() string
```
SmallText returns a short description about a services status




### <a name="Service.Sum">func</a> (\*Service) [Sum](/src/target/hits.go?s=2458:2498#L79)
``` go
func (s *Service) Sum() (float64, error)
```
Sum returns the added value Latency for all of the services successful hits.




### <a name="Service.ToJSON">func</a> (\*Service) [ToJSON](/src/target/services.go?s=2414:2447#L86)
``` go
func (s *Service) ToJSON() string
```
ToJSON will convert a service to a JSON string




### <a name="Service.TotalFailures">func</a> (\*Service) [TotalFailures](/src/target/failures.go?s=3270:3319#L116)
``` go
func (s *Service) TotalFailures() (uint64, error)
```
TotalFailures returns the total amount of failures for a service




### <a name="Service.TotalFailures24">func</a> (\*Service) [TotalFailures24](/src/target/failures.go?s=3071:3122#L110)
``` go
func (s *Service) TotalFailures24() (uint64, error)
```
TotalFailures24 returns the amount of failures for a service within the last 24 hours




### <a name="Service.TotalFailuresSince">func</a> (\*Service) [TotalFailuresSince](/src/target/failures.go?s=3544:3611#L124)
``` go
func (s *Service) TotalFailuresSince(ago time.Time) (uint64, error)
```
TotalFailuresSince returns the total amount of failures for a service since a specific time/date




### <a name="Service.TotalHits">func</a> (\*Service) [TotalHits](/src/target/hits.go?s=1889:1934#L63)
``` go
func (s *Service) TotalHits() (uint64, error)
```
TotalHits returns the total amount of successful hits a service has




### <a name="Service.TotalHitsSince">func</a> (\*Service) [TotalHitsSince](/src/target/hits.go?s=2134:2197#L71)
``` go
func (s *Service) TotalHitsSince(ago time.Time) (uint64, error)
```
TotalHitsSince returns the total amount of hits based on a specific time/date




### <a name="Service.TotalUptime">func</a> (\*Service) [TotalUptime](/src/target/services.go?s=8292:8330#L289)
``` go
func (s *Service) TotalUptime() string
```
TotalUptime returns the total uptime percent of a service




### <a name="Service.Update">func</a> (\*Service) [Update](/src/target/services.go?s=9766:9810#L342)
``` go
func (u *Service) Update(restart bool) error
```
Update will update a service in the database, the service's checking routine can be restarted by passing true




### <a name="Service.UpdateSingle">func</a> (\*Service) [UpdateSingle](/src/target/services.go?s=9541:9598#L337)
``` go
func (u *Service) UpdateSingle(attr ...interface{}) error
```
UpdateSingle will update a single column for a service




## <a name="ServiceOrder">type</a> [ServiceOrder](/src/target/core.go?s=4393:4435#L158)
``` go
type ServiceOrder []types.ServiceInterface
```
ServiceOrder will reorder the services based on 'order_id' (Order)










### <a name="ServiceOrder.Len">func</a> (ServiceOrder) [Len](/src/target/core.go?s=4437:4468#L160)
``` go
func (c ServiceOrder) Len() int
```



### <a name="ServiceOrder.Less">func</a> (ServiceOrder) [Less](/src/target/core.go?s=4567:4608#L162)
``` go
func (c ServiceOrder) Less(i, j int) bool
```



### <a name="ServiceOrder.Swap">func</a> (ServiceOrder) [Swap](/src/target/core.go?s=4497:4533#L161)
``` go
func (c ServiceOrder) Swap(i, j int)
```



## <a name="User">type</a> [User](/src/target/users.go?s=819:852#L26)
``` go
type User struct {
    *types.User
}

```






### <a name="AuthUser">func</a> [AuthUser](/src/target/users.go?s=2574:2628#L92)
``` go
func AuthUser(username, password string) (*User, bool)
```
AuthUser will return the User and a boolean if authentication was correct.
AuthUser accepts username, and password as a string


### <a name="ReturnUser">func</a> [ReturnUser](/src/target/users.go?s=911:947#L31)
``` go
func ReturnUser(u *types.User) *User
```
ReturnUser returns *core.User based off a *types.User


### <a name="SelectAllUsers">func</a> [SelectAllUsers](/src/target/users.go?s=2229:2267#L81)
``` go
func SelectAllUsers() ([]*User, error)
```
SelectAllUsers returns all users


### <a name="SelectUser">func</a> [SelectUser](/src/target/users.go?s=1025:1065#L36)
``` go
func SelectUser(id int64) (*User, error)
```
SelectUser returns the User based on the user's ID.


### <a name="SelectUsername">func</a> [SelectUsername](/src/target/users.go?s=1206:1257#L43)
``` go
func SelectUsername(username string) (*User, error)
```
SelectUser returns the User based on the user's username





### <a name="User.AfterFind">func</a> (\*User) [AfterFind](/src/target/database.go?s=3215:3253#L111)
``` go
func (u *User) AfterFind() (err error)
```
AfterFind for USer will set the timezone




### <a name="User.BeforeCreate">func</a> (\*User) [BeforeCreate](/src/target/database.go?s=4056:4097#L145)
``` go
func (u *User) BeforeCreate() (err error)
```
BeforeCreate for User will set CreatedAt to UTC




### <a name="User.Create">func</a> (\*User) [Create](/src/target/users.go?s=1790:1828#L64)
``` go
func (u *User) Create() (int64, error)
```
Create will insert a new user into the database




### <a name="User.Delete">func</a> (\*User) [Delete](/src/target/users.go?s=1434:1463#L51)
``` go
func (u *User) Delete() error
```
Delete will remove the user record from the database




### <a name="User.Update">func</a> (\*User) [Update](/src/target/users.go?s=1555:1584#L56)
``` go
func (u *User) Update() error
```
Update will update the user's record in database








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
