package folder

import (
	"github.com/gofrs/uuid"

	"strings"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.GetFoldersByOrgID(orgID)

	if len(folders) == 0 { // Empty case: return nil if folder is empty
		return nil
	}

	// Find the desired folder
	var path string
	for _, folder := range folders {
		if folder.Name == name {
			path = folder.Paths
			break
		}
	}

	// Not found case: return nil if folder is not found (maybe throw an error?)
	if path == "" {
		return nil
	}

	// Iterate through all folders, if they have path as a prefix, they are a child, append them
	// O(n^2) here, maybe try preprocessing all data into like a trie or smth
	children := []Folder{}
	for _, folder := range folders {
		if strings.HasPrefix(folder.Paths, path + ".") && folder.Name != name {
			children = append(children, folder)
		}
	}
	return children
}
