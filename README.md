# qhash

[![GitHub license](https://img.shields.io/github/license/EldersJavas/qhash?style=flat-square)](https://github.com/EldersJavas/qhash/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/EldersJavas/qhash?style=flat-square)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/EldersJavas/qhash?style=flat-square)](https://github.com/EldersJavas/qhash)
[![GoDoc](https://godoc.org/github.com/EldersJavas/qhash?status.svg)](https://pkg.go.dev/github.com/EldersJavas/qhash)
[![Go Report Card](https://goreportcard.com/badge/github.com/EldersJavas/qhash)](https://goreportcard.com/report/github.com/EldersJavas/qhash)


get hash quickly.

---
# package
## Get
```bash
go get github.com/EldersJavas/qhash
```
## Usage
```go
package main
import (
	"fmt"
	"github.com/EldersJavas/qhash"
)
func main()  {
	fmt.Println(qhash.MD5([]byte("Hello world!")))
}
```

# cmd
## Install

```bash
go install github.com/EldersJavas/qhash/cmd@latest
```


## Usage

```bash
Quick Hash
https://github.com/EldersJavas/qhash

Usage:                                                                                                  
	qhash (-h) [crypto] [type]                       

Options:
	[crypto]:
		-MD4
		-MD5
		-SHA1      
		-SHA224      
		-SHA256      
		-SHA384      
		-SHA512      
		-MD5SHA1                     
		-RIPEMD160                   
		-SHA3_224               
		-SHA3_256               
		-SHA3_384               
		-SHA3_512               
		-SHA512_224  
		-SHA512_256  
		-BLAKE2s_256               
		-BLAKE2b_256               
		-BLAKE2b_384               
		-BLAKE2b_512               
	[type]:
		-f file1             Get file
		--s="string"         Get string
		-fs file1 file2...   Get files

Example:
	qhash -MD5 -f hello.go
	qhash -MD5 --s="Hello"

```


