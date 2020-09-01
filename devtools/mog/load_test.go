package main

import (
	"go/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadSourceStructs(t *testing.T) {
	actual, err := loadSourceStructs("./internal/sourcepkg")
	require.NoError(t, err)
	require.Equal(t, []string{"GroupedSample", "Sample"}, actual.Names())
	_, ok := actual.structs["Sample"]
	require.True(t, ok)
	_, ok = actual.structs["GroupedSample"]
	require.True(t, ok)

	// TODO: check the value in structs map
}

func TestLoadTargetStructs(t *testing.T) {
	actual, err := loadTargetStructs([]string{"./internal/targetpkgone", "./internal/targetpkgtwo"})
	require.NoError(t, err)

	expected := map[string]targetPkg{
		"github.com/hashicorp/mog/internal/targetpkgone": {
			Name: "github.com/hashicorp/mog/internal/targetpkgone",
			Structs: map[string]targetStruct{
				"TheSample": {
					Name: "TheSample",
					Fields: []*types.Var{
						types.NewVar(0, nil, "BoolField", &types.Basic{}),
					},
				},
			},
		},
		"github.com/hashicorp/mog/internal/targetpkgtwo": {
			Name: "github.com/hashicorp/mog/internal/targetpkgtwo",
			Structs: map[string]targetStruct{
				"Lamp": {
					Name: "Lamp",
				},
				"Flood": {
					Name: "Flood",
				},
			},
		},
	}

	require.Equal(t, expected, actual)
}
