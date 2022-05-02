package hello_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/xxxxxchan/hello/testutil/keeper"
	"github.com/xxxxxchan/hello/testutil/nullify"
	"github.com/xxxxxchan/hello/x/hello"
	"github.com/xxxxxchan/hello/x/hello/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloKeeper(t)
	hello.InitGenesis(ctx, *k, genesisState)
	got := hello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
