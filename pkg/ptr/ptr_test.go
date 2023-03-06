package ptr

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestToPtr(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		value := "some string"
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("int", func(t *testing.T) {
		value := int(10)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("int8", func(t *testing.T) {
		value := int8(127)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("int16", func(t *testing.T) {
		value := int16(32767)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("int32", func(t *testing.T) {
		value := int32(2147483647)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("int64", func(t *testing.T) {
		value := int64(9223372036854775807)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("float32", func(t *testing.T) {
		value := float32(9.1111111111)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

	t.Run("float64", func(t *testing.T) {
		value := float64(99.11111111112222222222222)
		assert.DeepEqual(t, ToPtr(value), &value)
	})

}

func TestToValue2(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		value := "some string"
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("int", func(t *testing.T) {
		value := int(10)
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("int8", func(t *testing.T) {
		value := int8(127)
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("int16", func(t *testing.T) {
		value := int16(32767)
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("int32", func(t *testing.T) {
		value := int32(2147483647)
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("int64", func(t *testing.T) {
		value := int64(9223372036854775807)
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("float32", func(t *testing.T) {
		value := float32(9.1111111111)
		assert.DeepEqual(t, ToValue(&value), value)
	})

	t.Run("float64", func(t *testing.T) {
		value := float64(99.11111111112222222222222)
		assert.DeepEqual(t, ToValue(&value), value)
	})

}
