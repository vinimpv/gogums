package models

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

type Repository struct {
	Id     int64  `db:"id" json:"id,ommitempty"`
	Url    string `db:"url" json:"url"` // repo url
	Key    string `db:"key" json:"key"` //ssh key as plain text
	SiteId int64  `db:"site_id" json:"site_id"`
}

// commits all changes
func (r *Repository) commit() error {
	return nil
}
func (r *Repository) Dir() string {
	return fmt.Sprintf("repos/%v", r.Id)
}

// pulls from master
func (r *Repository) pull() error {
	ss, _ := ssh.ParsePrivateKey([]byte(r.Key))
	auth := &gitssh.PublicKeys{User: "git", Signer: ss}
	repo, _ := git.PlainOpen(r.Dir())
	w, _ := repo.Worktree()
	err := w.Pull(&git.PullOptions{RemoteName: "origin", Auth: auth, Progress: os.Stdout})
	if err != nil {
		panic(err)
	}
	return nil

}

// pushes changes to master
func (r *Repository) pushOrigin() error {
	return nil
}

func (r *Repository) Clone() error {
	ss, _ := ssh.ParsePrivateKey([]byte(r.Key))
	auth := &gitssh.PublicKeys{User: "git", Signer: ss}
	// dir
	repo, _ := git.PlainClone(r.Dir(), false, &git.CloneOptions{
		URL:               r.Url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})

	fmt.Println(repo)
	return nil
}
