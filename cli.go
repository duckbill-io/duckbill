package main

import (
	"github.com/duckbill-io/duckbill/cli"
)

const (
//	RepoUrl = "git@github.com:CaryWill/blog-docs.git"
	RepoDir = "_posts"
	DataDir = "_data"
)

func main() {
//	g := cli.NewGitter(RepoUrl, RepoDir)
//	g.Clone()
	p := cli.NewParser(RepoDir, DataDir)
	p.Run()
}
