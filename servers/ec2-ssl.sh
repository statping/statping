#!/usr/bin/env bash
exec > >(tee /var/log/user-data.log|logger -t user-data -s 2>/dev/console) 2>&1

## Domain for new SSL certificate
LETSENCRYPT_HOST="status.balancebadge.io"

## An email address that can recover SSL certificate from Lets Encrypt
LETSENCRYPT_EMAIL="info@socialeck.com"

###################################################
############# Leave this part alone ###############
###################################################

printf "Statup will try to create an SSL for domain: $LETSENCRYPT_HOST\n"
printf "\nexport LETSENCRYPT_HOST=$LETSENCRYPT_HOST\n" >> /home/ubuntu/.profile
printf "export LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL\n" >> /home/ubuntu/.profile
sudo printf "\nexport LETSENCRYPT_HOST=$LETSENCRYPT_HOST\n" >> /root/.profile
sudo printf "export LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL\n" >> /root/.profile
source /home/ubuntu/.profile
sudo /bin/bash -c "source /root/.profile"
LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL /home/ubuntu/init.sh
