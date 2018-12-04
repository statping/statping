#!/usr/bin/env bash
#
#   Steps for Automatic SSL with Statping
#
# 1. Create EC2 Instance while including this file has "--data-file"
# 2. Once EC2 is created, view System Logs in EC2 instance.
# 3. In System logs, scroll all the way down and you'll see useful DNS records to use
# 4. Edit your domains DNS to point to the EC2 server.
#
## Domain for new SSL certificate and email for Lets Encrypt SSL recovery

LETSENCRYPT_HOST="status.MYDOMAIN.com"
LETSENCRYPT_EMAIL="noreply@MYEMAIL.com"

###################################################
############# Leave this part alone ###############
###################################################

exec > >(tee /var/log/user-data.log|logger -t user-data -s 2>/dev/console) 2>&1
printf "Statping will try to create an SSL for domain: $LETSENCRYPT_HOST\n"
printf "\nexport LETSENCRYPT_HOST=$LETSENCRYPT_HOST\n" >> /home/ubuntu/.profile
printf "export LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL\n" >> /home/ubuntu/.profile
sudo printf "\nexport LETSENCRYPT_HOST=$LETSENCRYPT_HOST\n" >> /root/.profile
sudo printf "export LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL\n" >> /root/.profile
source /home/ubuntu/.profile
sudo /bin/bash -c "source /root/.profile"
LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL /home/ubuntu/init.sh
