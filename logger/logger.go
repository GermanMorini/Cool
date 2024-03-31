package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	iPfx string = "[INFO] "
	ePfx string = "[ERROR] "
)

type Logger interface {
	logResponse(res *http.Response, r_time time.Duration)
	logError(msg string, err error)
	logFatal(msg string, err error, code int)
	SetOut(o io.Writer)
}

type BaseLogger struct {
	Logs_out io.Writer
}
type JSONLogger struct {
	BaseLogger
}
type StderrLogger struct {
	BaseLogger
}

func (stdL *StderrLogger) logError(msg string, err error) {
	fmt.Fprintln(stdL.Logs_out, ePfx+msg, err)
}

func (stdL *StderrLogger) logFatal(msg string, err error, code int) {
	stdL.logError(msg, err)
	os.Exit(code)
}

func (stdL *StderrLogger) SetOut(o io.Writer) {
	stdL.Logs_out = o
}

func (jL *JSONLogger) logError(msg string, err error) {
	log := map[string]interface{}{
		"error": err.Error(),
		"msg":   msg,
	}

	jsonData, _ := json.Marshal(log)
	jL.Logs_out.Write(jsonData)
}

func (jL *JSONLogger) logFatal(msg string, err error, code int) {
	jL.logError(msg, err)
	os.Exit(code)
}

func (jL *JSONLogger) logResponse(res *http.Response, r_time time.Duration) {
	log := map[string]interface{}{
		"headers":  res.Header,
		"http_ver": res.Proto,
		"status":   res.Status,
		"time":     r_time.String(),
		"method":   memfn,
		"url":      url + path,
	}

	jsonData, _ := json.Marshal(log)
	jL.Logs_out.Write(jsonData)
}

func (jL *JSONLogger) SetOut(o io.Writer) {
	jL.Logs_out = o
}
func (stdL *StderrLogger) logResponse(res *http.Response, r_time time.Duration) {
	fmt.Fprintf(stdL.Logs_out, "%s %s:\n\tResponse time: %s\n\tHTTP ver.: %s\n\tStatus: %s\n", method, url+path, r_time.String(), res.Proto, res.Status)
	for k, v := range res.Header {
		fmt.Fprintf(stdL.Logs_out, "\t%s: %v\n", k, v)
	}
	stdL.Logs_out.Write([]byte("\n----------------------------------------------------------\n"))
}
