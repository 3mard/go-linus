package main

import (
	"errors"
	"strings"
)

func splitRepoOwner(repoOwner string) (repo, owner string, err error) {
	index := strings.Index(repoOwner, "/")
	if index <= 0 || index == len(repoOwner)-1 {
		return "", "", errors.New("wrong input")
	}

	return repoOwner[index+1:], repoOwner[0:index], nil

}
