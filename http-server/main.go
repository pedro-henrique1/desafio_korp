package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)



var (
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP",
		},
	)

	serviceUp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_up",
			Help: "Disponibilidade do serviço (1 up, 0 down)",
		},
	)
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func init() {
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(serviceUp)
	serviceUp.Set(1)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	httpRequestsTotal.Inc()

	response := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	println("Servidor iniciado na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}