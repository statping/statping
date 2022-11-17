module github.com/statping-ng/statping-ng

// +heroku goVersion go1.17
go 1.16

require (
	github.com/GeertJohan/go.rice v1.0.2
	github.com/aws/aws-sdk-go v1.30.20
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/structs v1.1.0
	github.com/foomo/simplecert v1.7.5
	github.com/foomo/tlsconfig v0.0.0-20180418120404-b67861b076c9
	github.com/getsentry/sentry-go v0.5.1
	github.com/go-mail/mail v2.3.1+incompatible
	github.com/gorilla/mux v1.7.4
	github.com/hako/durafmt v0.0.0-20200605151348-3a43fc422dd9
	github.com/jinzhu/gorm v1.9.12
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.1.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	github.com/statping-ng/emails v1.0.3
	github.com/stretchr/testify v1.6.1
	github.com/t-tiger/gorm-bulk-insert/v2 v2.0.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/grpc v1.28.1
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.3.0
)

require (
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-ping/ping v1.1.0
	github.com/google/uuid v1.3.0 // indirect
	golang.org/x/net v0.2.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)
