# Statup - Status Page [![Build Status](https://travis-ci.org/hunterlong/statup.svg?branch=master)](https://travis-ci.org/hunterlong/statup)
An easy to use Status Page for your websites and applications. Statup will automatically fetch the application and render a beautiful status page with tons of features 
for you to build an even better status page. 

# A Future-Proof Status Page
Statup strives to remain future-proof and remain intact if a failure is created. Your Statup service should not be running on the same instance you're trying to monitor. 
If your server crashes your Status Page should still remaining online to notify your users of downtime. 

## Run on Docker
Use the [Statup Docker Image](https://hub.docker.com/r/hunterlong/statup) to create a status page in seconds.
```
docker run -it -p 8080:8080 hunterlong/statup
```
There are multiple way to startup a Statup server. You want to make sure Statup is on it's own instance that is not on the same server as the applications you wish to monitor. 
It doesn't look good when your Status Page goes down, I recommend a small EC2 instance so you can set it, and forget it.

## Docker Compose
In this folder there is a standard docker-compose file that include nginx, postgres, and Statup. 
```$xslt
docker-compose up -d
```

## Docker Compose with Automatic SSL
You can automatically start a Statup server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statup server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.
```
LETSENCRYPT_HOST=mydomain.com \ 
    LETSENCRYPT_EMAIL=info@mydomain.com \ 
    docker-compose -f docker-compose-ssl.yml up -d
```
Once your instance has started, it will take a moment to get your SSL certificate. Make sure you have a A or CNAME record on your domain that points to the IP/DNS of your server running Statup.

## Run on AWS EC2
Running Statup on the smallest EC2 server is very quick using the AWS AMI Image: `ami-1f7c3567`.
```
aws ec2 run-instances \ 
    --image-id ami-1f7c3567 \
    --count 1 \ 
    --instance type t2.micro \ 
    --region us-west-2
    --key-name <key name> \ 
    --security-group-ids <your security group id here> \ 
    --subnet-id <your subnet id here> \ 
    --region <your region here>
```

## Email Nofitications
Statup includes email notification via SMTP if your services go offline. 

## User Created Plugins
Statup isn't just another Status Page for your applications, it's a framework that allows you to create your own plugins to interact with every element of your status page.
Plugin are created in Golang using the [statup/plugin](https://github.com/hunterlong/statup/tree/master/plugin) golang package. The plugin package has a list of 
interfaces/events to accept into your own plugin application. 