package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	method      string            = "GET"
	c_type      string            = ""
	headers     map[string]string = make(map[string]string)
	body_reader io.Reader
	body_out    io.Writer = os.Stdout
	no_out      bool      = false
	quiet       bool      = false
)

func parse_args() bool {
	var met string
	var body string = ""

	flag.StringVar(&met, "m", "", "Metodo a usar (se capitaliza automáticamente)")
	flag.StringVar(&c_type, "ct", c_type, "Contenido del header 'Content-Type'")
	flag.Func("H", "Headers de la consulta (se puede especificar varias veces)", func(H string) error {
		H = strings.TrimSpace(H)
		k := strings.TrimSpace(strings.SplitN(H, ":", 2)[0])
		v := strings.TrimSpace(strings.SplitN(H, ":", 2)[1])

		headers[k] = v
		return nil
	})
	flag.StringVar(&body, "b", body, "Body de la petición (usar @ para leer desde archivos. Ej: @endp3.json)")
	flag.BoolVar(&quiet, "q", quiet, "No imprimir headers e info (a stderr)")
	flag.BoolVar(&no_out, "Q", no_out, "No imprimir body (a stdout)")
	flag.Parse()

	met = strings.ToUpper(met)

	if body != "" {
		method = "POST"

		if body[0] == '@' {
			var err error

			body_reader, err = os.Open(body[1:])
			if err != nil {
				log.Fatal("error al abrir el archivo", err, 2)
			}
		} else {
			body_reader = strings.NewReader(body)
		}

		headers["Content-Type"] = "application/json"
	}

	if met != "" {
		method = met
	}

	if no_out {
		body_out = io.Discard
	}

	if quiet {
		log.SetOutput(io.Discard)
	}

	if c_type != "" {
		headers["Content-Type"] = c_type
	}

	headers["User-Agent"] = "cool/2.0"

	return flag.Parsed()
}

func make_request(url string) *http.Request {
	req, err := http.NewRequest(method, url, body_reader)
	if err != nil {
		return nil
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	return req
}

func main() {
	if !parse_args() {
		flag.PrintDefaults()
		os.Exit(1)
	}
	log.SetFlags(0)

	url := flag.Arg(0)
	if url == "" {
		url = "http://localhost:8080"
	}
	req := make_request(url)

	t := time.Now()
	res, err := http.DefaultClient.Do(req)
	r_t := time.Since(t)

	if err != nil {
		log.Fatal("error al realizar la solicitud", err, 2)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("error al leer el body", err, 2)
	}

	// imprimir headers, URL, método y demás info
	log.Printf("%s %s:\n\tResponse time: %s\n\tHTTP ver.: %s\n\tStatus: %s\n",
		method,
		req.URL,
		r_t.String(),
		res.Proto,
		res.Status,
	)
	for k, v := range res.Header {
		log.Printf("\t%s: %v\n", k, v)
	}
	log.Println("\n------------------------------------------------------------------")

	// imprimir el body
	body_out.Write(data)
}
