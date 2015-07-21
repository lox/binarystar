package main

import "regexp"

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

func NewMatcherSet() *MatcherSet {
	return &MatcherSet{[]Matcher{}, []Matcher{}}
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

var NullMatcher = MatcherFunc(func(p string) bool {
	return false
})

func NewRegexMatcher(pattern string) (Matcher, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return NullMatcher, err
	}
	return MatcherFunc(func(path string) bool {
		return r.MatchString(path)
	}), nil
}
