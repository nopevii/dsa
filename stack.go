package main

import (
	"errors"
)

var (
    errEmptyStack    = errors.New("the stack is empty")
    errNotFoundStack = errors.New("item not found in the stack")
)

type Stack []string

func (s Stack) isEmpty() bool {
    if len(s) != 0 {
        return false
    }

    return true
}

func (s *Stack) push(item string) {
    *s = append(*s, item)
}

func (s *Stack) pop() (string, error) {
    if s.isEmpty() {
        return "", errEmptyStack
    }

    tmp := *s
    item := tmp[len(tmp) - 1]
    tmp = tmp[:len(tmp) - 1]
    *s = tmp

    return item, nil
}

func (s Stack) peek() (string, error) {
    if s.isEmpty() {
        return "", errEmptyStack
    }

    item := s[len(s) - 1]

    return item, nil
}

func (s Stack) search(toFind string) (int, error) {
    if s.isEmpty() {
        return -1, errEmptyStack
    }

    for i, item := range s {
        if item == toFind {
            return len(s) - i, nil
        }
    }

    return -1, errNotFoundStack
}
