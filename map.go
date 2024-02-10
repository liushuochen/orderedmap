// Copyright 2022 The orderedmap Authors Shuochen Liu. All rights reserved.
// Use of this source code is governed by a BSD-style.
// License that can be found in the LICENSE file.

// Package orderedmap defines functions and methods for OrderedMap.
package orderedmap

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// Version defined package `orderedmap` version information.
	Version = "0.1.4"
)

// OrderedMap is like a Go map[any]any but is having order for each key add or delete.
// It is not concurrent security.
type OrderedMap struct {
	dirty map[any]any
	list  []any
}

// New function returns a pointer for OrderedMap.
func New() *OrderedMap {
	return &OrderedMap{
		dirty: make(map[any]any),
		list:  make([]any, 0),
	}
}

// Store sets the value for a key.
// If the value of the key is an unhashable value, a panic will be raised.
func (om *OrderedMap) Store(key, value any) {
	exist := false
	for _, k := range om.list {
		if k != key {
			continue
		}

		exist = true
		break
	}

	if !exist {
		om.list = append(om.list, key)
	}

	om.dirty[key] = value
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (om *OrderedMap) Load(key any) (any, bool) {
	value, ok := om.dirty[key]
	return value, ok
}

// Delete deletes the value for a key.
func (om *OrderedMap) Delete(key any) {
	newList := make([]any, 0)
	for _, oldKey := range om.list {
		if oldKey == key {
			continue
		}

		newList = append(newList, oldKey)
	}
	om.list = newList
	delete(om.dirty, key)
}

// String method used to realize "fmt.Stringer".
func (om *OrderedMap) String() string {
	results := make([]string, 0)
	length := om.Length()
	for index, key := range om.list {
		result := om.keyValueString(key)
		if index != length-1 {
			result = fmt.Sprintf("%s,", result)
		}
		results = append(results, result)
	}

	return "{" + strings.Join(results, " ") + "}"
}

// GoString method implements `fmt.GoStringer`.
func (om *OrderedMap) GoString() string {
	return om.String()
}

func (om *OrderedMap) keyValueString(key any) string {
	value, _ := om.Load(key)

	result := ""
	// TODO: Currently, only support string.
	switch reflect.TypeOf(key).Kind() {
	case reflect.String:
		result = fmt.Sprintf("\"%s\":", key)
	default:
		result = fmt.Sprintf("%v:", key)
	}

	// TODO: Currently, only support string.
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		result = fmt.Sprintf("%s \"%s\"", result, value)
	default:
		result = fmt.Sprintf("%s %v", result, value)
	}
	return result
}

// Length method return an integer number of OrderedMap's length.
func (om *OrderedMap) Length() int {
	return len(om.dirty)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// No key will be visited more than once. Unless the same key is deleted first and then added in iteration.
func (om *OrderedMap) Range(f func(key, value any) bool) {
	for _, key := range om.list {
		if !f(key, om.dirty[key]) {
			break
		}
	}
}
