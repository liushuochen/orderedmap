// Copyright 2022 The orderedmap Authors Shuochen Liu. All rights reserved.
// Use of this source code is governed by a BSD-style.
// License that can be found in the LICENSE file.

package orderedmap_test

import (
	"github.com/liushuochen/orderedmap"
	"testing"
)

func TestStore(t *testing.T) {
	o := orderedmap.New()
	o.Store("Name", "Bob")
	o.Store("Age", 12)

	length := o.Length()
	if length != 2 {
		t.Errorf("expect the length of orderedmap is 2, but %d got", length)
	}
}
