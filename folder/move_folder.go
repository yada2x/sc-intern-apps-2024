package folder

import (
	"errors"

	"strings"
)

// Move a source folder and its children into another folder
// Input: source folder name, destination folder name
// Output: slice of folders, IO errors
// Errors: Moving folders to a different organisation, moving a folder to its child
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Get indices of folders of interest
	start, dest, err := f.getFolderIndices(name, dst)
	if err != nil {
		return nil, err
	}

	nodeToMove := f.folders[start]
	destination := f.folders[dest]

	// Handle cases where folders are in different organisations or where one is a child of the other
	if nodeToMove.OrgId != destination.OrgId {
		return nil, errors.New("cannot move a folder to a different organisation")
	} else if strings.HasPrefix(destination.Paths, nodeToMove.Paths+".") {
		return nil, errors.New("cannot move folder to a child of itself")
	}

	// Update child nodes with new paths
	newPath := destination.Paths + "." + nodeToMove.Name
	oldPath := nodeToMove.Paths + "."
	f.folders[start].Paths = newPath
	f.updateFolderPaths(oldPath, newPath)

	return f.folders, nil
}

// Finds and returns the indices of the source and destination folder if valid
// Input: name of source folder, name of destination folder
// Output: index of source folder, index of destination folder, error
// Errors: Non-existent source folder, non-existent destination folder, moving a folder to itself
func (f *driver) getFolderIndices(name string, dst string) (int, int, error) {
	// Try to find corresponding folders and get their indices
	start := -1
	dest := -1
	for i := range f.folders {
		if f.folders[i].Name == name {
			start = i
		} 
		if f.folders[i].Name == dst {
			dest = i
		}
	}

	// Handle errors for non-existent folders or moving a folder to itself
	if start == -1 {
		return -1, -1, errors.New("source folder does not exist")
	} else if dest == -1 {
		return -1, -1, errors.New("destination folder does not exist")
	} else if start == dest {
		return -1, -1, errors.New("cannot move a folder to itself")
	}

	return start, dest, nil
}

// Update the paths of folders that contain the old path with the new path
// Input: original path of parent, new path of parent
// Output: None
func (f *driver) updateFolderPaths(oldPath string, newPath string) {
	// For each child part of an input path, change their path to a new path
	for i := range f.folders {
		if strings.HasPrefix(f.folders[i].Paths, oldPath) {
			leftover := strings.TrimPrefix(f.folders[i].Paths, oldPath)
			f.folders[i].Paths = newPath + "." + leftover
		}
	}
}
