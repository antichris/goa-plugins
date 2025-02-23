// Code generated by goa v3.5.5, DO NOT EDIT.
//
// archiver HTTP server
//
// Command:
// $ goa gen goa.design/plugins/v3/goakit/examples/fetcher/archiver/design -o
// $(GOPATH)/src/goa.design/plugins/goakit/examples/fetcher/archiver

package server

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	archiver "goa.design/plugins/v3/goakit/examples/fetcher/archiver/gen/archiver"
)

// Server lists the archiver service endpoint HTTP handlers.
type Server struct {
	Mounts  []*MountPoint
	Archive http.Handler
	Read    http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the archiver service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *archiver.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Archive", "POST", "/archive"},
			{"Read", "GET", "/archive/{id}"},
		},
		Archive: NewArchiveHandler(e.Archive, mux, decoder, encoder, errhandler, formatter),
		Read:    NewReadHandler(e.Read, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "archiver" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Archive = m(s.Archive)
	s.Read = m(s.Read)
}

// Mount configures the mux to serve the archiver endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountArchiveHandler(mux, h.Archive)
	MountReadHandler(mux, h.Read)
}

// Mount configures the mux to serve the archiver endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountArchiveHandler configures the mux to serve the "archiver" service
// "archive" endpoint.
func MountArchiveHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/archive", f)
}

// NewArchiveHandler creates a HTTP handler which loads the HTTP request and
// calls the "archiver" service "archive" endpoint.
func NewArchiveHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeArchiveRequest(mux, decoder)
		encodeResponse = EncodeArchiveResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "archive")
		ctx = context.WithValue(ctx, goa.ServiceKey, "archiver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountReadHandler configures the mux to serve the "archiver" service "read"
// endpoint.
func MountReadHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/archive/{id}", f)
}

// NewReadHandler creates a HTTP handler which loads the HTTP request and calls
// the "archiver" service "read" endpoint.
func NewReadHandler(
	endpoint endpoint.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeReadRequest(mux, decoder)
		encodeResponse = EncodeReadResponse(encoder)
		encodeError    = EncodeReadError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "read")
		ctx = context.WithValue(ctx, goa.ServiceKey, "archiver")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
