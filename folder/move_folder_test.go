package folder_test

import (
	"testing"
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
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
			testName: "Empty folder",
			start: "bravo",
			destination: "delta",
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

			fmt.Println("INPUT")
			for _, folder := range tt.folders {
				str := folder.Name + " | " + folder.Paths + " | "+ folder.OrgId.String() 
				fmt.Println(str)
			}

			fmt.Println("GOT")
			for _, folder := range get {
				str := folder.Name + " | " + folder.Paths + " | "+ folder.OrgId.String() 
				fmt.Println(str)
			}

			// assert.Equal(t, tt.want, get)
		})
	}
}
