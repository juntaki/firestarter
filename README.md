# firestarter


[![Docker Automated build](https://img.shields.io/docker/automated/juntaki/firestarter.svg)](https://hub.docker.com/r/juntaki/firestarter/)

Configurable slack bot with simple Web UI

## How to run

### Start with docker

Go http://localhost:8080 for configuration.
:3000 is for slack interactive message. Go Your Slack App setting page, set http://yourhostname:3000 to Interactive Components -> Request URL.

~~~
docker run \
 -v /path/to/config:/app/config \
 -p 3000:3000 -p 8080:8080 \
 -e SLACK_TOKEN=xoxb-xxxxxxxxxxxxxx \
 -e SLACK_VERIFICATION_TOKEN=xxxxxxxxxxxxxx \
 juntaki/firestarter
~~~

### Start with docker (SQS proxy mode)

Even if you are in a firewall, you can use interactive message w/o opening :3000.
See details for [juntaki/firestarter-sqs-proxy](https://github.com/juntaki/firestarter-sqs-proxy)

~~~
docker run \
 -v /path/to/config:/app/config \
 -p 8080:8080 \
 -e SLACK_TOKEN=xoxb-xxxxxxxxxxxxxx \
 -e SLACK_VERIFICATION_TOKEN=xxxxxxxxxxxxxx \
 -e AWS_ACCESS_KEY_ID=AKIAxxxx \
 -e AWS_SECRET_ACCESS_KEY=xxxxxxx \
 -e SQS_URL=https://sqs.xxxxxx.amazonaws.com/xxxxxxxxxx/xxxxxxxx \
 -e AWS_REGION=xxxxxxxx \
 juntaki/firestarter
~~~



### Start from local (for development)

Install dependency package and build.

~~~
go get -u -v
go build -o firestarter
~~~

~~~
cd admin
yarn install
yarn build
~~~

Set environment variables, and run.

~~~
export SLACK_TOKEN=xoxb-xxxxxxxxxxxxxx
export SLACK_VERIFICATION_TOKEN=xxxxxxxxxxxxxx
./firestarter
~~~

## Slack configuration

1. [Create Slack app](https://api.slack.com/apps)
2. Set Interactive Components -> Request URL to http://your-hostname:3000 (you should use https)

## Reference

* https://github.com/tcnksm/go-slack-interactive
