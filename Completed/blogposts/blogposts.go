package blogposts

import (
	"io/fs"
)

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err //clarify failing if one file fails or not
		}
		posts = append(posts, post)

	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName) //opened a file
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile) //
}
