package folder

import (
	"errors"

	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	if len(f.folders) == 0 {
		return nil, errors.New("provided folder is empty")
	}

	start, dest, err := f.getFolderIndices(name, dst)
	if (err != nil) {
		return nil, err
	}

	nodeToMove := f.folders[start]
	destination := f.folders[dest]

	if nodeToMove.OrgId != destination.OrgId {
		return nil, errors.New("cannot move a folder to a different organisation")
	} else if strings.HasPrefix(destination.Paths, nodeToMove.Paths + ".") {
		return nil, errors.New("cannot move folder to a child of itself")
	}

	newPath := destination.Paths + "." + nodeToMove.Name
	oldPath := nodeToMove.Paths + "."
	f.folders[start].Paths = newPath
	
	f.updateFolderPaths(oldPath, newPath)

	return f.folders, nil
}

func (f* driver) getFolderIndices(name string, dst string) (int, int, error) {
	start := -1
	dest := 1
	for i := range f.folders {
		if f.folders[i].Name == name {
			start = i
		} else if f.folders[i].Name == dst {
			dest = i
		}
	}

	if start == -1 {
		return -1, -1, errors.New("source folder does not exist")
	} else if dest == -1 {
		return -1, -1, errors.New("destination folder does not exist")
	} else if start == dest {
		return -1, -1, errors.New("cannot move a folder to itself")
	}

	return start, dest, nil
}

func (f* driver) updateFolderPaths(oldPath string, newPath string) {
	for i := range f.folders {
		if strings.HasPrefix(f.folders[i].Paths, oldPath) {
			leftover := strings.TrimPrefix(f.folders[i].Paths, oldPath)
			f.folders[i].Paths = newPath + "." + leftover
		}
	}
}