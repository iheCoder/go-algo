package encode_decode_tinyurl

import "strconv"

type Codec struct {
	m  map[string]string
	id int
}

func Constructor() Codec {
	return Codec{m: make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	v := strconv.Itoa(c.id)
	c.m[v] = longUrl
	c.id++
	return v
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.m[shortUrl]
}
