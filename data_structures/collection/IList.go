package main

type ICollection interface {
    IEnumerable

    Add(item interface{})
    Remove(index uint64)
    Get (index uint64) interface{}
}
type IList interface {
    ICollection
}
