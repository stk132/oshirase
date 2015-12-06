# oshirase

oshirase is very simple slack cli command

## configuration

please set below environment variable

``` shell
export SLACK_API_TOKEN=<your slack api token>
```

## usage

### post message

``` shell
oshirase message -c <channelID or groupID> -u <username> <message>
or
<some command> | oshirase message -c <channelID or groupID> -u <username>
```

### show channel information(channelName, ID)

``` shell
oshirase channel
```

### show group (groupName, ID)

``` shell
oshirase group
```


## How to build

### 1. install gom command

``` shell
go get github.com/mattn/gom
```

### 2. install dependency package

``` shell
gom install
```

### 3. build

``` shell
gom build
```
