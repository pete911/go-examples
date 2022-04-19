package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type multiResponseWriter struct {
	writer         io.Writer
	responseWriter http.ResponseWriter
}

func (m *multiResponseWriter) Write(p []byte) (int, error) {
	return m.writer.Write(p)
}

func (m *multiResponseWriter) Header() http.Header {
	return m.responseWriter.Header()
}

func (m *multiResponseWriter) WriteHeader(statusCode int) {
	m.responseWriter.WriteHeader(statusCode)
}

func main() {
	mux := http.NewServeMux()
	h := http.HandlerFunc(handler)
	mux.Handle("/", middleware(h))

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// copy request body and re-set it back on request, so it can be used in the handler
		var requestBody []byte
		if r.Body != nil {
			requestBody, _ = io.ReadAll(r.Body)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		// set response to multi writer so we can read response body
		responseBody := new(bytes.Buffer)
		mw := io.MultiWriter(w, responseBody)
		w = &multiResponseWriter{writer: mw, responseWriter: w}

		log.Printf("middleware request: %s", string(requestBody))
		next.ServeHTTP(w, r)
		log.Printf("middleware response: %s", responseBody.String())
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	request, _ := io.ReadAll(r.Body)
	log.Printf("handler request: %s", string(request))
	response := "OK"
	w.Write([]byte(response))
	log.Printf("handler response: %s", response)
}
