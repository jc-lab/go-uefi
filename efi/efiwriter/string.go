// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io> and contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package efiwriter

import (
	"bytes"
	"encoding/binary"
	"io"
	"unicode/utf16"
)

func LengthASCIINullBytes(in []byte) int {
	nullIndex := bytes.IndexByte(in, 0)
	if nullIndex < 0 {
		return len(in) + 1
	} else {
		return nullIndex + 1
	}
}

func WriteASCIINullBytes(w io.Writer, in []byte) (n int, err error) {
	nullIndex := bytes.IndexByte(in, 0)
	if nullIndex < 0 {
		n, err = w.Write(in)
		if err == nil {
			var tmp int
			tmp, err = w.Write([]byte{0})
			n += tmp
		}
	} else {
		n, err = w.Write(in[0 : nullIndex+1])
	}
	return
}

func StringToASCIIZBytes(s string) []byte {
	return append([]byte(s), 0)
}

func LengthUTF16NullBytes(in []byte) int {
	nullIndex := indexUtf16NullTerminate(in)
	if nullIndex < 0 {
		return len(in) + 2
	} else {
		return nullIndex + 2
	}
}

func WriteUTF16NullBytes(w io.Writer, in []byte) (n int, err error) {
	nullIndex := indexUtf16NullTerminate(in)
	if nullIndex < 0 {
		n, err = w.Write(in)
		if err == nil {
			var tmp int
			tmp, err = w.Write([]byte{0, 0})
			n += tmp
		}
	} else {
		n, err = w.Write(in[0 : nullIndex+2])
	}
	return
}

func StringToUTF16Bytes(s string) []byte {
	unicode := utf16.Encode([]rune(s))
	output := make([]byte, len(unicode)*2)
	j := 0
	for _, u := range unicode {
		binary.LittleEndian.PutUint16(output[j:j+2], u)
		j += 2
	}
	return output
}

func StringToUTF16ZBytes(s string) []byte {
	b := StringToUTF16Bytes(s)
	b = append(b, 0, 0)
	return b
}

func indexUtf16NullTerminate(in []byte) int {
	var offset int
	var nullIndex = -1
	for offset < len(in) && (nullIndex&1) == 1 {
		nullIndex = bytes.Index(in[offset:], []byte{0, 0})
		if nullIndex < 0 {
			break
		} else {
			nullIndex += offset
		}
		offset = nullIndex + 1
	}
	return nullIndex
}
