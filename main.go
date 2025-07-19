package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

func usage() {
	main := filepath.Base(os.Args[0])
	filesuffix := filepath.Ext(main)
	main = main[0 : len(main)-len(filesuffix)]
	fmt.Fprintf(os.Stderr, "usage: %s <ip> [-no-pretty]\n", main)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nexample: iplookup 8.8.8.8 -no-pretty\n")
}

var noPretty = flag.Bool("n", false, "no pretty json output")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		usage()
		return
	}
	rx := regexp.MustCompile(`\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b`)
	if !rx.MatchString(args[0]) {
		usage()
		return
	}
	ip := args[0]
	uri := &url.URL{Scheme: "https", Host: "ipinfo.io", Path: fmt.Sprintf("/%s/json", ip)}
	req := &http.Request{Method: http.MethodGet, URL: uri}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body := bytes.NewBufferString("")
	_, err = io.Copy(body, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if *noPretty {
		fmt.Println(body.String())
		return
	}
	// Format JSON output
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body.Bytes(), "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(prettyJSON.String())
}
