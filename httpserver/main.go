package main

import (
	"fmt"
	"log"
	"net/http"
        "math/rand"
	"os"
	"strings"
	"time"
        "httpserver/metrics"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/021", headerModify)
	mux.HandleFunc("/022", getEnv)
	mux.HandleFunc("/023", writeLog)
	mux.HandleFunc("/healthz", healthCheck)
	mux.HandleFunc("/delay", delay)
        mux.Handle("/metrics", promhttp.Handler())

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func headerModify(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		w.Header().Set("Req-"+key, strings.Join(value, "; "))
		fmt.Println(key, "---->", value)
	}
	w.Write([]byte(`接收客户端 request，并将 request 中带的 header 写入 response header`))
}

func getEnv(w http.ResponseWriter, r *http.Request) {
	v := os.Getenv("VERSION")
	//a := os.Getenv("OneDrive")
	//println(a)

	if v == "" {
		v = "0.0.0"
	}
	w.Header().Set("Env-Version", v)
	w.Write([]byte(`读取当前系统的环境变量中的 VERSION 配置，并写入 response header`))
}

func writeLog(w http.ResponseWriter, r *http.Request) {
	l := fmt.Sprintf("Time: %s Method: %s RequestURI: %s IP: %s Status: %d", time.Now().Format(time.UnixDate), r.Method, r.RequestURI, r.RemoteAddr, http.StatusOK)
	w.Write([]byte(l))
	println(l)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func delay(w http.ResponseWriter, r *http.Request) {
        timer := metrics.NewTimer()
        defer timer.ObserveTotal()
        randInt := rand.Intn(2000)
        time.Sleep(time.Millisecond * time.Duration(randInt))
        w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}
