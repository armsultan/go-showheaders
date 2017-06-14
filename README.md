# go-showheaders
A Simple Go app to Display hostname, IP address and Request Headers

For maximum lightweightness:

1. **app.go** was complied with:
	
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app.go .

1. Then added to a minimal [scratch](https://hub.docker.com/_/scratch/) Docker container

See:

* **app-source.go** for source code
* Dockerfile for build information (very basic)

Code modified from 
 [arjanschaaf.github.io](https://arjanschaaf.github.io/request-headers-webserver-in-go/) 