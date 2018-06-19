# Statup - Status Page [![Build Status](https://travis-ci.org/hunterlong/statup.svg?branch=master)](https://travis-ci.org/hunterlong/statup)
An easy to use Status Page for your websites and applications. Statup will automatically fetch the application and render a beautiful status page with tons of features 
for you to build an even better status page. 

## Run on Docker
Use the [Statup Docker Image](https://hub.docker.com/r/hunterlong/statup) to create a status page in seconds.
```
docker run -it -p 8080:8080 hunterlong/statup
```

### Install on Linux
```
bash <(curl -s https://statup.io/install.sh)
```

## User Created Plugins
Statup isn't just another Status Page for your applications, it's a framework that allows you to create your own plugins to interact with every element of your status page.
Plugin are created in Golang using the [statup/plugin](https://github.com/hunterlong/statup/tree/master/plugin) golang package. The plugin package has a list of 
interfaces/events to accept into your own plugin application. 