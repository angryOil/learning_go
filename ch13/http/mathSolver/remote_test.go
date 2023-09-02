package mathSolver

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type info struct {
	expression string
	code       int
	body       string
}

var ino info

func TestServer(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			expression := r.URL.Query().Get("expression")
			fmt.Println("ex?:" + ino.expression)
			fmt.Println("url?:" + r.URL.String())
			if expression != ino.expression {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid expression" + ino.expression))
				return
			}
			w.WriteHeader(ino.code)
			w.Write([]byte(ino.body))
		}),
	)
	defer server.Close()
	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	}
	data := []struct {
		name   string
		io     info
		result float64
	}{
		{"case1", info{"2+2*10", http.StatusOK, "22"}, 22},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			ino = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if err != nil {
				t.Error(err.Error())
			}
			if result != d.result {
				t.Errorf("ino `%f`,got `%f`", d.result, result)
			}

		})
	}
}
