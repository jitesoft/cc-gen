package template

import (
	"fmt"
	"github.com/jitesoft/cc-gen/gitwrapper/conventional"
	"io"
	"text/template"
)

type TagData struct {
	Commits   map[string][]*conventional.Commit
	Name       string

}

type Data struct {
	CommitUri string
	Extra     string
	Tags      []*TagData
}

var templates = map[string]*template.Template{}

func init() {
	templates["default"] = template.Must(template.New("markdown").Parse(`{{ range $tag := .Tags -}}
### {{ .Name }}
{{ range $type, $commits := .Commits }}
#### {{ $type }}
{{ range $commits }}
  * {{ if $.CommitUri }}[{{- slice .Hash 0 8 }}]({{ $.CommitUri }}{{ .Hash }}){{ else }}[{{- slice .Hash 0 8 }}]{{ end }} {{ .Header }} <small>By: {{ .Author }} @ {{ .Time }}</small>
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
