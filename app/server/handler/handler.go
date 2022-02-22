package handler

import (
	"errors"
	"net/http"

	"github.com/zchary-ma/proto-mux-template/app/pkg/rpcx"

	"github.com/zchary-ma/proto-mux-template/app/server"

	"github.com/zchary-ma/proto-mux-template/app/log"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	responseLimitSize        = 32 * 1024 * 1024
	HeaderAccept             = "Accept"
	HeaderContentType        = "Content-Type"
	MimeApplicationJSON      = "application/json"
	MimeApplicationXProtobuf = "application/x-protobuf"
)

var (
	errResponseLimitSize = errors.New("response data exceeds the limit size")
)

type handlerFunc func(w http.ResponseWriter, r *http.Request) (proto.Message, error)

func (fn handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	m, err := fn(w, r)
	if err != nil {
		if se, ok := err.(*server.Err); ok {
			log.Errorf("handler errors, status code: %d, message: %s, underlying err: %v")
			http.Error(w, se.Message, se.Code)
			return
		}

		s, ok := status.FromError(err)
		if !ok {
			log.Errorf("handler errors: %v", err)
			http.Error(w, "internal errors", http.StatusInternalServerError)
			return
		}
		statusCode := rpcx.HTTPStatusFromCode(s.Code())
		log.Errorf("handler errors: %v, status code: %d", err, statusCode)

		m = s.Proto()
		w.WriteHeader(statusCode)
	}

	if m == nil {
		return
	}
	// NOTE add trace here
	accept := r.Header.Get(HeaderAccept)
	switch accept {
	case MimeApplicationJSON:
		writeJSON(w, m)
	case MimeApplicationXProtobuf:
		writeProtobuf(w, m)
	// TODO add more content type support
	default:
		writeProtobuf(w, m)
	}
}

func writeJSON(w http.ResponseWriter, m proto.Message) {
	w.Header().Set(HeaderContentType, MimeApplicationJSON)
	jsonBytes, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(m)
	if err != nil {
		httpError(w, err, "failed to marshal proto message")
		return
	}
	if len(jsonBytes) > responseLimitSize {
		httpError(w, errResponseLimitSize, "failed to http response")
		return
	}
	if _, err := w.Write(jsonBytes); err != nil {
		httpError(w, err, "failed to write to http response")
	}
}

func writeProtobuf(w http.ResponseWriter, m proto.Message) {
	w.Header().Set(HeaderContentType, MimeApplicationXProtobuf)
	protoBytes, err := proto.Marshal(m)
	if err != nil {
		httpError(w, err, "failed to marshal proto message")
		return
	}
	if len(protoBytes) > responseLimitSize {
		httpError(w, errResponseLimitSize, "failed to http response")
		return
	}
	if _, err := w.Write(protoBytes); err != nil {
		httpError(w, err, "failed to write to http response")
	}
}

func httpError(w http.ResponseWriter, err error, msg string) {
	e := server.Err{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Cause:   err,
	}

	log.Errorf("handler errors, status code: %d, message: %s, underlying err: %v",
		e.Code, e.Message, e.Cause)
	http.Error(w, e.Message, e.Code)
}
