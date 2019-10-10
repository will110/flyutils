package flyutils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"lukechampine.com/frand"
	"math"
	"math/big"
)

/*
描述：查找字符串在另一个字符串中，第一次出现的位置，没有找到子串，则返回-1
注意：对大小写敏感
参数：
	str: 搜索的字符串
	sub: 要查找的字符串
 */
func StrPos(sub,str string) int {
	strList := []rune(str)
	subPos := 0
	isFind := false
	for k := range strList {
		if string(strList[k]) == sub {
			subPos = k
			isFind = true
			break
		}
	}

	if !isFind {
		return -1
	}

	return subPos
}

/**
描述：查找字符串在另一个字符串中，最后一次出现的位置, 没有找到子串，则返回-1
注意：对大小写敏感
参数：
	str: 搜索的字符串
	sub: 要查找的字符串
 */
func StrLastPos(sub,str string) int {
	strList := []rune(str)
	subPos := 0
	isFind := false
	for k := range strList {
		if string(strList[k]) == sub {
			subPos = k
			isFind = true
		}
	}

	if !isFind {
		return -1
	}

	return subPos
}

/**
描述：截取字符串，从start开始截取，截取length长度
参数：
	str: 要截取的字符串
	start: 开始的位置
	cutOutLen： 截取的长度
 */
func Substr(str string, start int, cutOutLen ...int) string {
	strList := []rune(str)
	strLen := len(strList)
	var result string
	if start >= 0 {
		if start > strLen {
			return ""
		}

		subLen := strLen
		if len(cutOutLen) > 0 {
			if cutOutLen[0] > strLen {
				cutOutLen[0] = strLen
			}

			subLen = start + cutOutLen[0]
			if subLen > strLen {
				subLen = strLen
			}
		}

		if start > 0 {
			start = start - 1
		}

		result = string(strList[start:subLen])
	} else {
		if math.Abs(float64(start)) > float64(strLen) {
			return ""
		}

		start = strLen + start + 1
		newStrList := strList[0:start]
		newStrLen := len(newStrList)
		sLen := newStrLen
		if len(cutOutLen) > 0 {
			if cutOutLen[0] > newStrLen {
				cutOutLen[0] = newStrLen
			}

			sLen = cutOutLen[0]
		}

		var n int
		for i := newStrLen; i >= 0; i-- {
			n++
			if n == sLen {
				if i > 0 {
					n = i - 1
				} else {
					n = i
				}
				break;
			}
		}

		result = string(newStrList[n:])
	}

	return result
}

/**
描述：字符串反转
参数：
	str: 要反转的字符串
 */
func StrReverse(str string) string {
	strList := []rune(str)
	strLen := len(strList)
	newStrList := make([]rune, 0, strLen)
	for i := strLen - 1; i >= 0; i-- {
		newStrList = append(newStrList, strList[i])
	}

	return string(newStrList)
}

/**
描述：字符串md5
参数：
	str: 字符串
 */
func StrMd5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

/**
描述：生成随机的字符串
参数：
	len: 生成字符串的长度
 */
func GenerateRandomStr(len int) string {
	var str = "ABC4DEP3QRST8UV7WX9YZ2FG5HIJ0KL6M1NO"
	var container bytes.Buffer
	newInt := big.NewInt(int64(bytes.NewBufferString(str).Len() - 1))
	for i := 0; i < len; i++ {
		randomInt := frand.BigIntn(newInt)
		container.WriteByte(str[randomInt.Int64()])
	}

	return container.String()
}