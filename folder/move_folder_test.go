package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	example1 := []folder.Folder{
		{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
		{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
		{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
		{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
		{ Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID },
		{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
		{ Name: "golf", Paths: "golf", OrgId: defaultOrdID },
	}

	tests := [...]struct {
		testName    string
		start		string
		destination	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			testName: "Example 1",
			start: "bravo",
			destination: "delta",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
				{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
				{ Name: "bravo", Paths: "alpha.delta.bravo", OrgId: defaultOrdID },
				{ Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: defaultOrdID },
				{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
				{ Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID },
				{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID },
				{ Name: "golf", Paths: "golf", OrgId: defaultOrdID },
			},
		},
		{
			testName: "Example 2",
			start: "bravo",
			destination: "golf",
			orgID: defaultOrdID,
			folders: example1,
			want: []folder.Folder{
				{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
				{ Name: "bravo", Paths: "golf.bravo", OrgId: defaultOrdID },
				{ Name: "charlie", Paths: "golf.bravo.charlie", OrgId: defaultOrdID },
				{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
				{ Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID },
				{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID },
				{ Name: "golf", Paths: "golf", OrgId: defaultOrdID },
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, _ := f.MoveFolder(tt.start, tt.destination)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_MoveFolder_Error(t *testing.T) {
	t.Parallel()

	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	example1 := []folder.Folder{
		{ Name: "alpha", Paths: "alpha", OrgId: defaultOrdID },
		{ Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID },
		{ Name: "charlie", Paths : "alpha.bravo.charlie", OrgId: defaultOrdID },
		{ Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID },
		{ Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID },
		{ Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
		{ Name: "golf", Paths: "golf", OrgId: defaultOrdID },
	}

	tests := [...]struct {
		testName    string
		start		string
		destination	string
		orgID   uuid.UUID
		folders []folder.Folder
		want    bool
	}{
		{
			testName: "Example 1",
			start: "bravo",
			destination: "charlie",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
		{
			testName: "Example 2",
			start: "bravo",
			destination: "bravo",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
		{
			testName: "Example 3",
			start: "bravo",
			destination: "foxtrot",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
		{
			testName: "Example 4",
			start: "invalid_folder",
			destination: "delta",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
		{
			testName: "Example 5",
			start: "bravo",
			destination: "invalid_folder",
			orgID: defaultOrdID,
			folders: example1,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			_, err := f.MoveFolder(tt.start, tt.destination)
			assert.Equal(t, tt.want, err != nil)
		})
	}
}