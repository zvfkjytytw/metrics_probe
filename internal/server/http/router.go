package metricshttpserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (h *HTTPServer) newRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	// root handler.
	r.Get("/", notImplementedYet)

	// ping handler.
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// metrics update.
	r.Route("/update", func(r chi.Router) {
		r.Get("/", notImplementedYet)
		r.Post("/", notImplementedYet)
		r.Route("/{metricType}/{metricName}/{metricValue}", func(r chi.Router) {
			r.Use(updateCtx)
			r.Post("/", h.putValue)
		})
	})

	// metrics receive.
	r.Route("/metric", func(r chi.Router) {
		r.Get("/", notImplementedYet)
		r.Post("/", notImplementedYet)
		r.Route("/{metricType}/{metricName}", func(r chi.Router) {
			r.Use(metricCtx)
			r.Get("/", h.getValue)
		})
	})

	return r
}

func notImplementedYet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet"))
}
