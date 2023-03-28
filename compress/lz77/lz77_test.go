package lz77

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestFindLongestMatch(t *testing.T) {
	data := "abcabcabc"
	c := DefaultLz77Compressor()
	match, off, l := c.findLongestMatch(data, 3)
	t.Logf("match %v offset %d len %d", match, off, l)
}

func TestCompress(t *testing.T) {
	data := "abcabcabc"
	c := DefaultLz77Compressor()
	r := c.compress(data)
	t.Log(r)
}

func TestDeCompress(t *testing.T) {
	data := "abcabcabc"
	c := DefaultLz77Compressor()
	r := c.compress(data)
	rr := c.decompress(r)
	t.Log(rr)
}

func TestBigLittleEndian(t *testing.T) {
	b16 := uint16(11265)
	buf := make([]byte, 2)
	bufW := bytes.NewBuffer(buf)
	binary.Write(bufW, binary.BigEndian, b16)

	var ruint16 uint16
	binary.Read(bufW, binary.LittleEndian, &ruint16)
	t.Logf("write big read little %d \n", ruint16)

	var ruint16b uint16
	binary.Read(bufW, binary.BigEndian, &ruint16b)
	t.Logf("write big read big %d \n", ruint16b)

	bufW = bytes.NewBuffer(buf)
	binary.Write(bufW, binary.LittleEndian, b16)

	binary.Read(bufW, binary.LittleEndian, &ruint16)
	t.Logf("write little read little %d \n", ruint16)

	binary.Read(bufW, binary.BigEndian, &ruint16b)
	t.Logf("write little read big %d \n", ruint16b)
}
