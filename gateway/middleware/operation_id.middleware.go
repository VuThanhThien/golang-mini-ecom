package middleware

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

// OperationID adds a unique operation ID to each request
func OperationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		operationID := uuid.New().String()
		ctx := context.WithValue(r.Context(), "operation-id", operationID)
		w.Header().Set("X-Operation-ID", operationID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Logging logs information about each request and response
func Logging(log zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			operationID, _ := r.Context().Value("operation-id").(string)

			// Log request details
			bodyBytes, _ := io.ReadAll(r.Body)
			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			log.Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Str("remote_addr", r.RemoteAddr).
				Str("operation-id", operationID).
				Str("params", r.URL.Query().Encode()).
				Str("body", string(bodyBytes)).
				Msg("Incoming request")

			crw := &customResponseWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(crw, r)

			// Log response details
			log.Info().
				Int("status", crw.status).
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Str("operation-id", operationID).
				Msg("Request completed")
		})
	}
}

type customResponseWriter struct {
	http.ResponseWriter
	status int
}

func (crw *customResponseWriter) WriteHeader(status int) {
	crw.status = status
	crw.ResponseWriter.WriteHeader(status)
}
