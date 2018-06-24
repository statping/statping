#!/usr/bin/env bash
cd /home/ubuntu
rm -f docker-compose.yml
curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose.yml
sudo service docker start
sudo docker system prune -af
sudo docker-compose pull
sudo docker-compose up -d