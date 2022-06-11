// Code generated {{.generator}}; DO NOT EDIT.

package iso4217

const (
	_   Currency = iota
{{- range .currencies}}
	{{.Alpha}}
    {{- if .Historic }}          // Deprecated: {{.Name}} ({{.Numeric}})
    {{- else}}          // {{.Name}} ({{.Numeric}})
    {{- end}}
{{- end}}
)

var currencies = [...]struct {
	// ISO 4217 three-letter alphabetic code
	alpha string
	// ISO 4217 three-digit numeric code
	numeric string
	// Number of decimals to express minor currency unit
	exponent int
	// English name
	name string
}{
{{- range .currencies}}
	{{.Alpha}}: {alpha: "{{.Alpha}}", numeric: "{{.Numeric}}", exponent: {{.Exponent}}, name: {{.Name | printf "%q"}}},
{{- end}}
}

var fromAlpha = map[string]Currency{
{{- range .currencies}}
    "{{.Alpha}}": {{.Alpha}},
{{- end}}
}

var fromNumeric = map[string]Currency{
{{- range .currencies}}
    {{- if not .Historic }} {{/* Ignore historic numbers due to duplicates */}}
    "{{.Numeric}}": {{.Alpha}},
    {{- end}}
{{- end}}
}