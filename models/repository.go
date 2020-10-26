package models

// this is going to be a model
// encrypt the password and save the key in plain text
// every site will have a repository

type Repository struct {
	url     string `db:"url"`      // repo url
	keyPath string `db:"key_path"` //ssh key as plain text
	keyPass string `db:"key_pass"` //encrypted ssh key pass
	siteId  int64  `db:"site_id"`
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
