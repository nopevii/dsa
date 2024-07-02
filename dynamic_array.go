package main

import "fmt"

type DynamicArray struct {
    data []string
    size int
    capacity int
}

func newDynamicArray() *DynamicArray {
    capacity := 5

    return &DynamicArray{
        data: make([]string, capacity),
        size: 0,
        capacity: capacity,
    }
}

func (da *DynamicArray) add(item string) {
    if da.size >= da.capacity {
        da.grow()
    }

    da.data[da.size] = item
    da.size += 1
}

func (da *DynamicArray) insert(index int, item string) {
    if da.size >= da.capacity {
        da.grow()
    }

    for i := da.size; i > index; i-- {
        da.data[i] = da.data[i - 1]
    }

    da.data[index] = item
    da.size += 1
}

func (da *DynamicArray) delete(givenItem string) {
    for i, item := range da.data {
        if item == givenItem {
            left := da.data[:i]
            right := da.data[i + 1:]
            da.data = append(left, right...)
            da.size -= 1

            if da.size <= int(da.capacity / 3) {
                da.shrink()
            }

            break
        }
    }
}

func (da *DynamicArray) search(givenItem string) int {
    for i, item := range da.data {
        if item == givenItem {
            return i
        }
    }

    return -1
}

func (da *DynamicArray) grow() {
    da.capacity = int(da.capacity * 2)
    tmp := make([]string, da.capacity)

    for i, item := range da.data {
        tmp[i] = item
    }

    da.data = tmp
}

func (da *DynamicArray) shrink() {
    da.capacity = int(da.capacity / 2)
    tmp := make([]string, da.capacity)

    for i, item := range da.data {
        tmp[i] = item
    }

    da.data = tmp
}

func (da *DynamicArray) isEmpty() bool {
    return da.size == 0
}

func (da *DynamicArray) show() {
    for i, item := range da.data {
        fmt.Printf("%d -> %v\n", i, item)
    }
}

func (da *DynamicArray) toString() string {
    str := ""

    for i, item := range da.data {
        if item == "" {
            continue
        }

        if i == (da.size - 1) {
            str += item
        } else {
            str += fmt.Sprintf("%s, ", item)
        }
    }

    str = fmt.Sprintf("[%s]", str)

    return str
}
