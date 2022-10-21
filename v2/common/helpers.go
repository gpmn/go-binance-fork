package common

import (
	"bytes"
	"errors"
	"math"
	"strconv"
)

// AmountToLotSize converts an amount to a lot sized amount
func AmountToLotSize(lot float64, precision int, amount float64) float64 {
	return math.Trunc(math.Floor(amount/lot)*lot*math.Pow10(precision)) / math.Pow10(precision)
}

// ToJSONList convert v to json list if v is a map
func ToJSONList(v []byte) []byte {
	if len(v) > 0 && v[0] == '{' {
		var b bytes.Buffer
		b.Write([]byte("["))
		b.Write(v)
		b.Write([]byte("]"))
		return b.Bytes()
	}
	return v
}

// Float64 :
type Float64 float64

// Int64 :
type Int64 int64

// UnmarshalJSON :
func (val *Float64) UnmarshalJSON(b []byte) (err error) {
	var slice []byte
	if len(b) >= 2 && b[0] == '"' { // "123.456" 和 ""格式的数字字符串
		if b[len(b)-1] != '"' {
			return errors.New(`missing '"' suffix`)
		}
		slice = b[1 : len(b)-1]
		if len(slice) == 0 {
			*val = 0
			return nil
		}
	} else { //	    123.456格式的字符串
		slice = b
	}
	tmp, err := strconv.ParseFloat(string(slice), 10)
	if nil != err {
		return err
	}
	*val = Float64(tmp)
	return nil
}

// UnmarshalJSON :
func (val *Int64) UnmarshalJSON(b []byte) (err error) {
	var slice []byte
	if len(b) >= 2 && b[0] == '"' { // "123.456" 和 ""格式的数字字符串
		if b[len(b)-1] != '"' {
			return errors.New(`missing '"' suffix`)
		}
		slice = b[1 : len(b)-1]
		if len(slice) == 0 {
			*val = 0
			return nil
		}
	} else { //	    123.456格式的字符串
		slice = b
	}
	tmp, err := strconv.ParseInt(string(slice), 10, 64)
	if nil != err {
		return err
	}
	*val = Int64(tmp)
	return nil
}
