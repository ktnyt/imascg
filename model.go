package imascg

import (
	"encoding/json"
	"net/url"
)

// Model defines the interface required for REST handling
type Model interface {
	Validate() error
	MakeKey(uint64) []byte
	Filter(url.Values) bool
	Merge(Model)
	Clone() Model
}

// Marshalable defines the interface required for Model marshalling
type Marshalable interface {
	GetModel() Model
	SetModel(Model)
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
	Clone() Marshalable
}

// MarshalableModel is the integrated type for Model and Marshalable
type MarshalableModel struct {
	m Marshalable
}

// Validate delegates the call to Marshalable.Model.Validate
func (m *MarshalableModel) Validate() error {
	return m.m.GetModel().Validate()
}

// MakeKey delegates the call to Marshalable.Model.MakeKey
func (m *MarshalableModel) MakeKey(n uint64) []byte {
	return m.m.GetModel().MakeKey(n)
}

// Filter delegates the call to Marshalable.Model.Filter
func (m *MarshalableModel) Filter(v url.Values) bool {
	return m.m.GetModel().Filter(v)
}

// Merge delegates the call to Marshalable.Model.Merge
func (m *MarshalableModel) Merge(other MarshalableModel) {
	m.m.GetModel().Merge(other.m.GetModel())
}

// Marshal delegates the call to Marshalable.Marshal
func (m *MarshalableModel) Marshal() ([]byte, error) {
	return m.m.Marshal()
}

// Unmarshal delegates the call to Marshalable.Unmarshal
func (m *MarshalableModel) Unmarshal(data []byte) error {
	return m.m.Unmarshal(data)
}

// MarshalJSON delegates the call to Marshalable.MarshalJSON
func (m *MarshalableModel) MarshalJSON() ([]byte, error) {
	return m.m.MarshalJSON()
}

// UnmarshalJSON delegates the call to Marshalable.UnmarshalJSON
func (m *MarshalableModel) UnmarshalJSON(data []byte) error {
	return m.m.UnmarshalJSON(data)
}

// Clonel delegates the call to Marshalable.Clone
func (m *MarshalableModel) Clone() MarshalableModel {
	return MarshalableModel{m: m.m.Clone()}
}

// JSONMarshalable implements Marshalable for JSON marshalling
type JSONMarshalable struct {
	model Model
}

// NewJSONModel returns a new JSONMarshalable MarshalableModel
func NewJSONModel(model Model) MarshalableModel {
	return MarshalableModel{m: &JSONMarshalable{model: model}}
}

// GetModel returns the model in JSONMarshalable
func (j *JSONMarshalable) GetModel() Model {
	return j.model
}

// SetModel sets the model in JSONMarshalable
func (j *JSONMarshalable) SetModel(m Model) {
	j.model = m
}

// Marshal marshalls the model to JSON
func (j *JSONMarshalable) Marshal() ([]byte, error) {
	return json.Marshal(&j.model)
}

// Unmarshal unmarshalls the model from JSON
func (j *JSONMarshalable) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &j.model)
}

// MarshalJSON marshalls the model to JSON
func (j *JSONMarshalable) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.model)
}

// UnmarshalJSON unmarshalls the model from JSON
func (j *JSONMarshalable) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.model)
}

// Clone will clone the JSONMarshalable instance
func (j *JSONMarshalable) Clone() Marshalable {
	return &JSONMarshalable{model: j.model.Clone()}
}
