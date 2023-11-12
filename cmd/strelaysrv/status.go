// Copyright (C) 2015 Audrius Butkevicius and Contributors.

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"
	"sync/atomic"
	"time"
)

var rc *rateCalculator

func statusService(addr string) {
	rc = newRateCalculator(360, 10*time.Second, &bytesProxied)

	handler := http.NewServeMux()
	handler.HandleFunc("/metrics", getStatus)
	if pprofEnabled {
		handler.HandleFunc("/debug/pprof/", pprof.Index)
	}

	srv := http.Server{
		Addr:        addr,
		Handler:     handler,
		ReadTimeout: 15 * time.Second,
	}
	srv.SetKeepAlivesEnabled(false)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	status := make(map[string]interface{})

	sessionMut.Lock()
	// This can potentially be double the number of pending sessions, as each session has two keys, one for each side.
	status["uptime"] = fmt.Sprintf("%d", time.Since(rc.startTime) / time.Second)
	status["pending_session_keys_nr"] = len(pendingSessions)
	status["active_sessions_nr"] = len(activeSessions)
	sessionMut.Unlock()
	status["connections_nr"] = numConnections.Load()
	status["proxies_nr"] = numProxies.Load()
	status["bytes_proxied"] = bytesProxied.Load()

	status["kbps{interval=\"10s\"}"] = rc.rate(1) * 8 / 1000
	status["kbps{interval=\"1m\"}"] = rc.rate(60/10) * 8 / 1000
	status["kbps{interval=\"5m\"}"] = rc.rate(5*60/10) * 8 / 1000
	status["kbps{interval=\"15m\"}"] = rc.rate(15*60/10) * 8 / 1000
	status["kbps{interval=\"30m\"}"] = rc.rate(30*60/10) * 8 / 1000
	status["kbps{interval=\"60m\"}"] = rc.rate(60*60/10) * 8 / 1000
	status["routine_nr"] = runtime.NumGoroutine()

	resp := ""

	for k, v := range status {
            resp += fmt.Sprintf("strelaysrv_%s %v\n", k, v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resp))
}

type rateCalculator struct {
	counter   *atomic.Int64
	rates     []int64
	prev      int64
	startTime time.Time
}

func newRateCalculator(keepIntervals int, interval time.Duration, counter *atomic.Int64) *rateCalculator {
	r := &rateCalculator{
		rates:     make([]int64, keepIntervals),
		counter:   counter,
		startTime: time.Now(),
	}

	go r.updateRates(interval)

	return r
}

func (r *rateCalculator) updateRates(interval time.Duration) {
	for {
		now := time.Now()
		next := now.Truncate(interval).Add(interval)
		time.Sleep(next.Sub(now))

		cur := r.counter.Load()
		rate := int64(float64(cur-r.prev) / interval.Seconds())
		copy(r.rates[1:], r.rates)
		r.rates[0] = rate
		r.prev = cur
	}
}

func (r *rateCalculator) rate(periods int) int64 {
	var tot int64
	for i := 0; i < periods; i++ {
		tot += r.rates[i]
	}
	return tot / int64(periods)
}
