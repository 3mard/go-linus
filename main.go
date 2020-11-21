package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"log"

	"math/rand"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	if len(os.Args) != 5 {
		log.Fatal("Usage go-linus <token> <owner> <repo> <prNumber>")
	}

	token := os.Args[1]
	owner := os.Args[2]
	repo := os.Args[3]
	prNumber, _ := strconv.Atoi(os.Args[4])

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomInit := r1.Intn(len(rants))
	body := rants[randomInit]
	client.Issues.CreateComment(context.TODO(), owner, repo, prNumber, &github.IssueComment{
		Body: &body,
	})

}
