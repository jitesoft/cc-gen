package template

import (
	"fmt"
	"github.com/jitesoft/cc-gen/gitwrapper/conventional"
	"io"
	"text/template"
)

// Data contains data which is used to render the template.
type Data struct {
	Commits   map[string][]*conventional.Commit
	Tag       string
	CommitUri string
	Extra     string
}

var templates = map[string]*template.Template{}

func init() {
	templates["default"] = template.Must(template.New("markdown").Parse(`# {{ .Tag }}
{{ range $key, $value := .Commits }}
{{- if $value }}
## {{ $key }}
{{ range $value }}
  * {{ if $.CommitUri }}[{{- .Hash }}]({{ $.CommitUri }}{{ .Hash }}){{ else }}[{{- .Hash }}]{{ end }} {{ .Header }} <small>By: {{ .Author }} @ {{ .Time }}</small>
{{- end }}
{{ end }}
{{- end }}
{{ .Extra }}
`))
}

func AddTemplate(temp string, name string) {
	templates[name] = template.Must(template.New(name).Parse(temp))
}

func RenderTemplate (data *Data, template string, output io.Writer) error {
	t, exists := templates[template]
	if !exists {
		return fmt.Errorf("failed to render template, template %s does not exist", template)
	}

	return t.Execute(output, data)
}
