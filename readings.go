package main

import (
  "github.com/asdine/storm/q"
  "strings"
)

type ReadingTuple struct {
  ID      string `json:"id"`
  Reading string `json:"reading"`
}

type readingSubstrMatcher struct {
  pattern string
  err error
}

func (m *readingSubstrMatcher) MatchField(value interface{}) (bool, error) {
  tuple := value.(ReadingTuple)

  return strings.Contains(
    strings.ToUpper(tuple.GetReading()),
    strings.ToUpper(m.pattern),
  ), nil
}

func (t *ReadingTuple) GetReading() string {
  return t.Reading
}

func ReadingSubstr(field string, pattern string) q.Matcher {
  return q.NewFieldMatcher(field, &readingSubstrMatcher{ pattern: pattern })
}
