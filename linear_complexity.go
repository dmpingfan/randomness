// Copyright (c) 2021 Quan guanyu
// randomness is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package randomness

import (
	"fmt"
	"math"
)

// LinearComplexityTest 线型复杂度检测
func LinearComplexityTest(bits []bool) float64 {
	n := len(bits)

	if n == 0 {
		fmt.Println("LinearComplexityTest:args wrong")
		return -1
	}
	m := 500

	N := n / m

	var v = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
	var pi = []float64{0.010417, 0.03125, 0.12500, 0.5000, 0.25000, 0.06250, 0.020833}
	var V float64 = 0.0
	var P float64 = 0

	var arr []bool
	var complexity int
	var T, miu float64
	arr = make([]bool, m)
	var _1_m float64
	if m%2 == 0 {
		_1_m = 1.0
	} else {
		_1_m = -1.0
	}
	miu = float64(m)/2.0 + (9.0+_1_m)/36.0 - (float64(m)/3.0+2.0/9.0)/math.Pow(2.0, float64(m))

	for i := 0; i < N; i++ {
		for j := 0; j < m; j++ {
			arr[j], bits = bits[0], bits[1:]
		}
		complexity = linearComplexity(arr, m)
		if m%2 == 0 {
			_1_m = 1.0
		} else {
			_1_m = -1.0
		}
		T = _1_m*(float64(complexity)-miu) + 2.0/9.0
		if T <= -2.5 {
			v[0]++
		} else if T <= -1.5 {
			v[1]++
		} else if T <= -0.5 {
			v[2]++
		} else if T <= 0.5 {
			v[3]++
		} else if T <= 1.5 {
			v[4]++
		} else if T <= 2.5 {
			v[5]++
		} else {
			v[6]++
		}
	}

	for i := 0; i < 7; i++ {
		V += math.Pow(v[i]-float64(N)*pi[i], 2.0) / (float64(N) * pi[i])
	}
	P = igamc(3.0, V/2.0)

	return P
}