<p align="center">
<img width="60%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statping.png">
</p>
<p align="center">
    <b>Statping - Web and App Status Monitoring for Any Type of Project</b><br>
<a href="https://github.com/statping/statping/wiki">View Wiki</a> | <a href="https://demo.statping.com">Demo</a> | <a href="https://itunes.apple.com/us/app/apple-store/id1445513219">iPhone</a> | <a href="https://play.google.com/store/apps/details?id=com.statping">Android</a> <br> <a href="http://docs.statping.com">API</a> | <a href="https://github.com/statping/statping/wiki/Docker">Docker</a> | <a href="https://github.com/statping/statping/wiki/AWS-EC2">EC2</a> | <a href="https://github.com/statping/statping/wiki/Mac">Mac</a> | <a href="https://github.com/statping/statping/wiki/Linux">Linux</a> | <a href="https://github.com/statping/statping/wiki/Windows">Windows</a>
</p>

# Statping - Status Page & Monitoring Server
An easy to use Status Page for your websites and applications. Statping will automatically fetch the application and render a beautiful status page with tons of features for you to build an even better status page. This Status Page generator allows you to use MySQL, Postgres, or SQLite on multiple operating systems.

![Master Release](https://github.com/statping/statping/workflows/Master%20Release/badge.svg?branch=master) [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/statping/statping) [![Slack](https://slack.statping.com/badge.svg)](https://slack.statping.com) [![](https://images.microbadger.com/badges/image/statping/statping.svg)](https://microbadger.com/images/statping/statping) [![Docker Pulls](https://img.shields.io/docker/pulls/statping/statping.svg)](https://hub.docker.com/r/statping/statping/builds/)

<br><br>
<img align="left" width="320" height="235" src="https://img.cjx.io/statupsiterun.gif">
<h2>A Future-Proof Status Page</h2>
Statping strives to remain future-proof and remain intact if a failure is created. Your Statping service should not be running on the same instance you're trying to monitor. If your server crashes your Status Page should still remaining online to notify your users of downtime.

<br><a href="https://labs.play-with-docker.com/?stack=https://raw.githubusercontent.com/statping/statping/master/dev/pwd-stack.yml"><img height=25 src="https://assets.statping.com/docker-pwd.png"></a> (dashboard login is `admin`, password `admin`)
<br><br><br>

<h2>No Requirements</h2>
Statping is built in Go Language so all you need is the precompile binary based on your operating system. You won't need to install anything extra once you have the Statping binary installed. You can even run Statping on a Raspberry Pi.
<br><br>
<p align="center">
    <a href="https://github.com/statping/statping/wiki/Linux"><img width="5%" src="https://img.cjx.io/linux.png"></a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://github.com/statping/statping/wiki/Mac"><img width="5%" src="https://img.cjx.io/apple.png"></a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://github.com/statping/statping/wiki/Windows"><img width="5%" src="https://img.cjx.io/windows.png"></a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://play.google.com/store/apps/details?id=com.statping"><img width="5%" src="https://img.cjx.io/android.png"></a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img width="5%" src="https://img.cjx.io/appstore.png"></a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://hub.docker.com/r/statping/statping"><img width="5%" src="https://img.cjx.io/dockericon.png"></a>
<br><br></p>


<img align="right" width="320" height="235" src="https://gitimgs.s3-us-west-2.amazonaws.com/slack-notifer.png">
<h2>Lightweight and Fast</h2>
Statping is a very lightweight application and is available for Linux, Mac, and Windows. The Docker image is only ~16Mb so you know that this application won't be filling up your hard drive space.
The Status binary for all other OS's is ~17Mb at most.
<br><br><br><br><br><br>

<img align="left" width="320" height="235" src="https://img.cjx.io/statping_iphone_bk.png">
<h2>Mobile App is Gorgeous</h2>
The Statping app is available on the App Store and Google Play for free. The app will allow you to view services, receive notifications when a service is offline, update groups, users, services, messages, and more! Start your own Statping server and then connect it to the app by scanning the QR code in settings.

<p align="center">
<a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a>
<a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a>
</p>

<br><br>

## Run on Any Server
Whether you're a Docker fan-boy or a [AWS EC2](https://github.com/statping/statping/wiki/AWS-EC2) master, Statping gives you multiple options to simply get running. Our Amazon AMI image is only 8Gb and will automatically update to the most stable version of Statping.
Running on an EC2 server might be the most cost effective way to host your own Statping Status Page. The server runs on the smallest EC2 instance (t2.nano) AWS has to offer, which only costs around $4.60 USD a month for your dedicated Status Page.
Want to run it on your own Docker server? Awesome! Statping has multiple docker-compose.yml files to work with. Statping can automatically create a SSL Certification for your status page.
<br><br><br><br>

<img align="left" width="320" height="205" src="https://img.cjx.io/statping_theme.gif">
<h2>Custom SASS Styling</h2>
Statping will allow you to completely customize your Status Page using SASS styling with easy to use variables. The Docker image actually contains a prebuilt SASS binary so you won't even need to setup anything!
<br><br><br><br>

## Slack, Email, Twilio and more
Statping includes email notification via SMTP and Slack integration using [Incoming Webhook](https://api.slack.com/incoming-webhooks). Insert the webhook URL into the Settings page in Statping and enable the Slack integration. Anytime a service fails, you're channel that you specified on Slack will receive a message.
<br><br><br><br>

<h2>User Created Notifiers</h2>
View the [Plugin Wiki](https://github.com/statping/statping/wiki/Statping-Plugins) to see detailed information about Golang Plugins. Statping isn't just another Status Page for your applications, it's a framework that allows you to create your own plugins to interact with every element of your status page. [Notifier's](https://github.com/statping/statping/wiki/Notifiers) can also be create with only 1 golang file.
<br><br><br><br>

<img align="center" width="100%" height="250" src="https://img.cjx.io/statupsc2.png">

<br><br><br><br>

<img align="right" width="320" height="235" src="https://img.cjx.io/statping_settings.gif">
<h2>Easy to use Dashboard</h2>
Having a straight forward dashboard makes Statping that much better. Monitor your websites and applications with a basic HTTP GET request, or add a POST request with your own JSON to post to the endpoint.
<br><br><br><br>

## Run on Docker
Use the [Statping Docker Image](https://hub.docker.com/r/statping/statping) to create a status page in seconds. Checkout the [Docker Wiki](https://github.com/statping/statping/wiki/Docker) to view more details on how to get started using Docker.
```bash
docker run -it -p 8080:8080 statping/statping
```
There are multiple ways to startup a Statping server. You want to make sure Statping is on it's own instance that is not on the same server as the applications you wish to monitor. It doesn't look good when your Status Page goes down, I recommend a small EC2 instance so you can set it, and forget it.
<br><br><br><br>

## Docker Compose
In this folder there is a standard docker-compose file that include nginx, postgres, and Statping.
```bash
docker-compose up -d
```
<br><br><br><br>

## Docker Compose with Automatic SSL
You can automatically start a Statping server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statping server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.
```bash
LETSENCRYPT_HOST=mydomain.com \
    LETSENCRYPT_EMAIL=info@mydomain.com \
    docker-compose -f docker-compose-ssl.yml up -d
```
Once your instance has started, it will take a moment to get your SSL certificate. Make sure you have a A or CNAME record on your domain that points to the IP/DNS of your server running Statping.
<br><br><br><br>

## Prometheus Exporter
Statping includes a [Prometheus Exporter](https://github.com/statping/statping/wiki/Prometheus-Exporter) so you can have even more monitoring power with your services. The Prometheus exporter can be seen on `/metrics`, simply create another exporter in your prometheus config. Use your Statping API Secret for the Authorization Bearer header, the `/metrics` URL is dedicated for Prometheus and requires the correct API Secret has `Authorization` header.
```yaml
scrape_configs:
  - job_name: 'statping'
    bearer_token: 'MY API SECRET HERE'
    static_configs:
      - targets: ['statping:8080']
```
<br><br><br><br>

## Run on EC2 Server
Running Statping on the smallest EC2 server is very quick using the AWS AMI Image. Checkout the [AWS Wiki](https://github.com/statping/statping/wiki/AWS-EC2) to see a step by step guide on how to get your EC2 Statping service online.
<br><br><br><br>

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
wget https://raw.githubusercontent.com/statping/statping/master/servers/ec2-ssl.sh
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
Statping accepts Push Requests to the `dev` branch! Feel free to add your own features and notifiers. You probably want to checkout the [Notifier Wiki](https://github.com/statping/statping/wiki/Notifiers) to get a better understanding on how to create your own notification methods for failing/successful services. Testing on Statping will test each function on MySQL, Postgres, and SQLite. I recommend running MySQL and Postgres Docker containers for testing. You can find multiple docker-compose files in the dev directory. 

![Dev Release](https://github.com/statping/statping/workflows/Dev%20Release/badge.svg?branch=dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/statping/statping)](https://goreportcard.com/report/github.com/statping/statping)
[![Build Status](https://travis-ci.com/statping/statping.svg?branch=master)](https://travis-ci.com/statping/statping) [![Cypress.io tests](https://img.shields.io/badge/cypress.io-tests-green.svg?style=flat-square)](https://dashboard.cypress.io/#/projects/bi8mhr/runs)
[![Docker Pulls](https://img.shields.io/docker/pulls/statping/statping.svg)](https://hub.docker.com/r/statping/statping/builds/) [![Godoc](https://godoc.org/github.com/statping/statping?status.svg)](https://godoc.org/github.com/statping/statping)[![Coverage Status](https://coveralls.io/repos/github/statping/statping/badge.svg?branch=master)](https://coveralls.io/github/statping/statping?branch=master)


<p align="center">
<a href="https://www.buymeacoffee.com/hunterlong" target="_blank">
<img height="55" src="https://img.cjx.io/buy-me-redbull.png" >
</a>
</p>
