package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// this is going to be a model
// encrypt the password and save the key in plain text
// every site will have a repository

type Repository struct {
	Id     int64  `db:"id" json:"id"`
	Url    string `db:"url" json:"url"` // repo url
	Key    string `db:"key" json:"key"` //ssh key as plain text
	SiteId int64  `db:"site_id" json:"site_id"`
}

// commits all changes
func (r *Repository) commit() error {
	return nil
}

// pulls from master
func (r *Repository) pull() error {
	return nil
}

// pushes changes to master
func (r *Repository) pushOrigin() error {
	return nil
}

func (r *Repository) clone() error {
	//TODO check errors
	keyPath, _ := r.getCreateKeyFile()
	cmd := fmt.Sprintf("ssh-agent bash -c 'ssh-add %s; git clone %s'", keyPath, r.Url)
	out, _ := exec.Command(cmd).Output()
	fmt.Println(out)
	return nil
}

// Returns the path to a key file that will be used to
// perform remote operations on the repo eg: clone, push, pull
func (r *Repository) getCreateKeyFile() (string, error) {
	// by default keys are saved in /tmp/keys/{repoId}
	fp := "/tmp/keys/" + string(r.Id)
	if _, err := os.Stat("/tmp/keys/"); err == nil {
		// create file if it doesnt exists
		d := []byte(r.Key)
		err := ioutil.WriteFile(fp, d, 0777)
		if err != nil {
			return "", err
		}
	}
	return fp, nil
}
