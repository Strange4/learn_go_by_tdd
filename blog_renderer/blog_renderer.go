package blog_renderer

import (
	"fmt"
	blogpost "hello/blog"
	"io"
)

func Render(writer io.Writer, post *blogpost.Post) error {
	_, err := fmt.Fprintf(writer, `<h1>%s</h1>
<p>%s</p>
Tags:
<ul>
`, post.Title, post.Description)

	if err != nil {
		return err
	}

	for _, tag := range post.Tags {
		_, err := fmt.Fprintf(writer, "\t<li>%s</li>\n", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(writer, "</ul>")
	if err != nil {
		return err
	}
	return nil
}
