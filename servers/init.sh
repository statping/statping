#!/usr/bin/env bash
cd /home/ubuntu
sudo rm -rf startup.sh
sudo curl -o startup.sh -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/startup.sh
sudo chmod +x startup.sh

sudo rm -f docker-compose.yml
if [ "$LETSENCRYPT_HOST" = "" ]
then
   sudo curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose.yml
else
   sudo curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose-ssl.yml
fi

sudo service docker start
sudo docker system prune -af

if [ "$LETSENCRYPT_HOST" = "" ]
then
   sudo docker-compose pull
   sudo docker-compose up -d
else
    sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose pull
    sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose up -d
fi


