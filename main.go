package main

import (
	blogpost "hello/blog"
	"log"
	"os"
)

func main() {
	posts, err := blogpost.NewPostsFromFS(os.DirFS("blog/posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
