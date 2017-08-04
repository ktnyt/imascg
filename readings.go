package main

import (
  "github.com/asdine/storm/q"
  "strings"
)

// ReadingTuple holds the reading information body.
type ReadingTuple struct {
  ID      string `json:"id"`
  Reading string `json:"reading"`
}

type readingSubstrMatcher struct {
  pattern string
  err error
}

// MatchField matches a ReadingTuple with a given string.
func (m *readingSubstrMatcher) MatchField(value interface{}) (bool, error) {
  tuple := value.(ReadingTuple)

  return strings.Contains(
    strings.ToUpper(tuple.Reading),
    strings.ToUpper(m.pattern),
  ), nil
}

// ReadingSubstr matcher, checks if a ReadingTuple contains the given value as a substring.
func ReadingSubstr(field string, pattern string) q.Matcher {
  return q.NewFieldMatcher(field, &readingSubstrMatcher{ pattern: pattern })
}
