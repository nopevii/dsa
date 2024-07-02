package main

import "fmt"

type Node struct {
    value     string
    next      *Node
}

type LinkedList struct {
    head *Node
    tail *Node
}

func (l *LinkedList) push(value string) {
    n := &Node{
        value: value,
        next:  nil,
    }

    if l.head == nil {
        l.head = n
        l.tail = n
    } else {
        l.tail.next = n
        l.tail = n
    }

}

func (l *LinkedList) insert(value string) {
    n := &Node{
        value: value,
        next:  nil,
    }

    if l.head == nil {
        l.head = n
    } else {
        p := l.head

        for p.next != nil {
            p = p.next
        }

        p.next = n
    }
}

func (l *LinkedList) show() {
    p := l.head

    for p != nil {
        fmt.Printf("-> %v ", p.value)
        p = p.next
    }
}

func (l *LinkedList) indexOf(value string) int {
    p := l.head

    count := 0

    for p != nil {
        if p.value == value {
            return count
        }

        count += 1
        p = p.next
    }

    return -1
}
