<p align="center">
<img width="60%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statping.png">
</p>
<p align="center">
    <b>Statping - Web and App Status Monitoring for Any Type of Project</b><br>
<a href="https://github.com/hunterlong/statping/wiki">View Wiki</a> | <a href="https://demo.statping.com">Demo</a> | <a href="https://itunes.apple.com/us/app/apple-store/id1445513219">iPhone</a> | <a href="https://play.google.com/store/apps/details?id=com.statping">Android</a> <br> <a href="https://github.com/hunterlong/statping/wiki/API">API</a> | <a href="https://github.com/hunterlong/statping/wiki/Docker">Docker</a> | <a href="https://github.com/hunterlong/statping/wiki/AWS-EC2">EC2</a> | <a href="https://github.com/hunterlong/statping/wiki/Mac">Mac</a> | <a href="https://github.com/hunterlong/statping/wiki/Linux">Linux</a> | <a href="https://github.com/hunterlong/statping/wiki/Windows">Windows</a> | <a href="https://github.com/hunterlong/statping/wiki/Statping-Plugins">Plugins</a>
</p>

# Statping - Status Page & Monitoring Server
An easy to use Status Page for your websites and applications. Statping will automatically fetch the application and render a beautiful status page with tons of features for you to build an even better status page. This Status Page generator allows you to use MySQL, Postgres, or SQLite on multiple operating systems.

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/hunterlong/statping) [![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/statping/general) [![](https://images.microbadger.com/badges/image/hunterlong/statping.svg)](https://microbadger.com/images/hunterlong/statping) [![Docker Pulls](https://img.shields.io/docker/pulls/hunterlong/statping.svg)](https://hub.docker.com/r/hunterlong/statping/builds/)

## A Future-Proof Status Page
Statping strives to remain future-proof and remain intact if a failure is created. Your Statping service should not be running on the same instance you're trying to monitor. If your server crashes your Status Page should still remaining online to notify your users of downtime.

<p align="center">
<img width="80%" src="https://img.cjx.io/statupsiterun.gif">
</p>

## Lightweight and Fast
Statping is a very lightweight application and is available for Linux, Mac, and Windows. The Docker image is only ~16Mb so you know that this application won't be filling up your hard drive space.
The Status binary for all other OS's is ~17Mb at most.

## No Requirements
Statping is built in Go Language so all you need is the precompile binary based on your operating system. You won't need to install anything extra once you have the Statping binary installed. You can even run Statping on a Raspberry Pi.

<p align="center">
    <a href="https://github.com/hunterlong/statping/wiki/Linux"><img width="5%" src="https://img.cjx.io/linux.png"></a>
    <a href="https://github.com/hunterlong/statping/wiki/Mac"><img width="5%" src="https://img.cjx.io/apple.png"></a>
    <a href="https://github.com/hunterlong/statping/wiki/Windows"><img width="5%" src="https://img.cjx.io/windows.png"></a>
    <a href="https://play.google.com/store/apps/details?id=com.statping"><img width="5%" src="https://img.cjx.io/android.png"></a>
    <a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img width="5%" src="https://img.cjx.io/appstore.png"></a>
    <a href="https://hub.docker.com/r/hunterlong/statping"><img width="5%" src="https://img.cjx.io/dockericon.png"></a>
</p>


## Mobile App is Gorgeous
The Statping app is available on the App Store and Google Play for free. The app will allow you to view services, receive notifications when a service is offline, update groups, users, services, messages, and more! Start your own Statping server and then connect it to the app by scanning the QR code in settings. 

<p align="center">
<img width="80%" src="https://img.cjx.io/statping_iphone_bk.png">
</p>

<p align="center">
<a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a>
<a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a>
</p>

## Run on Any Server
Whether you're a Docker fan-boy or a [AWS EC2](https://github.com/hunterlong/statping/wiki/AWS-EC2) master, Statping gives you multiple options to simply get running. Our Amazon AMI image is only 8Gb and will automatically update to the most stable version of Statping.
Running on an EC2 server might be the most cost effective way to host your own Statping Status Page. The server runs on the smallest EC2 instance (t2.nano) AWS has to offer, which only costs around $4.60 USD a month for your dedicated Status Page.
Want to run it on your own Docker server? Awesome! Statping has multiple docker-compose.yml files to work with. Statping can automatically create a SSL Certification for your status page.

## Custom SASS Styling
Statping will allow you to completely customize your Status Page using SASS styling with easy to use variables. The Docker image actually contains a prebuilt SASS binary so you won't even need to setup anything!

<p align="center">
<img width="80%" src="https://img.cjx.io/statping_theme.gif">
</p>

## Slack, Email, Twilio and more
Statping includes email notification via SMTP and Slack integration using [Incoming Webhook](https://api.slack.com/incoming-webhooks). Insert the webhook URL into the Settings page in Statping and enable the Slack integration. Anytime a service fails, you're channel that you specified on Slack will receive a message.

## User Created Plugins and Notifiers
View the [Plugin Wiki](https://github.com/hunterlong/statping/wiki/Statping-Plugins) to see detailed information about Golang Plugins. Statping isn't just another Status Page for your applications, it's a framework that allows you to create your own plugins to interact with every element of your status page. [Notifier's](https://github.com/hunterlong/statping/wiki/Notifiers) can also be create with only 1 golang file.
Plugin are created in Golang using the [statping/plugin](https://github.com/hunterlong/statping/tree/master/plugin) golang package. The plugin package has a list of interfaces/events to accept into your own plugin application.

<p align="center">
<img width="100%" src="https://img.cjx.io/statupsc2.png">
</p>

## Easy to use Dashboard
Having a straight forward dashboard makes Statping that much better. Monitor your websites and applications with a basic HTTP GET request, or add a POST request with your own JSON to post to the endpoint.
<p align="center">
<img width="80%" src="https://img.cjx.io/statping_settings.gif">
</p>

## Exporting Static HTML
If you want to use Statping as a CLI application without running a server, you can export your status page to a static HTML.
This export tool is very useful for people who want to export their HTML and upload/commit it to Github Pages or an FTP server.
```dash
statping export
```
###### `index.html` will be created in the current directory with CDN URL's for assets.

## Run on Docker
Use the [Statping Docker Image](https://hub.docker.com/r/hunterlong/statping) to create a status page in seconds. Checkout the [Docker Wiki](https://github.com/hunterlong/statping/wiki/Docker) to view more details on how to get started using Docker.
```bash
docker run -it -p 8080:8080 hunterlong/statping
```
There are multiple ways to startup a Statping server. You want to make sure Statping is on it's own instance that is not on the same server as the applications you wish to monitor. It doesn't look good when your Status Page goes down, I recommend a small EC2 instance so you can set it, and forget it.

## Docker Compose
In this folder there is a standard docker-compose file that include nginx, postgres, and Statping.
```bash
docker-compose up -d
```

## Docker Compose with Automatic SSL
You can automatically start a Statping server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statping server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.
```bash
LETSENCRYPT_HOST=mydomain.com \
    LETSENCRYPT_EMAIL=info@mydomain.com \
    docker-compose -f docker-compose-ssl.yml up -d
```
Once your instance has started, it will take a moment to get your SSL certificate. Make sure you have a A or CNAME record on your domain that points to the IP/DNS of your server running Statping.

## Prometheus Exporter
Statping includes a [Prometheus Exporter](https://github.com/hunterlong/statping/wiki/Prometheus-Exporter) so you can have even more monitoring power with your services. The Prometheus exporter can be seen on `/metrics`, simply create another exporter in your prometheus config. Use your Statping API Secret for the Authorization Bearer header, the `/metrics` URL is dedicated for Prometheus and requires the correct API Secret has `Authorization` header.
```yaml
scrape_configs:
  - job_name: 'statping'
    bearer_token: 'MY API SECRET HERE'
    static_configs:
      - targets: ['statping:8080']
```

## Run on EC2 Server
Running Statping on the smallest EC2 server is very quick using the AWS AMI Image. Checkout the [AWS Wiki](https://github.com/hunterlong/statping/wiki/AWS-EC2) to see a step by step guide on how to get your EC2 Statping service online.

##### Create Security Groups
Create the AWS Security Groups with the commands below, Statping will expose port 80 and 443.
```bash
aws ec2 create-security-group --group-name StatpingPublicHTTP \
     --description "Statping HTTP Server on port 80 and 443"
# will response back a Group ID. Copy ID and use it for --group-id below.

aws ec2 authorize-security-group-ingress \
     --group-id sg-7e8b830f --protocol tcp \
     --port 80 --cidr 0.0.0.0/0

aws ec2 authorize-security-group-ingress \
     --group-id sg-7e8b830f --protocol tcp \
     --port 443 --cidr 0.0.0.0/0
```
##### Create EC2 without SSL
```bash
aws ec2 run-instances \
    --image-id ami-7be8a103 \
    --count 1 --instance-type t2.nano \
    --key-name MYKEYHERE \
    --security-group-ids sg-7e8b830f
```
##### Create EC2 with Automatic SSL Certification
The AWS-CLI command below will automatically create an EC2 server and create an SSL certificate from Lets Encrypt, for free.
```bash
wget https://raw.githubusercontent.com/hunterlong/statping/master/servers/ec2-ssl.sh
# Edit ec2-ssl.sh and insert your domain you want to use, then run command below.
# Use the Security Group ID that you used above for --security-group-ids

aws ec2 run-instances \
    --user-data file://ec2-ssl.sh \
    --image-id ami-7be8a103 \
    --count 1 --instance-type t2.nano \
    --key-name MYKEYHERE \
    --security-group-ids sg-7e8b830f
```

## Contributing
Statping accepts Push Requests! Feel free to add your own features and notifiers. You probably want to checkout the [Notifier Wiki](https://github.com/hunterlong/statping/wiki/Notifiers) to get a better understanding on how to create your own notification methods for failing/successful services. Testing on Statping will test each function on MySQL, Postgres, and SQLite. I recommend you run a MySQL and a Postgres Docker image for testing.

[![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/statping)](https://goreportcard.com/report/github.com/hunterlong/statping)
[![Build Status](https://travis-ci.com/hunterlong/statping.svg?branch=master)](https://travis-ci.com/hunterlong/statping) [![Cypress.io tests](https://img.shields.io/badge/cypress.io-tests-green.svg?style=flat-square)](https://dashboard.cypress.io/#/projects/bi8mhr/runs)
[![Docker Pulls](https://img.shields.io/docker/pulls/hunterlong/statping.svg)](https://hub.docker.com/r/hunterlong/statping/builds/) [![Godoc](https://godoc.org/github.com/hunterlong/statping?status.svg)](https://godoc.org/github.com/hunterlong/statping)[![Coverage Status](https://coveralls.io/repos/github/hunterlong/statping/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/statping?branch=master)


<p align="center">
<a href="https://www.buymeacoffee.com/hunterlong" target="_blank">
<img height="55" src="https://img.cjx.io/buy-me-redbull.png" >
</a>
</p>
