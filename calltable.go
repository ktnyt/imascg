package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	rest "github.com/ktnyt/go-rest"
	"github.com/mitchellh/mapstructure"
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

// CalltableFilter applies the given filter to the model.
func CalltableFilter(ctx context.Context) rest.Filter {
	return func(value interface{}) bool {
		params := ctx.Value(rest.Params).(url.Values)

		caller := params.Get("caller")
		callee := params.Get("callee")

		calltable := value.(*Calltable)

		if len(caller) > 0 {
			if *calltable.Caller != caller {
				return false
			}
		}

		if len(callee) > 0 {
			if *calltable.Callee != callee {
				return false
			}
		}

		return true
	}
}

// CalltableConverter converts a value to a Model.
func CalltableConverter(value interface{}) rest.Model {
	calltable := new(Calltable)
	mapstructure.Decode(value, &calltable)
	return calltable
}

// NewCalltableService creates a default Calltable Service.
func NewCalltableService() rest.Service {
	return rest.NewDictService(NewCalltable, CalltableFilter, CalltableConverter)
}

func init() {
	name := "calltable"
	handler := NewFileIOHandler(name+".db", NewCalltableService)
	service := rest.NewIOService(handler)
	iface := rest.NewServiceInterface(service)
	register(name, router, iface)
}
