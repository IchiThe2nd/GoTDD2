package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	//"github.com/approvals/go-approval-test"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

type PostRender struct {
	templ *template.Template
}

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
