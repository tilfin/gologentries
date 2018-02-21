# gologentries
Redirecting stdin to Logentries for UNIX pipe

## Build

```
$ go build -o gologentries main.go
```

### for Raspberry pi

```
$ GOOS=linux GOARCH=arm go build -o gologentries main.go
```

## Usage

```
$ tail -F /var/log/application.log | LE_TOKEN=xxxx-xxxx-xxx-xxxx ./gologentries
```
