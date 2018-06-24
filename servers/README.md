# Statup Servers
There are multiple way to startup a Statup server. You want to make sure Statup is on it's own instance that is not on the same server as the applications you wish to monitor. 
It doesn't look good when your Status Page goes down, I recommend a small EC2 instance so you can set it, and forget it.

## AWS EC2 (`ami-1f7c3567`)
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

## Docker Compose
In this folder there is a standard docker-compose file that include nginx, postgres, and Statup. 
```$xslt
docker-compose up -d
```

## Docker Compose with Automatic SSL
You can automatically start a Statup server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statup server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.
```
DOMAIN=mydomain.com EMAIL=info@mydomain.com \
    docker-compose -f docker-compose-ssl.yml up -d
```
