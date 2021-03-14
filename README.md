# which-ip-fastest

test which ip of url will response fastest

## Installation

Install:

```shell
go get -u github.com/go-redis/redis
```

## Usage

```
Usage of which-ip-fastest:
  -ip string
    	the ip of the url's host, multiple will split by char , (default "13.114.40.48,52.192.72.89,52.69.186.44,15.164.81.167,52.78.231.108,13.234.176.102,13.234.210.38,13.229.188.59,13.250.177.223,52.74.223.119,13.236.229.21,13.237.44.5,52.64.108.95,18.228.52.138,18.228.67.229,18.231.5.6")
  -timeout duration
    	max request time (default 5s)
  -url string
    	the request url (default "https://github.com")
```

## Cross compiling

### Build for windows

```shell
GOARCH=amd64 GOOS=windows go build -o which-ip-fastest.exe
```

### Build for linux

```shell
GOARCH=amd64 GOOS=linux go build -o which-ip-fastest
```

### Build for mac

```shell
GOARCH=amd64 GOOS=darwin go build -o which-ip-fastest
```