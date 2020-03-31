/*
Copyright (c) 2020 Brian M. Kessler

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package streamvbyte

import (
	"testing"
)

func TestRoundTripUint32Scalar(t *testing.T) {
	testUniformAndRandomUint32(t, encodeUint32scalar, decodeUint32scalar)
}

func BenchmarkEncodeUint32Scalar(b *testing.B) {
	b.SetBytes(int64(4 * benchSize))
	for i := 0; i < b.N; i++ {
		benchEncodedSize = encodeUint32scalar(benchEncoded, benchUint32Data)
	}
}

func BenchmarkDecodeUint32Scalar(b *testing.B) {
	b.SetBytes(int64(4 * benchSize))
	benchEncodedSize = encodeUint32scalar(benchEncoded, benchUint32Data)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decodeUint32scalar(benchUint32Data, benchEncoded)
	}
}
