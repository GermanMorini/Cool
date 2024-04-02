package logging

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
	LogResponse(res *http.Response, r_time time.Duration)
	LogError(msg string, err error)
	LogFatal(msg string, err error, code int)
	SetOut(o io.Writer)
}

type BaseLogger struct {
	Logs_out          io.Writer
	Url, Path, Method string
}

type JSONLogger struct {
	BaseLogger
}

type StderrLogger struct {
	BaseLogger
}

type ResponseTimeLogger struct {
	BaseLogger
}

func (stdL *StderrLogger) LogError(msg string, err error) {
	fmt.Fprintln(stdL.Logs_out, ePfx+msg, err)
}

func (stdL *StderrLogger) LogFatal(msg string, err error, code int) {
	stdL.LogError(msg, err)
	os.Exit(code)
}

func (stdL *StderrLogger) LogResponse(res *http.Response, r_time time.Duration) {
	fmt.Fprintf(stdL.Logs_out, "%s %s:\n\tResponse time: %s\n\tHTTP ver.: %s\n\tStatus: %s\n", stdL.Method, stdL.Url+stdL.Path, r_time.String(), res.Proto, res.Status)
	for k, v := range res.Header {
		fmt.Fprintf(stdL.Logs_out, "\t%s: %v\n", k, v)
	}
	stdL.Logs_out.Write([]byte("\n----------------------------------------------------------\n"))
}

func (stdL *StderrLogger) SetOut(o io.Writer) {
	stdL.Logs_out = o
}

func (jL *JSONLogger) LogError(msg string, err error) {
	log := map[string]interface{}{
		"error": err.Error(),
		"msg":   msg,
	}

	jsonData, _ := json.Marshal(log)
	jL.Logs_out.Write(jsonData)
}

func (jL *JSONLogger) LogFatal(msg string, err error, code int) {
	jL.LogError(msg, err)
	os.Exit(code)
}

func (jL *JSONLogger) LogResponse(res *http.Response, r_time time.Duration) {
	log := map[string]interface{}{
		"headers":  res.Header,
		"http_ver": res.Proto,
		"status":   res.Status,
		"time":     r_time.String(),
		"method":   jL.Method,
		"url":      jL.Url + jL.Path,
	}

	jsonData, _ := json.Marshal(log)
	jL.Logs_out.Write(jsonData)
}

func (jL *JSONLogger) SetOut(o io.Writer) {
	jL.Logs_out = o
}

func (rtL *ResponseTimeLogger) LogResponse(res *http.Response, r_time time.Duration) {
	rtL.Logs_out.Write([]byte(r_time.String() + "\n"))
}

func (rtL *ResponseTimeLogger) LogError(msg string, err error) {
	fmt.Fprintln(rtL.Logs_out, ePfx+msg, err)
}

func (rtL *ResponseTimeLogger) LogFatal(msg string, err error, code int) {
	rtL.LogError(msg, err)
	os.Exit(code)
}

func (rtL *ResponseTimeLogger) SetOut(o io.Writer) {
	rtL.Logs_out = o
}
