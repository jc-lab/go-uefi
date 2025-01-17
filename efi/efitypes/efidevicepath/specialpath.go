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

package efidevicepath

import (
	"fmt"
	"github.com/jc-lab/go-uefi/efi/efireader"
	"github.com/jc-lab/go-uefi/efi/efiwriter"
	"io"
	"strings"

	"github.com/jc-lab/go-uefi/efi/efihex"
)

// EndOfPath terminates a Device Path.
type EndOfPath struct{ Head }

func (p *EndOfPath) ReadFrom(r io.Reader) (n int64, err error) { return 0, nil }
func (p *EndOfPath) WriteTo(w io.Writer) (n int64, err error)  { return 0, nil }
func (p *EndOfPath) GetHead() *Head                            { return &p.Head }
func (p *EndOfPath) UpdateHead() *Head {
	p.Head.Type = EndOfPathType
	p.Head.SubType = EndEntireSubType
	p.Head.Length = 4
	return &p.Head
}
func (p *EndOfPath) Text() string { return "" }

const (
	_ DevicePathSubType = iota

	// EndSingleSubType terminates one Device Path instance and
	// denotes the start of another.
	EndSingleSubType

	// EndEntireSubType terminates an entire Device Path.
	EndEntireSubType DevicePathSubType = 0xff
)

// UnrecognizedDevicePath represents a Device Path that is unimplemented.
type UnrecognizedDevicePath struct {
	Head

	Data []byte
}

func (p *UnrecognizedDevicePath) ReadFrom(r io.Reader) (n int64, err error) {
	p.Data, err = io.ReadAll(r)
	n = int64(len(p.Data))
	return
}

func (p *UnrecognizedDevicePath) WriteTo(w io.Writer) (n int64, err error) {
	return efiwriter.WriteFields(w, p.Data)
}

func (p *UnrecognizedDevicePath) GetHead() *Head {
	return &p.Head
}

func (p *UnrecognizedDevicePath) UpdateHead() *Head {
	p.Length = uint16(4 + len(p.Data))
	return &p.Head
}

func (p *UnrecognizedDevicePath) Text() string {
	if p.Head.Is(MessagingType, 11) {
		var ifType byte
		if len(p.Data) >= (36 - 4) {
			ifType = p.Data[36-4]
		}
		return fmt.Sprintf(
			"MAC(%s,%v)",
			strings.Trim(efihex.EncodeToString(p.Data), "0"),
			ifType,
		)
	}

	if p.Head.Is(MessagingType, 24) {
		return fmt.Sprintf(
			"Uri(%s)",
			efireader.ASCIIZBytesToString(p.Data),
		)
	}

	return fmt.Sprintf(
		"Path(%d,%d,%s)",
		p.Head.Type,
		p.Head.SubType,
		efihex.EncodeToString(p.Data),
	)
}

func ParseUnrecognizedDevicePath(r io.Reader, h Head) (p DevicePath, err error) {
	p = &UnrecognizedDevicePath{Head: h}
	_, err = p.ReadFrom(r)
	return
}
