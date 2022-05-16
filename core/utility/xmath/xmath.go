//go:binary-only-package
package xmath

import (
	"encoding/binary"
	_ "math/rand"
)

const (
	INT8_MIN   = -0x7f - 1
	INT16_MIN  = -0x7fff - 1
	INT32_MIN  = -0x7fffffff - 1
	INT64_MIN  = -0x7fffffffffffffff - 1
	INT8_MAX   = 0x7f
	INT16_MAX  = 0x7fff
	INT32_MAX  = 0x7fffffff
	INT64_MAX  = 0x7fffffffffffffff
	UINT8_MAX  = 0xff
	UINT16_MAX = 0xffff
	UINT32_MAX = 0xffffffff
	UINT64_MAX = 0xffffffffffffffff
)

var (
	BYTE_ORDER binary.ByteOrder = &binary.LittleEndian
)

func MaxValue(a int, b int) int {
	return 0
}
func MinValue(a int, b int) int {
	return 0
}
func RandInt(min int, max int) int {
	return 0
}
func Uint32ToBytes(value uint32) []byte {
	return []byte{}
}
func Uint32FromBytes(value []byte) uint32 {
	return 0
}
