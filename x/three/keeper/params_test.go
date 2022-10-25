package keeper_test

import (
	"testing"

	testkeeper "github.com/gregnuj/three/testutil/keeper"
	"github.com/gregnuj/three/x/three/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ThreeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
