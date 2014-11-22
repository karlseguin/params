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
p = p.Set("leto", "ghanima")
fmt.Println(p.Get("leto"))
```

**It's important to note that when you put more than the configured number of pairs, the object will convert itself to a `map[string]string`.** Therefore, you must re-assign the return value of `Set`, much like you do with Go's built-in `append`:

```go
p := New(2)
p = p.Set("leto", "ghanima")
p = p.Set("paul", "alia")
p = p.Set("feyd", "glossu")
//p is now a different type of object
```

## Delete
Delete isn't currently supported.

## Iteration
Use the `Each(func(key, value string))` function to iterate through the params:

```go
p := New(2)
p = p.Set("leto", "ghanima")
p = p.Set("paul", "alia")
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
