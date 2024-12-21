package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/IchiThe2nd/GoTDD2/blogposts"
	//blogposts "github.com/IchiThe2nd/GoTDD2/blogposts
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello worl.md": {Data: []byte("hi")},
		"hello-world2":  {Data: []byte("hola")},
	}
	posts := blogposts.NewPostFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts,wanred %d posts", len(posts), len(fs))
	}
}
