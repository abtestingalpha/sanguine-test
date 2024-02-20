// Code generated by "stringer -type=contractTypeImpl -linecomment"; DO NOT EDIT.

package testutil

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InterchainClient-1]
	_ = x[InterchainDB-2]
	_ = x[InterchainModuleMock-3]
	_ = x[InterchainAppMock-4]
	_ = x[OptionsLib-5]
	_ = x[ExecutionServiceMock-6]
	_ = x[ExecutionFeesMock-7]
}

const _contractTypeImpl_name = "SynapseModuleSynapseModuleInterchainModuleMockInterchainAppMockOptionsLibExecutionServiceMockExecutionFeesMock"

var _contractTypeImpl_index = [...]uint8{0, 13, 26, 46, 63, 73, 93, 110}

func (i contractTypeImpl) String() string {
	i -= 1
	if i < 0 || i >= contractTypeImpl(len(_contractTypeImpl_index)-1) {
		return "contractTypeImpl(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _contractTypeImpl_name[_contractTypeImpl_index[i]:_contractTypeImpl_index[i+1]]
}
