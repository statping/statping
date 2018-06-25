# Statup - Status Page [![Build Status](https://travis-ci.org/hunterlong/statup.svg?branch=master)](https://travis-ci.org/hunterlong/statup)
An easy to use Status Page for your websites and applications. Statup will automatically fetch the application and render a beautiful status page with tons of features 
for you to build an even better status page. 

## A Future-Proof Status Page
Statup strives to remain future-proof and remain intact if a failure is created. Your Statup service should not be running on the same instance you're trying to monitor. 
If your server crashes your Status Page should still remaining online to notify your users of downtime. 

## Lightweight and Fast
Statup is a very lightweight application and is available for Linux, Mac, and Windows. The Docker image is only ~16Mb so you know that this application won't be filling up your hard drive space. 
The Status binary for all other OS's is ~17Mb at most. 

## 3 Different Databases
This Status Page generator allows you to use MySQL, Postgres, or SQLite. 

## No Requirements
Statup is built in Go Language so all you need is the precompile binary based on your operating system. You won't need to install anything extra once you have the Statup binary installed. 

## Run on Any Server
Whether you're a Docker fan-boy or a AWS EC2 master, Statup gives you multiple options to simply get running. Our Amazon AMI image (`ami-7be8a103`) is only 8Gb and will automatically update to the most stable version of Statup. 
Running on an EC2 server might be the most cost effective way to host your own Statup Status Page. The server runs on the smallest EC2 instance (t2.nano) AWS has to offer, which only costs around $4.60 USD a month for your dedicated Status Page.
Want to run it on your own Docker server? Awesome! Statup has multiple docker-compose.yml files to work with. Statup can automatically create a SSL Certification for your status page.

## Email Nofitications
Statup includes email notification via SMTP if your services go offline. 

## User Created Plugins
Statup isn't just another Status Page for your applications, it's a framework that allows you to create your own plugins to interact with every element of your status page.
Plugin are created in Golang using the [statup/plugin](https://github.com/hunterlong/statup/tree/master/plugin) golang package. The plugin package has a list of 
interfaces/events to accept into your own plugin application. 

## Exporting Static HTML
If you want to use Statup as a CLI application without running a server, you can export your status page to a static HTML. 
This export tool is very useful for people who want to export their HTML and upload/commit it to Github Pages or an FTP server.
```dash
statup export
```
###### `index.html` will be created in the current directory with CDN URL's for assets.

## Run on Docker
Use the [Statup Docker Image](https://hub.docker.com/r/hunterlong/statup) to create a status page in seconds.
```bash
docker run -it -p 8080:8080 hunterlong/statup
```
There are multiple way to startup a Statup server. You want to make sure Statup is on it's own instance that is not on the same server as the applications you wish to monitor. 
It doesn't look good when your Status Page goes down, I recommend a small EC2 instance so you can set it, and forget it.

## Docker Compose
In this folder there is a standard docker-compose file that include nginx, postgres, and Statup. 
```bash
docker-compose up -d
```

## Docker Compose with Automatic SSL
You can automatically start a Statup server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statup server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.
```bash
LETSENCRYPT_HOST=mydomain.com \ 
    LETSENCRYPT_EMAIL=info@mydomain.com \ 
    docker-compose -f docker-compose-ssl.yml up -d
```
Once your instance has started, it will take a moment to get your SSL certificate. Make sure you have a A or CNAME record on your domain that points to the IP/DNS of your server running Statup.

## Run on EC2 Server
Running Statup on the smallest EC2 server is very quick using the AWS AMI Image: `ami-7be8a103`.

##### Create Security Groups
```bash
aws ec2 create-security-group --group-name StatupPublicHTTP --description "Statup HTTP Server on port 80 and 443"
# will response back a Group ID. Copy ID and use it for --group-id below.
aws ec2 authorize-security-group-ingress --group-id sg-7e8b830f --protocol tcp --port 80 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id sg-7e8b830f --protocol tcp --port 443 --cidr 0.0.0.0/0
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
```bash
wget https://raw.githubusercontent.com/hunterlong/statup/master/servers/ec2-ssl.sh
# Edit ec2-ssl.sh and insert your domain you want to use, then run command below.
# Use the Security Group ID that you used above for --security-group-ids
aws ec2 run-instances \ 
    --user-data file://ec2-ssl.sh \ 
    --image-id ami-7be8a103 \ 
    --count 1 --instance-type t2.nano \ 
    --key-name MYKEYHERE \ 
    --security-group-ids sg-7e8b830f
```

## Prometheus Exporter
Statup includes a prometheus exporter so you can have even more monitoring power with your services. The prometheus exporter can be seen on `/metrics`, simply create another exporter in your prometheus config.
```yaml
scrape_configs:
  - job_name: 'statup'
    static_configs:
      - targets: ['statup:8080']
```



