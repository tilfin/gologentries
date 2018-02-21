# gologentries
Redirecting stdin to Logentries for UNIX/Linux pipe

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

## Setup a log with token

Make a Token TCP referencing [Adding a Log | Logentries](https://docs.logentries.com/docs/setup-first-log)

## Logentries screenshot

Logentries receiving raw logs from stdin

```
$ echo "test log" | LE_TOKEN=xxxx-xxxx-xxx-xxxx ./gologentries
$ tail -F /var/log/application.log | LE_TOKEN=xxxx-xxxx-xxx-xxxx ./gologentries
```

![Screenshot](https://user-images.githubusercontent.com/519017/36490864-a866329c-176c-11e8-9f7c-7d2a4f52dd2b.png)
