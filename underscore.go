package underscore;

import(
	"reflect"
)

type underscore struct {
}


type iteratee func(v interface{})
func Each(array interface{}, iteratee iteratee) {
	switch reflect.TypeOf(array).Kind() {
		case reflect.Slice:
		    arr := reflect.ValueOf(array)
		    for i := 0; i < arr.Len(); i++ {
		        iteratee(arr.Index(i))
		    }
    }
}