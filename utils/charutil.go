package utils

import (
	"bytes"
	"math/rand"
	"time"
)

// 指定长度随机中文字符(包含复杂字符)
func GenFixedLengthChineseChars(length int) string {

	var buf bytes.Buffer

	for i := 0; i < length; i++ {
		buf.WriteRune(rune(RandInt(19968, 40869)))
	}
	return buf.String()
}

// 指定范围随机中文字符
func GenRandomLengthChineseChars(start, end int) string {
	length := RandInt(start, end)
	return GenFixedLengthChineseChars(length)
}

// 随机英文小写字母
func RandStr(len int) string {
	rand.Seed(time.Now().UnixNano())
	data := make([]byte, len)
	for i := 0; i < len; i++ {
		data[i] = byte(rand.Intn(26) + 97)
	}
	return string(data)
}

// 指定范围随机 int
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// 指定范围随机 int64
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

// 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// CardBin Metadata 生成(仅供测试生成代码)
//func CreateCardBinMetadata(){
//
//	f,_:=os.Open("../resources/cardBin")
//	r:=bufio.NewReader(f)
//
//	name:=""
//	length:= 0
//	cardType :=""
//	prefixes := []int{}
//	t,_:=template.ParseFiles("../resources/cardBin.tpl")
//	for  {
//		s,err:= r.ReadString('\n')
//		if err == io.EOF {
//			break
//		}
//		strs:=strings.Fields(s)
//		if name == strs[0] {
//			cardBin,_:=strconv.Atoi(strs[1])
//			prefixes = append(prefixes,cardBin)
//		}else {
//
//			if name!="" {
//				c:=metadata.CardBin{
//					Name:     name,
//					Length:   length,
//					CardType: cardType,
//					Prefixes: prefixes,
//				}
//				data,_:=os.OpenFile("../resources/metadata_code",os.O_CREATE|os.O_RDWR|os.O_APPEND,0644)
//				t.Execute(data,c)
//				prefixes = prefixes[:0]
//			}
//
//
//			name= strs[0]
//			length,_=strconv.Atoi(strs[2])
//			cardType=strs[3]
//			cardBin,_:=strconv.Atoi(strs[1])
//			prefixes = append(prefixes,cardBin)
//		}
//	}
//}
