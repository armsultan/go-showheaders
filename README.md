# go-showheaders
A Simple Go app to Display hostname, IP address and Request Headers


 ![Screenshot](https://raw.githubusercontent.com/armsultan/go-showheaders/master/Extras/screenshot.png)

For maximum lightweightness (< 6MB):

1. Two builds are available, under `/go` and `/nginx-unit`. [Nginx Unit](http://unit.nginx.org/installation/) builds 
import `"nginx/unit"` and use `unit.ListenAndServe` function instead of `"net/http"`. [See here](http://unit.nginx.org/installation/) 
for more details on building Go applications for Nginx Unit

1. go apps can be complied with:

	`CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app.go .`

1. go apps for **Nginx Unit** can be complied from the unit server using:

	`GOPATH=/usr/share/gocode go build -o /tmp/app-v1-green-u.go /go-showheaders/Extras/nginx-unit/source/v1-green/app-v1-source.go`

1. To containerize add to a minimal [scratch](https://hub.docker.com/_/scratch/) Docker container, or
2. run using Nginx Unit [See here](http://unit.nginx.org/installation/)

See:

* Up on [Docker Hub](https://hub.docker.com/r/armsultan/go-showheaders/)
* /Extras/**app-source.go** for source code
* Dockerfile for build information (very basic)
* Code modified from [arjanschaaf.github.io](https://arjanschaaf.github.io/request-headers-webserver-in-go/)
* [Nginx Unit](http://unit.nginx.org/installation/) 



