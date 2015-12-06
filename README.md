# oshirase

oshirase is very simple slack cli command

## configuration

please set below environment variable

``` shell
export SLACK_API_TOKEN=<your slack api token>
```

## usage

post message

``` shell
oshirase message -c <channelID or groupID> -u <username> <message>
```

show channel information(channelName, ID)

``` shell
oshirase channel
```

show group (groupName, ID)

``` shell
oshirase group
```
