# URL shortener (written in Go)

This repository contains the source code of a simple URL shortener. 

## Installation
```
go build -i -o url_shortener
```

## Usage
Execute `./url_shortener show` to display all shortened links, `./url_shortener add <short URL> <target URL>` and `./url_shortener delete <short URL>` to add or delete a link. <br>

`./url_shortener` starts the server, which is available at `http://localhost:8484` (nginx could be used as a reverse proxy, otherwise go to `server.go` and change `SERVER_PORT` to 80 and `SERVER_HOST` to your custom domain)