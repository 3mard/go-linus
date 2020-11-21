package main

import (
	"context"
	"fmt"
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

	fmt.Println(os.Environ())

	token := os.Args[1]

	r := os.Args[3]
	repo, owner, err := splitRepoOwner(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Args)
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
	fmt.Println(client.Issues.CreateComment(context.TODO(), owner, repo, prNumber, &github.IssueComment{
		Body: &body,
	}))

}
