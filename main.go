package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)


func main() {
	Arg := os.Args
	var Hs hash.Hash
	var UseByte []byte
	var HsResult string
	log.Printf("Arg: %v\n", Arg)
	switch Arg[2] {
	case "-f":
		log.Println(Arg[2:])
	case "fs":
		log.Fatal("Wait next version")
	default:
		if string(Arg[2][2]) != "s" {log.Fatal("Unknown Commond")}
		str := strings.TrimRight(strings.TrimPrefix(Arg[2], "--s\""), "\"")
		UseByte = []byte(str)
	}
	switch Arg[1] {
	case "-h":
		PrintHelp(Arg)
	case "-MD4":
		Hs = md4.New()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-MD5":
		Hs = md5.New()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA1":
		Hs = sha1.New()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA224":
		Hs = sha256.New224()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA256":
		Hs = sha256.New()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA384":
		Hs = sha512.New384()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA512":
		Hs = sha512.New()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-MD5SHA1":
		Hs1 := md4.New()
		Hs1.Write(UseByte)
		Hs2 := sha1.New()
		Hs2.Write(UseByte)
		HsResult = fmt.Sprint(Hs1.Sum(nil), "+", Hs2.Sum(nil))
	case "-RIPEMD160":
		Hs = ripemd160.New()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA3_224":
		Hs = sha3.New224()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA3_256":
		Hs = sha3.New256()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA3_384":
		Hs = sha3.New384()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA3_512":
		Hs = sha3.New512()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA512_224":
		Hs = sha512.New512_224()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-SHA512_256":
		Hs = sha512.New512_256()
		Hs.Write(UseByte)
		HsResult = string(Hs.Sum(nil))
	case "-BLAKE2s_256":
		blake := blake2s.Sum256(UseByte)
		HsResult = fmt.Sprint(blake)
	case "-BLAKE2b_256":
		blake := blake2b.Sum256(UseByte)
		HsResult = fmt.Sprint(blake)
	case "-BLAKE2b_384":
		blake := blake2b.Sum384(UseByte)
		HsResult = fmt.Sprint(blake)
	case "-BLAKE2b_512":
		blake := blake2b.Sum512(UseByte)
		HsResult = fmt.Sprint(blake)
	default:
		PrintHelp(Arg)
	}

	fmt.Print(HsResult)

}

/*
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b
*/

func PrintHelp(Arg []string) {
	fmt.Printf(`
Quick Hash
https://github.com/EldersJavas/qhash

Usage:                                                                                                  
	%s (-h) [crypto] [type]                       

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
	%s -MD5 -f hello.go
	%s -MD5 --s="Hello"


	`, Arg[0], Arg[0], Arg[0])
}
