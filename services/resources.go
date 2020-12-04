package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"vinimpv/gogums/models"
)

func parseListFile(filePath string, rg *models.ResourcesGroup) error {
	//fields := rg.ListTemplate.Fields
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	text := []string{}
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	// index where content starts
	var contentIndex int
	fieldMap := make(map[string]string)

	// ignores first ---
	for i, line := range text[1:] {
		// ignores comments
		if string(line[0]) == "#" {
			continue
		}
		if line == "---" {
			contentIndex = i + 1
			break
		}
		arr := strings.SplitN(line, ": ", 2)
		fieldMap[arr[0]] = arr[1]
	}
	content := strings.Join(text[contentIndex+1:], "\n")
	listResource := &models.Resource{Name: file.Name(), Type: "list", Path: filePath, Content: content}
	for _, f := range rg.ListTemplate.Fields {
		// removes double quotes
		f.Value = strings.ReplaceAll(fieldMap[f.Name], `"`, "")
		listResource.Fields = append(listResource.Fields, f)
	}
	rg.List = append(rg.List, listResource)

	return nil

}

// opens up site repo dir
// parses the templates
// iterate through the names, read files in folder convert to Resource
func parseGroup(group string, sr *models.SiteResources, repoDir string) {
	groupFile, _ := os.Open(fmt.Sprintf("%s/.gogums/%s.json", repoDir, group))
	groupBytes, _ := ioutil.ReadAll(groupFile)
	rg := &models.ResourcesGroup{}
	json.Unmarshal(groupBytes, rg)
	sr.ResourcesGroups = append(sr.ResourcesGroups, rg)
	groupPath := fmt.Sprintf("%s/%s/%s", repoDir, sr.ContentFolder, group)
	files, _ := ioutil.ReadDir(groupPath)
	for _, file := range files {
		switch file.Name() {
		case "_index.md":
			// TODO case for _index templates
			continue
		default:
			filePath := fmt.Sprintf("%s/%s", groupPath, file.Name())
			parseListFile(filePath, rg)

		}

	}

}

func ParseResources(site *models.Site) (*models.SiteResources, error) {
	// read base, get  the list of resources and parse each resource template
	baseFile, err := os.Open(fmt.Sprintf("%s/.gogums/base.json", site.Repository.Dir()))
	if err != nil {
		return nil, err
	}
	defer baseFile.Close()
	baseBytes, _ := ioutil.ReadAll(baseFile)
	siteResources := &models.SiteResources{}
	json.Unmarshal(baseBytes, siteResources)
	for _, group := range siteResources.ResourcesGroupsNames {
		parseGroup(group, siteResources, site.Repository.Dir())
	}
	return siteResources, nil
}
