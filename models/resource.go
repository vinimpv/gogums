package models

type Resources struct {
	imgsPath  string
	postsPath string
	Pages     []Page
	Posts     []Posts
	Images    []Images
}
type Page struct{}
type Images struct{}
type Posts struct{}
