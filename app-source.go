// modified from https://arjanschaaf.github.io/request-headers-webserver-in-go/
package main

import (
    "fmt"
    "net/http"
    "sort"
    "net"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {

// Show Hostname
name, err := os.Hostname()

         if err != nil {
                 panic(err)
         }

        fmt.Fprintln(w, "<h2>Hostname", name, "</h2>")


// Show IP Address
addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
                fmt.Fprintln(w, "<h2>Hostname", ipnet.IP.String(), "</h2>")
			}
		}
	}

// Show Request Headers
    var keys []string
    for k := range r.Header {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    fmt.Fprintln(w, "<b>Request Headers:</b></br>", r.URL.Path[1:])
    for _, k := range keys {
        fmt.Fprintln(w, k, ":", r.Header[k], "</br>", r.URL.Path[1:])
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}