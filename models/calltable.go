package models

import (
	"fmt"
	"strings"

	rest "github.com/ktnyt/go-rest"
)

// CallData holds individual called data.
type CallData struct {
	Called string `json:"called"`
	Remark string `json:"remark"`
}

// Calltable is the model for calltable
type Calltable struct {
	ID     string     `json:"id"`
	Caller *string    `json:"caller, omitempty"`
	Callee *string    `json:"callee, omitempty"`
	Data   []CallData `json:"data"`
}

// NewCalltable creates a new empty Calltable.
func NewCalltable() rest.Model {
	return &Calltable{}
}

// Validate the calltable entry fields
func (c *Calltable) Validate() error {
	missing := make([]string, 0)

	if c.Caller == nil {
		missing = append(missing, "'caller'")
	}

	if c.Callee == nil {
		missing = append(missing, "'callee'")
	}

	if c.Data == nil {
		missing = append(missing, "'data'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

// MakeKey for a new calltable entry
func (c *Calltable) MakeKey(i int) string {
	c.ID = *c.Caller + *c.Callee
	return c.ID
}

// Merge another calltable entry into this calltable entry
func (c *Calltable) Merge(m interface{}) error {
	other := m.(*Calltable)
	if other.Caller != nil {
		c.Caller = other.Caller
	}
	if other.Callee != nil {
		c.Callee = other.Callee
	}
	if other.Data != nil {
		c.Data = other.Data
	}
	return nil
}
