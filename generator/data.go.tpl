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
	// ISO 4217 alphabetic code
	alpha string
	// ISO 4217 numeric code
	numeric string
	// Number of decimals
	exponent int
	// English name
	name string
}{
{{- range .currencies}}
	{{.Alpha}}: {alpha: "{{.Alpha}}", numeric: "{{.Numeric}}", exponent: {{.Exponent}}, name: {{.Name | printf "%q"}}},
{{- end}}
}
