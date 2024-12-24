package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/approvals/go-approval-tests"
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
