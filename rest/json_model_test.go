package rest

import (
	"encoding/json"
	"net/url"
	"testing"
)

type test struct {
	field string
}

func (t *test) Validate() error {
	return nil
}

func (t *test) MakeKey(n uint64) []byte {
	return nil
}

func (t *test) Filter(url.Values) bool {
	return true
}

func (t *test) Merge(Model) error {
	return nil
}

func (t *test) Clone() Model {
	return &test{
		field: t.field,
	}
}

func TestNewJSONModel(t *testing.T) {
	m := &test{
		field: "42",
	}

	NewJSONModel(m)
}

func TestJSONModelMarshalUnmarshal(t *testing.T) {
	m := &test{
		field: "42",
	}

	j := NewJSONModel(m)

	data, err := j.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	err = j.Unmarshal(data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJSONModelJSONMarshalUnmarshal(t *testing.T) {
	m := &test{
		field: "42",
	}

	j := NewJSONModel(m)

	data, err := json.Marshal(j)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(data, &j)
	if err != nil {
		t.Fatal(err)
	}
}
