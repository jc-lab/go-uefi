# uefi

[![License: APACHE-2.0](https://img.shields.io/badge/license-APACHE--2.0-blue?style=flat-square)](https://www.apache.org/licenses/)

A UEFI library written in go to interact with efivars. Compatible with Windows and Linux.

This library tries its best to follow the UEFI 2.9 specification outlined [here](https://uefi.org/sites/default/files/resources/UEFI_Spec_2_9_2021_03_18.pdf).


## 📦 Installation

```console
$ go get -u github.com/jc-lab/go-uefi@latest
```


## 🤔 Usage

```go
package main

import (
	"fmt"

	"github.com/jc-lab/go-uefi/efi/efivario"
	"github.com/jc-lab/go-uefi/efi/efivars"
)

func main() {
	c := efivario.NewDefaultContext()

	if err := efivars.BootNext.Set(c, 1); err != nil {
		fmt.Println(err)
	}
}
```

For a more in-depth example of how to use this library take a look at [efibootcfg](https://github.com/0x5a17ed/efibootcfg).


## 💡 Features
- Works on both Linux and on Windows exposing the same API
- Extensible
- Simple API
- Reading individual Boot options
- Setting next Boot option
- Managing Boot order
