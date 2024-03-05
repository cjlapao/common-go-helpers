package guard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	TestString string
}

func TestFatalEmptyOrNil(t *testing.T) {
	emptyStruc := TestStruct{}
	var nilStruc TestStruct
	var nilInterface interface{}
	emptyString := ""
	nonEmptyString := "foo"
	nonEmptyStruct := TestStruct{
		TestString: "bar",
	}

	assert.PanicsWithValuef(t, "Value guard.TestStruct cannot be nil", func() { FatalEmptyOrNil(emptyStruc) }, "Value should issue a panic with a specific message")
	assert.PanicsWithValuef(t, "Value guard.TestStruct cannot be nil", func() { FatalEmptyOrNil(nilStruc) }, "Value should issue a panic with a specific message")
	assert.PanicsWithValuef(t, "Value <nil> cannot be nil", func() { FatalEmptyOrNil(nilInterface) }, "Value should issue a panic with a specific message")
	assert.PanicsWithValuef(t, "Value string cannot be nil", func() { FatalEmptyOrNil(emptyString) }, "Value should issue a panic with a specific message")

	assert.PanicsWithValuef(t, "Value emptyStruc of type guard.TestStruct cannot be nil", func() { FatalEmptyOrNil(emptyStruc, "emptyStruc") }, "Value should issue a panic with a specific message")
	assert.PanicsWithValuef(t, "Value nilStruc of type guard.TestStruct cannot be nil", func() { FatalEmptyOrNil(nilStruc, "nilStruc") }, "Value should issue a panic with a specific message")
	assert.PanicsWithValuef(t, "Value nilInterface of type <nil> cannot be nil", func() { FatalEmptyOrNil(nilInterface, "nilInterface") }, "Value should issue a panic with a specific message")
	assert.PanicsWithValuef(t, "Value emptyString of type string cannot be nil", func() { FatalEmptyOrNil(emptyString, "emptyString") }, "Value should issue a panic with a specific message")

	assert.NotPanics(t, func() { FatalEmptyOrNil(nonEmptyString) })
	assert.NotPanics(t, func() { FatalEmptyOrNil(nonEmptyStruct) })
}

func TestEmptyOrNil(t *testing.T) {
	// Arrange
	emptyStruc := TestStruct{}
	var nilStruc TestStruct
	var nilInterface interface{}

	emptyString := ""
	nonEmptyString := "foo"
	nonEmptyStruct := TestStruct{
		TestString: "bar",
	}

	// Act + Assert
	assert.EqualErrorf(t, EmptyOrNil(emptyStruc), "Value guard.TestStruct cannot be nil", "Empty Struct should issue an error")
	assert.EqualErrorf(t, EmptyOrNil(nilStruc), "Value guard.TestStruct cannot be nil", "nil Struct should issue an error")
	assert.EqualErrorf(t, EmptyOrNil(nilInterface), "Value <nil> cannot be nil", "nil interface should issue an error")
	assert.EqualErrorf(t, EmptyOrNil(emptyString), "Value string cannot be nil", "Empty string should issue an error")

	assert.EqualErrorf(t, EmptyOrNil(emptyStruc, "emptyStruc"), "Value emptyStruc of type guard.TestStruct cannot be nil", "Empty Struct should issue an error with variable name")
	assert.EqualErrorf(t, EmptyOrNil(nilStruc, "nilStruc"), "Value nilStruc of type guard.TestStruct cannot be nil", "nil Struct should issue an error with variable name")
	assert.EqualErrorf(t, EmptyOrNil(nilInterface, "nilInterface"), "Value nilInterface of type <nil> cannot be nil", "nil interface should issue an error with variable name")
	assert.EqualErrorf(t, EmptyOrNil(emptyString, "emptyString"), "Value emptyString of type string cannot be nil", "Empty string should issue an error with variable name")

	assert.NoErrorf(t, EmptyOrNil(nonEmptyString), "Non Empty string should not issue an error")
	assert.NoErrorf(t, EmptyOrNil(nonEmptyStruct), "Non Empty struct should not issue an error")
}

func TestIsFalse(t *testing.T) {
	// Arrange
	trueValue := true
	falseValue := false

	// Act + Assert
	assert.EqualErrorf(t, IsFalse(falseValue), "Value bool cannot be false", "Empty Struct should issue an error")
	assert.EqualErrorf(t, IsFalse(falseValue, "falseValue"), "Value falseValue cannot be false", "Empty Struct should issue an error")
	assert.NoErrorf(t, IsFalse(trueValue), "Non Empty struct should not issue an error")
}

func TestEmptyOrNilWithMessage(t *testing.T) {
	emptyStruc := TestStruct{}
	nilStruc := TestStruct{}
	nilInterface := interface{}(nil)
	emptyString := ""
	nonEmptyString := "foo"
	nonEmptyStruct := TestStruct{
		TestString: "bar",
	}

	err := EmptyOrNilWithMessage(emptyStruc, "Value guard.TestStruct cannot be nil")
	assert.EqualError(t, err, "Value guard.TestStruct cannot be nil", "Empty struct should return an error")

	err = EmptyOrNilWithMessage(nilStruc, "Value guard.TestStruct cannot be nil")
	assert.EqualError(t, err, "Value guard.TestStruct cannot be nil", "Nil struct should return an error")

	err = EmptyOrNilWithMessage(nilInterface, "Value <nil> cannot be nil")
	assert.EqualError(t, err, "Value <nil> cannot be nil", "Nil interface should return an error")

	err = EmptyOrNilWithMessage(emptyString, "Value string cannot be nil")
	assert.EqualError(t, err, "Value string cannot be nil", "Empty string should return an error")

	err = EmptyOrNilWithMessage(emptyStruc, "Value emptyStruc of type guard.TestStruct cannot be nil")
	assert.EqualError(t, err, "Value emptyStruc of type guard.TestStruct cannot be nil", "Empty struct with variable name should return an error")

	err = EmptyOrNilWithMessage(nilStruc, "Value nilStruc of type guard.TestStruct cannot be nil")
	assert.EqualError(t, err, "Value nilStruc of type guard.TestStruct cannot be nil", "Nil struct with variable name should return an error")

	err = EmptyOrNilWithMessage(nilInterface, "Value nilInterface of type <nil> cannot be nil")
	assert.EqualError(t, err, "Value nilInterface of type <nil> cannot be nil", "Nil interface with variable name should return an error")

	err = EmptyOrNilWithMessage(emptyString, "Value emptyString of type string cannot be nil")
	assert.EqualError(t, err, "Value emptyString of type string cannot be nil", "Empty string with variable name should return an error")

	err = EmptyOrNilWithMessage(nonEmptyString, "Value nonEmptyString of type string cannot be nil")
	assert.NoError(t, err, "Non-empty string should not return an error")

	err = EmptyOrNilWithMessage(nonEmptyStruct, "Value nonEmptyStruct of type guard.TestStruct cannot be nil")
	assert.NoError(t, err, "Non-empty struct should not return an error")
}

func TestIsNil(t *testing.T) {
	// Test nil values
	assert.True(t, IsNil(nil), "nil value should be considered nil")

	// Test string values
	assert.False(t, IsNil(""), "empty string should not be considered nil")
	assert.False(t, IsNil("foo"), "non-empty string should not be considered nil")

	// Test bool values
	assert.False(t, IsNil(true), "true should not be considered nil")
	assert.False(t, IsNil(false), "false should not be considered nil")

	// Test integer values
	assert.False(t, IsNil(0), "0 should not be considered nil")
	assert.False(t, IsNil(42), "non-zero integer should not be considered nil")

	// Test float values
	assert.False(t, IsNil(0.0), "0.0 should not be considered nil")
	assert.False(t, IsNil(3.14), "non-zero float should not be considered nil")

	// Test uint values
	assert.False(t, IsNil(uint(0)), "0 should not be considered nil")
	assert.False(t, IsNil(uint(42)), "non-zero uint should not be considered nil")

	// Test struct values
	type TestStruct struct {
		Foo string
		Bar int
	}
	assert.False(t, IsNil(TestStruct{}), "empty struct should not be considered nil")
	assert.False(t, IsNil(TestStruct{Foo: "foo", Bar: 42}), "non-empty struct should not be considered nil")

	// Test other types
	var ptr *int
	assert.True(t, IsNil(ptr), "nil pointer should be considered nil")

	var iface interface{}
	assert.True(t, IsNil(iface), "nil interface should be considered nil")
}
