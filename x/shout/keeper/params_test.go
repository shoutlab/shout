package keeper_test

import (
	"testing"

	testkeeper "github.com/shoutlab/shout/testutil/keeper"
	"github.com/shoutlab/shout/x/shout/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ShoutKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
