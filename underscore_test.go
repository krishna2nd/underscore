package underscore_test

import (
	_ "fmt"
	. "github.com/krishna2nd/underscore"
	"testing"
)

func TestEach(t *testing.T) {
	type test struct {
		s    string
		rune rune
		out  int
	}
	type test2 struct {
		Title       string
		Description string
	}
	var (
		values_string, values_result = []string{"Title1", "Description1"}, []string{}
		values                       = test2{Title: values_string[0], Description: values_string[1]}
		_int_result, _int            = []int{}, []int{1, 2, 3, 4}
		_string_result, _string      = []string{}, []string{"test1", "test1", "test1", "test1"}
		_struct_result, _struct      = []test{}, []test{
			{"a A x", 'A', 2},
			{"some_text=some_value", '=', 9},
			{"☺a", 'a', 3},
			{"a☻☺b", '☺', 4},
		}
	)

	// Each test for integer values
	Each(_int, func(v interface{}, i interface{}, array interface{}) {
		_int_result = append(_int_result, v.(int))
	})

	if !Equal(_int, _int_result) {
		t.Errorf("Each iteration failed", _int, _int_result)
	}

	// Each test for string values
	Each(_string, func(v interface{}, i interface{}, array interface{}) {
		_string_result = append(_string_result, v.(string))
	})

	if !Equal(_string, _string_result) {
		t.Errorf("Each iteration failed", _string, _string_result)
	}

	// Each test for struct values
	Each(_struct, func(v interface{}, i interface{}, array interface{}) {
		_struct_result = append(_struct_result, v.(test))
	})

	if !Equal(_struct_result, _struct) {
		t.Errorf("Each iteration failed", _struct_result, _struct)
	}

	Each(values, func(v interface{}, i interface{}, array interface{}) {
		values_result = append(values_result, v.(string))
	})

	if !Equal(values_string, values_result) {
		t.Errorf("Each iteration failed", values_string, values_result)
	}
}

func BenchmarkEach(b *testing.B) {

}
