package main

import (
	"encoding/json"
	"net/url"
)

type Model interface {
	Validate() error
	MakeKey(uint64) []byte
	Filter(url.Values) bool
	Merge(Model)
	Clone() Model
}

type Marshalable interface {
	GetModel() Model
	SetModel(Model)
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
	Clone() Marshalable
}

type MarshalableModel struct {
	m Marshalable
}

func (m *MarshalableModel) Validate() error {
	return m.m.GetModel().Validate()
}

func (m *MarshalableModel) MakeKey(n uint64) []byte {
	return m.m.GetModel().MakeKey(n)
}

func (m *MarshalableModel) Filter(v url.Values) bool {
	return m.m.GetModel().Filter(v)
}

func (m *MarshalableModel) Merge(other MarshalableModel) {
	m.m.GetModel().Merge(other.m.GetModel())
}

func (m *MarshalableModel) Marshal() ([]byte, error) {
	return m.m.Marshal()
}

func (m *MarshalableModel) Unmarshal(data []byte) error {
	return m.m.Unmarshal(data)
}

func (m *MarshalableModel) MarshalJSON() ([]byte, error) {
	return m.m.MarshalJSON()
}

func (m *MarshalableModel) UnmarshalJSON(data []byte) error {
	return m.m.UnmarshalJSON(data)
}

func (m *MarshalableModel) Clone() MarshalableModel {
	return MarshalableModel{m:m.m.Clone()}
}

type JSONMarshalable struct {
	model Model
}

func NewJSONModel(model Model) MarshalableModel {
	return MarshalableModel{m:&JSONMarshalable{model:model}}
}

func (j *JSONMarshalable) GetModel() Model {
	return j.model
}

func (j *JSONMarshalable) SetModel(m Model) {
	j.model = m
}

func (j *JSONMarshalable) Marshal() ([]byte, error) {
	return json.Marshal(&j.model)
}

func (j *JSONMarshalable) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &j.model)
}

func (j *JSONMarshalable) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.model)
}

func (j *JSONMarshalable) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.model)
}

func (j *JSONMarshalable) Clone() Marshalable {
	return &JSONMarshalable{model:j.model.Clone()}
}
