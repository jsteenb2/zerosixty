package zerosixty_test

import "testing"

func TestVariableDeclarationsToZeroValues(t *testing.T) {
	// All variables are automatically assigned to their zero value (not default value)
	// creates the memory allocation for you
	var boolean bool

	var aString string

	var (
		anInt   int
		anInt8  int8
		anInt32 int32
		anInt64 int64
		uInt    uint64
	)

	var (
		aFloat32 float32
		aFloat64 float64
	)

	t.Logf("anInt: type=%T, value=%v", anInt, anInt)
	t.Logf("anInt8: type=%T, value=%v", anInt8, anInt8)
	t.Logf("anInt32: type=%T, value=%v", anInt32, anInt32)
	t.Logf("anInt64: type=%T, value=%v", anInt64, anInt64)
	t.Logf("uInt: type=%T, value=%v", uInt, uInt)
	t.Logf("aFloat32: type=%T, value=%v", aFloat32, aFloat32)
	t.Logf("aFloat64: type=%T, value=%v", aFloat64, aFloat64)
	t.Logf("boolean: type=%T, value=%v", boolean, boolean)
	t.Logf("aString: type=%T, value=%v", aString, aString)
}

func TestVariableDelcarationAssignmentToNonZeroValue(t *testing.T) {
	var maxInt uint64 = 1<<64 - 1
	var falsey bool = false

	t.Logf("maxInt: type=%T, value=%v", maxInt, maxInt)
	t.Logf("falsey: type=%T, value=%v", falsey, falsey)
}

func TestVariableDeclarationShorthand(t *testing.T) {
	anInt := 2017  // int type default
	aFloat := 32.0 // float64 type default
	boolean := true

	t.Logf("anInt: type=%T, value=%v", anInt, anInt)
	t.Logf("aFloat: type=%T, value=%v", aFloat, aFloat)
	t.Logf("boolean: type=%T, value=%v", boolean, boolean)
}

func TestVariableConversions(t *testing.T) {
	// there are no type casts in Go, only type conversions
	// conversion creates all new memory allocations
	var anInt8 int8 = 2
	var aFloat32 float32 = 32.00

	var anInt16 int16 = 2
	var aFloat64 float64 = 32.00

	//if anInt8 == anInt16 { // doesn't compile, cannot do operations on different types
	if convertedInt8 := int16(anInt8); convertedInt8 == anInt16 {
		// convertedInt8 only exists in this if blocks's scope
		t.Logf("convertedInt8: type=%T, value=%v", convertedInt8, convertedInt8)
	}

	if convertedFloat64 := float64(aFloat32); convertedFloat64 == aFloat64 {
		t.Logf("convertedFloat64: type=%T, value=%v", convertedFloat64, convertedFloat64)
	}
}
