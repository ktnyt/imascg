package rest

import "net/url"

// Model defines the interface required for REST handling
type Model interface {
	Validate() error
	MakeKey(uint64) []byte
	Filter(url.Values) bool
	Merge(Model) error
}

// Marshalable defines the interface required for Model marshalling
type Marshalable interface {
	GetModel() Model
	SetModel(Model)
	Marshal() ([]byte, error)
	Unmarshal([]byte) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
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
func (m *MarshalableModel) Merge(other MarshalableModel) error {
	return m.m.GetModel().Merge(other.m.GetModel())
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
