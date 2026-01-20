package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
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

// FormatType represents output format type
type FormatType string

const (
	FormatTable FormatType = "table"
	FormatJSON  FormatType = "json"
	FormatCSV   FormatType = "csv"
)

// Formatter handles output formatting
type Formatter struct {
	format FormatType
	writer io.Writer
}

// NewFormatter creates a new formatter with specified format
func NewFormatter(w io.Writer, format string) *Formatter {
	ft := FormatTable
	switch strings.ToLower(format) {
	case "json":
		ft = FormatJSON
	case "csv":
		ft = FormatCSV
	case "table":
		ft = FormatTable
	}
	return &Formatter{format: ft, writer: w}
}

// FormatSingle formats a single field value (for single-field commands like idgen name)
func (f *Formatter) FormatSingle(field string, values []string) error {
	switch f.format {
	case FormatJSON:
		return f.formatSingleJSON(field, values)
	case FormatCSV:
		return f.formatSingleCSV(field, values)
	default:
		return f.formatSingleTable(field, values)
	}
}

// FormatIdentities formats multiple identity records
func (f *Formatter) FormatIdentities(identities []Identity) error {
	switch f.format {
	case FormatJSON:
		return f.formatJSON(identities)
	case FormatCSV:
		return f.formatCSV(identities)
	default:
		return f.formatTable(identities)
	}
}

func (f *Formatter) formatSingleJSON(field string, values []string) error {
	if len(values) == 1 {
		data := map[string]string{field: values[0]}
		enc := json.NewEncoder(f.writer)
		enc.SetIndent("", "  ")
		return enc.Encode(data)
	}
	var data []map[string]string
	for _, v := range values {
		data = append(data, map[string]string{field: v})
	}
	enc := json.NewEncoder(f.writer)
	enc.SetIndent("", "  ")
	return enc.Encode(data)
}

func (f *Formatter) formatSingleCSV(field string, values []string) error {
	w := csv.NewWriter(f.writer)
	defer w.Flush()

	if err := w.Write([]string{field}); err != nil {
		return err
	}
	for _, v := range values {
		if err := w.Write([]string{v}); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) formatSingleTable(field string, values []string) error {
	if len(values) == 1 {
		_, err := fmt.Fprintln(f.writer, values[0])
		return err
	}

	// Calculate max width
	maxWidth := len(field)
	for _, v := range values {
		if w := displayWidth(v); w > maxWidth {
			maxWidth = w
		}
	}

	// Print table
	border := "+" + strings.Repeat("-", maxWidth+4) + "+"
	_, _ = fmt.Fprintln(f.writer, border)
	_, _ = fmt.Fprintf(f.writer, "|  %-*s  |\n", maxWidth, field)
	_, _ = fmt.Fprintln(f.writer, border)
	for _, v := range values {
		padding := maxWidth - displayWidth(v)
		_, _ = fmt.Fprintf(f.writer, "|  %s%s  |\n", v, strings.Repeat(" ", padding))
	}
	_, _ = fmt.Fprintln(f.writer, border)
	return nil
}

func (f *Formatter) formatJSON(identities []Identity) error {
	enc := json.NewEncoder(f.writer)
	enc.SetIndent("", "  ")
	if len(identities) == 1 {
		return enc.Encode(identities[0])
	}
	return enc.Encode(identities)
}

func (f *Formatter) formatCSV(identities []Identity) error {
	w := csv.NewWriter(f.writer)
	defer w.Flush()

	// Header
	if err := w.Write([]string{"name", "idno", "mobile", "bank", "email", "address"}); err != nil {
		return err
	}

	// Data rows
	for _, id := range identities {
		row := []string{id.Name, id.IDNo, id.Mobile, id.Bank, id.Email, id.Address}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func (f *Formatter) formatTable(identities []Identity) error {
	if len(identities) == 1 {
		// Single record - simple key-value format
		id := identities[0]
		_, _ = fmt.Fprintf(f.writer, "Name:    %s\n", id.Name)
		_, _ = fmt.Fprintf(f.writer, "IDNo:    %s\n", id.IDNo)
		_, _ = fmt.Fprintf(f.writer, "Mobile:  %s\n", id.Mobile)
		_, _ = fmt.Fprintf(f.writer, "Bank:    %s\n", id.Bank)
		_, _ = fmt.Fprintf(f.writer, "Email:   %s\n", id.Email)
		_, _ = fmt.Fprintf(f.writer, "Address: %s\n", id.Address)
		return nil
	}

	// Multiple records - table format
	// Calculate column widths
	cols := []struct {
		header string
		width  int
		getter func(Identity) string
	}{
		{"#", 3, nil},
		{"Name", 8, func(id Identity) string { return id.Name }},
		{"IDNo", 18, func(id Identity) string { return id.IDNo }},
		{"Mobile", 11, func(id Identity) string { return id.Mobile }},
		{"Bank", 19, func(id Identity) string { return id.Bank }},
		{"Email", 20, func(id Identity) string { return id.Email }},
		{"Address", 30, func(id Identity) string { return id.Address }},
	}

	// Update widths based on data
	for i, id := range identities {
		numWidth := len(fmt.Sprintf("%d", i+1))
		if numWidth > cols[0].width {
			cols[0].width = numWidth
		}
		for j := 1; j < len(cols); j++ {
			if w := displayWidth(cols[j].getter(id)); w > cols[j].width {
				cols[j].width = w
			}
		}
	}

	// Build border and format strings
	var borderParts []string
	for _, col := range cols {
		borderParts = append(borderParts, strings.Repeat("-", col.width+2))
	}
	border := "+" + strings.Join(borderParts, "+") + "+"

	// Print header
	_, _ = fmt.Fprintln(f.writer, border)
	var headerParts []string
	for _, col := range cols {
		headerParts = append(headerParts, fmt.Sprintf(" %-*s ", col.width, col.header))
	}
	_, _ = fmt.Fprintln(f.writer, "|"+strings.Join(headerParts, "|")+"|")
	_, _ = fmt.Fprintln(f.writer, border)

	// Print data rows
	for i, id := range identities {
		var rowParts []string
		rowParts = append(rowParts, fmt.Sprintf(" %-*d ", cols[0].width, i+1))
		for j := 1; j < len(cols); j++ {
			val := cols[j].getter(id)
			padding := cols[j].width - displayWidth(val)
			rowParts = append(rowParts, fmt.Sprintf(" %s%s ", val, strings.Repeat(" ", padding)))
		}
		_, _ = fmt.Fprintln(f.writer, "|"+strings.Join(rowParts, "|")+"|")
	}
	_, _ = fmt.Fprintln(f.writer, border)

	return nil
}

// displayWidth calculates the display width of a string (handling CJK characters)
func displayWidth(s string) int {
	width := 0
	for _, r := range s {
		if r >= 0x4E00 && r <= 0x9FFF || // CJK Unified Ideographs
			r >= 0x3400 && r <= 0x4DBF || // CJK Unified Ideographs Extension A
			r >= 0xFF00 && r <= 0xFFEF { // Fullwidth Forms
			width += 2
		} else {
			width += 1
		}
	}
	return width
}

// ToClipboardText formats identities for clipboard
func ToClipboardText(identities []Identity) string {
	if len(identities) == 1 {
		id := identities[0]
		return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
			id.Name, id.IDNo, id.Mobile, id.Bank, id.Email, id.Address)
	}

	var lines []string
	for i, id := range identities {
		lines = append(lines, fmt.Sprintf("[%d] %s | %s | %s | %s | %s | %s",
			i+1, id.Name, id.IDNo, id.Mobile, id.Bank, id.Email, id.Address))
	}
	return strings.Join(lines, "\n")
}

// SingleToClipboardText formats single field values for clipboard
func SingleToClipboardText(values []string) string {
	return strings.Join(values, "\n")
}
