// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package bridge

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DepositEvent-0]
	_ = x[RedeemEvent-1]
	_ = x[WithdrawEvent-2]
	_ = x[MintEvent-3]
	_ = x[DepositAndSwapEvent-4]
	_ = x[MintAndSwapEvent-5]
	_ = x[RedeemAndSwapEvent-6]
	_ = x[RedeemAndRemoveEvent-7]
	_ = x[WithdrawAndRemoveEvent-8]
	_ = x[RedeemV2Event-9]
	_ = x[CircleRequestSentEvent-10]
	_ = x[CircleRequestFulfilledEvent-11]
}

const _EventType_name = "DepositEventRedeemEventWithdrawEventMintEventDepositAndSwapEventMintAndSwapEventRedeemAndSwapEventRedeemAndRemoveEventWithdrawAndRemoveEventRedeemV2EventCircleRequestSentEventCircleRequestFulfilledEvent"

var _EventType_index = [...]uint8{0, 12, 23, 36, 45, 64, 80, 98, 118, 140, 153, 175, 202}

func (i EventType) String() string {
	if i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
