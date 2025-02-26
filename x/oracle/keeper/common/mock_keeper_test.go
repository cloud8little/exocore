// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ExocoreNetwork/exocore/x/oracle/keeper/common (interfaces: KeeperOracle)
//
// Generated by this command:
//
//	mockgen -destination mock_keeper_test.go -package common github.com/ExocoreNetwork/exocore/x/oracle/keeper/common KeeperOracle
//

// Package common is a generated GoMock package.
package common

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/ExocoreNetwork/exocore/x/oracle/types"
	types0 "github.com/cometbft/cometbft/abci/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "go.uber.org/mock/gomock"
)

// MockKeeperOracle is a mock of KeeperOracle interface.
type MockKeeperOracle struct {
	ctrl     *gomock.Controller
	recorder *MockKeeperOracleMockRecorder
}

// MockKeeperOracleMockRecorder is the mock recorder for MockKeeperOracle.
type MockKeeperOracleMockRecorder struct {
	mock *MockKeeperOracle
}

// NewMockKeeperOracle creates a new mock instance.
func NewMockKeeperOracle(ctrl *gomock.Controller) *MockKeeperOracle {
	mock := &MockKeeperOracle{ctrl: ctrl}
	mock.recorder = &MockKeeperOracleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeeperOracle) EXPECT() *MockKeeperOracleMockRecorder {
	return m.recorder
}

// GetAllRecentMsgAsMap mocks base method.
func (m *MockKeeperOracle) GetAllRecentMsgAsMap(arg0 types1.Context) map[int64][]*types.MsgItem {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRecentMsgAsMap", arg0)
	ret0, _ := ret[0].(map[int64][]*types.MsgItem)
	return ret0
}

// GetAllRecentMsgAsMap indicates an expected call of GetAllRecentMsgAsMap.
func (mr *MockKeeperOracleMockRecorder) GetAllRecentMsgAsMap(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRecentMsgAsMap", reflect.TypeOf((*MockKeeperOracle)(nil).GetAllRecentMsgAsMap), arg0)
}

// GetAllRecentParamsAsMap mocks base method.
func (m *MockKeeperOracle) GetAllRecentParamsAsMap(arg0 types1.Context) map[uint64]*types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRecentParamsAsMap", arg0)
	ret0, _ := ret[0].(map[uint64]*types.Params)
	return ret0
}

// GetAllRecentParamsAsMap indicates an expected call of GetAllRecentParamsAsMap.
func (mr *MockKeeperOracleMockRecorder) GetAllRecentParamsAsMap(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRecentParamsAsMap", reflect.TypeOf((*MockKeeperOracle)(nil).GetAllRecentParamsAsMap), arg0)
}

// GetIndexRecentMsg mocks base method.
func (m *MockKeeperOracle) GetIndexRecentMsg(arg0 types1.Context) (types.IndexRecentMsg, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexRecentMsg", arg0)
	ret0, _ := ret[0].(types.IndexRecentMsg)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetIndexRecentMsg indicates an expected call of GetIndexRecentMsg.
func (mr *MockKeeperOracleMockRecorder) GetIndexRecentMsg(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexRecentMsg", reflect.TypeOf((*MockKeeperOracle)(nil).GetIndexRecentMsg), arg0)
}

// GetIndexRecentParams mocks base method.
func (m *MockKeeperOracle) GetIndexRecentParams(arg0 types1.Context) (types.IndexRecentParams, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexRecentParams", arg0)
	ret0, _ := ret[0].(types.IndexRecentParams)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetIndexRecentParams indicates an expected call of GetIndexRecentParams.
func (mr *MockKeeperOracleMockRecorder) GetIndexRecentParams(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexRecentParams", reflect.TypeOf((*MockKeeperOracle)(nil).GetIndexRecentParams), arg0)
}

// GetLastTotalPower mocks base method.
func (m *MockKeeperOracle) GetLastTotalPower(arg0 types1.Context) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastTotalPower", arg0)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// GetLastTotalPower indicates an expected call of GetLastTotalPower.
func (mr *MockKeeperOracleMockRecorder) GetLastTotalPower(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastTotalPower", reflect.TypeOf((*MockKeeperOracle)(nil).GetLastTotalPower), arg0)
}

// GetParams mocks base method.
func (m *MockKeeperOracle) GetParams(arg0 types1.Context) types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", arg0)
	ret0, _ := ret[0].(types.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockKeeperOracleMockRecorder) GetParams(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockKeeperOracle)(nil).GetParams), arg0)
}

// GetValidatorByConsAddr mocks base method.
func (m *MockKeeperOracle) GetValidatorByConsAddr(arg0 types1.Context, arg1 types1.ConsAddress) (types2.Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types2.Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidatorByConsAddr indicates an expected call of GetValidatorByConsAddr.
func (mr *MockKeeperOracleMockRecorder) GetValidatorByConsAddr(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorByConsAddr", reflect.TypeOf((*MockKeeperOracle)(nil).GetValidatorByConsAddr), arg0, arg1)
}

// GetValidatorUpdateBlock mocks base method.
func (m *MockKeeperOracle) GetValidatorUpdateBlock(arg0 types1.Context) (types.ValidatorUpdateBlock, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorUpdateBlock", arg0)
	ret0, _ := ret[0].(types.ValidatorUpdateBlock)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidatorUpdateBlock indicates an expected call of GetValidatorUpdateBlock.
func (mr *MockKeeperOracleMockRecorder) GetValidatorUpdateBlock(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorUpdateBlock", reflect.TypeOf((*MockKeeperOracle)(nil).GetValidatorUpdateBlock), arg0)
}

// GetValidatorUpdates mocks base method.
func (m *MockKeeperOracle) GetValidatorUpdates(arg0 types1.Context) []types0.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorUpdates", arg0)
	ret0, _ := ret[0].([]types0.ValidatorUpdate)
	return ret0
}

// GetValidatorUpdates indicates an expected call of GetValidatorUpdates.
func (mr *MockKeeperOracleMockRecorder) GetValidatorUpdates(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorUpdates", reflect.TypeOf((*MockKeeperOracle)(nil).GetValidatorUpdates), arg0)
}

// IterateBondedValidatorsByPower mocks base method.
func (m *MockKeeperOracle) IterateBondedValidatorsByPower(arg0 types1.Context, arg1 func(int64, types2.ValidatorI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateBondedValidatorsByPower", arg0, arg1)
}

// IterateBondedValidatorsByPower indicates an expected call of IterateBondedValidatorsByPower.
func (mr *MockKeeperOracleMockRecorder) IterateBondedValidatorsByPower(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateBondedValidatorsByPower", reflect.TypeOf((*MockKeeperOracle)(nil).IterateBondedValidatorsByPower), arg0, arg1)
}

// RemoveRecentMsg mocks base method.
func (m *MockKeeperOracle) RemoveRecentMsg(arg0 types1.Context, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveRecentMsg", arg0, arg1)
}

// RemoveRecentMsg indicates an expected call of RemoveRecentMsg.
func (mr *MockKeeperOracleMockRecorder) RemoveRecentMsg(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRecentMsg", reflect.TypeOf((*MockKeeperOracle)(nil).RemoveRecentMsg), arg0, arg1)
}

// RemoveRecentParams mocks base method.
func (m *MockKeeperOracle) RemoveRecentParams(arg0 types1.Context, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveRecentParams", arg0, arg1)
}

// RemoveRecentParams indicates an expected call of RemoveRecentParams.
func (mr *MockKeeperOracleMockRecorder) RemoveRecentParams(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRecentParams", reflect.TypeOf((*MockKeeperOracle)(nil).RemoveRecentParams), arg0, arg1)
}

// SetIndexRecentMsg mocks base method.
func (m *MockKeeperOracle) SetIndexRecentMsg(arg0 types1.Context, arg1 types.IndexRecentMsg) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIndexRecentMsg", arg0, arg1)
}

// SetIndexRecentMsg indicates an expected call of SetIndexRecentMsg.
func (mr *MockKeeperOracleMockRecorder) SetIndexRecentMsg(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndexRecentMsg", reflect.TypeOf((*MockKeeperOracle)(nil).SetIndexRecentMsg), arg0, arg1)
}

// SetIndexRecentParams mocks base method.
func (m *MockKeeperOracle) SetIndexRecentParams(arg0 types1.Context, arg1 types.IndexRecentParams) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIndexRecentParams", arg0, arg1)
}

// SetIndexRecentParams indicates an expected call of SetIndexRecentParams.
func (mr *MockKeeperOracleMockRecorder) SetIndexRecentParams(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIndexRecentParams", reflect.TypeOf((*MockKeeperOracle)(nil).SetIndexRecentParams), arg0, arg1)
}

// SetRecentMsg mocks base method.
func (m *MockKeeperOracle) SetRecentMsg(arg0 types1.Context, arg1 types.RecentMsg) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRecentMsg", arg0, arg1)
}

// SetRecentMsg indicates an expected call of SetRecentMsg.
func (mr *MockKeeperOracleMockRecorder) SetRecentMsg(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRecentMsg", reflect.TypeOf((*MockKeeperOracle)(nil).SetRecentMsg), arg0, arg1)
}

// SetRecentParams mocks base method.
func (m *MockKeeperOracle) SetRecentParams(arg0 types1.Context, arg1 types.RecentParams) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRecentParams", arg0, arg1)
}

// SetRecentParams indicates an expected call of SetRecentParams.
func (mr *MockKeeperOracleMockRecorder) SetRecentParams(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRecentParams", reflect.TypeOf((*MockKeeperOracle)(nil).SetRecentParams), arg0, arg1)
}

// SetValidatorUpdateBlock mocks base method.
func (m *MockKeeperOracle) SetValidatorUpdateBlock(arg0 types1.Context, arg1 types.ValidatorUpdateBlock) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetValidatorUpdateBlock", arg0, arg1)
}

// SetValidatorUpdateBlock indicates an expected call of SetValidatorUpdateBlock.
func (mr *MockKeeperOracleMockRecorder) SetValidatorUpdateBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetValidatorUpdateBlock", reflect.TypeOf((*MockKeeperOracle)(nil).SetValidatorUpdateBlock), arg0, arg1)
}
