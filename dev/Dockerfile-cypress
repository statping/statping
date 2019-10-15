FROM cypress/browsers:chrome67
LABEL maintainer="Hunter Long (https://github.com/hunterlong)"
# Statping 'test' image for running a full test using the production environment

WORKDIR $HOME/statping
ADD dev/test .

RUN npm install node-sass
ENV SASS=node-sass
RUN npm install

ADD ./statping-linux-amd64 /usr/local/bin/statping
RUN statping version

RUN npm run test-docker
