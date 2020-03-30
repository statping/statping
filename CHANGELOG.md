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
