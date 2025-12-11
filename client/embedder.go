package client

import (
	"fmt"

	"github.com/kelindar/search"
)

func SetupEmbeddingClient() *search.Vectorizer {
	m, err := search.NewVectorizer("./dist/all-minilm-l6-v2-q8_0.gguf", 1)
	if err != nil {
		fmt.Println("error setting up embedding client:", err)
		// handle error
		panic(err)
	}

	return m
}
