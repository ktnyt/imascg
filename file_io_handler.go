package main

import (
	"encoding/json"
	"os"

	rest "github.com/ktnyt/go-rest"
	"github.com/pkg/errors"
	flock "github.com/theckman/go-flock"
)

type FileIOHandler struct {
	lock  *flock.Flock
	build rest.ServiceBuilder
}

func NewFileIOHandler(path string, build rest.ServiceBuilder) rest.IOHandler {
	lock := flock.NewFlock(path)
	handler := FileIOHandler{lock: lock, build: build}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		service := build()
		if err := handler.Save(service); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return handler
}

func (h FileIOHandler) Save(service rest.Service) error {
	if err := h.lock.Lock(); err != nil {
		return errors.Wrap(err, "in FileIOHandler Save")
	}

	file, err := os.Create(h.lock.Path())
	if err != nil {
		return errors.Wrap(err, "in FileIOHandler Save")
	}

	if err := json.NewEncoder(file).Encode(&service); err != nil {
		return errors.Wrap(err, "in FileIOHandler Save")
	}

	if err := h.lock.Unlock(); err != nil {
		return errors.Wrap(err, "in FileIOHandler Save")
	}

	return nil
}

func (h FileIOHandler) Load() (rest.Service, error) {
	if err := h.lock.Lock(); err != nil {
		return nil, errors.Wrap(err, "in FileIOHandler Load")
	}

	file, err := os.Open(h.lock.Path())
	if err != nil {
		return nil, errors.Wrap(err, "in FileIOHandler Load")
	}

	service := h.build()
	if err := json.NewDecoder(file).Decode(&service); err != nil {
		return nil, errors.Wrap(err, "in FileIOHandler Load")
	}

	if err := h.lock.Unlock(); err != nil {
		return nil, errors.Wrap(err, "in FileIOHandler Load")
	}

	return service, nil
}
