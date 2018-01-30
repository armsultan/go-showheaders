// modified from https://arjanschaaf.github.io/request-headers-webserver-in-go/
package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sort"
	"nginx/unit"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// Set Content-Type Header
	w.Header().Set("Content-Type", "text/html")

	// Static HTML and CSS styles
	top := `<!DOCTYPE html lang="en">
<head>
  <meta charset="utf-8">
  <title>Show me Headers! V1</title>
  <meta name="description" content="Version1: A Simple application to show Hostname, IP Address and Request Headers">
  <meta name="author" content="Armand Sultantono">
  </head>
  <body>
  <style type="text/css">
		body {
			background-color: blue;
		}
        header {
            font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
            margin: 2em;
            width: auto;
            padding: 5px;
        }
        section {
            border: 1px dashed green;
            margin: 2em;
            width: auto;
            padding: 5px;
            font-family: Courier New, Courier, monospace;
        }
        section ul {
            list-style-type: none;
        }
    </style>
	<header>
	`

	fmt.Fprintln(w, top)

	// Show Hostname
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}
	// Show the text "version 1", this app can be used to show handling of different
	// versions of an application i.e. used show blue/green, A-B and Nginx Unit testing
	fmt.Fprintln(w, "<h1>Version 2</h1>")
	fmt.Fprintln(w, "<h2>Hostname:", name, "</h1>")

	// Show IP Address
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Fprintln(w, "<h2>IP Address:", ipnet.IP.String(), "</h2>")
			}
		}
	}
	fmt.Fprintln(w, "</header>")

	// Show Request Headers
	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Fprintln(w, "<section>")
	fmt.Fprintln(w, "<h3>Request Headers:</h3><ul>", r.URL.Path[1:])
	for _, k := range keys {
		fmt.Fprintln(w, "<li><strong>", k, "</strong>:", r.Header[k], "</li>", r.URL.Path[1:])
	}

	// Static HTML
	bottom := `</ul>
	</section>
	</body>
	</html>`

	fmt.Fprintln(w, bottom)

}

func main() {
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":80", nil)
	unit.ListenAndServe(":80", nil) // For Nginx Unit Compatibility
}
