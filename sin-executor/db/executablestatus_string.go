// Code generated by "stringer -type=ExecutableStatus"; DO NOT EDIT.

package db

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Seen-1]
}

const _ExecutableStatus_name = "Seen"

var _ExecutableStatus_index = [...]uint8{0, 4}

func (i ExecutableStatus) String() string {
	i -= 1
	if i >= ExecutableStatus(len(_ExecutableStatus_index)-1) {
		return "ExecutableStatus(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ExecutableStatus_name[_ExecutableStatus_index[i]:_ExecutableStatus_index[i+1]]
}
