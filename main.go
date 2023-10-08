package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Content string `json:"content"`
}

type Prompt struct {
	Context  string    `json:"context"`
	Examples []string  `json:"examples"`
	Messages []Message `json:"messages"`
}

type RequestBody struct {
	Model          string  `json:"model"`
	Temperature    float32 `json:"temperature"`
	CandidateCount int     `json:"candidate_count"`
	TopK           int     `json:"top_k"`
	TopP           float32 `json:"top_p"`
	Prompt         Prompt  `json:"prompt"`
}

type Candidate struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type ResponseBody struct {
	Candidates []Candidate `json:"candidates"`
	Messages   []Message   `json:"messages"`
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY environment variable not set")
		return
	}

	url := "https://generativelanguage.googleapis.com/v1beta2/models/chat-bison-001:generateMessage?key=" + apiKey

	fmt.Print("Enter your question: ")
	var input string
	fmt.Scanln(&input)

	messages := []Message{{Content: input}}

	requestBody := &RequestBody{
		Model:          "models/chat-bison-001",
		Temperature:    0.25,
		CandidateCount: 1,
		TopK:           40,
		TopP:           0.95,
		Prompt:         Prompt{Context: "", Examples: []string{}, Messages: messages},
	}

	jsonValue, _ := json.Marshal(requestBody)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := io.ReadAll(resp.Body)
		var responseBody ResponseBody
		json.Unmarshal(data, &responseBody)
		fmt.Println("response: ", responseBody.Candidates[0].Content)
	}
}
