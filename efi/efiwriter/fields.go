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
	"encoding/binary"
	"fmt"
	"io"
)

type FieldWriter struct {
	writer io.Writer
	offset *int64
}

func (w *FieldWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	*w.offset += int64(n)
	return
}

func (w *FieldWriter) Offset() int64 {
	return *w.offset
}

func (w *FieldWriter) WriteFields(fields ...any) (err error) {
	for i, d := range fields {
		if err = binary.Write(w, binary.LittleEndian, d); err != nil {
			err = fmt.Errorf("field #%d: %w", i, err)
			return
		}
	}
	return
}

func NewFieldWriter(writer io.Writer, offset *int64) *FieldWriter {
	if offset == nil {
		offset = new(int64)
	}
	return &FieldWriter{writer: writer, offset: offset}
}

func WriteFields(w io.Writer, fields ...any) (n int64, err error) {
	err = NewFieldWriter(w, &n).WriteFields(fields...)
	return
}
