{
    "{{ .Name }}",
    {{ .Length }},
    "{{ .CardType }}",
    []int{
        {{ range .Prefixes}}{{ . }},{{end}}
    },
},
