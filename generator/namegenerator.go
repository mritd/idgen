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
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/util"
)

// 生成姓名
func NameGenerate() string {
	if util.DBExist() {
		var firstName string
		util.CheckAndExit(util.DB().QueryRow(util.FirstNameSQL, util.RandInt(0, util.FirstNameSum)).Scan(&firstName))
		return metadata.LastName[util.RandInt(0, len(metadata.LastName))] + firstName
	} else {
		return metadata.LastName[util.RandInt(0, len(metadata.LastName))] + util.GenRandomLengthChineseChars(1, 3)
	}
}
