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
