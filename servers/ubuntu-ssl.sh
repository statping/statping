#!/usr/bin/env bash

LETSENCRYPT_HOST="status.balancebadge.io"
LETSENCRYPT_EMAIL="info@socialeck.com"

sudo apt-get update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get install docker-ce -y
sudo curl -L https://github.com/docker/compose/releases/download/1.21.2/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo docker-compose --version
sudo systemctl enable docker
mkdir statup
cd statup
rm -f docker-compose.yml
curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose-ssl.yml
sudo service docker start
sudo docker system prune -af
sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose pull
sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose up -d