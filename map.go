// Copyright 2022 The orderedmap Authors Shuochen Liu. All rights reserved.
// Use of this source code is governed by a BSD-style.
// License that can be found in the LICENSE file.

// Package orderedmap defines functions and methods for OrderedMap.
package orderedmap

import (
	"fmt"
	"strings"
)

const (
	// Version defined package `orderedmap` version information.
	Version = "0.1.3"
)

// OrderedMap is like a Go map[interface{}]interface{} but is has order for each key add or delete.
// It is not concurrent security.
type OrderedMap struct {
	dirty map[interface{}]interface{}
	list  []interface{}
}

// New function returns a pointer for OrderedMap.
func New() *OrderedMap {
	return &OrderedMap{
		dirty: make(map[interface{}]interface{}),
		list:  make([]interface{}, 0),
	}
}

// Store sets the value for a key.
func (om *OrderedMap) Store(key, value interface{}) {
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
func (om *OrderedMap) Load(key interface{}) (interface{}, bool) {
	value, ok := om.dirty[key]
	return value, ok
}

// Delete deletes the value for a key.
func (om *OrderedMap) Delete(key interface{}) {
	newList := make([]interface{}, 0)
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
	for _, key := range om.list {
		results = append(results, fmt.Sprintf("%v: %v", key, om.dirty[key]))
	}

	return "{" + strings.Join(results, " ") + "}"
}

// Length method return an integer number of OrderedMap's length.
func (om *OrderedMap) Length() int {
	return len(om.dirty)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// No key will be visited more than once. Unless the same key is deleted first and then added in iteration.
func (om *OrderedMap) Range(f func(key, value interface{}) bool) {
	for _, key := range om.list {
		if !f(key, om.dirty[key]) {
			break
		}
	}
}
