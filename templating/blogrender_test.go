package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	//"github.com/approvals/go-approval-tests"
	"github.com/IchiThe2nd/GoTDD2/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "this is a post",
			Description: "this is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("it converts a single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "this is a post",
			Description: "this is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogrenderer.Render(io.Discard, aPost)
	}

}
