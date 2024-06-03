package blog_renderer_test

import (
	"bytes"
	"hello/assertions"
	blogpost "hello/blog"
	"hello/blog_renderer"
	"strings"
	"testing"
)

const wantedHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <title>My amazing blog!</title>
    <meta charset="UTF-8"/>
    <meta name="description" content="oh no i am under the water, please help me" lang="en"/>
</head>
<body>
<nav role="navigation">
    <div>
        <h1>Strange's blog</h1>
        <ul>
            <li><a href="/">home</a></li>
            <li><a href="about">about</a></li>
            <li><a href="archive">archive</a></li>
        </ul>
    </div>
</nav>
<main>

<h1>How to make real time graphics renderer</h1>
<p>this is not a real advice</p>
Tags:
<ul>
    <li>gpu</li>
    <li>shader</li>
    <li>BRDF</li>
</ul>

</main>
<footer>
    <ul>
        <li><a href="https://github.com/Strange4">GitHub</a></li>
    </ul>
</footer>
</body>
</html>`

func TestRender(t *testing.T) {
	post := blogpost.Post{
		Title:       "How to make real time graphics renderer",
		Body:        "Well, it's kind of hard. But start with learning a book",
		Description: "this is not a real advice",
		Tags:        []string{"gpu", "shader", "BRDF"},
	}

	buf := bytes.Buffer{}
	renderer, err := blog_renderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	err = renderer.Render(&buf, &post)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.ReplaceAll(buf.String(), "\r", "")
	want := wantedHTML
	assertions.AssertString(t, got, want)
}
