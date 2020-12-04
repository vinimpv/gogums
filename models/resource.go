package models

type SiteResources struct {
	ImagesFolder         string            `json:"images_folder"`
	ContentFolder        string            `json:"content_folder"`
	PreviewCommand       string            `json:"preview_command"`
	ResourcesGroups      []*ResourcesGroup `json:"resources_groups"`
	ResourcesGroupsNames []string          `json:"groups"`
}

// ResourcesGroup represents a folder on your "content" folder
// it contains the definition for an _index.md and for a possible
// list of definitions (for a list of posts for an example).
type ResourcesGroup struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Index       *Resource   `json:"index"`
	List        []*Resource `json:"list"`

	MainTemplate *ResourceTemplate `json:"index_template"`
	ListTemplate *ResourceTemplate `json:"list_template"`
}

// ResourcesDefinition represents a md file with multiple Resources
type Resource struct {
	Type        string  `json:"type"` //list or single
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Path        string  `json:"path"` //can be a path to a single file or to a folder, in case of list type, the file path is Path/slugfied-name
	Content     string  `json:"content"`
	Fields      []Field `json:"fields"`
}

// Resource represents a single entry on a md file ex {type: text, value: My site}
type Field struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Fixed       bool   `json:"fixed"`

	MinLength int `json:"min_length"`
	MaxLength int `json:"max_length"`

	MinItens int `json:"min_itens"`
	MaxItens int `json:"max_itens"`

	Type  string `json:"type"`  //should be enum
	Value string `json:"value"` //in case of lists and content, this is going to be parsed
}
type ResourceTemplate struct {
	Description    string  `json:"description"`
	Fields         []Field `json:"fields"`
	ContentEnabled bool    `json:"content_enabled"`
}
