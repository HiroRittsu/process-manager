package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"../utils"
)

// GPUSstatus GPUの情報を配信
func GPUSstatus(w http.ResponseWriter, r *http.Request) {
	flusher, _ := w.(http.Flusher)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// 1秒おきにデータを流す
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()
	go func() {
		for {
			select {
			case <-t.C:
				fmt.Fprintf(w, "data: %f\n\n", utils.GetUsedVRAM()/utils.GetTotalVRAM())
				flusher.Flush()
			}
		}
	}()
	<-r.Context().Done()
	log.Println("コネクションが閉じました")
}
