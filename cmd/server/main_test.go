package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// verifica se retorna 200
func TestStatusOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/projeto-korp", nil)
	if err != nil {
		t.Fatalf("erro ao criar req: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handlerProjetoKorp).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperava 200, veio %d", rr.Code)
	}
}

// verifica o content type da resposta
func TestContentTypeJson(t *testing.T) {
	req, err := http.NewRequest("GET", "/projeto-korp", nil)
	if err != nil {
		t.Fatalf("erro ao criar req: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handlerProjetoKorp).ServeHTTP(rr, req)

	ct := rr.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("content-type errado: %s", ct)
	}
}

// verifica os campos do json
func TestCamposResposta(t *testing.T) {
	req, err := http.NewRequest("GET", "/projeto-korp", nil)
	if err != nil {
		t.Fatalf("erro ao criar req: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handlerProjetoKorp).ServeHTTP(rr, req)

	var resp Resposta
	err = json.NewDecoder(rr.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("erro ao decodar json: %v", err)
	}

	if resp.Nome != "Projeto Korp" {
		t.Errorf("nome errado: %s", resp.Nome)
	}

	if resp.Horario == "" {
		t.Error("horario vazio")
	}
}

// verifica se o horario ta em UTC
func TestHorarioUTC(t *testing.T) {
	req, err := http.NewRequest("GET", "/projeto-korp", nil)
	if err != nil {
		t.Fatalf("erro ao criar req: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handlerProjetoKorp).ServeHTTP(rr, req)

	var resp Resposta
	json.NewDecoder(rr.Body).Decode(&resp)

	horario, err := time.Parse(time.RFC3339, resp.Horario)
	if err != nil {
		t.Errorf("formato de horario invalido: %v", err)
	}

	if horario.Location().String() != "UTC" {
		t.Errorf("nao esta em UTC: %s", horario.Location().String())
	}
}

// testa o health
func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("erro ao criar req: %v", err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handlerHealth).ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperava 200, veio %d", rr.Code)
	}

	if rr.Body.String() != "ok" {
		t.Errorf("esperava ok, veio %s", rr.Body.String())
	}
}

// verifica se o horario muda a cada requisicao
func TestHorarioMuda(t *testing.T) {
	req1, _ := http.NewRequest("GET", "/projeto-korp", nil)
	rr1 := httptest.NewRecorder()
	http.HandlerFunc(handlerProjetoKorp).ServeHTTP(rr1, req1)

	time.Sleep(1 * time.Second)

	req2, _ := http.NewRequest("GET", "/projeto-korp", nil)
	rr2 := httptest.NewRecorder()
	http.HandlerFunc(handlerProjetoKorp).ServeHTTP(rr2, req2)

	var r1, r2 Resposta
	json.NewDecoder(rr1.Body).Decode(&r1)
	json.NewDecoder(rr2.Body).Decode(&r2)

	if r1.Horario == r2.Horario {
		t.Error("horario nao mudou entre requisicoes")
	}
}
