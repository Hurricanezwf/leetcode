package main

import (
	"log"
	"math"
)

func main() {
	if v := myAtoi("0"); v != 0 {
		log.Fatalf("%d != 0", v)
	}
	if v := myAtoi("9"); v != 9 {
		log.Fatalf("%d != 9", v)
	}
	if v := myAtoi("10"); v != 10 {
		log.Fatalf("%d != 10", v)
	}
	if v := myAtoi("-10"); v != -10 {
		log.Fatalf("%d != -10", v)
	}
	if v := myAtoi(" -11"); v != -11 {
		log.Fatalf("%d != -11", v)
	}
	if v := myAtoi(" -11 2"); v != -11 {
		log.Fatalf("%d != -11", v)
	}
	if v := myAtoi(" w22 2"); v != 0 {
		log.Fatalf("%d != 0", v)
	}
	if v := myAtoi(" w22w2"); v != 0 {
		log.Fatalf("%d != 0", v)
	}
	if v := myAtoi("- w22w2"); v != 0 {
		log.Fatalf("%d != 0", v)
	}
	if v := myAtoi("4193 with words"); v != 4193 {
		log.Fatalf("%d != 4193", v)
	}
	if v := myAtoi("words and 987"); v != 0 {
		log.Fatalf("%d != 0", v)
	}
	if v := myAtoi("99999999999999999"); v != math.MaxInt32 {
		log.Fatalf("%d != %d", v, math.MaxInt32)
	}
	if v := myAtoi("-99999999999999999"); v != math.MinInt32 {
		log.Fatalf("%d != %d", v, math.MinInt32)
	}
	if v := myAtoi("-91283472332"); v != -2147483648 {
		log.Fatalf("%d != %d", v, -2147483648)
	}
	if v := myAtoi("9223372036854775808"); v != 2147483647 {
		log.Fatalf("%d != %d", v, 2147483647)
	}
	log.Printf("OK\n")
}

type CharType uint8

const (
	Symbol = iota + 1
	Number
	Space
	Other
)

func myAtoi(str string) int {
	var (
		value        int64
		isNegative   bool
		startConvert bool
	)
	for i := 0; i < len(str); i++ {
		if !startConvert {
			switch charType(str[i]) {
			case Symbol:
				startConvert = true
				if str[i] == '-' {
					isNegative = true
				}
			case Number:
				startConvert = true
				value = int64(str[i] - byte('0'))
			case Space:
				// do nothing and continue
			case Other:
				return 0
			default:
				log.Fatalf("unknown char type %v\n", charType(str[i]))
			}
		} else {
			switch charType(str[i]) {
			case Symbol, Space, Other:
				return toInt(value, isNegative)
			case Number:
				value = value*10 + charToInt64(str[i])
				if value > (-1 * math.MinInt32) {
					return toInt(value, isNegative)
				}
			default:
				log.Fatalf("unknown char type %v\n", charType(str[i]))
			}
		}
	}
	return toInt(value, isNegative)
}

func toInt(value int64, isNegative bool) int {
	if isNegative {
		value = value * -1
	}
	// 范围筛选
	if value < math.MinInt32 {
		return math.MinInt32
	}
	if value > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(value)
}

func charToInt64(c byte) int64 {
	return int64(c - byte('0'))
}

func charType(c byte) CharType {
	if isSymbol(c) {
		return Symbol
	}
	if isNumber(c) {
		return Number
	}
	if isSpace(c) {
		return Space
	}
	return Other
}

func isSymbol(c byte) bool {
	return c == '-' || c == '+'
}

func isNumber(c byte) bool {
	return byte('0') <= c && c <= byte('9')
}

func isSpace(c byte) bool {
	return c == ' '
}
