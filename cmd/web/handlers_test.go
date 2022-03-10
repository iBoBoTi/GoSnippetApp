package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestPing(t *testing.T) {
//	rr := httptest.NewRecorder()
//
//	r, err := http.NewRequest("GET", "/", nil)
//	if err != nil {
//		t.Fatal(err)
//	}
//	ping(rr, r)
//	rs := rr.Result()
//	if rs.StatusCode != http.StatusOK {
//		t.Errorf("want %v, got %v", http.StatusOK, rs.StatusCode)
//	}
//
//	defer rs.Body.Close()
//	body, err := ioutil.ReadAll(rs.Body)
//	if err != nil {
//		t.Fatal(err)
//	}
//	if string(body) != "OK" {
//		t.Errorf("want %v, got %v", "OK", string(body))
//	}
//}

func TestPing(t *testing.T) {
	app := &application{
		errorLog: log.New(ioutil.Discard, "", 0),
		infoLog:  log.New(ioutil.Discard, "", 0),
	}

	ts := httptest.NewTLSServer(app.Routes())
	defer ts.Close()
	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %v, got %v", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "OK" {
		t.Errorf("want %v, got %v", "OK", string(body))
	}
}
