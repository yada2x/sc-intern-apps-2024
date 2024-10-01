package folder

import (
	"errors"

	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	// find node to move, find destination
	// get old path of node, change it to destination path + "." + node.name
	// for children nodes with prefix old path, change to new path + "." + child.name
	folders := f.folders

	if len(folders) == 0 {
		return nil, errors.New("provided folder is empty")
	}

	var start int
	var dest int
	for i := range folders {
		if folders[i].Name == name {
			start = i
		} else if folders[i].Name == dst {
			dest = i
		}
	}

	nodeToMove := folders[start]
	destination := folders[dest]

	if nodeToMove.Name == "" {
		return nil, errors.New("invalid start node")
	} else if destination.Name == "" {
		return nil, errors.New("invalid destination node")
	} else if nodeToMove.OrgId != destination.OrgId {
		return nil, errors.New("start and destination nodes belong to different organisations")
	} else if strings.HasPrefix(destination.Paths, nodeToMove.Paths) {
		return nil, errors.New("cannot move node to its child")
	}
	
	newPath := destination.Paths + "." + nodeToMove.Name
	prefix := nodeToMove.Paths + "."
	nodeToMove.Paths = newPath

	for i := range folders {
		if strings.HasPrefix(folders[i].Paths, prefix) {
			leftover := strings.TrimPrefix(folders[i].Paths, prefix)
			folders[i].Paths = newPath + "." + leftover
		}
	}

	return folders, nil
}
