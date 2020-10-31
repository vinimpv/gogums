package models

// ResourcesGroup represents a folder on your "content" folder
// it contains the definition for an _index.md and for a possible
// list of definitions (for a list of posts for an example).
type ResourcesGroup struct {
	Name        string
	Description string
	FolderPath  string
	Main        *ResourcesDefinition
	List        []ResourcesDefinition
}

// ResourcesDefinition represents a md file with multiple Resources
type ResourcesDefinition struct {
	Type        string //list or single
	Name        string
	Description string
	Path        string //can be a path to a single file or to a folder, in case of list type, the file path is Path/slugfied-name
	Resouces    []Resource
}

// Resource represents a single entry on a md file ex {type: text, value: My site}
type Resource struct {
	Name        string
	Description string

	Type  string
	Value string //in case of lists and content, this is going to be parsed
}
