#!/usr/bin/env bash
cd /home/ubuntu
source /home/ubuntu/.profile
sudo rm -rf startup.sh > /dev/null
sudo curl -o startup.sh -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/startup.sh > /dev/null
sudo chmod +x startup.sh > /dev/null
sudo rm -f docker-compose.yml > /dev/null

EC2_ENDPOINT=$(curl -s http://169.254.169.254/latest/meta-data/public-hostname)
EC_IP=$(curl -s http://169.254.169.254/latest/meta-data/public-ipv4)

if [ "$LETSENCRYPT_HOST" = "" ]
then
   printf "\n              Status Status Page on EC2\n"
   printf "================================================================================================================\n"
   printf "Point your domain's DNS records to one of these endpoints.\n"
   printf "  A   =>   EC2 Server IP:     $EC_IP   \n"
   printf "CNAME =>   EC2 Public DNS:    $EC2_ENDPOINT\n"
   printf "================================================================================================================\n"
   sudo curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose.yml > /dev/null
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

   sudo curl -o docker-compose.yml -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/docker-compose-ssl.yml > /dev/null
fi

sudo service docker start

if [ "$LETSENCRYPT_HOST" = "" ]
then
   sudo docker-compose pull
   sudo docker-compose up -d
else
    sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose pull
    sudo LETSENCRYPT_HOST=$LETSENCRYPT_HOST LETSENCRYPT_EMAIL=$LETSENCRYPT_EMAIL docker-compose up -d
fi

sudo docker system prune -af

logger "Status Process complete. IP: $EC_IP | DNS: $EC2_ENDPOINT\n"
printf "Status Process complete. IP: $EC_IP | DNS: $EC2_ENDPOINT\n"

# commands to reverse installation and only leave init.sh
#sudo docker rm -f $(sudo docker ps -a -q)
#sudo docker rmi $(sudo docker images -q)
#rm -rf docker-compose.yml
#sudo rm -rf statup
#sudo rm startup.sh
#sudo docker system prune -af
#sudo service docker stop
