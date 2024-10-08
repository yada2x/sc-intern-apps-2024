package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	// Organisation IDs for testing
	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	// Testing data
	example1 := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: defaultOrdID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrdID},
		{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
		{Name: "golf", Paths: "golf", OrgId: defaultOrdID},
	}

	tests := [...]struct {
		testName    string
		start       string
		destination string
		orgID       uuid.UUID
		folders     []folder.Folder
		want        []folder.Folder
	}{
		{
			testName:    "Example 1: Move inner folder to another inner folder",
			start:       "bravo",
			destination: "delta",
			orgID:       defaultOrdID,
			folders:     example1,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrdID},
				{Name: "bravo", Paths: "alpha.delta.bravo", OrgId: defaultOrdID},
				{Name: "charlie", Paths: "alpha.delta.bravo.charlie", OrgId: defaultOrdID},
				{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
				{Name: "golf", Paths: "golf", OrgId: defaultOrdID},
			},
		},
		{
			testName:    "Example 2: Move inner folder to folder with no children",
			start:       "bravo",
			destination: "golf",
			orgID:       defaultOrdID,
			folders:     example1,
			want: []folder.Folder{
				{Name: "alpha", Paths: "alpha", OrgId: defaultOrdID},
				{Name: "bravo", Paths: "golf.bravo", OrgId: defaultOrdID},
				{Name: "charlie", Paths: "golf.bravo.charlie", OrgId: defaultOrdID},
				{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID},
				{Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID},
				{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
				{Name: "golf", Paths: "golf", OrgId: defaultOrdID},
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

	// Organisation IDs for testing
	defaultOrdID := uuid.FromStringOrNil(folder.DefaultOrgID)
	secondaryOrdID := uuid.Must(uuid.NewV4())

	// Testing data
	example1 := []folder.Folder{
		{Name: "alpha", Paths: "alpha", OrgId: defaultOrdID},
		{Name: "bravo", Paths: "alpha.bravo", OrgId: defaultOrdID},
		{Name: "charlie", Paths: "alpha.bravo.charlie", OrgId: defaultOrdID},
		{Name: "delta", Paths: "alpha.delta", OrgId: defaultOrdID},
		{Name: "echo", Paths: "alpha.delta.echo", OrgId: defaultOrdID},
		{Name: "foxtrot", Paths: "foxtrot", OrgId: secondaryOrdID},
		{Name: "golf", Paths: "golf", OrgId: defaultOrdID},
	}

	tests := [...]struct {
		testName    string
		start       string
		destination string
		orgID       uuid.UUID
		folders     []folder.Folder
		want        string
	}{
		{
			testName:    "Example 1: Move folder to a child of itself",
			start:       "bravo",
			destination: "charlie",
			orgID:       defaultOrdID,
			folders:     example1,
			want:        "cannot move folder to a child of itself",
		},
		{
			testName:    "Example 2: Move a folder to itself",
			start:       "bravo",
			destination: "bravo",
			orgID:       defaultOrdID,
			folders:     example1,
			want:        "cannot move a folder to itself",
		},
		{
			testName:    "Example 3: Move a folder to a different organisation",
			start:       "bravo",
			destination: "foxtrot",
			orgID:       defaultOrdID,
			folders:     example1,
			want:        "cannot move a folder to a different organisation",
		},
		{
			testName:    "Example 4: Source folder does not exist",
			start:       "invalid_folder",
			destination: "delta",
			orgID:       defaultOrdID,
			folders:     example1,
			want:        "source folder does not exist",
		},
		{
			testName:    "Example 5: Destination folder does not exist",
			start:       "bravo",
			destination: "invalid_folder",
			orgID:       defaultOrdID,
			folders:     example1,
			want:        "destination folder does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			_, err := f.MoveFolder(tt.start, tt.destination)
			assert.ErrorContains(t, err, tt.want)
		})
	}
}
