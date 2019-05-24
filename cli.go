package main

import (
	"github.com/duckbill-io/duckbill/cli"
)

const (
	RepoUrl = "git@github.com:duckbill-io/blog-docs.git"
	RepoDir = "_posts"
	DataDir = "_data"
)

func main() {
	g := cli.NewGitter(RepoUrl, RepoDir)
	g.Clone()
	p := cli.NewParser(RepoDir, DataDir)
	p.Fire()
}
