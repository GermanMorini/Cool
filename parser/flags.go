package parser

import (
	"os"
)

type Flags struct {
	Url           string    = os.Getenv("COOL_URL")
	Path          string    = ""
	Method        string    = "GET"
	C_type        string    = "application/json"
	Body          string    = ""
	Quiet         bool      = false
	No_out        bool      = false
	Json_logs     bool      = false
	Response_time bool      = false
}

func (f *Flags) ParseArgs() bool {
	var met string

	flag.StringVar(&f.Url, "u", url, "Dirección URL")
	flag.StringVar(&f.Path, "p", path, "Path de la dirección (se le adiciona a la URL base)")
	flag.StringVar(&f.Met, "m", "", "Metodo a usar (GET, POST, PUT, DELETE)")
	flag.StringVar(&f.C_type, "ct", c_type, "Content type")
	flag.StringVar(&f.Body, "b", body, "Body de la petición (usar @ para leer desde archivos. Ej: @endp3.json)")
	flag.BoolVar(&f.Quiet, "q", quiet, "No imprimir headers e info (a stderr)")
	flag.BoolVar(&f.No_out, "Q", no_out, "No imprimir body (a stdout)")
	flag.BoolVar(&f.Response_time, "rt", response_time, "Mide el tiempo de respuesta (no imprime el body)")
	flag.BoolVar(&f.Json_logs, "j", json_logs, "Logs en formato json")
	flag.Parse()

	if f.Url == "" {
		f.Url = "http://localhost:8080"
	}

	if f.Body != "" {
		f.Method = "POST"

		if f.Body[0] == '@' {
			file, err := os.Open(body[1:])
			if err != nil {
				logger.LogFatal("error al abrir el archivo", err, 2)
			}

			data, err := io.ReadAll(file)
			if err != nil {
				logger.LogFatal("error al leer el archivo", err, 2)
			}
			file.Close()

			body = string(data)
		}
	}
	if met != "" {
		method = met
	}
	if no_out {
		body_out = io.Discard
	}

	base_logger := log.BaseLogger{
		Logs_out: os.Stderr,
		Url:      url,
		Path:     path,
		Method:   method,
	}

	switch {
	case response_time:
		logger = &log.ResponseTimeLogger{BaseLogger: base_logger}
		body_out = io.Discard
	case json_logs:
		logger = &log.JSONLogger{BaseLogger: base_logger}
	default:
		logger = &log.StderrLogger{BaseLogger: base_logger}
	}

	if quiet {
		logger.SetOut(io.Discard)
	}

	return flag.Parsed()
}
