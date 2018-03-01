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

package server

import (
	"encoding/json"
	"fmt"
	"github.com/mritd/idgen/generator"
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/models"
	"github.com/mritd/idgen/util"
	"html/template"
	"net"
	"net/http"
)

func Start(mode string, addr net.Addr) {
	fmt.Println("starting server at", addr.String())
	switch mode {
	case "html":
		http.HandleFunc("/", htmlServer)
		fmt.Println("html server starting...")
	case "json":
		http.HandleFunc("/", jsonServer)
		fmt.Println("json server starting...")
	default:
		http.HandleFunc("/", htmlServer)
		http.HandleFunc("/api", jsonServer)
		fmt.Println("html and json server starting...")
	}

	err := http.ListenAndServe(addr.String(), nil)
	util.CheckAndExit(err)
}

func htmlServer(w http.ResponseWriter, _ *http.Request) {
	tpl, err := template.New("htmlTpl").Parse(metadata.HtmlTpl)
	util.CheckAndExit(err)
	tpl.Execute(w, gen())
}

func jsonServer(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(gen())
	util.CheckAndExit(err)
	fmt.Fprintf(w, string(b))
}

func gen() models.GenData {
	name := generator.NameGenerate()
	idNo := generator.IDCardGenerate()
	mobile := generator.MobileGenerate()
	bank := generator.BankGenerate()
	email := generator.EmailGenerate()
	addr := generator.AddrGenerate()

	return models.GenData{
		Name:   name,
		Mobile: mobile,
		IdNo:   idNo,
		Bank:   bank,
		Email:  email,
		Addr:   addr,
	}
}
