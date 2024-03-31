package main

import (
	"cool/logger"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	url       string    = os.Getenv("COOL_URL")
	path      string    = ""
	method    string    = "GET"
	c_type    string    = "application/json"
	body      string    = ""
	quiet     bool      = false
	no_out    bool      = false
	json_logs bool      = false
	body_out  io.Writer = os.Stdout
	logger    Logger    = &StderrLogger{BaseLogger{os.Stderr}}
)

func parse_args() bool {
	var met string

	flag.StringVar(&url, "url", url, "Dirección HTTP")
	flag.StringVar(&path, "path", path, "Path de la dirección (se le adiciona a la URL base)")
	flag.StringVar(&met, "m", "", "Metodo a usar (GET, POST, PUT, DELETE)")
	flag.StringVar(&c_type, "ct", c_type, "Content type")
	flag.StringVar(&body, "b", body, "Body de la petición (usar @ para leer desde archivos. Ej: @endp3.json)")
	flag.BoolVar(&quiet, "q", quiet, "No imprimir headers e info (a stderr)")
	flag.BoolVar(&no_out, "Q", no_out, "No imprimir body (a stdout)")
	flag.BoolVar(&json_logs, "j", json_logs, "Logs en formato json")
	flag.Parse()

	if url == "" {
		url = "http://localhost:8080"
	}

	if json_logs {
		logger = &JSONLogger{BaseLogger{os.Stderr}}
	}

	if body != "" {
		method = "POST"

		if body[0] == '@' {
			file, err := os.Open(body[1:])
			if err != nil {
				logger.logFatal("error al abrir el archivo", err, 1)
			}

			data, err := io.ReadAll(file)
			if err != nil {
				logger.logFatal("error al leer el archivo", err, 2)
			}
			file.Close()

			body = string(data)
		}
	}
	if met != "" {
		method = met
	}
	if quiet {
		logger.SetOut(io.Discard)
	}
	if no_out {
		body_out = io.Discard
	}

	return flag.Parsed()
}

func make_request() *http.Request {
	req, err := http.NewRequest(method, url+path, strings.NewReader(body))
	if err != nil {
		return nil
	}
	if body != "" {
		req.Header.Set("Content-Type", c_type)
	}

	return req
}

func main() {
	if !parse_args() {
		flag.PrintDefaults()
		os.Exit(1)
	}

	req := make_request()
	t := time.Now()
	response, err := http.DefaultClient.Do(req)
	r_t := time.Since(t)
	if err != nil {
		logger.logFatal("error al realizar la solicitud", err, 2)
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		logger.logFatal("error al leer el body", err, 3)
	}

	logger.logResponse(response, r_t)

	body_out.Write(data)
}
