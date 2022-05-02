package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/xxxxxchan/hello/testutil/keeper"
	"github.com/xxxxxchan/hello/x/hello/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HelloKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
