package convert

import (
	"encoding/json"
	"math"
	"strconv"
	"strings"
)

// String2Int 字符串转Int
// intStr：数字的字符串
func String2Int(intStr string) (intNum int) {
	intNum, _ = strconv.Atoi(intStr)
	return
}

// String2Int64 字符串转Int64
// intStr：数字的字符串
func String2Int64(intStr string) (int64Num int64) {
	intNum, _ := strconv.Atoi(intStr)
	int64Num = int64(intNum)
	return
}

// String2Float64 字符串转Float64
// floatStr：小数点数字的字符串
func String2Float64(floatStr string) (floatNum float64) {
	floatNum, _ = strconv.ParseFloat(floatStr, 64)
	return
}

// String2Float32 字符串转Float32
// floatStr：小数点数字的字符串
func String2Float32(floatStr string) (floatNum float32) {
	floatNum64, _ := strconv.ParseFloat(floatStr, 32)
	floatNum = float32(floatNum64)
	return
}

// Int2String Int转字符串
// intNum：数字字符串
func Int2String(intNum int) (intStr string) {
	intStr = strconv.Itoa(intNum)
	return
}

// Int642String Int64转字符串
// intNum：数字字符串
func Int642String(intNum int64) (int64Str string) {
	//10, 代表10进制
	int64Str = strconv.FormatInt(intNum, 10)
	return
}

// Float64ToString Float64转字符串
// floatNum：float64数字
// prec：精度位数（不传则默认float数字精度）
func Float64ToString(floatNum float64, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(floatNum, 'f', prec[0], 64)
		return
	}
	floatStr = strconv.FormatFloat(floatNum, 'f', -1, 64)
	return
}

// Float32ToString Float32转字符串
// floatNum：float32数字
// prec：精度位数（不传则默认float数字精度）
func Float32ToString(floatNum float32, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(float64(floatNum), 'f', prec[0], 32)
		return
	}
	floatStr = strconv.FormatFloat(float64(floatNum), 'f', -1, 32)
	return
}

// BinaryToDecimal 二进制转10进制
func BinaryToDecimal(bit string) (num int) {
	fields := strings.Split(bit, "")
	lens := len(fields)
	var tempF = 0.0
	for i := 0; i < lens; i++ {
		floatNum := String2Float64(fields[i])
		tempF += floatNum * math.Pow(2, float64(lens-i-1))
	}
	num = int(tempF)
	return
}

func ToString(v any) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}

func ToBytes(v any) (bs []byte) {
	if v == nil {
		return nil
	}
	bs, _ = json.Marshal(v)
	return
}
