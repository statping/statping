#!/usr/bin/env bash
LETSENCRYPT_HOST="status.balancebadge.io"
LETSENCRYPT_EMAIL="info@socialeck.com"

rm -f docker-compose.yml
curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statu$
sudo service docker start
sudo docker system prune -af
sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose pull
sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose up -d