package main

import "net/url"

// Model represents the backend model interface
type Model interface {
	Validate() error
	MakeKey(uint64) []byte
	ToBytes() ([]byte, error)
	FromBytes([]byte) error
	ToJSON() ([]byte, error)
	FromJSON([]byte) error
	Filter(url.Values) bool
	Clone() Model
	Merge(Model)
}
