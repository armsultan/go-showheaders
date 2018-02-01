#!/bin/bash
# Requirement: The Go module for NGINX Unit needs to be installed.
# Usage: sudo ./install-app.sh

service unit stop
rm -rf /tmp/my-go-app
rm -rf /tmp/app-*
git clone https://github.com/armsultan/go-showheaders.git /tmp/my-go-app
GOPATH=/usr/share/gocode go build -o /tmp/app-u.go /tmp/my-go-app/Extras/nginx-unit/source/app/app-source.go
GOPATH=/usr/share/gocode go build -o /tmp/app-v1-green-u.go /tmp/my-go-app/Extras/nginx-unit/source/v1-green/app-v1-source.go
GOPATH=/usr/share/gocode go build -o /tmp/app-v2-blue-u.go /tmp/my-go-app/Extras/nginx-unit/source/v2-blue/app-v2-source.go

service unit restart
service unit loadconfig /tmp/my-go-app/Extras/nginx-unit/examples/unit.config
curl -s http://localhost:8500/ | grep title
curl -s http://localhost:8600/ | grep title
curl -s http://localhost:8700/ | grep title