package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func requestTimeWithIP(reqUrl, ip string, timeLimit time.Duration) (time.Duration, error) {
	var tr = http.Transport{
		DialContext: func(ctx context.Context, network, address string) (net.Conn, error) {
			var u, err = url.Parse(reqUrl)
			if err != nil {
				return nil, err
			}
			var port = u.Port()
			if port == "" {
				if u.Scheme == "https" {
					port = "443"
				} else {
					port = "80"
				}
			}

			return net.DialTimeout(network, ip+":"+port, timeLimit)
		},
	}

	var ctx, cancel = context.WithTimeout(context.Background(), timeLimit)
	defer cancel()
	var req, err = http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return 0, err
	}

	var startTime = time.Now()
	var httpClient = http.Client{Transport: &tr, Timeout: timeLimit}
	resp, err := httpClient.Do(req)
	if err == nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}

	return time.Now().Sub(startTime), err
}

func whichIPFastest(reqUrl string, ips []string, timeLimit time.Duration) {
	type result struct {
		ip       string
		takeTime time.Duration
		err      error
	}

	var resultChan = make(chan *result)
	for _, ip := range ips {
		go func(ip string) {
			var takeTime, err = requestTimeWithIP(reqUrl, ip, timeLimit)
			resultChan <- &result{ip, takeTime, err}
		}(ip)
	}

	for range ips {
		var result = <-resultChan
		if result.err != nil {
			fmt.Printf("%s: %s, err: %s\n", result.ip, result.takeTime, result.err)
		} else {
			fmt.Printf("%s: %s\n", result.ip, result.takeTime)
		}
	}

}

func main() {
	var reqUrl = flag.String("url", "https://github.com", "the request url")

	// see: https://api.github.com/meta
	var ips = flag.String("ip", "13.114.40.48,52.192.72.89,52.69.186.44,15.164.81.167,52.78.231.108,13.234.176.102,13.234.210.38,13.229.188.59,13.250.177.223,52.74.223.119,13.236.229.21,13.237.44.5,52.64.108.95,18.228.52.138,18.228.67.229,18.231.5.6",
		"the ip of the url's host, multiple will split by char ,")
	var requestTimeout = flag.Duration("timeout", time.Second*5, "max request time")
	flag.Parse()
	whichIPFastest(*reqUrl, strings.Split(*ips, ","), *requestTimeout)
}
