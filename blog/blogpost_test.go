package blogpost_test

import (
	"hello/assertions"
	blogpost "hello/blog"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	const (
		file1 = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
BBL Drizzy`
		file2 = `Title: Second post
Description: Description 2
Tags: rust, bevy
---
TanTan`
	)
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(file1)},
		"hello world2.md": {Data: []byte(file2)},
	}

	posts, err := blogpost.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	got := posts[0]
	want := blogpost.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body:        "BBL Drizzy",
	}

	assertions.AssertEqual(t, got, want)
}
