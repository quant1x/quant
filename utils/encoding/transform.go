package encoding

import (
	"github.com/mymmsc/gox/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset = string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GBK     = Charset("GBK")
)

func GetBytes(data []byte, charset Charset) (cdata []byte, err error) {
	tagCoder := encoding.NewDecoder(charset)
	_, cdata, err = tagCoder.Translate(data, true)
	return
}

func GetBytes0(data []byte, charset Charset) ([]byte, error) {
	switch charset {
	case GBK:
		return simplifiedchinese.GBK.NewDecoder().Bytes(data)
	case GB18030:
		return simplifiedchinese.GB18030.NewDecoder().Bytes(data)
	case UTF8:
		fallthrough
	default:
		return data, nil
	}
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := encoding.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := encoding.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
