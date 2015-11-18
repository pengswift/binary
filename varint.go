/***********************************************************
*
* varints 变长编码 (压缩整型数据, 使用更少的字节来表示整数)
*
*   var a int64 = 300  一共64位
*   varints 300: 10101100 00000010  16位， 内存节省3倍
*
*   在变长编码中每个字节的最高位被当作最高有效位(msb), 如果为1，表示整数还没有结束
*
*   10101100 00000010  -> 300 推导:
*   10101100 -> msb 为1，剩余7位0101100, 继续读下一个字节
*   00000010 -> msb 位0, 整数结束，取余下7位 0000010
*   0000010 ++ 0101100
*
*   因为是采用低字节序
*   实际字节 000 0010 010 1100 -> 1 0010 1100  = 256 + 32 + 8 + 4 = 300
*   如果用数组来存储10101100 00000010
*   b[0] = 10101100 存储低位  b[1] = 00000010 存储高位
*   则在解析时要颠倒下顺序
*
 */

package binary

import (
	"encoding/binary"
	"io"
)

const (
	MaxVarintLen16 = binary.MaxVarintLen16
	MaxVarintLen32 = binary.MaxVarintLen32
	MaxVarintLen64 = binary.MaxVarintLen64
)

func UvarintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func VarintSize(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return UvarintSize(ux)
}

func GetUvarint(b []byte) (uint64, int) {
	return binary.Uvarint(b)
}

func PutUvarint(b []byte, v uint64) int {
	return binary.PutUvarint(b, v)
}

func GetVarint(b []byte) (int64, int) {
	return binary.Varint(b)
}

func PutVarint(b []byte, v int64) int {
	return binary.PutVarint(b, v)
}

func ReadUvarint(r io.ByteReader) (uint64, error) {
	return binary.ReadUvarint(r)
}

func ReadVarint(r io.ByteReader) (int64, error) {
	return binary.ReadVarint(r)
}
