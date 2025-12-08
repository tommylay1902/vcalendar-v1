package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/kelindar/search"
	"github.com/qdrant/go-client/qdrant"
)

func IgnoreAudioWarnings() {
	devNull, _ := os.OpenFile("/dev/null", os.O_WRONLY, 0666)
	syscall.Dup2(int(devNull.Fd()), int(os.Stderr.Fd()))
	devNull.Close()
}

func SetupQdrantClient() *qdrant.Client {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	if err != nil {
		fmt.Println("ERROR SETTING UP QDRANT CLIENT:", err)
		panic(err)
	}

	return client
}

func SetupEmbeddingClient() *search.Vectorizer {
	m, err := search.NewVectorizer("./dist/all-minilm-l6-v2-q8_0.gguf", 1)
	if err != nil {
		fmt.Println("error setting up embedding client:", err)
		// handle error
		panic(err)
	}

	return m
}
