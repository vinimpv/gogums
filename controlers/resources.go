package controlers

type ResourcesDefinitionSerializer struct{}

type ResourceSerializer struct{}

type ResourcesGroupSerializer struct {
	Name        string
	Description string
	FolderPath  string
	Main        *ResourcesDefinitionSerializer
	List        []ResourcesDefinitionSerializer
}

// opens up site repo dir
// parses the templates
// iterate through the names, read files in folder convert to Resource
