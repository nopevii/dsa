package main

import "errors"

var (
    errEmptyQueue    = errors.New("the queue is empty")
    errNotFoundQueue = errors.New("the element is not in the queue")
)

type Queue []string

func (q *Queue) isEmpty() error {
    if len(*q) == 0 {
        return errEmptyQueue
    }

    return nil
}

func (q *Queue) size() int {
    return len(*q)
}

func (q *Queue) enqueue(element string) {
    *q = append(*q, element)
}

func (q *Queue) dequeue() (string, error) {
    if err := q.isEmpty(); err != nil {
        return "", err
    }

    tmp := *q
    element := tmp[0]
    tmp = tmp[1:]
    *q = tmp

    return element, nil
}

func (q *Queue) peek() (string, error) {
    if err := q.isEmpty(); err != nil {
        return "", err
    }

    tmp := *q
    element := tmp[0]

    return element, nil
}

func (q *Queue) contains(element string) (bool, error) {
    if err := q.isEmpty(); err != nil {
        return false, err
    }

    for _, e := range *q {
        if e == element {
            return true, nil
        }
    }

    return false, errNotFoundQueue
}
