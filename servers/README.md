# Statup Server Setup
There are multiple way to startup a Statup server. You want to make sure Statup is on it's own instance that is not on the same server as the applications you wish to monitor. 
It doesn't look good when your Status Page goes down, I recommend a small EC2 instance so you can set it, and forget it.

> It's a good idea to have a Status Page not on the same server as your applications. 

# Install on Linux
Installing on Linux is a pretty easy task. Simply run the command below to have the `statup` command ready to rock. 
```bash
bash <(curl -s https://statup.io/install.sh)
statup version
```

# Install on Mac
Installing on Mac/OSX is also very easy, 
```bash
bash <(curl -s https://statup.io/install.sh)
statup version
```

# Install on Windows
Go to the [Latest Releases](https://github.com/hunterlong/statup/releases/latest) page for Statup and simply download `statup-windows-x64`! 
Statup only supports Windows 64-bit currently, sorry 32-bit users. Rename the file to `statup` for ease of use!

# Install on Docker
This Docker image uses Alpine Linux to keep it ultra small. 
```bash
docker run -it -p 8080:8080 hunterlong/statup
```
#### Or use Docker Compose
This Docker Compose file inlcudes NGINX, Postgres, and Statup.
```bash
wget https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose.yml
docker-compose up -d
```

#### Docker Compose with Automatic SSL
You can automatically start a Statup server with automatic SSL encryption using this docker-compose file. First point your domain's DNS to the Statup server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.
```
wget https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose-ssl.yml

LETSENCRYPT_HOST=mydomain.com \ 
    LETSENCRYPT_EMAIL=info@mydomain.com \
    docker-compose -f docker-compose-ssl.yml up -d
```

# Install on EC2
Running Statup on the smallest EC2 server is very quick using the AWS AMI Image: `ami-7be8a103`.

### Create Security Groups
```bash
aws ec2 create-security-group --group-name StatupPublicHTTP --description "Statup HTTP Server on port 80 and 443"
# will response back a Group ID. Copy ID and use it for --group-id below.
aws ec2 authorize-security-group-ingress --group-id sg-7e8b830f --protocol tcp --port 80 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id sg-7e8b830f --protocol tcp --port 443 --cidr 0.0.0.0/0
```
### Create EC2 without SSL
```bash
aws ec2 run-instances \ 
    --image-id ami-7be8a103 \ 
    --count 1 --instance-type t2.nano \ 
    --key-name MYKEYHERE \ 
    --security-group-ids sg-7e8b830f
```
### Create EC2 with Automatic SSL Certification
```bash
wget https://raw.githubusercontent.com/hunterlong/statup/master/servers/ec2-ssl.sh
```
Edit ec2-ssl.sh and insert your domain you want to use, then run command below. Use the Security Group ID that you used above for --security-group-ids
```
aws ec2 run-instances \ 
    --user-data file://ec2-ssl.sh \ 
    --image-id ami-7be8a103 \ 
    --count 1 --instance-type t2.nano \ 
    --key-name MYKEYHERE \ 
    --security-group-ids sg-7e8b830f
```

#### EC2 Server Specs
- t2.nano ($4.60 monthly)
- 8gb SSD Memory
- 0.5gb RAM
- Docker with Docker Compose installed
- Running Statup, NGINX, and Postgres
- boot scripts to automatically clean unused containers.


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

## Systemd
```$xslt
/etc/systemd/system/statup.service


```