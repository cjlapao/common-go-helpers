package reflect

import (
	"reflect"
	"testing"

	"github.com/cjlapao/common-go-helpers/test"
	"github.com/stretchr/testify/assert"
)

type RemoveFieldStruct struct {
	ID          string
	Name        string
	IsOld       bool
	Age         int
	AndSomeMore interface{}
	OneMap      map[string]interface{}
	Sub         RemoveFieldSub
}

type RemoveFieldSub struct {
	Hello string
}

func TestIsNilOrEmpty(t *testing.T) {
	// Arrange
	emptyString := ""
	nonEmptyString := "foo"
	zeroVal := 0
	falseVal := false
	int64Val := int64(0)
	floatVal := float64(0)
	emptyStructValue := test.TestStructure{}
	var nilStructValue test.TestStructure
	var nilInterfaceValue interface{}
	nonEmptyStruct := test.TestStructure{
		TestString: "bar",
	}

	// Act + Assert
	assert.True(t, IsNilOrEmpty(nilInterfaceValue))
	assert.True(t, IsNilOrEmpty(emptyString))
	assert.False(t, IsNilOrEmpty(nonEmptyString))
	assert.False(t, IsNilOrEmpty(zeroVal))
	assert.False(t, IsNilOrEmpty(falseVal))
	assert.False(t, IsNilOrEmpty(int64Val))
	assert.False(t, IsNilOrEmpty(floatVal))
	assert.True(t, IsNilOrEmpty(emptyStructValue))
	assert.True(t, IsNilOrEmpty(nilStructValue))
	assert.False(t, IsNilOrEmpty(nonEmptyStruct))
}

func TestRemoveField(t *testing.T) {
	// Arrange
	test := RemoveFieldStruct{
		ID:          "ID",
		Name:        "SomeName",
		IsOld:       true,
		Age:         20,
		AndSomeMore: "yep",
		OneMap:      make(map[string]interface{}),
		Sub: RemoveFieldSub{
			Hello: "world",
		},
	}
	test.OneMap["testing"] = "something"

	result := RemoveField(test, "ID")

	assert.Equal(t, result["Name"], "SomeName")
}

func TestGetFieldTag(t *testing.T) {
	type testStruct struct {
		Field1 string `json:"field1,omitempty"`
		Field2 int    `json:"field2,omitempty"`
		Field3 bool   `json:"field3,omitempty"`
	}

	tests := []struct {
		name     string
		field    reflect.StructField
		key      string
		expected *FieldTag
	}{
		{
			name:     "Empty tag",
			field:    reflect.StructField{},
			key:      "json",
			expected: nil,
		},
		{
			name: "Tag with matching key",
			field: reflect.StructField{
				Tag: reflect.StructTag(`json:"field1,omitempty"`),
			},
			key: "json",
			expected: &FieldTag{
				Type:    "json",
				Name:    "field1",
				Options: []string{"omitempty"},
			},
		},
		{
			name: "Tag with non-matching key",
			field: reflect.StructField{
				Tag: reflect.StructTag(`json:"field1,omitempty"`),
			},
			key:      "xml",
			expected: nil,
		},
		{
			name: "Tag with non-matching key",
			field: reflect.StructField{
				Tag: reflect.StructTag(`json:"field1"`),
			},
			key: "json",
			expected: &FieldTag{
				Type:    "json",
				Name:    "field1",
				Options: []string{},
			},
		},
		{
			name: "Tag with multiple tags",
			field: reflect.StructField{
				Tag: reflect.StructTag(`json:"field1,omitempty" xml:"field2"`),
			},
			key: "xml",
			expected: &FieldTag{
				Type:    "xml",
				Name:    "field2",
				Options: []string{},
			},
		},
		{
			name: "Tag with multiple values",
			field: reflect.StructField{
				Tag: reflect.StructTag(`json:"field1,omitempty" xml:"field2,omitempty"`),
			},
			key: "xml",
			expected: &FieldTag{
				Type:    "xml",
				Name:    "field2",
				Options: []string{"omitempty"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetFieldTag(tt.field, tt.key)
			assert.Equal(t, tt.expected, result)
		})
	}
}
