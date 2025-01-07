package avs

import (
	"fmt"

	exocmn "github.com/ExocoreNetwork/exocore/precompiles/common"
	avstypes "github.com/ExocoreNetwork/exocore/x/avs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	cmn "github.com/evmos/evmos/v16/precompiles/common"
	"golang.org/x/xerrors"
)

func (p Precompile) GetAVSParamsFromInputs(_ sdk.Context, args []interface{}) (*avstypes.AVSRegisterOrDeregisterParams, error) {
	if len(args) != len(p.ABI.Methods[MethodRegisterAVS].Inputs) {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, len(p.ABI.Methods[MethodRegisterAVS].Inputs), len(args))
	}
	avsParams := &avstypes.AVSRegisterOrDeregisterParams{}
	//	we'd better not use evm.Origin but let the precompile caller pass in the sender address,
	//	since tx.origin has some security issue and might not be supported
	//	in a long term: https://docs.soliditylang.org/en/latest/security-considerations.html#tx-origin
	callerAddress, ok := args[0].(common.Address)
	if !ok || (callerAddress == common.Address{}) {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 0, "common.Address", callerAddress)
	}
	avsParams.CallerAddress = callerAddress[:]
	avsName, ok := args[1].(string)
	if !ok || avsName == "" {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 1, "string", avsName)
	}
	avsParams.AvsName = avsName
	// When creating tasks in AVS, check the minimum requirements,minStakeAmount at least greater than 0
	minStakeAmount, ok := args[2].(uint64)
	if !ok || minStakeAmount == 0 {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 2, "uint64", minStakeAmount)
	}
	avsParams.MinStakeAmount = minStakeAmount

	taskAddr, ok := args[3].(common.Address)
	if !ok || taskAddr == (common.Address{}) {
		return nil, xerrors.Errorf(exocmn.ErrContractInputParamOrType, 3, "common.Address", taskAddr)
	}
	avsParams.TaskAddr = taskAddr

	slashContractAddr, ok := args[4].(common.Address)
	if !ok || (slashContractAddr == common.Address{}) {
		return nil, xerrors.Errorf(exocmn.ErrContractInputParamOrType, 4, "common.Address", slashContractAddr)
	}
	avsParams.SlashContractAddr = slashContractAddr

	rewardContractAddr, ok := args[5].(common.Address)
	if !ok || (rewardContractAddr == common.Address{}) {
		return nil, xerrors.Errorf(exocmn.ErrContractInputParamOrType, 5, "common.Address", rewardContractAddr)
	}
	avsParams.RewardContractAddr = rewardContractAddr

	// bech32
	avsOwnerAddress, ok := args[6].([]common.Address)
	if !ok || avsOwnerAddress == nil {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 6, "[]common.Address", avsOwnerAddress)
	}
	exoAddresses := make([]string, len(avsOwnerAddress))
	for i, addr := range avsOwnerAddress {
		var accAddr sdk.AccAddress = addr[:]
		exoAddresses[i] = accAddr.String()
	}
	avsParams.AvsOwnerAddress = exoAddresses
	// bech32
	whitelistAddress, ok := args[7].([]common.Address)
	if !ok || whitelistAddress == nil {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 7, "[]common.Address", whitelistAddress)
	}
	exoWhiteAddresses := make([]string, len(whitelistAddress))
	for i, addr := range whitelistAddress {
		var accAddr sdk.AccAddress = addr[:]
		exoWhiteAddresses[i] = accAddr.String()
	}
	avsParams.WhitelistAddress = exoWhiteAddresses
	// string, since it is the address_id representation
	assetID, ok := args[8].([]string)
	if !ok || len(assetID) == 0 {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 8, "[]string", assetID)
	}
	avsParams.AssetID = assetID

	unbondingPeriod, ok := args[9].(uint64)
	if !ok || unbondingPeriod == 0 {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 9, "uint64", unbondingPeriod)
	}
	avsParams.UnbondingPeriod = unbondingPeriod

	minSelfDelegation, ok := args[10].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 10, "uint64", minSelfDelegation)
	}
	avsParams.MinSelfDelegation = minSelfDelegation

	epochIdentifier, ok := args[11].(string)
	if !ok || epochIdentifier == "" {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 11, "string", epochIdentifier)
	}
	avsParams.EpochIdentifier = epochIdentifier

	// The parameters below are used when creating tasks, to ensure that the minimum criteria are met by the set
	// of operators.

	taskParam, ok := args[12].([]uint64)
	if !ok || taskParam == nil || len(taskParam) != 4 {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 12, "[]string", taskParam)
	}
	minOptInOperators := taskParam[0]
	avsParams.MinOptInOperators = minOptInOperators

	minTotalStakeAmount := taskParam[1]
	avsParams.MinTotalStakeAmount = minTotalStakeAmount

	avsReward := taskParam[2]
	avsParams.AvsReward = avsReward

	avsSlash := taskParam[3]
	avsParams.AvsSlash = avsSlash

	return avsParams, nil
}

func (p Precompile) GetAVSParamsFromUpdateInputs(_ sdk.Context, args []interface{}) (*avstypes.AVSRegisterOrDeregisterParams, error) {
	if len(args) != len(p.ABI.Methods[MethodRegisterAVS].Inputs) {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, len(p.ABI.Methods[MethodRegisterAVS].Inputs), len(args))
	}
	avsParams := &avstypes.AVSRegisterOrDeregisterParams{}
	callerAddress, ok := args[0].(common.Address)
	if !ok || (callerAddress == common.Address{}) {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 0, "common.Address", callerAddress)
	}
	avsParams.CallerAddress = callerAddress[:]

	avsName, ok := args[1].(string)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 1, "string", avsName)
	}
	avsParams.AvsName = avsName
	// When creating tasks in AVS, check the minimum requirements,minStakeAmount at least greater than 0
	minStakeAmount, ok := args[2].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 2, "uint64", minStakeAmount)
	}
	avsParams.MinStakeAmount = minStakeAmount

	taskAddr, ok := args[3].(common.Address)
	if !ok || taskAddr == (common.Address{}) {
		return nil, xerrors.Errorf(exocmn.ErrContractInputParamOrType, 3, "common.Address", taskAddr)
	}
	avsParams.TaskAddr = taskAddr

	slashContractAddr, ok := args[4].(common.Address)
	if !ok {
		return nil, xerrors.Errorf(exocmn.ErrContractInputParamOrType, 4, "common.Address", slashContractAddr)
	}
	avsParams.SlashContractAddr = slashContractAddr

	rewardContractAddr, ok := args[5].(common.Address)
	if !ok {
		return nil, xerrors.Errorf(exocmn.ErrContractInputParamOrType, 5, "common.Address", rewardContractAddr)
	}
	avsParams.RewardContractAddr = rewardContractAddr

	// bech32
	avsOwnerAddress, ok := args[6].([]common.Address)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 6, "[]common.Address", avsOwnerAddress)
	}
	exoAddresses := make([]string, len(avsOwnerAddress))
	for i, addr := range avsOwnerAddress {
		var accAddr sdk.AccAddress = addr[:]
		exoAddresses[i] = accAddr.String()
	}
	avsParams.AvsOwnerAddress = exoAddresses
	// bech32
	whitelistAddress, ok := args[7].([]common.Address)
	if !ok || whitelistAddress == nil {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 7, "[]common.Address", whitelistAddress)
	}
	exoWhiteAddresses := make([]string, len(whitelistAddress))
	for i, addr := range whitelistAddress {
		var accAddr sdk.AccAddress = addr[:]
		exoWhiteAddresses[i] = accAddr.String()
	}
	avsParams.WhitelistAddress = exoWhiteAddresses
	// string, since it is the address_id representation
	assetID, ok := args[8].([]string)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 8, "[]string", assetID)
	}
	avsParams.AssetID = assetID

	unbondingPeriod, ok := args[9].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 9, "uint64", unbondingPeriod)
	}
	avsParams.UnbondingPeriod = unbondingPeriod

	minSelfDelegation, ok := args[10].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 10, "uint64", minSelfDelegation)
	}
	avsParams.MinSelfDelegation = minSelfDelegation

	epochIdentifier, ok := args[11].(string)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 11, "string", epochIdentifier)
	}
	avsParams.EpochIdentifier = epochIdentifier

	// The parameters below are used when creating tasks, to ensure that the minimum criteria are met by the set
	// of operators.

	taskParam, ok := args[12].([]uint64)
	if !ok || taskParam == nil || len(taskParam) != 4 {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 12, "[]string", taskParam)
	}
	minOptInOperators := taskParam[0]
	avsParams.MinOptInOperators = minOptInOperators

	minTotalStakeAmount := taskParam[1]
	avsParams.MinTotalStakeAmount = minTotalStakeAmount

	avsReward := taskParam[2]
	avsParams.AvsReward = avsReward

	avsSlash := taskParam[3]
	avsParams.AvsSlash = avsSlash

	return avsParams, nil
}

func (p Precompile) GetTaskParamsFromInputs(_ sdk.Context, args []interface{}) (*avstypes.TaskInfoParams, error) {
	if len(args) != len(p.ABI.Methods[MethodCreateAVSTask].Inputs) {
		return nil, fmt.Errorf(cmn.ErrInvalidNumberOfArgs, len(p.ABI.Methods[MethodCreateAVSTask].Inputs), len(args))
	}
	taskParams := &avstypes.TaskInfoParams{}
	callerAddress, ok := args[0].(common.Address)
	if !ok || (callerAddress == common.Address{}) {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 0, "common.Address", callerAddress)
	}
	taskParams.CallerAddress = callerAddress[:]
	name, ok := args[1].(string)
	if !ok || name == "" {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 1, "string", name)
	}
	taskParams.TaskName = name

	hash, ok := args[2].([]byte)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 2, "[]byte", hash)
	}
	taskParams.Hash = hash

	taskResponsePeriod, ok := args[3].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 3, "uint64", taskResponsePeriod)
	}
	taskParams.TaskResponsePeriod = taskResponsePeriod

	taskChallengePeriod, ok := args[4].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 4, "uint64", taskChallengePeriod)
	}
	taskParams.TaskChallengePeriod = taskChallengePeriod

	thresholdPercentage, ok := args[5].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 5, "uint64", thresholdPercentage)
	}
	taskParams.ThresholdPercentage = thresholdPercentage

	taskStatisticalPeriod, ok := args[6].(uint64)
	if !ok {
		return nil, fmt.Errorf(exocmn.ErrContractInputParamOrType, 6, "uint64", taskStatisticalPeriod)
	}
	taskParams.TaskStatisticalPeriod = taskStatisticalPeriod
	return taskParams, nil
}
