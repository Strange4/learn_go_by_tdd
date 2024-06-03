package blogpost

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func NewPostsFromFS(fileSys fs.FS) ([]Post, error) {
	var posts []Post

	dir, err := fs.ReadDir(fileSys, ".")
	if err != nil {
		return nil, err
	}
	for _, entry := range dir {
		post, err := getPost(fileSys, entry.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSys fs.FS, fileName string) (Post, error) {
	file, err := fileSys.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	return newPost(file)
}

func newPost(reader io.Reader) (Post, error) {
	scanner := bufio.NewScanner(reader)
	const (
		titleSeparator       = "Title: "
		descriptionSeparator = "Description: "
		tagSeparator         = "Tags: "
		bodySeparator        = "---"
	)

	title := readMetaLine(titleSeparator, scanner)
	description := readMetaLine(descriptionSeparator, scanner)
	tags := strings.Split(readMetaLine(tagSeparator, scanner), ", ")

	_ = readMetaLine(bodySeparator, scanner) // ignore the body separator

	// scan the rest
	body := readEntireScanner(scanner)

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

func readMetaLine(metaTagName string, scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), metaTagName)
}

func readEntireScanner(scanner *bufio.Scanner) string {
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}
	return buffer.String()
}
