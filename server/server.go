package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"net/http"

	"github.com/mritd/idgen/metadata"

	"github.com/gobuffalo/packr/v2"

	"github.com/mritd/idgen/generator"
)

var indexStr string

func init() {
	var err error
	box := packr.New("resources", "../resources")
	indexStr, err = box.FindString("index.tpl")
	if err != nil {
		panic(err)
	}
}

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
	if err != nil {
		panic(err)
	}
}

func htmlServer(w http.ResponseWriter, _ *http.Request) {

	tpl, err := template.New("").Parse(indexStr)
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.Execute(w, getResponse())
	if err != nil {
		fmt.Println(err)
	}
}

func jsonServer(w http.ResponseWriter, _ *http.Request) {
	b, err := json.MarshalIndent(getResponse(), "", "    ")
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err = fmt.Fprintf(w, string(b))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getResponse() metadata.Response {
	return metadata.Response{
		Name:   generator.GetName(),
		Mobile: generator.GetMobile(),
		IdNo:   generator.GetIDNo(),
		Bank:   generator.GetBank(),
		Email:  generator.GetEmail(),
		Addr:   generator.GetAddress(),
	}
}
