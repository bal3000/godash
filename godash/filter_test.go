package godash

import (
	"testing"
)

func TestFilterInt(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		selectorFunc func(a int) bool
		expected     []int
	}{
		{
			name:  "select all numbers greater than 4",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			selectorFunc: func(a int) bool {
				return a > 4
			},
			expected: []int{5, 6, 7, 8, 9, 10},
		},
		{
			name:  "select all contained in slice",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			selectorFunc: func(a int) bool {
				wanted := []int{3, 7, 10}
				return Contains(wanted, a)

			},
			expected: []int{3, 7, 10},
		},
		{
			name:  "select nothing if nothing matches",
			input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			selectorFunc: func(a int) bool {
				return a > 10
			},
			expected: []int{},
		},
		{
			name:  "select nothing if empty",
			input: []int{},
			selectorFunc: func(a int) bool {
				return a > 1
			},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Filter(test.input, test.selectorFunc)
			if !IsEqual(actual, test.expected) {
				t.Errorf("expected %d, but got %d\n", test.expected, actual)
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	tests := []struct {
		name         string
		input        []string
		selectorFunc func(a string) bool
		expected     []string
	}{
		{
			name:  "select all strings with length greater than 4",
			input: []string{"hello", "there", "!", "bye"},
			selectorFunc: func(a string) bool {
				return len(a) > 4
			},
			expected: []string{"hello", "there"},
		},
		{
			name:  "select all with wanted string",
			input: []string{"hello", "there", "!", "bye"},
			selectorFunc: func(a string) bool {
				return a == "!"
			},
			expected: []string{"!"},
		},
		{
			name:  "select nothing if nothing matches",
			input: []string{"hello", "there", "!", "bye"},
			selectorFunc: func(a string) bool {
				return a == "test"
			},
			expected: []string{},
		},
		{
			name:  "select nothing if empty",
			input: []string{},
			selectorFunc: func(a string) bool {
				return a == "test"
			},
			expected: []string{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := Filter(test.input, test.selectorFunc)
			if !IsEqual(actual, test.expected) {
				t.Errorf("expected %s, but got %s\n", test.expected, actual)
			}
		})
	}
}

// func TestFilterStruct(t *testing.T) {
// 	type example struct {
// 		field1 string
// 		field2 bool
// 		field3 int
// 		field4 map[string]bool
// 		field5 struct {
// 			field6 string
// 		}
// 	}

// 	exampleData := []example{
// 		{
// 			field1: "test",
// 			field2: true,
// 			field3: 1,
// 			field4: map[string]bool{"hey": true},
// 			field5: struct{ field6 string }{field6: "there"},
// 		},
// 		{
// 			field1: "test2",
// 			field2: false,
// 			field3: 2,
// 			field4: map[string]bool{"hey": false},
// 			field5: struct{ field6 string }{field6: "there"},
// 		},
// 		{
// 			field1: "test3",
// 			field2: false,
// 			field3: 3,
// 			field4: map[string]bool{"hey": true},
// 			field5: struct{ field6 string }{field6: "there"},
// 		},
// 	}

// 	tests := []struct {
// 		name         string
// 		input        []example
// 		selectorFunc func(a example) bool
// 		expected     []example
// 	}{
// 		{
// 			name:  "select all that have field2 set to true",
// 			input: exampleData,
// 			selectorFunc: func(a example) bool {
// 				return a.field2
// 			},
// 			expected: []example{
// 				{
// 					field1: "test",
// 					field2: true,
// 					field3: 1,
// 					field4: map[string]bool{"hey": true},
// 					field5: struct{ field6 string }{field6: "there"},
// 				},
// 			},
// 		},
// 		{
// 			name:  "select all with field4 value is hey:false",
// 			input: exampleData,
// 			selectorFunc: func(a example) bool {
// 				val, ok := a.field4["hey"]
// 				return ok && !val
// 			},
// 			expected: []example{
// 				{
// 					field1: "test2",
// 					field2: false,
// 					field3: 2,
// 					field4: map[string]bool{"hey": false},
// 					field5: struct{ field6 string }{field6: "there"},
// 				},
// 			},
// 		},
// 		{
// 			name:  "select nothing if nothing matches",
// 			input: exampleData,
// 			selectorFunc: func(a example) bool {
// 				return a.field3 == 4
// 			},
// 			expected: []example{},
// 		},
// 		{
// 			name:  "select nothing if empty",
// 			input: []example{},
// 			selectorFunc: func(a example) bool {
// 				return a.field1 == "test"
// 			},
// 			expected: []example{},
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			actual := Filter(test.input, test.selectorFunc)
// 			if !IsEqual(actual, test.expected) {
// 				t.Errorf("expected %s, but got %s\n", test.expected, actual)
// 			}
// 		})
// 	}
// }
