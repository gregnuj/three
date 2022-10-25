package three_test

import (
	"testing"

	keepertest "github.com/gregnuj/three/testutil/keeper"
	"github.com/gregnuj/three/testutil/nullify"
	"github.com/gregnuj/three/x/three"
	"github.com/gregnuj/three/x/three/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ThreeKeeper(t)
	three.InitGenesis(ctx, *k, genesisState)
	got := three.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
