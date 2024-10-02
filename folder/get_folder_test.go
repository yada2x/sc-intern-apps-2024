package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// Organisation IDs for testing
	defaultOrgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrgID := uuid.Must(uuid.NewV4())

	// Testing data
	example1 := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
		{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
		{Name: "echo", Paths: "echo", OrgId: defaultOrgID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrgID},
	}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name:    "Empty folder",
			orgID:   defaultOrgID,
			folders: []folder.Folder{},
			want:    []folder.Folder{},
		},
		{
			name:  "One folder",
			orgID: defaultOrgID,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: defaultOrgID, Paths: "alpha"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: defaultOrgID, Paths: "alpha"},
			},
		},
		{
			name:    "Multiple folders",
			orgID:   defaultOrgID,
			folders: example1,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
				{Name: "echo", Paths: "echo", OrgId: defaultOrgID},
			},
		},
		{
			name:  "Folders in a different empty organisation",
			orgID: secondaryOrgID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
				{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
				{Name: "echo", Paths: "echo", OrgId: defaultOrgID},
			},
			want: []folder.Folder{},
		},
		{
			name:    "Folders from different organisations",
			orgID:   secondaryOrgID,
			folders: example1,
			want: []folder.Folder{
				{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrgID},
			},
		},
		{
			name:  "No folders in desired organisation",
			orgID: defaultOrgID,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: secondaryOrgID, Paths: "alpha"},
			},
			want: []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	// Organisation IDs for testing
	defaultOrgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrgID := uuid.Must(uuid.NewV4())

	// Testing data
	example1 := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
		{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
		{Name: "echo", Paths: "echo", OrgId: defaultOrgID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrgID},
	}

	tests := [...]struct {
		testName string
		parent   string
		orgID    uuid.UUID
		folders  []folder.Folder
		want     []folder.Folder
	}{
		{
			testName: "Single folder",
			parent:   "alpha",
			orgID:    defaultOrgID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
			},
			want: []folder.Folder{},
		},
		{
			testName: "Same folder names, different organisations",
			parent:   "alpha",
			orgID:    defaultOrgID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
				{Name: "beta", Paths: "alpha.beta", OrgId: defaultOrgID},
				{Name: "alpha", Paths: "alpha", OrgId: secondaryOrgID},
				{Name: "echo", Paths: "alpha.echo", OrgId: secondaryOrgID},
			},
			want: []folder.Folder{
				{Name: "beta", Paths: "alpha.beta", OrgId: defaultOrgID},
			},
		},
		{
			testName: "Similar root folder names",
			parent:   "alpha",
			orgID:    defaultOrgID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
				{Name: "beta", Paths: "alpha.beta", OrgId: defaultOrgID},
				{Name: "alphaa", Paths: "alphaa", OrgId: defaultOrgID},
				{Name: "echo", Paths: "alphaa.echo", OrgId: defaultOrgID},
			},
			want: []folder.Folder{
				{Name: "beta", Paths: "alpha.beta", OrgId: defaultOrgID},
			},
		},
		{
			testName: "Similar child folder names",
			parent:   "alpha",
			orgID:    defaultOrgID,
			folders: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
				{Name: "beta", Paths: "alpha.beta", OrgId: defaultOrgID},
				{Name: "betaa", Paths: "alpha.betaa", OrgId: defaultOrgID},
			},
			want: []folder.Folder{
				{Name: "beta", Paths: "alpha.beta", OrgId: defaultOrgID},
				{Name: "betaa", Paths: "alpha.betaa", OrgId: defaultOrgID},
			},
		},
		{
			testName: "Example 1: Get children from root folder",
			parent:   "alpha",
			orgID:    defaultOrgID,
			folders:  example1,
			want: []folder.Folder{
				{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
				{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
			},
		},
		{
			testName: "Example 2: Get children from inner folder",
			parent:   "bravo",
			orgID:    defaultOrgID,
			folders:  example1,
			want: []folder.Folder{
				{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
			},
		},
		{
			testName: "Example 3: Get children from leaf folder",
			parent:   "charlie",
			orgID:    defaultOrgID,
			folders:  example1,
			want:     []folder.Folder{},
		},
		{
			testName: "Example 4: Get children from folder with no children",
			parent:   "echo",
			orgID:    defaultOrgID,
			folders:  example1,
			want:     []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.parent)
			assert.Equal(t, tt.want, get)
			assert.ErrorIs(t, err, nil)
		})
	}
}

func Test_folder_GetAllChildFolders_Error(t *testing.T) {
	t.Parallel()

	defaultOrgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrgID := uuid.Must(uuid.NewV4())

	example1 := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: defaultOrgID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrgID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrgID},
		{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrgID},
		{Name: "echo", Paths: "echo", OrgId: defaultOrgID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrgID},
	}

	tests := [...]struct {
		testName      string
		parent        string
		orgID         uuid.UUID
		folders       []folder.Folder
		want          string
	}{
		{
			testName:      "Empty list",
			parent:        "alpha",
			orgID:         defaultOrgID,
			folders:       []folder.Folder{},
			want: "folder does not exist in the specified organisation",
		},
		{
			testName:      "Example 5: Folder does not exist",
			parent:        "invalid_folder",
			orgID:         defaultOrgID,
			folders:       example1,
			want: "folder does not exist in the specified organisation",
		},
		{
			testName:      "Example 6: Folder does not exist in specified organisation",
			parent:        "foxtrot",
			orgID:         defaultOrgID,
			folders:       example1,
			want: "folder does not exist in the specified organisation",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			_, err := f.GetAllChildFolders(tt.orgID, tt.parent)
			assert.ErrorContains(t, err, tt.want)
		})
	}
}
