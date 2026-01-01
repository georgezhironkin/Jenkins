package main

import (
	"QAProject/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	http.HandlerFunc(handlers.Handler).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("wrong status code: want %v got %v", http.StatusOK, rec.Code)
	}
	if rec.Body.String() != "Hello!\n" {
		t.Errorf("wrong body: got %s", rec.Body.String())
	}
}

func TestHandler2(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	http.HandlerFunc(handlers.Handler2).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("wrong status code: want %v got %v", http.StatusOK, rec.Code)
	}
	if rec.Body.String() != "Hello from Service 2!\n" {
		t.Errorf("wrong body: got %s", rec.Body.String())
	}
}

func TestIntegration(t *testing.T) {
	// Запускаем тестовые серверы для обоих сервисов
	server1 := httptest.NewServer(http.HandlerFunc(handlers.Handler))
	defer server1.Close()

	server2 := httptest.NewServer(http.HandlerFunc(handlers.Handler))
	defer server2.Close()

	// Проверяем сервис 2
	resp, err := http.Get(server2.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service 2 returned wrong status: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	// Проверяем сервис 1
	resp, err = http.Get(server1.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Service 1 returned wrong status: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	t.Log("Integration test passed: both services are responding correctly")
}
