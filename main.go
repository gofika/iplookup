package main

import (
	"bytes"
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
	fmt.Printf("usage: %s <ip>\nexample: iplookup 8.8.8.8\n", main)
}

func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}
	rx := regexp.MustCompile(`\b((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4}\b`)
	if !rx.MatchString(os.Args[1]) {
		usage()
		return
	}
	ip := os.Args[1]
	q := make(url.Values)
	q.Add("action", "get_user_info_data")
	q.Add("ip", ip)
	uri := &url.URL{Scheme: "https", Host: "nordvpn.com", Path: "/wp-admin/admin-ajax.php", RawQuery: q.Encode()}
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
	fmt.Printf("%s", body.String())
}
