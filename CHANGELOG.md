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