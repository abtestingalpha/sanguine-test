// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {FastBridgeV2, IFastBridgeV2} from "../contracts/FastBridgeV2.sol";

import {FastBridge, FastBridgeTest} from "./FastBridge.t.sol";

contract FastBridgeV2ParityTest is FastBridgeTest {
    FastBridgeV2 fastBridgeV2;

    function deployFastBridge() internal override {
        fastBridgeV2 = new FastBridgeV2(owner);
        fastBridge = FastBridge(address(fastBridgeV2));
    }

    // ═══════════════════════════════════════════ OVERRIDES FOR CHECKS ════════════════════════════════════════════════

    function checkNotCompletedTx(bytes32 transactionId) internal virtual override {
        super.checkNotCompletedTx(transactionId);
        assertEq(fastBridgeV2.getDestinationRelayer(transactionId), address(0));
    }

    function checkCompletedTx(bytes32 transactionId, address expectedRelayer) internal virtual override {
        super.checkCompletedTx(transactionId, expectedRelayer);
        assertEq(fastBridgeV2.getDestinationRelayer(transactionId), expectedRelayer);
    }

    // ════════════════════════════════════════════ OVERRIDES FOR TESTS ════════════════════════════════════════════════

    // ════════════════════════════════════════ OVERRIDES FOR CUSTOM ERRORS ════════════════════════════════════════════

    // FastBridgeV2 uses a different naming convention for errors, so we need to override the error selectors

    function amountIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__AmountIncorrect.selector;
    }

    function chainIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__ChainIncorrect.selector;
    }

    function deadlineExceededSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DeadlineExceeded.selector;
    }

    function deadlineNotExceededSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DeadlineNotExceeded.selector;
    }

    function deadlineTooShortSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DeadlineTooShort.selector;
    }

    function disputePeriodNotPassedSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DisputePeriodNotPassed.selector;
    }

    function disputePeriodPassedSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__DisputePeriodPassed.selector;
    }

    function msgValueIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__MsgValueIncorrect.selector;
    }

    function senderIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__SenderIncorrect.selector;
    }

    function statusIncorrectSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__StatusIncorrect.selector;
    }

    function transactionRelayedSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__TransactionRelayed.selector;
    }

    function zeroAddressSelector() internal pure virtual override returns (bytes4) {
        return IFastBridgeV2.FastBridge__ZeroAddress.selector;
    }
}
