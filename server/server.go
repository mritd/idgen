package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/mritd/idgen/generator"
	"github.com/mritd/idgen/metadata"
	"github.com/mritd/idgen/models"
	"github.com/mritd/idgen/utils"
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
	utils.CheckAndExit(err)
}

func htmlServer(w http.ResponseWriter, _ *http.Request) {
	tpl, err := template.New("htmlTpl").Parse(metadata.HtmlTpl)
	utils.CheckAndExit(err)
	tpl.Execute(w, gen())
}

func jsonServer(w http.ResponseWriter, _ *http.Request) {
	b, err := json.Marshal(gen())
	utils.CheckAndExit(err)
	fmt.Fprintf(w, string(b))
}

func gen() models.GenData {
	name := generator.GetName()
	idNo := generator.GetIDCard()
	mobile := generator.GetMobile()
	bank := generator.GetBank()
	email := generator.GetEmail()
	addr := generator.GetAddress()

	return models.GenData{
		Name:   name,
		Mobile: mobile,
		IdNo:   idNo,
		Bank:   bank,
		Email:  email,
		Addr:   addr,
	}
}
