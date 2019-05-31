// cli 负责克隆并解析博客远程仓库数据到本地
package cli

import (
	"fmt"
	"os"
	"os/exec"
)

type Gitter struct {
	Repo string // 远程仓库地址
	Dir  string // 本地仓库地址
}

// NewGitter 创建一个Gitter实例
func NewGitter(repo, dir string) *Gitter {
	return &Gitter{
		Repo: repo,
		Dir:  dir,
	}
}

// Clone 克隆远程仓库g.Repo中的文件到本地文件夹g.Dir(目前是全量更新)
func (g *Gitter) Clone() {
	g.clearDir()
	checkerror(err)
	err = g.cloneRepo()
	checkerror(err)
}

// clear 清空g.Dir文件夹
func (g *Gitter) clearDir() {
	err := os.RemoveAll(g.Dir)
	checkerror(err)
}

// cloneRepo 克隆g.Repo文件到本地
func (g *Gitter) cloneRepo() error {
	cmd := exec.Command("git", "clone", g.Repo, g.Dir)
	err := cmd.Run()
	checkerror(err)
}
