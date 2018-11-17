# Statup Help
Statup is an easy to use Status Page monitor for your websites and applications. Statup is developed in Go Language and you are able to create custom plugins with it!

<p>
    <a href="https://github.com/hunterlong/statup"><img src="https://img.shields.io/github/stars/hunterlong/statup.svg?style=social&label=Stars"></a>
    <a href="https://github.com/hunterlong/statup"><img src="https://img.shields.io/docker/build/hunterlong/statup.svg"></a>
    <a href="https://github.com/hunterlong/statup"><img src="https://img.shields.io/github/release/hunterlong/statup.svg"></a>
</p>

# Core Elements
Statup is continuing being updated and has many awesome features to help you catch issues if your servers go down.

## Services
For each website and application you want to add a new Service. Each Service will require a URL endpoint to test your applications status.
You can also add expected HTTP responses (regex allow), expected HTTP response codes, and other fields to make sure your service is online or offline.

## Settings
Changing variables for your Statup instance is fairly simple and quick. You can change the footer HTML/text, domain of server, and many other aspects.
The guide below will explain each setting feature.

#### Export Assets
The single Statup binary file includes all assets used for the web page. Go the the Theme Editor in Settings and click Enable Assets.
This will create a 'assets' folder in the working directory, it will create all the assets used into their own folders.

#### Custom Design
After you've exported the assets you can edit the CSS directly or use the Theme Editor to customize the SASS design.
Statup uses sass to generate CSS files from SASS. You can install sass with a command below.

- node: `npm install sass -g`
- ruby: `gem install sass`

#### CDN Assets
If you want to host the Statup assets from our CDN rather than from your local instance, enable "Use CDN" toggle switch on the Settings page.

## Notifications
Statup includes a few notification methods to receive alerts when a service is online/offline. Each notifier is different, users can create your own notifier and send a Push Request to github.

## Users
Administrators can add, update, and remove all elements on your Statup instance. Other users can only view the status page and

## Plugins
Creating a plugin for Statup is not that difficult, if you know a little bit of Go Language you can create any type of application to be embedded into the Status framework.
Checkout the example plugin that includes all the interfaces, information, and custom HTTP routing at <a href="https://github.com/hunterlong/statup_plugin">https://github.com/hunterlong/statup_plugin</a>.
Anytime there is an action on your status page, all of your plugins will be notified of the change with the values that were changed or created.
<p></p>
Using the statup/plugin Golang package you can quickly implement the event listeners. Statup uses <a href="https://github.com/upper/db">upper.io/db.v3</a> for the database connection.
You can use the database inside of your plugin to create, update, and destroy tables/data. <b>Please only use respectable plugins!</b>

# API Usage
Statup includes a RESTFUL API so you can view, update, and edit your services with easy to use routes. You can currently view, update and delete services, view, create, update users, and get detailed information about the Statup instance. To make life easy, try out a Postman or Swagger JSON file and use it on your Statup Server.

<p align="center">
<a href="/files/postman.json">Postman Export</a> | <a href="/files/swagger.json">Swagger Export</a> | <a href="https://app.swaggerhub.com/apis/statup/statup/1">Swagger Hub</a>
</p>

### API Authentication
Authentication uses the Statup API Secret to accept remote requests. You can find the API Secret in the Settings page of your Statup server. To send requests to your Statup API, include a Authorization Header when you send the request. The API will accept any one of the headers below.

- HTTP Header: `Authorization: API SECRET HERE`
- HTTP Header: `Authorization: Bearer API SECRET HERE`

# Prometheus Exporter
Statup includes a prometheus exporter so you can have even more monitoring power with your services. The prometheus exporter can be seen on `/metrics`, simply create another exporter in your prometheus config. Use your Statup API Secret for the Authorization Bearer header, the `/metrics` URL is dedicated for Prometheus and requires the correct API Secret has `Authorization` header.

# Grafana Dashboard
Statup has a Grafana Dashboard that you can quickly implement if you've added your Statup service to Prometheus. Import Dashboard ID: `6950` into your Grafana dashboard and watch the metrics come in!

<p align="center">
<a href="https://grafana.com/dashboards/6950">Grafana Dashboard</a> | <a href="/files/grafana.json">Grafana JSON Export</a>
<br>
<img width="80%" src="https://img.cjx.io/statupgrafana.png">
</p>

## Basic Prometheus Exporter
If you have Statup and the Prometheus server in the same Docker network, you can use the yaml config below.
``` yaml
scrape_configs:
  - job_name: 'statup'
    scrape_interval: 30s
    bearer_token: 'SECRET API KEY HERE'
    static_configs:
      - targets: ['statup:8080']
```

## Remote URL Prometheus Exporter
This exporter yaml below has `scheme: https`, which you can remove if you arn't using HTTPS.
``` yaml
scrape_configs:
  - job_name: 'statup'
    scheme: https
    scrape_interval: 30s
    bearer_token: 'SECRET API KEY HERE'
    static_configs:
      - targets: ['status.mydomain.com']
```

### `/metrics` Output
```
statup_total_failures 206
statup_total_services 4
statup_service_failures{id="1" name="Google"} 0
statup_service_latency{id="1" name="Google"} 12
statup_service_online{id="1" name="Google"} 1
statup_service_status_code{id="1" name="Google"} 200
statup_service_response_length{id="1" name="Google"} 10777
statup_service_failures{id="2" name="Statup.io"} 0
statup_service_latency{id="2" name="Statup.io"} 3
statup_service_online{id="2" name="Statup.io"} 1
statup_service_status_code{id="2" name="Statup.io"} 200
statup_service_response_length{id="2" name="Statup.io"} 2
```

# Static HTML Exporter
You might have a server that won't allow you to run command that run longer for 60 seconds, or maybe you just want to export your status page to a static HTML file. Using the Statup exporter you can easily do this with 1 command.

```
statup export
```
###### 'index.html' is created in current directory with static CDN url's.

## Push to Github
Once you have the `index.html` file, you could technically send it to an FTP server, Email it, Pastebin it, or even push to your Github repo for Status updates directly from repo.

```bash
git add index.html
git commit -m "Updated Status Page"
git push -u origin/master
```

# Config with .env File
It may be useful to load your environment using a `.env` file in the root directory of your Statup server. The .env file will be automatically loaded on startup and will overwrite all values you have in config.yml.

If you have the `DB_CONN` environment variable set Statup will bypass all values in config.yml and will require you to have the other DB_* variables in place. You can pass in these environment variables without requiring a .env file.

### `.env` File
```bash
DB_CONN=postgres
DB_HOST=0.0.0.0
DB_PORT=5432
DB_USER=root
DB_PASS=password123
DB_DATABASE=root

NAME=Demo
DESCRIPTION=This is an awesome page
DOMAIN=https://domain.com
ADMIN_USER=admin
ADMIN_PASS=admin
ADMIN_EMAIL=info@admin.com
USE_CDN=true

IS_DOCKER=false
IS_AWS=false
SASS=/usr/local/bin/sass
CMD_FILE=/bin/bash
```
This .env file will include additional variables in the future, subscribe to this repo to keep up-to-date with changes and updates.

# Makefile
Here's a simple list of Makefile commands you can run using `make`. The [Makefile](https://github.com/hunterlong/statup/blob/master/Makefile) may change often, so i'll try to keep this Wiki up-to-date.

- Ubuntu `apt-get install build-essential`
- MacOSX `sudo xcode-select -switch /Applications/Xcode.app/Contents/Developer`
- Windows [Install Guide for GNU make utility](http://gnuwin32.sourceforge.net/packages/make.htm)
- CentOS/RedHat `yum groupinstall "Development Tools"`

### Commands
``` bash
make build                         # build the binary
make install
make run
make test
make coverage
make docs
# Building Statup
make build-all
make build-alpine
make docker
make docker-run
make docker-dev
make docker-run-dev
make databases
make dep
make dev-deps
make clean
make compress
make cypress-install
make cypress-test
```

# Testing
* If you want to test your updates with the current golang testing units, you can follow the guide below to run a full test process. Each test for Statup will run in MySQL, Postgres, and SQlite to make sure all database types work correctly.

## Create Docker Databases
The easiest way to run the tests on all 3 databases is by starting temporary databases servers with Docker. Docker is available for Linux, Mac and Windows. You can download/install it by going to the [Docker Installation](https://docs.docker.com/install/) site.

``` bash
docker run -it -d \
   -p 3306:3306 \
   -env MYSQL_ROOT_PASSWORD=password123 \
   -env MYSQL_DATABASE=root mysql
```

``` bash
docker run -it -d \
   -p 5432:5432 \
   -env POSTGRES_PASSWORD=password123 \
   -env POSTGRES_USER=root \
   -env POSTGRES_DB=root postgres
```

Once you have MySQL and Postgres running, you can begin the testing. SQLite database will automatically create a `statup.db` file and will delete after testing.

## Run Tests
Insert the database environment variables to auto connect the the databases and run the normal test command: `go test -v`. You'll see a verbose output of each test. If all tests pass, make a push request! 💃
``` bash
DB_DATABASE=root \
   DB_USER=root \
   DB_PASS=password123 \
   DB_HOST=localhost \
   go test -v
```
