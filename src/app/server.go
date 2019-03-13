package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func echoEnv(w http.ResponseWriter, r *http.Request) {
	envVal := os.Getenv("HOGE")
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Hello!! Your HOGE's value is '%s'\n at %s", envVal, nowStr)
	log.Printf("[%s] %s\n", nowStr, envVal)
	return
}

func main() {
	logfile, err := os.OpenFile("/var/logs/server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("Cannot open log file: %s" + err.Error())
	}
	log.SetOutput(logfile)
	defer logfile.Close()

	http.HandleFunc("/", echoEnv)
	http.ListenAndServe(":9200", nil)
}
