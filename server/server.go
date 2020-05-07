package server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/mritd/chinaid"
)

type Response struct {
	Name   string
	Mobile string
	IdNo   string
	Bank   string
	Email  string
	Addr   string
}

func Start(mode string, addr net.Addr) {
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

	fmt.Printf("server listen at %s...\n", addr.String())
	err := http.ListenAndServe(addr.String(), nil)
	if err != nil {
		panic(err)
	}
}

func htmlServer(w http.ResponseWriter, _ *http.Request) {
	resp := htmlStr
	resp = strings.ReplaceAll(resp, "{{ .Name }}", chinaid.Name())
	resp = strings.ReplaceAll(resp, "{{ .Mobile }}", chinaid.Mobile())
	resp = strings.ReplaceAll(resp, "{{ .IDNo }}", chinaid.IDNo())
	resp = strings.ReplaceAll(resp, "{{ .BankNo }}", chinaid.BankNo())
	resp = strings.ReplaceAll(resp, "{{ .Email }}", chinaid.Email())
	resp = strings.ReplaceAll(resp, "{{ .Address }}", chinaid.Address())
	_, err := fmt.Fprint(w, resp)
	if err != nil {
		fmt.Println(err)
	}
}

func jsonServer(w http.ResponseWriter, _ *http.Request) {
	b, err := json.MarshalIndent(
		Response{
			Name:   chinaid.Name(),
			Mobile: chinaid.Mobile(),
			IdNo:   chinaid.IDNo(),
			Bank:   chinaid.BankNo(),
			Email:  chinaid.Email(),
			Addr:   chinaid.Address(),
		}, "", "    ")
	if err != nil {
		_, err = fmt.Fprint(w, err.Error())
		if err != nil {
			fmt.Println(err)
		}
	} else {
		_, err = fmt.Fprint(w, string(b))
		if err != nil {
			fmt.Println(err)
		}
	}
}
