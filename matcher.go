package main

import (
	"fmt"
	"github.com/asdine/storm/q"
	strings "strings"
)

func Ss(field string, pattern string) q.Matcher {
  return q.NewFieldMatcher(field, &substrMatcher{ p: pattern })
}

type substrMatcher struct {
  p string
  err error
}

func (m *substrMatcher) MatchField(v interface{}) (bool, error) {
  if m.err != nil {
    return false, m.err
  }

  switch fieldValue := v.(type) {
  case string:
    return strings.Contains(strings.ToUpper(fieldValue), strings.ToUpper(m.p)), nil
  default:
    return false, fmt.Errorf("Only string supported for substr matcher, got %T", fieldValue)
  }
}
