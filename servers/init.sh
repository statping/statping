#!/usr/bin/env bash
cd /home/ubuntu
curl -o /home/ubuntu/startup.sh -H 'Cache-Control: no-cache' https://raw.githubusercontent.com/hunterlong/statup/master/servers/startup.sh
chmod +x /home/ubuntu/startup.sh
./startup.sh