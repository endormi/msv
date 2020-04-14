package tools

import (
    "context"
    "net/http"
    "github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)

func HTTPServer(ctx context.Context, endpoint Endpoint) http.Handler {
	router := mux.NewRouter()
	router.Use(cmMiddleware)

	router.Methods("POST").Path("/lang").Handler(httptransport.NewServer(
		endpoint.CreateLang,
		decodeLangReq,
		encodeResp,
	))

	router.Methods("GET").Path("/lang/{id}").Handler(httptransport.NewServer(
		endpoints.GetLang,
		decodeNameReq,
		encodeResp,
	))

	return router
}

func cmMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
