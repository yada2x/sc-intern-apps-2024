package folder

import (
	"errors"

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

// Retrieves the a slice of the children of a folder specified by organisation ID and name
// Assumes unique folder names within an organisations
// If this assumption is false, if there are many folders with the same name, 
// the first folder with the name will be selected.
// Input: organisation ID, folder name
// Output: slice of child folders, IO errors
// Errors: Invalid folder
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folders := f.GetFoldersByOrgID(orgID)

	// Find the desired folder
	var path string
	for _, folder := range folders {
		if folder.Name == name {
			path = folder.Paths
			break
		}
	}

	// Not found case: return nil if folder is not found
	if path == "" {
		return nil, errors.New("folder does not exist in the specified organisation")
	}

	// Find all children and append them to a result slice
	children := []Folder{}
	for _, folder := range folders {
		// Check if current folder has the parent path as a prefix of their own path
		if strings.HasPrefix(folder.Paths, path+".") {
			children = append(children, folder)
		}
	}
	return children, nil
}
