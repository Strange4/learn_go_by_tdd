package blog_renderer_test

import (
	"bytes"
	blogpost "hello/blog"
	"hello/blog_renderer"
	"strings"
	"testing"
)

const wantedHTML = `<h1>How to make real time graphics renderer</h1>
<p>this is not a real advice</p>
Tags:
<ul>
    <li>gpu</li>
    <li>shader</li>
    <li>BRDF</li>
</ul>`

func TestRender(t *testing.T) {
	post := blogpost.Post{
		Title:       "How to make real time graphics renderer",
		Body:        "Well, it's kind of hard. But start with learning a book",
		Description: "this is not a real advice",
		Tags:        []string{"gpu", "shader", "BRDF"},
	}

	buf := bytes.Buffer{}
	err := blog_renderer.Render(&buf, &post)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.ReplaceAll(buf.String(), "\r", "")
	want := wantedHTML
	if want != got {
		t.Errorf("wanted post '%q' but only got '%q'", want, got)
	}
}
