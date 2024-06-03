package blog_renderer

import (
	"embed"
	blogpost "hello/blog"
	"io"
	"text/template"
)

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	tmpl *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{tmpl}, nil
}

func (p *PostRenderer) Render(writer io.Writer, post *blogpost.Post) error {

	if err := p.tmpl.ExecuteTemplate(writer, "postTemplate.gohtml", post); err != nil {
		return err
	}

	return nil
}
