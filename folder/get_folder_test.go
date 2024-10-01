package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	example1 := []folder.Folder{
		{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
		{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
		{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
		{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
		{ Name: "echo", Paths: "echo", OrgId: defaultOrdID },
		{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
	}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "Empty folder",
			orgID: defaultOrdID,
			folders: []folder.Folder{},
			want: []folder.Folder{},
		},
		{
			name: "One folder",
			orgID: defaultOrdID,
			folders: []folder.Folder{
				{ Name: "alpha", OrgId: defaultOrdID, Paths: "alpha" },
			},
			want: []folder.Folder{
				{ Name: "alpha", OrgId: defaultOrdID, Paths: "alpha" },
			},
		},
		{
			name: "Many folders",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
				{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
				{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
				{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
				{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
				{ Name: "echo", Paths: "echo", OrgId: defaultOrdID },
			},
		},
		{
			name: "No folders in desired organisation",
			orgID: defaultOrdID,
			folders: []folder.Folder{
				{ Name: "alpha", OrgId: secondaryOrdID, Paths: "alpha" },
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

	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	example1 := []folder.Folder{
		{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
		{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
		{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
		{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
		{ Name: "echo", Paths: "echo", OrgId: defaultOrdID },
		{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
	}

	tests := [...]struct {
		testName    string
		parent	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			testName: "One folder",
			parent: "alpha",
			orgID: defaultOrdID,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: defaultOrdID, Paths: "alpha"},
				{Name: "beta", OrgId: defaultOrdID, Paths: "alpha.beta"},
			},
			want: []folder.Folder{
				{Name: "beta", OrgId: defaultOrdID, Paths: "alpha.beta"},
			},
		},
		{
			testName: "Get children from root folder",
			parent: "alpha",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
				{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
				{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
				{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
			},
		},
		{
			testName: "Get children from inner folder",
			parent: "bravo",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
				{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
			},
		},
		{
			testName: "Get children from leaf folder",
			parent: "charlie",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
			},
		},
		{
			testName: "Get children from root folder with no children",
			parent: "echo",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, _ := f.GetAllChildFolders(tt.orgID, tt.parent)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders_Error(t *testing.T) {
	t.Parallel()

	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	example1 := []folder.Folder{
		{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
		{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
		{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
		{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
		{ Name: "echo", Paths: "echo", OrgId: defaultOrdID },
		{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
	}

	tests := [...]struct {
		testName    string
		parent	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    bool
	}{
		{
			testName: "Folder does not exist",
			parent: "idonotexist",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
		{
			testName: "Folder does not exist in specified organisation",
			parent: "foxtrot",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			_, err := f.GetAllChildFolders(tt.orgID, tt.parent)
			assert.Equal(t, tt.want, err != nil)
		})
	}
}
