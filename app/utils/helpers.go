package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/revel/revel"
	"io"
	"math/rand"
	"time"
	"strconv"
)

type TrackResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	Alnum = iota
	Alpha
	Numeric
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphas = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

func RandString(stringType, n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	if stringType == Alnum {
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
	} else if stringType == Numeric {
		for i := range b {
			b[i] = numbers[rand.Intn(len(numbers))]
		}
	}else if stringType == Alpha {
		for i := range b {
			b[i] = alphas[rand.Intn(len(alphas))]
		}
	}
	return string(b)
}

func Token(uid int,ip string) string {
	return Md5String(strconv.Itoa(uid)+RandString(Alnum,6)+time.Now().String())
}

func Salt() string {
	return RandString(Alnum, 16)
}

func Sms_code() string {
	return RandString(Numeric, 6)
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func FormatNow(format string) string{
	return time.Now().Format(format)
}

// 加密密码,转成md5
func EncryptPassword(salt, password string) string {
	return Md5String(password+salt)
}

func Md5String(src string) string{
	h := md5.New()
	io.WriteString(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func WrapJsonResult(c int, d interface{}, m string) *TrackResult {
	return &TrackResult{
		Code:    c,
		Data:    d,
		Message: m,
	}
}

func WrapOKJsonResult(data interface{}) *TrackResult {
	return WrapJsonResult(200, data, "")
}

func WrapFailJsonResult(message string) *TrackResult {
	data := WrapJsonResult(501, nil, message)
	return data
}

func WrapValidationFailJsonResult(code int, e []*revel.ValidationError) *TrackResult {
	data := WrapJsonResult(code, nil, ValidationErrorToString(e))
	return data
}

func ValidationErrorToString(e []*revel.ValidationError) string {
	if nil != e {
		return e[0].Message
	}
	return ""
}
