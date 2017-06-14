# go-showheaders
A Simple Go app to Display hostname, IP address and Request Headers

Using Docker [scratch](https://hub.docker.com/_/scratch/) and a compiled application

**app.go** was complied with:
	
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app.go .

* See **app-source.go** for source code
* See Dockerfile for build information (very basic)

Code modified from 
 [arjanschaaf.github.io](https://arjanschaaf.github.io/request-headers-webserver-in-go/) 