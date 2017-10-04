package binarystar

import (
	"regexp"
	"strings"
)

type Matcher interface {
	Match(path string) bool
}

type MatcherFunc func(path string) bool

func (mf MatcherFunc) Match(path string) bool {
	return mf(path)
}

type MatcherSet struct {
	include []Matcher
	exclude []Matcher
}

func NewMatcherSet(include ...Matcher) *MatcherSet {
	return &MatcherSet{include: include, exclude: []Matcher{}}
}

func (ms *MatcherSet) Match(path string) bool {
	for _, ex := range ms.exclude {
		if ex.Match(path) {
			return false
		}
	}
	if len(ms.include) == 0 {
		return true
	}
	for _, inc := range ms.include {
		if inc.Match(path) {
			return true
		}
	}
	return false
}

func (ms *MatcherSet) Exclude(m Matcher) {
	ms.exclude = append(ms.exclude, m)
}

var MatchAll = MatcherFunc(func(p string) bool {
	return true
})

var MatchNone = MatcherFunc(func(p string) bool {
	return false
})

func MatchPrefix(prefix string) Matcher {
	return MatcherFunc(func(s string) bool {
		return strings.HasPrefix(s, prefix)
	})
}

func Match(exactly string) Matcher {
	return MatcherFunc(func(s string) bool {
		return s == exactly
	})
}

func MatchPattern(pattern string) Matcher {
	r := regexp.MustCompile(pattern)
	return MatcherFunc(func(path string) bool {
		return r.MatchString(path)
	})
}
