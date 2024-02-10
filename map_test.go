// Copyright 2022 The orderedmap Authors Shuochen Liu. All rights reserved.
// Use of this source code is governed by a BSD-style.
// License that can be found in the LICENSE file.

package orderedmap_test

import (
	"fmt"
	"github.com/liushuochen/orderedmap"
	"testing"
)

// Store more than one element, and verify map's length.
func TestStore(t *testing.T) {
	o := orderedmap.New()
	o.Store("Name", "Bob")
	o.Store("Age", 12)

	length := o.Length()
	if length != 2 {
		t.Errorf("expect the length of orderedmap is 2, but %d got", length)
	}
}

// Store same key-value pair more than one times, and verify map's length whether equal to 1.
func TestStoreWithDuplicateKey(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", 1)
	o.Store("A", 2)

	length := o.Length()
	if length != 1 {
		t.Errorf("expect the length of orderedmap is 1, but %d got", length)
	}
}

// Load a key from map.
func TestLoad(t *testing.T) {
	o := orderedmap.New()
	o.Store(1, "Mike")

	name, ok := o.Load(1)
	if !ok {
		t.Error("expect the result of bool flag is true, but false got")
	}

	nameValue, ok := name.(string)
	if !ok {
		t.Errorf("expect the type of name is a string type, but %T got", name)
	}

	if nameValue != "Mike" {
		t.Errorf("expect the value of result is Mike, but %s got", nameValue)
	}
}

// Load a nonexistent key from a map.
func TestLoadWithNonexistentKey(t *testing.T) {
	o := orderedmap.New()
	value, ok := o.Load(true)
	if ok {
		t.Error("expect the result of bool flag is false, but true got")
	}

	if value != nil {
		t.Errorf("expect the value of returned value is nil, but %v got", value)
	}
}

// Load a same key from a map more than one times.
func TestLoadASameKeyMoreThanOneTimes(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", "a")

	for i := 0; i < 3; i++ {
		_, _ = o.Load("A")
	}

	if o.Length() != 1 {
		t.Errorf("expected the length is 1, but %d got", o.Length())
	}

	value, ok := o.Load("A")
	if value.(string) != "a" {
		t.Errorf("expected the value of key \"A\" is \"a\", but \"%s\" got", value.(string))
	}

	if !ok {
		t.Error("expect the assert value is true, but false got")
	}
}

// Delete a key from a map.
func TestDelete(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", "a")

	o.Delete("A")
	_, ok := o.Load("A")
	if ok {
		t.Error("expected the key of \"A\" is not exist after deleting, but load success")
	}

	if o.Length() != 0 {
		t.Errorf("expected the length of map is 0, but %d got", o.Length())
	}
}

// Delete a nonexistent key from a map.
func TestDeleteNonexistentKey(t *testing.T) {
	o := orderedmap.New()
	o.Delete("A")
}

// Delete a key from a map more than one times.
func TestDeleteSameKeyMoreThanOneTimes(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", "a")

	for i := 0; i < 3; i++ {
		o.Delete("A")
	}
	_, ok := o.Load("A")
	if ok {
		t.Error("expected the key of \"A\" is not exist after deleting, but load success")
	}

	if o.Length() != 0 {
		t.Errorf("expected the length of map is 0, but %d got", o.Length())
	}
}

// Verify format string for an empty map.
func TestStringForEmptyMap(t *testing.T) {
	o := orderedmap.New()
	expectedString := "{}"
	if fmt.Sprint(o) != expectedString {
		t.Errorf("expected %s, but got %s", expectedString, o)
	}
}

// Verify format string with single key.
func TestStringForMapWithSingleKey(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", 1)
	expectedFormatString := "{\"A\": 1}"
	if fmt.Sprint(o) != expectedFormatString {
		t.Errorf("expected format string is %s, but %s got", expectedFormatString, fmt.Sprint(o))
	}
}

// Verify format string with multiple keys.
func TestStringForMapWithMultipleKeys(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", "a")
	o.Store("B", "b")
	expectedFormatString := "{\"A\": \"a\", \"B\": \"b\"}"
	if fmt.Sprint(o) != expectedFormatString {
		t.Errorf("expected format string is %s, but %s got", expectedFormatString, fmt.Sprint(o))
	}
}

// Test length for an empty map.
func TestEmptyMapLength(t *testing.T) {
	o := orderedmap.New()
	if o.Length() != 0 {
		t.Errorf("expected the length of empty map is 0, but %d got", o.Length())
	}
}

// Test `Range` method.
func TestRange(t *testing.T) {
	o := orderedmap.New()
	o.Store("A", "a")
	o.Store("B", "b")
	o.Store("C", 3)

	step := 0
	rangeMethod := func(key, value any) bool {
		step += 1
		return true
	}
	o.Range(rangeMethod)

	if step != o.Length() {
		t.Errorf("expected number of cycles is %d, but %d got", o.Length(), step)
	}
}
