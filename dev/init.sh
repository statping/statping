#!/usr/bin/env bash
#
#
#     Statping Status Page Server
#            for EC2
#
# This script run everytime the EC2 is booted. It will maintain the most latest version of Statping
# while removing old docker containers and images to reduce hard drive usage for long term use.
#
#

cd /home/ubuntu
source /home/ubuntu/.profile
sudo rm -rf startup.sh > /dev/null
sudo curl -o startup.sh -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statping/master/dev/startup.sh > /dev/null
sudo chmod +x startup.sh > /dev/null
sudo rm -f docker-compose.yml > /dev/null

EC2_ENDPOINT=$(curl -s http://169.254.169.254/latest/meta-data/public-hostname)
EC_IP=$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4)

if [ "$LETSENCRYPT_HOST" = "" ]
then
   sudo curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statping/master/dev/docker-compose-single.yml > /dev/null
else
   printf "                    \n\n\n\nDomain found for SSL certificate - $LETSENCRYPT_HOST\n"
   printf "================================================================================================================\n"
   printf "You must set the domains DNS records to point to this server!\n"
   printf "   EC2 Server IP:     $EC_IP\n"
   printf "   EC2 Public DNS:    $EC2_ENDPOINT\n"
   printf "================================================================================================================\n"
   printf " CNAME    $LETSENCRYPT_HOST   =>   $EC2_ENDPOINT       (if not using Elastic IP)\n"
   printf "   A      $LETSENCRYPT_HOST   =>   $EC_IP                                        (or use A record if you are using an Elastic IP)\n"
   printf "================================================================================================================\n\n\n"

   sudo curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statping/dev/docker-compose-ssl.yml > /dev/null
fi

sudo service docker start > /dev/null

if [ "$LETSENCRYPT_HOST" = "" ]
then
   sudo docker-compose pull > /dev/null
   sudo docker-compose up -d > /dev/null
else
    sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose pull > /dev/null
    sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose up -d > /dev/null
fi

sudo docker system prune -af > /dev/null

sudo curl https://raw.githubusercontent.com/hunterlong/statping/dev/init.sh > /home/ubuntu/init.sh
sudo chmod +x /home/ubuntu/init.sh > /dev/null

printf "\n\n\n\n\n              Statping Status Page on EC2\n"
printf "================================================================================================================\n"
if [ "$LETSENCRYPT_HOST" = "" ]
then
    printf "Point your domain's DNS records to one of these endpoints.\n"
    printf "A RECORD     =>   $EC_IP   \n"
    printf "CNAME RECORD =>   $EC2_ENDPOINT\n"
    printf "================================================================================================================\n"
    printf "Your Statping Server is ready! Go to the URL below to begin.\n"
    printf "Statping URL: $EC2_ENDPOINT\n"
    printf "================================================================================================================\n"
else
   printf "Domain found for SSL certificate - $LETSENCRYPT_HOST\n"
   printf "================================================================================================================\n"
   printf "You must set the domains DNS records to point to this server!\n"
   printf "A RECORD     =>   $EC_IP   \n"
   printf "CNAME RECORD =>   $EC2_ENDPOINT\n"
   printf "================================================================================================================\n"
   printf "Once you set your DNS records, Lets Encrypt will automatically\n"
   printf "create a SSL certificate for you and redirect you to HTTPS\n\n"
   printf "================================================================================================================\n"
   printf "Your Statping Server is ready! Go to the URL below to begin.\n"
   printf "Statping URL: $EC2_ENDPOINT\n"
   printf "SSL Domain: $LETSENCRYPT_HOST\n"
   printf "================================================================================================================\n"
fi


# commands to reverse installation and only leave init.sh
#sudo docker rm -f $(sudo docker ps -a -q)
#sudo docker rmi $(sudo docker images -q)
#rm -rf docker-compose.yml
#sudo rm -rf statping
#sudo rm startup.sh
#sudo docker system prune -af
#sudo service docker stop
