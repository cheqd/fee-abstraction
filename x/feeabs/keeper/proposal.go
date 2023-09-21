package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/fee-abstraction/v7/x/feeabs/types"
)

func (k Keeper) AddHostZoneProposal(ctx sdk.Context, p *types.AddHostZoneProposal) error {
	_, found := k.GetHostZoneConfig(ctx, p.HostChainConfig.IbcDenom)
	if found {
		return types.ErrDuplicateHostZoneConfig
	}

	err := k.SetHostZoneConfig(ctx, *p.HostChainConfig)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) DeleteHostZoneProposal(ctx sdk.Context, p *types.DeleteHostZoneProposal) error {
	_, found := k.GetHostZoneConfig(ctx, p.IbcDenom)
	if !found {
		return types.ErrHostZoneConfigNotFound
	}

	err := k.DeleteHostZoneConfig(ctx, p.IbcDenom)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetHostZoneProposal(ctx sdk.Context, p *types.SetHostZoneProposal) error {
	_, found := k.GetHostZoneConfig(ctx, p.HostChainConfig.IbcDenom)
	if !found {
		return types.ErrHostZoneConfigNotFound
	}

	err := k.SetHostZoneConfig(ctx, *p.HostChainConfig)
	if err != nil {
		return err
	}

	return nil
}
