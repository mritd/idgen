package generator

import (
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/util"
	"time"
)

// 生成签发机关：XXX公安局/XX区分局
func GenerateIssueOrg() string {
	return metadata.CityName[util.RandInt(0, len(metadata.ProvinceCity))]+ "公安局某某分局"
}

// 生成有效期限：20150906-20350906
func GenerateValidPeriod() string{
	begin :=RandDate()
	end := begin.AddDate(20,0,0)
	return begin.Format("20060102")+"-"+end.Format("20060102")
}



// 随机时间
func RandDate() time.Time {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2000-01-01 00:00:00")
	return time.Unix(util.RandInt64(begin.Unix(),end.Unix()),0)
}
