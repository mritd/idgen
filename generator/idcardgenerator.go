package generator

import (
    "github.com/mritd/idgen/metadata"
    "github.com/mritd/idgen/util"
    "time"
    "strconv"
    "fmt"
)

// 生成签发机关：XXX公安局/XX区分局
func GenerateIssueOrg() string {
    return metadata.CityName[util.RandInt(0, len(metadata.CityName))] + "公安局某某分局"
}

// 生成有效期限：20150906-20350906
func GenerateValidPeriod() string {
    begin := RandDate()
    end := begin.AddDate(20, 0, 0)
    return begin.Format("20060102") + "-" + end.Format("20060102")
}

// 身份证号生成
func IDCardGenerate() string {

    // AreaCode 随机一个+4位随机数字(不够左填充0)
    areaCode := metadata.AreaCode[util.RandInt(0, len(metadata.AreaCode))] +
        fmt.Sprintf("%0*d", 4, util.RandInt(1, 9999))

    birthday := RandDate().Format("20060102")
    randomCode := fmt.Sprintf("%0*d", 3, util.RandInt(0, 999))
    pre := areaCode + birthday + randomCode
    pre += VerifyCode(pre)

    return pre

}

// 获取 VerifyCode
func VerifyCode(cardId string) string {
    tmp := 0
    for i, v := range metadata.Wi {
        t, _ := strconv.Atoi(string(cardId[i]))
        tmp += t * v
    }
    return metadata.ValCodeArr[tmp%11]
}

// 随机时间 1970-2000
func RandDate() time.Time {
    begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
    end, _ := time.Parse("2006-01-02 15:04:05", "2000-01-01 00:00:00")
    return time.Unix(util.RandInt64(begin.Unix(), end.Unix()), 0)
}
