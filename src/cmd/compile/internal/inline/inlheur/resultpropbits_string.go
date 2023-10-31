// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by "stringer -bitset -type ResultPropBits"; DO NOT EDIT.

package inlheur

import (
	"bytes"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ResultNoInfo-0]
	_ = x[ResultIsAllocatedMem-2]
	_ = x[ResultIsConcreteTypeConvertedToInterface-4]
	_ = x[ResultAlwaysSameConstant-8]
	_ = x[ResultAlwaysSameFunc-16]
	_ = x[ResultAlwaysSameInlinableFunc-32]
}

var _ResultPropBits_value = [...]uint64{
	0x0,  /* ResultNoInfo */
	0x2,  /* ResultIsAllocatedMem */
	0x4,  /* ResultIsConcreteTypeConvertedToInterface */
	0x8,  /* ResultAlwaysSameConstant */
	0x10, /* ResultAlwaysSameFunc */
	0x20, /* ResultAlwaysSameInlinableFunc */
}

const _ResultPropBits_name = "ResultNoInfoResultIsAllocatedMemResultIsConcreteTypeConvertedToInterfaceResultAlwaysSameConstantResultAlwaysSameFuncResultAlwaysSameInlinableFunc"

var _ResultPropBits_index = [...]uint8{0, 12, 32, 72, 96, 116, 145}

func (i ResultPropBits) String() string {
	var b bytes.Buffer

	remain := uint64(i)
	seen := false

	for k, v := range _ResultPropBits_value {
		x := _ResultPropBits_name[_ResultPropBits_index[k]:_ResultPropBits_index[k+1]]
		if v == 0 {
			if i == 0 {
				b.WriteString(x)
				return b.String()
			}
			continue
		}
		if (v & remain) == v {
			remain &^= v
			x := _ResultPropBits_name[_ResultPropBits_index[k]:_ResultPropBits_index[k+1]]
			if seen {
				b.WriteString("|")
			}
			seen = true
			b.WriteString(x)
		}
	}
	if remain == 0 {
		return b.String()
	}
	return "ResultPropBits(0x" + strconv.FormatInt(int64(i), 16) + ")"
}