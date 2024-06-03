package blog_renderer

import (
	"embed"
	blogpost "hello/blog"
	"io"
	"text/template"
)

//go:embed "templates/*"
var postTemplates embed.FS

func Render(writer io.Writer, post *blogpost.Post) error {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err = tmpl.Execute(writer, post); err != nil {
		return err
	}

	return nil
}
