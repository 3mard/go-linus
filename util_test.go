package main

import (
	"errors"
	"testing"
)

func TestSplitRepoOwner(t *testing.T) {
	cases := []struct {
		input string
		repo  string
		owner string
		err   error
	}{
		{"3mard/go-linus", "go-linus", "3mard", nil},
		{"3mard/go-linus/test", "go-linus/test", "3mard", nil},
		{"3mard-go-linux", "", "", errors.New("wrong input")},
		{"/3mard-go=linux", "", "", errors.New("wrong input")},
		{"3mard-go=linux/", "", "", errors.New("wrong input")},
	}

	for _, testCase := range cases {
		repo, owner, err := splitRepoOwner(testCase.input)
		if repo != testCase.repo || owner != testCase.owner || (testCase.err != nil && testCase.err.Error() != err.Error()) {
			t.Errorf("expected %v ; got  %s %s %v", testCase, repo, owner, err)
		}

	}
}
