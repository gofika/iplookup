package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func usage() {
	main := filepath.Base(os.Args[0])
	filesuffix := filepath.Ext(main)
	main = main[0 : len(main)-len(filesuffix)]
	fmt.Fprintf(os.Stderr, "usage: %s <ip|domain> [-n]\n", main)
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nexamples:\n")
	fmt.Fprintf(os.Stderr, "  %s 8.8.8.8\n", main)
	fmt.Fprintf(os.Stderr, "  %s google.com\n", main)
	fmt.Fprintf(os.Stderr, "  %s github.com -n\n", main)
}

var noPretty = flag.Bool("n", false, "no pretty json output")

// isValidIP checks if the input is a valid IP address
func isValidIP(input string) bool {
	return net.ParseIP(input) != nil
}

// isValidDomain checks if the input looks like a domain name
func isValidDomain(input string) bool {
	// Basic domain validation - must contain at least one dot and valid characters
	if !strings.Contains(input, ".") {
		return false
	}
	// Domain regex pattern
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(input)
}

// resolveToIP converts domain to IP address
func resolveToIP(input string) (string, error) {
	// If it's already an IP, return it
	if isValidIP(input) {
		return input, nil
	}

	// If it looks like a domain, try to resolve it
	if isValidDomain(input) {
		ips, err := net.LookupIP(input)
		if err != nil {
			return "", fmt.Errorf("failed to resolve domain %s: %v", input, err)
		}
		if len(ips) == 0 {
			return "", fmt.Errorf("no IP addresses found for domain %s", input)
		}
		// Return the first IPv4 address, or the first IP if no IPv4 found
		for _, ip := range ips {
			if ip.To4() != nil {
				return ip.String(), nil
			}
		}
		return ips[0].String(), nil
	}

	return "", fmt.Errorf("invalid input: %s is neither a valid IP address nor a domain name", input)
}

func main() {
	// Custom parsing to handle "iplookup <ip> -n" format
	var input string
	noPrettyValue := false
	
	// Check if we have at least one argument
	if len(os.Args) < 2 {
		usage()
		return
	}
	
	// Parse arguments manually to support flexible order
	nonFlagArgs := []string{}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-n" {
			noPrettyValue = true
		} else if !strings.HasPrefix(os.Args[i], "-") {
			nonFlagArgs = append(nonFlagArgs, os.Args[i])
		} else {
			// Unknown flag
			usage()
			return
		}
	}
	
	// We should have exactly one non-flag argument (the IP/domain)
	if len(nonFlagArgs) != 1 {
		usage()
		return
	}
	
	input = nonFlagArgs[0]
	*noPretty = noPrettyValue

	// Resolve input to IP address
	ip, err := resolveToIP(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// If input was a domain, show what IP we resolved it to
	if input != ip && !*noPretty {
		fmt.Printf("Resolved %s to %s\n\n", input, ip)
	}

	// Query ipinfo.io
	uri := &url.URL{Scheme: "https", Host: "ipinfo.io", Path: fmt.Sprintf("/%s/json", ip)}
	req := &http.Request{Method: http.MethodGet, URL: uri}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error querying IP info: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body := bytes.NewBufferString("")
	_, err = io.Copy(body, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		os.Exit(1)
	}

	if *noPretty {
		fmt.Println(body.String())
		return
	}

	// Format JSON output
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body.Bytes(), "", "\t")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf(prettyJSON.String())
	fmt.Println()
}
