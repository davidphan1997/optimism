package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ethereum-optimism/optimism/l2geth/core/types"
	"github.com/ethereum-optimism/optimism/l2geth/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = ":9100"
	}

	dialAddress := os.Getenv("DIAL_ADDRESS")
	if dialAddress == "" {
		dialAddress = "wss://ws-goerli.optimism.io"
	}

	log.Root().SetHandler(log.CallerFileHandler(log.StdoutHandler))

	go runSubscribeCallStatusChecker(dialAddress)

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

func runSubscribeCallStatusChecker(address string) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		err := getSubscribeCallStatus(address)
		if err != nil {
			wsSubscribeCallStatus.WithLabelValues("error", "true").Inc()
			log.Error("getSubscribeCallStatus error", "error", err)
		}
		wsSubscribeCallStatus.WithLabelValues("error", "false").Inc()
		<-ticker.C
	}
}

func getSubscribeCallStatus(address string) error {
	log.Info("starting getSubscribeCallStatus")

	client, err := ethclient.Dial(address)
	if err != nil {
		log.Error("dial error", "error", err)
		return err
	}
	defer client.Close()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Error("subscribe error", "error", err)
		return err
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Info("finished getSubscribeCallStatus")
			sub.Unsubscribe()
			return nil
		case err := <-sub.Err():
			log.Error("subscription error", "error", err)
			return err
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
		}
	}
}
