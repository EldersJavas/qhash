package qhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"hash"
)

var Hs hash.Hash

func MD4(UseByte []byte) (HsResult string) {
	Hs = md4.New()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func MD5(UseByte []byte) (HsResult string) {
	Hs = md5.New()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA1(UseByte []byte) (HsResult string) {
	Hs = sha1.New()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA224(UseByte []byte) (HsResult string) {
	Hs = sha256.New224()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA256(UseByte []byte) (HsResult string) {
	Hs = sha256.New()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA384(UseByte []byte) (HsResult string) {
	Hs = sha512.New384()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA512(UseByte []byte) (HsResult string) {
	Hs = sha512.New()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func MD5SHA1(UseByte []byte) (HsResult string) {
	Hs1 := md4.New()
	Hs1.Write(UseByte)
	Hs2 := sha1.New()
	Hs2.Write(UseByte)
	HsResult = fmt.Sprintf("%x+%x", Hs1.Sum(nil), Hs2.Sum(nil))
	return
}
func RIPEMD160(UseByte []byte) (HsResult string) {
	Hs = ripemd160.New()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA3_224(UseByte []byte) (HsResult string) {
	Hs = sha3.New224()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA3_256(UseByte []byte) (HsResult string) {
	Hs = sha3.New256()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA3_384(UseByte []byte) (HsResult string) {
	Hs = sha3.New384()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA3_512(UseByte []byte) (HsResult string) {
	Hs = sha3.New512()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA512_224(UseByte []byte) (HsResult string) {
	Hs = sha512.New512_224()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func SHA512_256(UseByte []byte) (HsResult string) {
	Hs = sha512.New512_256()
	Hs.Write(UseByte)
	HsResult = fmt.Sprintf("%x", Hs.Sum(nil))
	return
}
func BLAKE2s_256(UseByte []byte) (HsResult string) {
	blake := blake2s.Sum256(UseByte)
	HsResult = fmt.Sprintf("%x", blake)
	return
}
func BLAKE2b_256(UseByte []byte) (HsResult string) {
	blake := blake2b.Sum256(UseByte)
	HsResult = fmt.Sprintf("%x", blake)
	return
}
func BLAKE2b_384(UseByte []byte) (HsResult string) {
	blake := blake2b.Sum384(UseByte)
	HsResult = fmt.Sprintf("%x", blake)
	return
}
func BLAKE2b_512(UseByte []byte) (HsResult string) {
	blake := blake2b.Sum512(UseByte)
	HsResult = fmt.Sprintf("%x", blake)
	return
}
