package middlewares

import (
	"github.com/VyacheslavKuzharov/gophermart/pkg/logger"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func Logger(l *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logFn := func(w http.ResponseWriter, r *http.Request) {
			l.Logger.Info().
				Str("method", r.Method).
				Str("uri", r.RequestURI).
				Str("query", r.URL.RawQuery).
				Msg("Request Started")

			start := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				l.Logger.Info().
					Int("status", ww.Status()).
					Str("method", r.Method).
					Str("uri", r.RequestURI).
					Str("query", r.URL.RawQuery).
					Dur("duration", time.Since(start)).
					Int("bytes", ww.BytesWritten()).
					Msg("Request Completed")
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(logFn)
	}
}
