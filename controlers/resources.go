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
