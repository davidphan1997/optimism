package main

import (
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":9100"
	}

	log.Root().SetHandler(log.CallerFileHandler(log.StdoutHandler))

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
		<head><title>edge-proxyd-exporter</title></head>
		<body>
		<h1>edge-proxyd-exporter</h1>
		<p><a href="/metrics">Metrics</a></p>
		</body>
		</html>`))
	})

	log.Info("Program starting", "listenAddress", listenAddress)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Error("Can't start http server", "error", err)
	}

}
