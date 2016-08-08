package utils

import (
	"io/ioutil"
	//"net"
	"net/http"
	"sort"
	"strings"
	//"time"
)

const (
	HTTP_TIMEOUT = 2
)

var (
	transport = http.Transport{
		//Dial: func(network, addr string) (net.Conn, error) {
		//	deadline := time.Now().Add((HTTP_TIMEOUT + 2) * time.Second)
		//	c, err := net.DialTimeout(network, addr, HTTP_TIMEOUT*time.Second)
		//	if err != nil {
		//		return nil, err
		//	}
		//	c.SetDeadline(deadline)
		//	return c, nil
		//},
		DisableKeepAlives: true,
	}
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() HttpClient {
	return HttpClient{
		client: &http.Client{Transport: &transport},
	}
}

//map to query string & sort by key
func httpBuildQuery(params map[string]string) string {
	list := make([]string, 0, len(params))
	buffer := make([]string, 0, len(params))
	for key := range params {
		list = append(list, key)
	}
	sort.Strings(list)
	for _, key := range list {
		value := params[key]
		buffer = append(buffer, key)
		buffer = append(buffer, "=")
		buffer = append(buffer, value)
		buffer = append(buffer, "&")
	}
	buffer = buffer[:len(buffer)-1]
	return strings.Join(buffer, "")
}

func (b *HttpClient) Get(url string) ([]byte, error) {
	resp, err := b.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
