# firestarter

Bot generator

## Start with docker

Go http://localhost:8080 for configuration.

~~~
docker run \
 -v /path/to/config:/app/config \
 -p 3000:3000 -p 8080:8080 \
 -e SLACK_VERIFICATION_TOKEN=xxxxxxxxxxxxxx \
 -e SLACK_TOKEN=xoxb-xxxxxxxxxxxxxx \
 juntaki/firestarter
~~~

## Reference

* https://github.com/tcnksm/go-slack-interactive
