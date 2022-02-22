package handler

import (
	"errors"
	"net/http"

	"github.com/zchary-ma/proto-mux-template/app/schema"

	"google.golang.org/protobuf/proto"
)

func Hello(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
	q := r.URL.Query().Get("m")
	return &schema.Hello{
		Name: q + "Hello, World!",
	}, nil
}

func SayHello(w http.ResponseWriter, r *http.Request) (proto.Message, error) {
	return nil, errors.New("not implemented")
}
