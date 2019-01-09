#!/usr/bin/env sh

COMMAND="rm -rf /app/statping.db && reboot"

echo "* * * * * echo $COMMAND >> /test_file 2>&1" > /etc/crontabs/root

statping