module github.com/statping/statping

// +heroku goVersion go1.14
go 1.14

require (
	github.com/GeertJohan/go.rice v1.0.0
	github.com/ararog/timeago v0.0.0-20160328174124-e9969cf18b8d
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dustin/go-humanize v1.0.0
	github.com/fatih/structs v1.1.0
	github.com/foomo/simplecert v1.7.5
	github.com/foomo/tlsconfig v0.0.0-20180418120404-b67861b076c9
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/getsentry/sentry-go v0.5.1
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/securecookie v1.1.1
	github.com/hako/durafmt v0.0.0-20200605151348-3a43fc422dd9
	github.com/jinzhu/gorm v1.9.12
	github.com/magiconair/properties v1.8.1
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/mitchellh/mapstructure v1.2.2 // indirect
	github.com/pelletier/go-toml v1.7.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/russross/blackfriday/v2 v2.0.1
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.5.1
	github.com/t-tiger/gorm-bulk-insert/v2 v2.0.1
	golang.org/x/crypto v0.0.0-20200420201142-3c4aac89819a
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.28.1
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.2.8
)
