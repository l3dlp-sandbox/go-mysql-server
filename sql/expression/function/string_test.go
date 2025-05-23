// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package function

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/types"
)

func TestAsciiFunc(t *testing.T) {
	f := sql.Function1{Name: "ascii", Fn: NewAscii}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding(uint8(0), "")
	tf.AddSucceeding(uint8(115), "string")
	tf.AddSucceeding(uint8(49), true)
	tf.AddSucceeding(uint8(50), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	tf.AddSignedVariations(uint8(48), 0)
	tf.AddUnsignedVariations(uint8(48), 0)
	tf.AddFloatVariations(uint8(54), 6.0)
	tf.Test(t, nil, nil)
}

func TestOrdFunc(t *testing.T) {
	f := sql.Function1{Name: "ord", Fn: NewOrd}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding(int64(0), "")
	tf.AddSucceeding(int64(115), "string")
	tf.AddSucceeding(int64(49826), "¢")
	tf.AddSucceeding(int64(49838), "®®")
	tf.AddSucceeding(int64(49), true)
	tf.AddSucceeding(int64(50), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	tf.AddSignedVariations(int64(48), 0)
	tf.AddUnsignedVariations(int64(48), 0)
	tf.AddFloatVariations(int64(54), 6.0)
	tf.Test(t, nil, nil)
}

func TestHexFunc(t *testing.T) {
	f := sql.Function1{Name: "hex", Fn: NewHex}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding("8F", []byte("\x8f"))
	tf.AddSignedVariations("FFFFFFFFFFFFFFFF", -1)
	tf.AddUnsignedVariations("5", 5)
	tf.AddFloatVariations("5", 4.5)
	tf.AddFloatVariations("5", 5.4)
	tf.AddSucceeding("FFFFFFFFFFFFFFFF", uint64(math.MaxUint64))
	tf.AddSucceeding("74657374", "test")
	tf.AddSignedVariations("FFFFFFFFFFFFFFF0", -16)
	tf.AddSignedVariations("FFFFFFFFFFFFFF00", -256)
	tf.AddSignedVariations("FFFFFFFFFFFFFE00", -512)
	tf.AddFloatVariations("FFFFFFFFFFFFFFFF", -0.5)
	tf.AddFloatVariations("FFFFFFFFFFFFFFFF", -1.4)
	tf.AddSucceeding("323032302D30322D30342031343A31303A33322E35", time.Date(2020, 2, 4, 14, 10, 32, 500000000, time.UTC))
	tf.AddSucceeding("323032302D30322D30342031343A31303A33322E30303035", time.Date(2020, 2, 4, 14, 10, 32, 500000, time.UTC))
	tf.AddSucceeding("323032302D30322D30342031343A31303A33322E303030303035", time.Date(2020, 2, 4, 14, 10, 32, 5000, time.UTC))
	tf.AddSucceeding("323032302D30322D30342031343A31303A3332", time.Date(2020, 2, 4, 14, 10, 32, 500, time.UTC))

	tf.Test(t, nil, nil)
}

func TestUnhexFunc(t *testing.T) {
	f := sql.Function1{Name: "unhex", Fn: NewUnhex}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding([]byte("MySQL"), "4D7953514C")
	tf.AddSucceeding([]byte{0x1, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}, "0123456789abcdef")
	tf.AddSucceeding([]byte{0x8f}, "8F")
	tf.AddSucceeding([]byte{0x8f}, "8f")
	tf.AddSucceeding([]byte{0x0b}, "B")
	tf.AddSucceeding(nil, "gh")
	tf.AddSignedVariations([]byte{0x35}, 35)
	tf.AddSignedVariations([]byte{0x01}, 1)
	tf.AddSignedVariations(nil, -1)
	tf.AddUnsignedVariations([]byte{0x35}, 35)
	tf.AddFloatVariations(nil, 35.5)
	tf.AddSucceeding(nil, time.Now())

	tf.Test(t, nil, nil)
}

func TestHexRoundTrip(t *testing.T) {
	tests := []struct {
		val interface{}
		typ sql.Type
		out string
	}{
		{"1B", types.Text, "1B"},
		{"C", types.Text, "0C"},
		{"8F", types.Text, "8F"},
		{"ABCD", types.Text, "ABCD"},
		{int64(1), types.Int64, "01"},
		{int8(11), types.Int64, "11"},
		{uint16(375), types.Int64, "0375"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v %s", test.val, test.typ.String()), func(t *testing.T) {
			lit := expression.NewLiteral(test.val, test.typ)
			f := NewHex(NewUnhex(lit))
			res, err := f.Eval(sql.NewEmptyContext(), nil)
			require.NoError(t, err)
			require.Equal(t, test.out, res)
		})
	}
}

func TestBinFunc(t *testing.T) {
	f := sql.Function1{Name: "bin", Fn: NewBin}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding("1100", "12")
	tf.AddSucceeding("0", "TEST")
	tf.AddSucceeding("11111100100", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	tf.AddSignedVariations("1100", 12)
	tf.AddUnsignedVariations("1100", 12)
	tf.AddFloatVariations("1100", 12.5)
	tf.AddSignedVariations("1111111111111111111111111111111111111111111111111111111111110100", -12)
	tf.AddFloatVariations("1111111111111111111111111111111111111111111111111111111111110100", -12.5)
	tf.Test(t, nil, nil)
}

func TestBitLength(t *testing.T) {
	f := sql.Function1{Name: "bin", Fn: NewBitlength}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding(32, "test")
	tf.AddSucceeding(8, true)
	tf.AddSucceeding(8, int8(0))
	tf.AddSucceeding(8, uint8(0))
	tf.AddSucceeding(16, int16(0))
	tf.AddSucceeding(16, uint16(0))
	tf.AddSucceeding(32, uint32(0))
	tf.AddSucceeding(32, int32(0))
	tf.AddSucceeding(32, uint(0))
	tf.AddSucceeding(32, 0)
	tf.AddSucceeding(64, uint64(0))
	tf.AddSucceeding(64, int64(0))
	tf.AddSucceeding(64, float64(0))
	tf.AddSucceeding(32, float32(0))
	tf.AddSucceeding(128, time.Now())
	tf.Test(t, nil, nil)
}

func TestQuote(t *testing.T) {
	f := sql.Function1{Name: "quote", Fn: NewQuote}
	tf := NewTestFactory(f.Fn)
	tf.AddSucceeding(nil, nil)
	tf.AddSucceeding("'test'", "test")
	tf.AddSucceeding("'0'", false)
	tf.AddSucceeding("'1'", true)
	tf.AddSucceeding("'12345'", 12345)
	tf.AddSucceeding("'\\\\, \\', \\0, \\\032'", "\\, ', \000, \032")
	tf.Test(t, nil, nil)
}
