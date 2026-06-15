package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metricas
var (
	totalRequisicoes = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "conta o total de requisicoes recebidas",
		},
		[]string{"path", "method"},
	)

	servicoOnline = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_up",
			Help: "se o servico ta online ou nao, 1 = online",
		},
	)
)

// struct do json que o endpoint retorna
type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func init() {
	prometheus.MustRegister(totalRequisicoes)
	prometheus.MustRegister(servicoOnline)
	servicoOnline.Set(1)
}

func handlerProjetoKorp(w http.ResponseWriter, r *http.Request) {
	totalRequisicoes.With(prometheus.Labels{
		"path":   "/projeto-korp",
		"method": r.Method,
	}).Inc()

	// pega horario atual em UTC
	agora := time.Now().UTC().Format(time.RFC3339)

	resp := Resposta{
		Nome:    "Projeto Korp",
		Horario: agora,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("erro ao gerar resposta: %v", err)
	}
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	totalRequisicoes.With(prometheus.Labels{
		"path":   "/health",
		"method": r.Method,
	}).Inc()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/projeto-korp", handlerProjetoKorp)
	http.HandleFunc("/health", handlerHealth)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("subindo servidor na porta 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("erro ao subir o servidor: %v", err)
	}
}
