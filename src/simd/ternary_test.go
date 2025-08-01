// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && amd64

package simd_test

import (
	"simd"
	"testing"
)

func TestFMA(t *testing.T) {
	if simd.HasAVX512() {
		testFloat32x4Ternary(t, simd.Float32x4.FusedMultiplyAdd, fmaSlice[float32])
		testFloat32x8Ternary(t, simd.Float32x8.FusedMultiplyAdd, fmaSlice[float32])
		testFloat32x16Ternary(t, simd.Float32x16.FusedMultiplyAdd, fmaSlice[float32])
		testFloat64x2Ternary(t, simd.Float64x2.FusedMultiplyAdd, fmaSlice[float64])
		testFloat64x4Ternary(t, simd.Float64x4.FusedMultiplyAdd, fmaSlice[float64])
		testFloat64x8Ternary(t, simd.Float64x8.FusedMultiplyAdd, fmaSlice[float64])
	}
}
