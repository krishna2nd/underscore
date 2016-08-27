# underscore.go

underscore.js implementation in golang

## Each, Equal

### Each

  Iterator functions can be applied on struct fields, arrays, maps

#### Eg: with `[]int`
```go
Each([]int{1, 2, 3, 4},
  func(v interface{}, i interface{}, array interface{}) {
	fmt.Println(v.(int))
})
```

#### Eg: with `[]string`
```go
Each([]string{"test1", "test1", "test1", "test1"},
  func(v interface{}, i interface{}, array interface{}) {
	fmt.Println(v.(string))
})
```

#### Eg: with `[]struct`

Iterator execute against all `struct` in `[]struct`
```go
Each([]struct {
  s    string
  rune rune
  out  int
}{
  {"a A x", 'A', 2},
  {"some_text=some_value", '=', 9},
  {"☺a", 'a', 3},
  {"a☻☺b", '☺', 4},
  }, func(v interface{}, i interface{}, array interface{}) {
		fmt.Println(v)
	})
```

#### Eg: with `struct`	

Iterator function will be executed against all fields.
```go
Each(struct {
  s    string
  rune rune
  out  int
}{"some_text=some_value", '=', 9},
func(v interface{}, i interface{}, array interface{}) {
	fmt.Println(v)
})
```

#### Eg: with `map`	

Iterator function will be executed against all fields.
```go
Each(map[string]string {
  "1": "One",
  "2" "Two"
},
func(v interface{}, i interface{}, array interface{}) {
	fmt.Println(v)
})
```

### Equal

Equality check for objects, arrays, any value type and reference type

eg 1: 
Equality check with two  arrays
```go
if !Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}) {
	t.Errorf("Arrays are not equal")
}
```

Equality check with unlimited inputs.
```go 

_int := []int{1, 2, 3, 4}

if !Equal(_int,_int,_int) {
	t.Errorf("Arrays are not equal",)
}


if !Equal(_int,_int,_int,_int ....) {
	t.Errorf("Arrays are not equal",)
}
```
Equality with strings

```go
if !Equal([]string{"test1", "test1", "test1", "test1"}, []string{"test1", "test1", "test1", "test1"}}) {
	t.Errorf("Arrays are not equal")
}

```

#### Full eg
```go
package main

import (
	_ "reflect"
	"fmt"
	. "github.com/krishna2nd/underscore"
)

func main() {
	type test struct {
		Foo string
		Bar int
	}
	x := test{"foo", 2}
	y := test{"foo", 2}
	z := test{"foo", 2}
	m := map[string]string { "1": "one", "2": "two", "3": "three" }

	Each(x, func(v interface{}, i interface{}, obj interface{}) {
		fmt.Println(v, i, obj)
	})

	Each([]test{x, y, z}, func(v interface{}, i interface{}, obj interface{}) {
		fmt.Println(v, i.(int), obj)
	})

	Each([]string{"1", "2", "3"}, func(v interface{}, i interface{}, obj interface{}) {
		fmt.Println(v, i.(int), obj)
	})

	Each(m, func(v interface{}, i interface{}, obj interface{}) {
		fmt.Println(v, i.(string), obj)
	})

	if Equal(x, y) {
		fmt.Println("equals")
	} else {
		fmt.Println("Not equals")
	}

	if Equal(x, y, z) {
		fmt.Println("equals")
	} else {
		fmt.Println("Not equals")
	}

}
```

