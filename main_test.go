package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSum(t *testing.T) {
	s1 := sum(1, 2)
	if s1 != 3 {
		t.Errorf("Expected sum %d does not match %d", 3, s1)
	}

	s2 := sum(-1, -2)
	if s2 != -3 {
		t.Errorf("Expected sum %d does not match %d", -3, s2)
	}
}

func TestMissingParams(t *testing.T) {
	testGet(t, http.StatusBadRequest, map[string]string{
	}, nil)
}

func TestNonIntParams(t *testing.T) {
	testGet(t, http.StatusBadRequest, map[string]string{
		"x": "a",
		"y": "b",
	}, nil)
}

func TestGoodRequest(t *testing.T) {
	testGet(t, http.StatusOK, map[string]string{
		"x": "5",
		"y": "3",
	}, &message{
		X:   5,
		Y:   3,
		Sum: 8,
	})
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(1, 2)
	}
}

func testGet(t *testing.T, status int, args map[string]string, msg *message) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	q := req.URL.Query()
	for k, v := range args {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	handler(rec, req)
	res := rec.Result()
	if res.StatusCode != status {
		t.Errorf("Expected status %d does not match %d", status, res.StatusCode)
	}

	if msg == nil && status != http.StatusOK {
		return
	}

	resMsg := message{}
	err := json.NewDecoder(res.Body).Decode(&resMsg)
	if err != nil {
		t.Errorf("Expected valid JSON response does not match actual response")
	}

	if resMsg.X != msg.X || resMsg.Y != msg.Y || resMsg.Sum != msg.Sum {
		t.Errorf("Expected sum %d + %d = %d does match %d + %d = %d",
			msg.X, msg.Y, msg.Sum, resMsg.X, resMsg.Y, resMsg.Sum)
	}
}
