@[TOC](Content)



# orderedmap

Unlike an unordered map, the orderedmap package is an orderedmap type. Orderedmap is an inherited Map maps keys to values. Orderedmap provides useful methods e.g.  `Store`, `Load`, `Delete` , `Range` and so on. The remaining methods are order-aware. Big-O running times for all methods are the same as regular maps.



# API

## Create orderedmap instance

```go
func New() *OrderedMap
```

For example:

```go
package main

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
)

func main() {
	o := orderedmap.New()
	fmt.Printf("%T", o)
	// output: *orderedmap.OrderedMap
}

```



## Store the data

Store method used to set the value for a key.

```go
func (om *OrderedMap) Store(key, value interface{})
```

For example:

```go
package main

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
)

func main() {
	o := orderedmap.New()
	o.Store("name", "Pizza")
	o.Store("price", 50)
	o.Store("size", "10#")
	fmt.Println(o)
	// output: {name: Piza price: 50 size: 10#}
}

```



## Load the data

Load method used to get the value from a key. If the key is not exist in the map, a nil and a false will be returned.

```go
func (om *OrderedMap) Load(key interface{}) (interface{}, bool)
```

For example:

```go
package main

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
)

func main() {
	o := orderedmap.New()
	o.Store("name", "Pizza")
	o.Store("price", 50)
	o.Store("size", "10#")

	size, ok := o.Load("size")
	fmt.Println(size, ok)
	// output: 10# true

	color, ok := o.Load("color")
	fmt.Println(color, ok)
	// output: <nil> false
}

```



## Delete the data

Delete method used to deletes the value for a key.

```go
func (om *OrderedMap) Delete(key interface{})
```

For example:

```go
package main

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
)


func main() {
	o := orderedmap.New()
	o.Store("name", "Pizza")
	o.Store("price", 50)
	o.Store("size", "10#")
	fmt.Println("Before using delete method, the content of orderedmap is: ", o)
	// Before using delete method, the content of orderedmap is:  {name: Pizza price: 50 size: 10#}

	o.Delete("size")
	fmt.Println("After using delete method, the content of orderedmap is: ", o)
	// After using delete method, the content of orderedmap is:  {name: Pizza price: 50}
}

```



## Range map

You can use range method to visit each key and value. Range method needs a argument which type is `f func(key, value interface{}) bool`. If f returns false, range stops the iteration.

```go
func (om *OrderedMap) Range(f func(key, value interface{}) bool)
```

For example:

```go
package main

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
)

func main() {
	o := orderedmap.New()
	o.Store("A", "a")
	o.Store("B", "b")

	f := func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	}

	o.Range(f)
	// outputs:
	// key: A, value: a
	// key: B, value: b
}

```



## Get the length of orderedmap

Using length method to get OrderedMap's length.

```go
func (om *OrderedMap) Length() int
```

For example:

```go
package main

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
)

func main() {
	o := orderedmap.New()
	o.Store("A", "a")
	o.Store("B", "b")

	fmt.Println(o.Length())
	// output: 2
}

```

