package sentence_similarity_3

import (
	"reflect"
	"strings"
)

// 可以插入啊
func areSentencesSimilarLRSub(sentence1 string, sentence2 string) bool {
	var match, target string
	if len(sentence1) > len(sentence2) {
		target = sentence1
		match = sentence2
	} else if len(sentence2) > len(sentence1) {
		target = sentence2
		match = sentence1
	} else {
		return sentence1 == sentence2
	}

	ml := len(match)
	return match == target[:ml] || match == target[len(target)-ml:]
}

func areSentencesSimilarCompose(sentence1 string, sentence2 string) bool {
	s1s := strings.Split(sentence1, " ")
	s2s := strings.Split(sentence2, " ")
	var target string
	var ts, ms []string
	// aaa 和a a，前者len是要超过后者的！！！
	if len(s1s) > len(s2s) {
		ts = s1s
		target = sentence1
		ms = s2s
	} else if len(s2s) > len(s1s) {
		target = sentence2
		ts = s2s
		ms = s1s
	} else {
		return sentence1 == sentence2
	}

	//ts := strings.Split(target, " ")
	//ms := strings.Split(match, " ")
	if reflect.DeepEqual(ms, ts[:len(ms)]) || reflect.DeepEqual(ms, ts[len(ts)-len(ms):]) {
		return true
	}

	composeStrs := make([]string, 0)
	var j int
	for i := 0; i < len(ts); i++ {
		if ms[j] == ts[i] {
			j++
			if j >= len(ms) {
				// target后面的没有计入到compose
				i++
				for ; i < len(ts); i++ {
					composeStrs = append(composeStrs, ts[i])
				}
				break
			}
		} else {
			composeStrs = append(composeStrs, ts[i])
		}
	}
	if j < len(ms) {
		return false
	}

	compose := strings.Join(composeStrs, " ")

	for k := 1; k < len(ms); k++ {
		// append 之后ms会改变
		try := make([]string, len(ms[:k]))
		copy(try, ms[:k])
		try = append(try, compose)
		try = append(try, ms[k:]...)
		if strings.Join(try, " ") == target {
			return true
		}
	}
	return false
}
