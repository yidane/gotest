package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"log"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		counter := prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "testNamespace",
			Subsystem: "testSubsystem",
			Name:      "testName",
		})
		counter.Add(1)
		_, err := writer.Write([]byte("hello world"))
		writer.WriteHeader(200)
		if err != nil {
			writer.WriteHeader(500)
			fmt.Println(err)
		}
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
