package imascg

import "encoding/json"

// NewJSONModel returns a new jsonMarshalable MarshalableModel
func NewJSONModel(model Model) MarshalableModel {
	return MarshalableModel{m: &jsonMarshalable{model: model}}
}

type jsonMarshalable struct {
	model Model
}

func (j *jsonMarshalable) GetModel() Model {
	return j.model
}

func (j *jsonMarshalable) SetModel(m Model) {
	j.model = m
}

func (j *jsonMarshalable) Marshal() ([]byte, error) {
	return json.Marshal(&j.model)
}

func (j *jsonMarshalable) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &j.model)
}

func (j *jsonMarshalable) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.model)
}

func (j *jsonMarshalable) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.model)
}

func (j *jsonMarshalable) Clone() Marshalable {
	return &jsonMarshalable{model: j.model.Clone()}
}
