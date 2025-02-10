package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/iolave/go-proxmox/pkg/errors"
)

const content_type_json = "application/json"

func jsonHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[INF]", r.Method, r.URL.Path, "started")
		b, err := io.ReadAll(r.Body)
		if err != nil {
			err := errors.NewHTTPError(
				http.StatusInternalServerError,
				"unable to read request body",
				err,
			)
			log.Println("[ERR]", r.Method, r.URL.Path, "failed", err.Message, err.Original)
			err.WriteResponse(w)
			return
		}
		contentType := r.Header.Get("Content-Type")
		if contentType != content_type_json {
			err := errors.NewHTTPError(
				http.StatusBadRequest,
				"request body is not valid json",
				nil,
			)
			log.Println("[ERR]", r.Method, r.URL.Path, "failed", err.Message, err.Original)
			err.WriteResponse(w)
			return
		}

		var body any
		if err := json.Unmarshal(b, &body); err != nil {
			err := errors.NewHTTPError(
				http.StatusBadRequest,
				"request body is not valid json",
				err,
			)
			log.Println("[ERR]", r.Method, r.URL.Path, "failed", err.Message, err.Original)
			err.WriteResponse(w)
			return
		}

		w.Header().Set("Content-Type", content_type_json)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
		log.Println("[INF]", r.Method, r.URL.Path, "success")
		return
	})
}
