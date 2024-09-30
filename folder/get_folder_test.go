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

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{
			name: "One folder",
			orgID: defaultOrdID,
			folders: []folder.Folder{
				{Name: "alpha", OrgId: defaultOrdID, Paths: "alpha"},
			},
			want: []folder.Folder{
				{Name: "alpha", OrgId: defaultOrdID, Paths: "alpha"},
			},
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
