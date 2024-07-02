package main

import (
    "fmt"
)

func main() {
    // s := Stack{}

    // s.push("test")
    // s.push("test2")
    // s.push("test3")
    // s.push("test4")

    // item, err := s.pop()
    // if err != nil {
    //     fmt.Println(err)
    // }

    // item, err := s.peek()
    // if err != nil {
    //     fmt.Println(err)
    // }

    // fmt.Println(item)

    // position, err := s.search("test")
    // if err != nil {
    //     fmt.Println(err)
    // }

    // fmt.Println(position)

    // fmt.Println(s)

    // q := Queue{}

    // q.enqueue("test")
    // q.enqueue("test1")
    // q.enqueue("test2")
    // q.enqueue("test3")
    // q.enqueue("test4")

    // fmt.Println(q)

    // e, err := q.peek()
    // if err != nil {
    //     fmt.Println(err)
    // } else {
    //     fmt.Println(e)
    // }

    // fmt.Println(q)
    // fmt.Println(q.size())

    // ll := LinkedList{}

    // ll.push("test")
    // ll.push("test2")
    // ll.push("test3")

    // ll.show()

    // fmt.Println(ll.indexOf("test3"))

    da := newDynamicArray()

    // fmt.Println(da.capacity)

    for i := range 10 {
        da.add(fmt.Sprintf("test%d", i))
    }

    // da.insert(1, "test2000")

    fmt.Println(da.search("test0"))

    da.delete("test2")
    da.delete("test5")

    // fmt.Println(da.data)
    // fmt.Println(da.toString())

    // da.show()
}
