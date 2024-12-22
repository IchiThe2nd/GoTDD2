package main

import (
	"log"
	"os"

	"github.com/IchiThe2nd/GoTDD2/blogposts"
)

func main() {
	posts, err := blogposts.NewPostFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
