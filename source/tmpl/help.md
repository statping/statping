# Statup Help
Statup is an easy to use Status Page monitor for your websites and applications. Statup is developed in Go Language and you are able to create custom plugins with it!

<p>
    <a href="https://github.com/hunterlong/statup"><img src="https://img.shields.io/github/stars/hunterlong/statup.svg?style=social&label=Stars"></a>
    <a href="https://github.com/hunterlong/statup"><img src="https://img.shields.io/docker/build/hunterlong/statup.svg"></a>
    <a href="https://github.com/hunterlong/statup"><img src="https://img.shields.io/github/release/hunterlong/statup.svg"></a>
</p>

# Services
For each website and application you want to add a new Service. Each Service will require a URL endpoint to test your applications status.
You can also add expected HTTP responses (regex allow), expected HTTP response codes, and other fields to make sure your service is online or offline.

# Statup Settings
You can change multiple settings in your Statup instance.

# Users
Users can access the Statup Dashboard to add, remove, and view services.

# Plugins
Creating a plugin for Statup is not that difficult, if you know a little bit of Go Language you can create any type of application to be embedded into the Status framework.
Checkout the example plugin that includes all the interfaces, information, and custom HTTP routing at <a href="https://github.com/hunterlong/statup_plugin">https://github.com/hunterlong/statup_plugin</a>.
Anytime there is an action on your status page, all of your plugins will be notified of the change with the values that were changed or created.
<p></p>
Using the statup/plugin Golang package you can quickly implement the event listeners. Statup uses <a href="https://github.com/upper/db">upper.io/db.v3</a> for the database connection.
You can use the database inside of your plugin to create, update, and destroy tables/data. <b>Please only use respectable plugins!</b>

# Custom Stlying
On Statup Status Page server can you create your own custom stylesheet to be rendered on the index view of your status page. Go to <a href="/settings">Settings</a> and click on Custom Styling.

# API Endpoints
Statup includes a RESTFUL API so you can view, update, and edit your services with easy to use routes. You can currently view, update and delete services, view, create, update users, and get detailed information about the Statup instance. To make life easy, try out a Postman or Swagger JSON file and use it on your Statup Server.

<p align="center">
<a href="https://github.com/hunterlong/statup/blob/master/dev/postman.json">Postman JSON Export</a> | <a href="https://github.com/hunterlong/statup/blob/master/dev/swagger.json">Swagger Export</a>
</p>

## Authentication
Authentication uses the Statup API Secret to accept remote requests. You can find the API Secret in the Settings page of your Statup server. To send requests to your Statup API, include a Authorization Header when you send the request. The API will accept any one of the headers below.

- HTTP Header: `Authorization: API SECRET HERE`
- HTTP Header: `Authorization: Bearer API SECRET HERE`

## Main Route `/api`
The main API route will show you all services and failures along with them.

## Services
The services API endpoint will show you detailed information about services and will allow you to edit/delete services with POST/DELETE http methods.

### Viewing All Services
- Endpoint: `/api/services`
- Method: `GET`
- Response: Array of [Services](https://github.com/hunterlong/statup/wiki/API#service-response)
- Response Type: `application/json`
- Request Type: `application/json`

### Viewing Service
- Endpoint: `/api/services/{id}`
- Method: `GET`
- Response: [Service](https://github.com/hunterlong/statup/wiki/API#service-response)
- Response Type: `application/json`
- Request Type: `application/json`

### Updating Service
- Endpoint: `/api/services/{id}`
- Method: `POST`
- Response: [Service](https://github.com/hunterlong/statup/wiki/API#service-response)
- Response Type: `application/json`
- Request Type: `application/json`

POST Data:
``` json
{
    "name": "Updated Service",
    "domain": "https://google.com",
    "expected": "",
    "expected_status": 200,
    "check_interval": 15,
    "type": "http",
    "method": "GET",
    "post_data": "",
    "port": 0,
    "timeout": 10,
    "order_id": 0
}
```

### Deleting Service
- Endpoint: `/api/services/{id}`
- Method: `DELETE`
- Response: [Object Response](https://github.com/hunterlong/statup/wiki/API#object-response)
- Response Type: `application/json`
- Request Type: `application/json`

Response:
``` json
{
    "status": "success",
    "id": 4,
    "type": "service",
    "method": "delete"
}
```

## Users
The users API endpoint will show you users that are registered inside your Statup instance.

### View All Users
- Endpoint: `/api/users`
- Method: `GET`
- Response: Array of [Users](https://github.com/hunterlong/statup/wiki/API#user-response)
- Response Type: `application/json`
- Request Type: `application/json`

### Viewing User
- Endpoint: `/api/users/{id}`
- Method: `GET`
- Response: [User](https://github.com/hunterlong/statup/wiki/API#user-response)
- Response Type: `application/json`
- Request Type: `application/json`

### Creating New User
- Endpoint: `/api/users`
- Method: `POST`
- Response: [User](https://github.com/hunterlong/statup/wiki/API#user-response)
- Response Type: `application/json`
- Request Type: `application/json`

POST Data:
``` json
{
    "username": "newadmin",
    "email": "info@email.com",
    "password": "password123",
    "admin": true
}
```

### Updating User
- Endpoint: `/api/users/{id}`
- Method: `POST`
- Response: [User](https://github.com/hunterlong/statup/wiki/API#user-response)
- Response Type: `application/json`
- Request Type: `application/json`

POST Data:
``` json
{
    "username": "updatedadmin",
    "email": "info@email.com",
    "password": "password123",
    "admin": true
}
```

### Deleting User
- Endpoint: `/api/services/{id}`
- Method: `DELETE`
- Response: [Object Response](https://github.com/hunterlong/statup/wiki/API#object-response)
- Response Type: `application/json`
- Request Type: `application/json`

Response:
``` json
{
    "status": "success",
    "id": 3,
    "type": "user",
    "method": "delete"
}
```

# Service Response
``` json
{
    "id": 8,
    "name": "Test Service 0",
    "domain": "https://status.coinapp.io",
    "expected": "",
    "expected_status": 200,
    "check_interval": 1,
    "type": "http",
    "method": "GET",
    "post_data": "",
    "port": 0,
    "timeout": 30,
    "order_id": 0,
    "created_at": "2018-09-12T09:07:03.045832088-07:00",
    "updated_at": "2018-09-12T09:07:03.046114305-07:00",
    "online": false,
    "latency": 0.031411064,
    "24_hours_online": 0,
    "avg_response": "",
    "status_code": 502,
    "last_online": "0001-01-01T00:00:00Z",
    "dns_lookup_time": 0.001727175,
    "failures": [
        {
            "id": 5187,
            "issue": "HTTP Status Code 502 did not match 200",
            "created_at": "2018-09-12T10:41:46.292277471-07:00"
        },
        {
            "id": 5188,
            "issue": "HTTP Status Code 502 did not match 200",
            "created_at": "2018-09-12T10:41:47.337659862-07:00"
        }
    ]
}
```

# User Response
``` json
{
    "id": 1,
    "username": "admin",
    "api_key": "02f324450a631980121e8fd6ea7dfe4a7c685a2f",
    "admin": true,
    "created_at": "2018-09-12T09:06:53.906398511-07:00",
    "updated_at": "2018-09-12T09:06:54.972440207-07:00"
}
```

# Object Response
``` json
{
    "type": "service",
    "id": 19,
    "method": "delete",
    "status": "success"
}
```

# Main API Response
``` json
{
    "name": "Awesome Status",
    "description": "An awesome status page by Statup",
    "footer": "This is my custom footer",
    "domain": "https://demo.statup.io",
    "version": "v0.56",
    "migration_id": 1536768413,
    "created_at": "2018-09-12T09:06:53.905374829-07:00",
    "updated_at": "2018-09-12T09:07:01.654201225-07:00",
    "database": "sqlite",
    "started_on": "2018-09-12T10:43:07.760729349-07:00",
    "services": [
        {
            "id": 1,
            "name": "Google",
            "domain": "https://google.com",
            "expected": "",
            "expected_status": 200,
            "check_interval": 10,
            "type": "http",
            "method": "GET",
            "post_data": "",
            "port": 0,
            "timeout": 10,
            "order_id": 0,
            "created_at": "2018-09-12T09:06:54.97549122-07:00",
            "updated_at": "2018-09-12T09:06:54.975624103-07:00",
            "online": true,
            "latency": 0.09080986,
            "24_hours_online": 0,
            "avg_response": "",
            "status_code": 200,
            "last_online": "2018-09-12T10:44:07.931990439-07:00",
            "dns_lookup_time": 0.005543935
        }
    ]
}
```

# Prometheus Exporter
Statup includes a prometheus exporter so you can have even more monitoring power with your services. The prometheus exporter can be seen on `/metrics`, simply create another exporter in your prometheus config. Use your Statup API Secret for the Authorization Bearer header, the `/metrics` URL is dedicated for Prometheus and requires the correct API Secret has `Authorization` header.

# Grafana Dashboard
Statup has a [Grafana Dashboard](https://grafana.com/dashboards/6950) that you can quickly implement if you've added your Statup service to Prometheus. Import Dashboard ID: `6950` into your Grafana dashboard and watch the metrics come in!

<p align="center"><img width="80%" src="https://img.cjx.io/statupgrafana.png"></p>

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

## `.env` File
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

## Testing
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
