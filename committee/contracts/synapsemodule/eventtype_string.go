// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package synapsemodule

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VerificationRequestedEvent-1]
	_ = x[EntryVerificationEvent-2]
}

const _EventType_name = "VerificationRequestedEventEntryVerificationEvent"

var _EventType_index = [...]uint8{0, 26, 48}

func (i EventType) String() string {
	i -= 1
	if i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
