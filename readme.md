# Params

An efficient way to represent small map[string]string

## Overview

Params stores key-value pairs in a []struct{key, value string}. For a small number of keys, this is more efficient than a map[string]string (in terms of size and get speed).

By using an array, we can also use a pool.

The approach is discussed [here](http://openmymind.net/Using-Small-Arrays-Instead-Of-Small-Dictionary/)

## Usage

```go
// create a params object that can hold 3 pairs
p := New(3)
p.Set("leto", "ghanima")
fmt.Println(p.Get("leto"))
```

## Delete
An item can be deleted from the set using `Delete(key string)`. The returned values are the same as `Get`, that is, the value and a boolean indicating if the value existing.

## Iteration
Use the `Each(func(key, value string))` function to iterate through the params:

```go
p := New(2)
p.Set("leto", "ghanima")
p.Set("paul", "alia")
p.Each(func(key, value string){
  fmt.Println("%s = %s", key, value)
})
```

## Pool
A main advantage of using an array is to have a pool:

```go
//create a pool that holds 100 items, each capabable of holding 10 elements
//(adding more elemenets will convert them, see `Set` above).
//The pool is thread-safe
var pool := NewPool(10, 100)
...
p := pool.Checkout()
...
p.Release()
```
