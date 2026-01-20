package server

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mritd/chinaid/v2"
	log "github.com/sirupsen/logrus"
)

// Identity represents a generated identity record
type Identity struct {
	Name    string `json:"name"`
	IDNo    string `json:"idno"`
	Mobile  string `json:"mobile"`
	Bank    string `json:"bank"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

// BatchResponse represents the batch generation response
type BatchResponse struct {
	Count int        `json:"count"`
	Data  []Identity `json:"data"`
}

func personToIdentity(p *chinaid.Person) Identity {
	return Identity{
		Name:    p.Name(),
		IDNo:    p.IDNo(),
		Mobile:  p.Mobile(),
		Bank:    p.BankNo(),
		Email:   p.Email(),
		Address: p.Address(),
	}
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	p := chinaid.NewPerson().Build()
	identity := personToIdentity(p)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(identity); err != nil {
		log.Errorf("Failed to encode response: %v", err)
	}
}

func handleBatch(w http.ResponseWriter, r *http.Request) {
	count := parseCount(r)

	var identities []Identity
	for _, p := range chinaid.NewPerson().BuildN(count) {
		identities = append(identities, personToIdentity(p))
	}

	resp := BatchResponse{
		Count: count,
		Data:  identities,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Errorf("Failed to encode response: %v", err)
	}
}

func handleExport(w http.ResponseWriter, r *http.Request) {
	count := parseCount(r)

	var identities []Identity
	for _, p := range chinaid.NewPerson().BuildN(count) {
		identities = append(identities, personToIdentity(p))
	}

	filename := fmt.Sprintf("idgen_export_%s.csv", time.Now().Format("20060102_150405"))
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	// Write BOM for Excel compatibility
	_, _ = w.Write([]byte{0xEF, 0xBB, 0xBF})

	csvWriter := csv.NewWriter(w)
	defer csvWriter.Flush()

	// Header
	if err := csvWriter.Write([]string{"name", "idno", "mobile", "bank", "email", "address"}); err != nil {
		log.Errorf("Failed to write CSV header: %v", err)
		return
	}

	// Data rows
	for _, id := range identities {
		row := []string{id.Name, id.IDNo, id.Mobile, id.Bank, id.Email, id.Address}
		if err := csvWriter.Write(row); err != nil {
			log.Errorf("Failed to write CSV row: %v", err)
			return
		}
	}
}
