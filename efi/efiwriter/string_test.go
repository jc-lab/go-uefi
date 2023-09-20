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
	"reflect"
	"testing"
)

func TestWriteUTF16NullString(t *testing.T) {
	tests := []struct {
		name    string
		inp     []byte
		wantOut []byte
		wantErr bool
	}{
		{
			name:    "golden path",
			inp:     []byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00, 0x00, 0x00},
			wantOut: []byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name:    "extra after string",
			inp:     []byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00, 0x00, 0x00, 0x01, 0x02},
			wantOut: []byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00, 0x00, 0x00},
			wantErr: false,
		},
		{
			name:    "unterminated",
			inp:     []byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00},
			wantOut: []byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00, 0x00, 0x00},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bytes.Buffer
			_, err := WriteUTF16NullBytes(&got, tt.inp)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteUTF16NullBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Bytes(), tt.wantOut) {
				t.Errorf("WriteUTF16NullBytes() gotOut = %v, want %v", got.Bytes(), tt.wantOut)
			}
		})
	}
}

func TestStringToUTF16NullBytes(t *testing.T) {
	tests := []struct {
		name string
		want []byte
		inp  string
	}{
		{
			"golden path",
			[]byte{0x61, 0x00, 0x73, 0x00, 0x64, 0x00, 0x00, 0x00},
			"asd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUTF16ZBytes(tt.inp); !bytes.Equal(got, tt.want) {
				t.Errorf("StringToUTF16ZBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToUTF16Bytes(t *testing.T) {
	tests := []struct {
		name     string
		expected []byte
		provided string
	}{
		{"", []byte{0x74, 0x00, 0x65, 0x00, 0x73, 0x00, 0x74, 0x00}, "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToUTF16Bytes(tt.provided); !bytes.Equal(got, tt.expected) {
				t.Errorf("StringToUTF16Bytes() = %v, want %v", got, tt.expected)
			}
		})
	}
}
