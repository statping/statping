# 0.90.75 (12-20-2020)
- Removed favicons and PNG files from assets, now using base64 images
- Cleaned up some issues with UI

# 0.90.74 (12-18-2020)
- Fixed issue with favicon/manifest.json throwing 404 errors
- Modified language go:generate script to slowdown for 429 errors
- Removed Sentry error logging functionality

# 0.90.73 (12-15-2020)
- Removed complexity in code for http server
- Removed internal cache functionality (not useful, needs refactor)
- Merged PR https://github.com/statping/statping/pull/909
- Merged PR https://github.com/statping/statping/pull/880
- Merged PR https://github.com/statping/statping/pull/859

# 0.90.72 (10-28-2020)
- Fixed issue with graphs becoming stuck on reload

# 0.90.71 (10-13-2020)
- Reverted Docker user in Dockerfile

# 0.90.70 (10-1-2020)
- Merged PR #806 - Enhance GRPC Monitoring
- Merged PR #692 - When login fields are autofilled the sign in button should be enabled
- Modified multiple Vue forms to use number models for integer inputs
- Fixed page freeze issue for incidents https://github.com/statping/statping/issues/842
- Modified cache routine from 5 seconds to 60 seconds

# 0.90.69 (09-18-2020)
- Fixed issue with service view not loading. #808 #811 #800

# 0.90.68 (09-17-2020)
- Added DB_DSN env for mysql, postgres or sqlite DSN database connection string
- Added READ_ONLY env for a read only connection to the database
- Added Custom OAuth OpenID toggle switch in settings (appends 'openid' in scope)
- Fixed Custom OAuth response_type issue
- Added Configs tab in Settings to edit the config.yml from frontend

# 0.90.67 (09-14-2020)
- Modified core settings to update config.yml on save
- Modified Theme Editor to restart the HTTP router on create/delete (fixing 404's)

# 0.90.66 (09-08-2020)
- Added Import and Export views in Dashboard
- Modified services list sparkline to use start/end of day timestamp
- Modified i18n language files, added go generate script to automatically translate

# 0.90.65 (09-01-2020)
- Fixed issue with dashboard not logging in (notifier panic)
- Modified static email templates to github.com/statping/emails
- Modified Regenerate API function to keep API_SECRET env
- Added DEMO_MODE env variable, if true, 'admin' cannot be deleted
- Modified Service sparklines on Dashboard
- Added modal popup for UI deletes/edits

# 0.90.64 (08-18-2020)
- Modified max-width for container to 1012px, larger UI
- Added failure sparklines in the Services list view
- Added "Update Available" alert on the top of Settings if new version is available
- Added Version and Github Commit hash to left navigation on Settings page
- Added "reason" for failures (will be used for more custom notification messages) [regex, lookup, timeout, connection, close, status_code]
- Added Help page that is generated from Statping's Wiki repo on build
- Modified Service Group failures on index page to show 90 days of failures
- Modified Service view page, updated Latency and Ping charts, added failures below
- Modified Service chart on index page to show ping data along with latency
- Added AWS SNS Notifier
- Modified dashboard services UI
- Modified service.Failures API to include 32 failures (max)

# 0.90.63 (08-17-2020)
- Modified build process to use xgo for all arch builds
- Modified Statping's Push Notifications server notifier to match with Firebase/gorush params

# 0.90.62 (08-07-2020)
- Added Notification logs
- Fixed issues with Notifer After (x) failures for notifications
- Modified notifications to not send on initial startup
- Updated Incident UI
- Added additional testing for notifications
- Modified SCSS/SASS files to be generated from 1, main.scss to main.css
- Modified index page to use /assets directory for assets, (main.css, style.css)
- Modified index page to use CDN asset paths
- Fixed New Checkin form
- Modified email notifier template to be rendered from MJML (using go generate)
- Modified database relationships with services using gorm
- Modified "statping env" command to show user/group ID
- Removed "js" folder when exporting assets, js files are always version of release, not static JS files

# 0.90.61 (07-22-2020)
- Modified sass layouts, organized and split up sections
- Modified Checkins to seconds rather than milliseconds (for cronjob)
- Modified Service View page to show data inside cards
- Fixed issue with uptime_data sending incorrect start/end timestamps
- Modified http cache to bypass if url has a "v" query param
- Added "Static Services" (a fake service that requires you to update the online/offline status)
- Added Update Static Service PATCH route (/api/services/{id})
- Modified SASS api endpoints (base, layout, forms, mixins, mobile, variables)
- Added additional testing
- Modified node version from 10.x to 12.18.2
- Modified Notifier's struct values to be NullString and NullInt to allow empty values
- Added Search ability to Logs in UI
- Fixed issue with Incidents and Checkins not being deleted once service is deleted

# 0.90.60 (07-15-2020)
- Added LETSENCRYPT_ENABLE (boolean) env to enable/disable letsencrypt SSL

# 0.90.59 (07-14-2020)
- Added LetsEncrypt SSL Generator by using LETSENCRYPT_HOST and LETSENCRYPT_EMAIL envs.
- Modified JWT token key to be sha256 of API Secret
- Modified github actions to build multi-arch Docker images
- Added "update" command to install latest version
- Fixed dashboard uptime_data API request to request correct start/time timestamp

# 0.90.58 (07-09-2020)
- Fixed ICMP latency/ping durations
- Fixed webhook notifier
- Modified file structure for Vue admin dashboard components.
- Added Gotify notifier

# 0.90.57 (07-04-2020)
- Fixed login issue

# 0.90.56 (06-25-2020)
- Modified metrics now include service name for each service metric
- Added switch for true/false notifier values
- Added list for notifiers that have static values (in drop down)
- Fixed oAuth form saving
- Fixed some HTTP Cookie issues
- Added error if Theme Editor returns an error from API
- Added Pushover priority and sounds
- Added HTTP headers for outgoing requests (includes User-Agent=Statping and Statping-Version=0.90.55)
- Fixed Google oAuth handling
- Added Google oAuth email/domain user restrictions
- Modified notifiers to use dereferenced services and failures
- Added core.Example() function for testing
- Added Custom oAuth Authentication method
- Fixed setup form not creating user from values inputted in form
- Fixed issues with Telegram Notifier
- Modified notifier test handler to return notifier based on URL, not JSON payload

# 0.90.55 (06-18-2020)
- Added 404 page
- Modified Statping's PR process, dev -> master
- Fixed Discord notifier
- Modified email template for SMTP emails
- Added OnSave() method for all notifiers

# 0.90.54 (06-17-2020)
- Fixed Slack Notifier's failure/success data saving issue
- Added additional i18n Languages (help needed!)

# 0.90.53 (06-16-2020)
- Modified most of the key's for prometheus metrics
- Added Database Stats in prometheus metrics
- Added object query counts in prometheus metrics

# 0.90.52 (06-15-2020)
- Fixed NOT NULL sql field

# 0.90.51 (06-15-2020)
- Fix Theme Editor codemirror inputs to show on load
- Added favicon folder for local assets can be used without remote access
- Modified Notifier's to return the response as a string for the frontend
- Modified Notifiers so they can use custom data for their request
- Added Notifier OnSuccess and onFailure custom data on frontend

# 0.90.50 (06-13-2020)
- Removed PORT, replaced with SERVER_PORT
- Removed HOST/IP, replaced with SERVER_IP

# 0.90.49 (06-12-2020)
- Added additional prometheus /metrics for better debugging

# 0.90.48 (06-11-2020)
- Modified shutdown routine to make command exit with code 0
- Modified install.sh for correct installation

# 0.90.47 (06-10-2020)
- Fixed Urgent bug taking 100% of CPU (Timer)
- Modified HTTP server, now in it's own go routine/channel
- Fixed Service form for editing
- Added pprof golang debugging http server if `GO_ENV` == "test"
- Added `HOST` env variable (hostname for http server)
- Added `DISABLE_HTTP` env variable (defaults to false, disables the http server)
- Added `DISABLE_COLORS` env variable (default to false, disables color encoding for logs)
- Added `LOGS_MAX_COUNT`
- Added `LOGS_MAX_AGE`
- Added `LOGS_MAX_SIZE`
- Added `DEBUG` (starts a pprof golang debugging http server on port 9090, defaults to false)
- Confirmed `DISABLE_LOGS` is working
- Modified Mobile Notifier to fit new push notification server endpoint
- PR Merged: Fix time conversion in overview and charts #645
- PR Merged: Wait for cmd reads to complete before calling Wait() #626
- PR Merged: separate command options and option arguments #623

# 0.90.46 (06-04-2020)
- Add i18n language translations for frontend
- Added PR for heatmap https://github.com/statping/statping/pull/589
- Added Statping newsletter option during /setup mode
- Fix for disabling logs with `DISABLE_LOGS` env

# 0.90.45 (06-01-2020)
- Merged PR [#612](https://github.com/statping/statping/pull/612) for edit/create service issue.

# 0.90.44 (05-25-2020)
- Modified Makefile to include "netgo" tag during golang build

# 0.90.43 (05-21-2020)
- Fixed service TLS checkbox form for edit and create
- Modified ICMP ping's to use system's "ping" command (doesn't need root access)

# 0.90.42 (05-20-2020)
- Fixed TCP services that dont use TLS.

# 0.90.41 (05-20-2020)
- Added TLS Client Cert/Key feature for HTTP and TCP/UDP services
- Replaced environment variable ADMIN_PASS to ADMIN_PASSWORD.

# 0.90.40 (05-18-2020)
- Fixed issues with MySQL and Postgres taking forever to insert sample data (now run in bulk)
- Removed API Authentication for /api/logout route
- Modified Core Sample/Upstart row to include NAME, DESCRIPTION, and DOMAIN environment vars (also added default values)

# 0.90.39 (05-15-2020)
- Modified some SCSS designs for services failures in group
- Fixed Twilio notifier and tests

# 0.90.38 (05-10-2020)
- Added service timeframe/interval on index charts
- Added --config flag to specify config.yml file
- Modified multiple files for simple UX fixes

# 0.90.37 (05-04-2020)
- Fixed authentication issues dealing with cookies
- Modified build process, arm/arm64 couldnt run sqlite

# 0.90.36 (05-02-2020)
- Fixed Notifier golang templating func to use correct variables

# 0.90.35 (05-01-2020)
- Fixed issue with API endpoints cannot accepting Authorization header
- Fixed issue with sass executable not being found, SASS environment var re-implemented
- Added additional Postman API doc endpoints

# 0.90.34 (04-28-2020)
- Added missing information to Mail notification ([#472](https://github.com/statping/statping/issues/472))
- Added service.yml file to auto create services (https://github.com/statping/statping/wiki/services.yml)
- Removed Core API_KEY, (unused code, use API_SECRET)

# 0.90.33 (04-24-2020)
- Fixed config loading method

# 0.90.32 (04-23-2020)
- Modified the saving and loading process config.yml

# 0.90.31 (04-21-2020)
- Version bump for github actions

# 0.90.30 (04-19-2020)
- Attempt to fix Github Actions build process
- Fix for empty database connection string, and not starting in setup mode

# 0.90.29 (04-19-2020)
- Added HTTP Redirects for services
- Removed use of SASS environment variable, now finds path or sends error
- Modified Makefile to create new snapcraft versions
- Fixed issue when logs are not initiated yet. Issue #502
- Fixed issue when SQLite (statping.db) is not found Issue #499
- Modified port flag in Docker image
- Fixed issue on startup without config.yml file not starting in setup mode

# 0.90.28 (04-16-2020)
- Fixed postgres timestamp grouping
- Added postman (newman) API testing
- Added Viper and Cobra config/env parsing package
- Added more golang tests
- Modified handlers to use a more generic find method
- Added 'env' command to show variables used in config
- Added 'reset' command that will delete files and backup .db file for a fresh install
- Added error type that has common errors with http status code based on error

# 0.90.27 (04-15-2020)
- Fixed postgres database table creation process
- Modified go build process, additional ARCHs
- Added 'SAMPLE_DATA' environment variable to disable example data on startup. (default: true)

# 0.90.26 (04-13-2020)
- Fixed Delete Failures button/function
- Removed timezone field from Settings (core)
- Modified CDN asset URL
- Fixed single Service view, more complex charts

# 0.90.25
- Added string response on OnTest for Notifiers
- Modified UI to show user the response for a Notifier.
- Modified some Notifiers title's
- Added more Cypress e2e testing
- Modified Incidents form and UX.
- Added /api/services/{id}/uptime_data API endpoint to show online/offline durations as a series for charts.
- Modified index page to automatically refresh Service details on interval

# 0.90.24
- Fixed login form from not showing

# 0.90.23
- Added Incident Reporting
- Added Cypress tests
- Added Github and Google OAuth login (beta)
- Added Delete All Failures
- Added Checkin form
- Added Pushover notifier

# 0.90.22
- Added range input types for integer form fields
- Modified Sentry error logging details
- Modified form field layouts for better UX.
- Modified Notifier form
- Fixed Notifier Test form and logic

# 0.90.21
- Fixed BASE_PATH when using a path for Statping
- Added Cypress testing
- Modified SQLite golang package
- Modified SQLite connection limit, and idle limit. (defaults to 25)
- Fixed installation to use project name and description from form post

# 0.90.20
- Fixed Service Form from sending integer values as strings to API
- Added Cypress e2e testing (working on adding more)

# 0.90.19
- Fixed private Services from showing in API (/api/services and /api/services/{id})
- Removed unused code

# 0.90.18
- Added service type gRPC, you can now check on gRPC services. (limited)

# 0.90.17
- Fixed notification fields for frontend
- Fixed notification JSON form to send integer if value is an integer.
- Added testing for notifiers

# 0.90.16
- Added Notify After (int) field for Services. Will send notifications after x amount of failures.
- Added new method in utils package for replacing `{{.Service.*}}` and `{{.Failure.*}}` variables from string to it's true value
- Fixed Notifer get endpoint
- Cleaned Notifier methods
- Updated recommended changes from [sonarcloud.io](https://sonarcloud.io/organizations/statping/projects)
- Organized utils package files

# 0.90.15
- Fixed /dashboard authentication state to show admin tabs if your an admin. [Issue #438](https://github.com/statping/statping/issues/438)
- Fixed Cache JS error on Dashboard

# 0.90.14
- Updated SCSS compiling, and confirmed it works.
- Added `$container-color` SCSS variable.
- Fixed issue with JWT token (nil pointer) for the Cookie name

# 0.90.13
- Added new function `utils.RenameDirectory` to rename directory
- Added new function `(*DbConfig) BackupAssets` to backup a customized theme and place into a directory named `assets_backup`. Only for migration 0.80 to 0.90+, entirely new frontend.
- Updated JS function `convertToChartData` to return an empty chart data if API response was empty.
- Updated `banner.png` to make a bit smaller, (680px)
- Fixed method that returns `no such table: services` on startup, check table first.
- Fixed version from not being added into Core table. [Issue #436](https://github.com/statping/statping/issues/436)

# 0.90.12
- Fixed MySQL timestamp formatting. (issue #432)
