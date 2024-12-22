package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/IchiThe2nd/GoTDD2/blogposts"
	//blogposts "github.com/IchiThe2nd/GoTDD2/blogposts
)

type StubFailinngFS struct {
}

func (s StubFailinngFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh nos I always fail")
}

func TestNewBlogPost(t *testing.T) {
	const ( // make some semi constants
		firstBody = `Title: Post 1
Description : Description 1`
		secondBody = `Title: Post 2
Description : Description 2`
	)

	fs := fstest.MapFS{ //make  a filesystem
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostFromFS(fs) //pull post from filesystm
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts,wanted %d posts", len(posts), len(fs))
	}
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
	})
}

///assertiions

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v,wanted %+v", got, want)

	}
}
