package lz77

import (
	"bytes"
	"encoding/binary"
)

// https://github.com/manassra/LZ77-Compressor/blob/master/src/LZ77.py
// https://blog.51cto.com/u_15127629/2873305
// chatgpt

const (
	defaultWindowSize          = 4
	defaultMinMatchSize        = 2
	defaultLookAheadBufferSize = 4
	matchFlag                  = 1
	missMatchFlag              = 0
)

type lz77Compressor struct {
	winSize                int
	maxLookAheadBufferSize int
	minMatchSize           int
}

func (c *lz77Compressor) Compress(inputFilePath, outputFilePath string) {

}

func NewLz77Compressor(lookAheadBufferSize, minMatchSize, winSize int) *lz77Compressor {
	return &lz77Compressor{
		winSize:                winSize,
		maxLookAheadBufferSize: lookAheadBufferSize,
		minMatchSize:           minMatchSize,
	}
}

func DefaultLz77Compressor() *lz77Compressor {
	return NewLz77Compressor(defaultLookAheadBufferSize, defaultMinMatchSize, defaultWindowSize)
}

func (c *lz77Compressor) compress(data string) []byte {
	bitBuffer := bytes.Buffer{}
	var pos int
	for pos < len(data) {
		match, offset, l := c.findLongestMatch(data, pos)
		if match {
			bitBuffer.WriteByte(matchFlag)
			binary.Write(&bitBuffer, binary.BigEndian, uint16(offset))
			bitBuffer.WriteByte(byte(l))
			pos += l
		} else {
			bitBuffer.WriteByte(missMatchFlag)
			bitBuffer.WriteByte(data[pos])
			pos++
		}
	}

	return bitBuffer.Bytes()
}

func (c *lz77Compressor) decompress(data []byte) string {
	output := bytes.Buffer{}
	var pos int
	var offset uint16
	var l byte
	var flag byte

	for pos < len(data) {
		flag = data[pos]
		pos++
		if flag == matchFlag {
			dataBuffer := bytes.NewBuffer(data[pos : pos+2])
			binary.Read(dataBuffer, binary.BigEndian, &offset)
			pos += 2
			l = data[pos]
			ob := output.Bytes()
			output.Write(ob[len(ob)-int(offset) : len(ob)-int(offset)+int(l)])
		} else if flag == missMatchFlag {
			output.WriteByte(data[pos])
		} else {
			panic("wrong flag")
		}
		pos++
	}

	return output.String()
}

func (c *lz77Compressor) findLongestMatch(data string, currentPos int) (bool, int, int) {
	endPos := minInt(currentPos+c.maxLookAheadBufferSize, len(data)+1)
	var bestMatchDistance int
	var bestMatchLength int
	// 遍历前向缓冲区
	for i := currentPos + c.minMatchSize; i < endPos; i++ {
		startIndex := maxInt(0, currentPos-c.winSize)
		// 获取当前前向缓冲区
		lookAheadBuffer := data[currentPos:i]

		// 遍历搜索缓冲区
		for j := startIndex; j < currentPos; j++ {
			repetitions := len(lookAheadBuffer) / (currentPos - j)
			last := len(lookAheadBuffer) % (currentPos - j)
			var matchedStr string
			// 在前向缓冲区大于搜索缓冲区的时候可以重复组合搜索缓冲区尝试是否等于前向缓冲区
			for k := 0; k < repetitions; k++ {
				matchedStr += data[j:currentPos]
			}
			matchedStr += data[j : j+last]

			// 在搜索缓冲区组合与前向缓冲区匹配且大于之前匹配的长度，那么更新匹配偏移量与长度
			if matchedStr == lookAheadBuffer && len(lookAheadBuffer) > bestMatchLength {
				bestMatchDistance = currentPos - j
				bestMatchLength = len(lookAheadBuffer)
			}
		}
	}

	if bestMatchDistance > 0 && bestMatchLength > 0 {
		return true, bestMatchDistance, bestMatchLength
	}
	return false, 0, 0
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
