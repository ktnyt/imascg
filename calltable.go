package imascg

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/ktnyt/imascg/rest"
)

// Calltable is the model for calltable
type Calltable struct {
	ID     string  `json:"id"`
	Caller *string `json:"caller"`
	Callee *string `json:"callee"`
	Called *string `json:"called"`
	Remark *string `json:"remark"`
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

	if c.Called == nil {
		missing = append(missing, "'called'")
	}

	if c.Remark == nil {
		missing = append(missing, "'remark'")
	}

	if len(missing) > 0 {
		return fmt.Errorf("Bad Request: %s", strings.Join(missing, ", "))
	}

	return nil
}

// MakeKey for a new calltable entry
func (c *Calltable) MakeKey(n uint64) []byte {
	t, err := time.Now().MarshalBinary()
	if err != nil {
		// If time.Now() generated time cannot be marshalled, what can?
		panic(err)
	}
	key := make([]byte, 0)
	key = append(key, []byte(*c.Caller)...)
	key = append(key, []byte(*c.Callee)...)
	key = append(key, t...)
	return key
}

// Filter calltable entry based on url values
func (c *Calltable) Filter(values url.Values) bool {
	caller := values.Get("caller")
	callee := values.Get("callee")
	called := values.Get("called")
	remark := values.Get("remark")

	if len(caller) > 0 {
		if *c.Caller != caller {
			return false
		}
	}

	if len(callee) > 0 {
		if *c.Callee != callee {
			return false
		}
	}

	if len(called) > 0 {
		if !strings.Contains(*c.Called, called) {
			return false
		}
	}

	if len(remark) > 0 {
		if !strings.Contains(*c.Remark, remark) {
			return false
		}
	}

	return true
}

// Merge another calltable entry into this calltable entry
func (c *Calltable) Merge(m rest.Model) error {
	return mergo.MergeWithOverwrite(c, m)
}

// Clone the calltable entry instance
func (c *Calltable) Clone() rest.Model {
	caller := *c.Caller
	callee := *c.Callee
	called := *c.Called
	remark := *c.Remark
	return &Calltable{
		ID:     c.ID,
		Caller: &caller,
		Callee: &callee,
		Called: &called,
		Remark: &remark,
	}
}
