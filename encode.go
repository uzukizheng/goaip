package goaip

import "github.com/axgle/mahonia"

func Utf8ToGBK(txt string) []byte {
	enc := mahonia.NewEncoder("gbk")
	return []byte(enc.ConvertString(txt))
}

func GBKToUtf8(txt string) string {
	result := mahonia.NewDecoder("gbk").ConvertString(txt)
	utf8Decoder := mahonia.NewDecoder("utf-8")
	return utf8Decoder.ConvertString(result)
}
