package convert_a_number_to_hex

const hexMask uint32 = 4026531840

var ten2hex = map[uint32]byte{
	0:        '0',
	1 << 28:  '1',
	2 << 28:  '2',
	3 << 28:  '3',
	4 << 28:  '4',
	5 << 28:  '5',
	6 << 28:  '6',
	7 << 28:  '7',
	8 << 28:  '8',
	9 << 28:  '9',
	10 << 28: 'a',
	11 << 28: 'b',
	12 << 28: 'c',
	13 << 28: 'd',
	14 << 28: 'e',
	15 << 28: 'f',
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	number := uint32(num)
	var sb []byte
	var hexNum uint32
	for i := 0; i < 8; i++ {
		hexNum = number << (4 * i) & hexMask
		if hexNum == 0 && len(sb) == 0 {
			continue
		}
		sb = append(sb, ten2hex[hexNum])
	}
	return string(sb)
}
