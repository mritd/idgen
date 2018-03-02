// Copyright © 2018 mritd <mritd1234@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package generator

import (
	"fmt"
	"math"
	"strconv"

	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/util"
)

// 随机生成银行卡号
func BankGenerate() string {

	// 随机选中银行卡卡头
	bank := metadata.CardBins[util.RandInt(0, len(metadata.CardBins))]

	// 获取 卡前缀(cardBin)
	prefixes := bank.Prefixes

	// 获取当前银行卡正确长度
	cardNoLength := bank.Length

	// 生成 长度-1 位卡号
	preCardNo := strconv.Itoa(prefixes[util.RandInt(0, len(prefixes))]) + fmt.Sprintf("%0*d", cardNoLength-7, util.RandInt64(0, int64(math.Pow10(cardNoLength-7))))

	// LUHN 算法处理
	return LUHNProcess(preCardNo)

}

// LUHN 合成卡号
func LUHNProcess(preCardNo string) string {

	checkSum := 0
	tmpCardNo := util.ReverseString(preCardNo)
	for i, s := range tmpCardNo {

		tmp, err := strconv.Atoi(string(s))
		util.CheckAndExit(err)

		// 由于卡号实际少了一位，所以反转后卡号第一位一定为偶数位
		// 同时 i 正好也是偶数，此时 i 将和卡号奇偶位同步
		if i%2 == 0 {
			// 偶数位 *2 是否为两位数(>9)
			if tmp*2 > 9 {
				// 如果为两位数则 -9
				checkSum += tmp*2 - 9
			} else {
				// 否则直接相加即可
				checkSum += tmp * 2
			}
		} else {
			// 奇数位直接相加
			checkSum += tmp
		}
	}

	if checkSum%10 != 0 {
		return preCardNo + strconv.Itoa(10-checkSum%10)
	} else {
		// 如果不巧生成的前 卡长度-1 位正好符合 LUHN 算法
		// 那么需要递归重新生成(需要符合 cardBind 中卡号长度)
		return BankGenerate()
	}
}
