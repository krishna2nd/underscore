package underscore;

import(
	"reflect"
	"github.com/fatih/structs"
	"fmt"
)

type underscore struct {
}


type iteratee func(v interface{}, i interface{}, array interface{})

func arrEach(arg *interface{}, iteratee *iteratee) {
	switch reflect.TypeOf(*arg).Kind() {
	case reflect.Map:
		m := reflect.ValueOf(*arg)
		keys := m.MapKeys()
		for _, key := range keys {
			(*iteratee)(m.MapIndex(key), key.Interface(), arg)
		}
		break;
	case reflect.Slice:
	    arr := reflect.ValueOf(*arg)
	    for i := 0; i < arr.Len(); i++ {
	        (*iteratee)(arr.Index(i).Interface(), i, arg)
	    }
	    break;
	}
}
func structEach(arg *interface{}, iteratee *iteratee) {
	fmt.Println(*arg);
	fields := structs.Fields(*arg)

	for _, f := range fields {
	    (*iteratee)(f.Value(), f.Name(), arg)
	}
}

func Each(arg interface{}, iteratee iteratee) {
	if structs.IsStruct(arg) {
		structEach(&arg, &iteratee)
	} else {
		arrEach(&arg, &iteratee)
	}

}

func Equal(args ...interface{}) bool {
	if nil == args || len(args) < 2  {
		return false
	}

	for index := 1; index < len(args); index++ {
		if !reflect.DeepEqual(args[index -1 ], args[index]) {
			return false;
		}
	}

	return true;
}