# firestarter

Configurable slack bot with simple Web UI

## How to run

### Start with docker

Go http://localhost:8080 for configuration.

~~~
docker run \
 -v /path/to/config:/app/config \
 -p 3000:3000 -p 8080:8080 \
 -e SLACK_TOKEN=xoxb-xxxxxxxxxxxxxx \
 -e SLACK_VERIFICATION_TOKEN=xxxxxxxxxxxxxx \
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
