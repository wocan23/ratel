package common

//import "golang.org/x/text/encoding/simplifiedchinese"

var Semaphore = make(chan int8, 1)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

/**
合并两个map,相同key以第二个为主
*/
func CombineMap(m1 map[string]string, m2 map[string]string) map[string]string {
	var m = make(map[string]string)
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m

}

//func ConvertByte2String(byte []byte, charset Charset) string {
//	var str string
//	switch charset {
//	case GB18030:
//		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
//		str= string(decodeBytes)
//	case UTF8:
//		fallthrough
//	default:
//		str = string(byte)
//	}
//	return str
//}
