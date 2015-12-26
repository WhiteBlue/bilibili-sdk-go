package main

import (
	"net/http"
	"log"
	"flag"
	"os"
	"regexp"
)


type Proxy struct {
	LocalPort string //local listen port
	Debug     bool
	Reg       *regexp.Regexp
}

func (this *Proxy) printReq(req *http.Request) {
	if !this.Reg.MatchString(req.URL.String()) {
		log.Println("Request : method [" + req.Method + "] host [" + req.URL.String() + "]")
	}
}


func (this *Proxy) serve(w http.ResponseWriter, req *http.Request) {

	req.RequestURI = ""
	req.RemoteAddr = ""
	req.Header.Del("Proxy-Connection")

	if (this.Debug) {
		this.printReq(req)
	}

	//do formal request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "请求并不合法哦~", 410)
		return
	}
	defer resp.Body.Close()

	//move headers
	h := w.Header()
	for k, v := range resp.Header {
		for _, m := range v {
			h.Add(k, m)
		}
	}
	h.Del("Connection")
	w.WriteHeader(resp.StatusCode)

	data := make([]byte, 8192)
	readCount, readCent, length := int64(0), 0, resp.ContentLength
	if length <= 0 {
		//to max size
		length = (1 << 63) - 1
	}
	for {
		readCent, err = resp.Body.Read(data)
		readCount += int64(readCent)
		w.Write(data[:readCent])
		if err != nil {
			break
		}
		if readCount >= length {
			break
		}
	}
}

func (this *Proxy) Serve() {
	http.HandleFunc("/", this.serve)

	err := http.ListenAndServe(":" + this.LocalPort, nil)
	if (err != nil) {
		log.Panic("Listen: %v", err)
	}
}


func main() {
	var port string;
	var debug bool;
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.StringVar(&port, "listen", "8888", "Listen for connections on this address.")
	flags.BoolVar(&debug, "debug", false, "The logging level.")
	flags.Parse(os.Args[1:])

	log.Println("Serve on localAddr: ", port)
	p := &Proxy{LocalPort:port, Debug:debug, Reg:regexp.MustCompile(".{0,100}.jpg|.png")}
	p.Serve()

}