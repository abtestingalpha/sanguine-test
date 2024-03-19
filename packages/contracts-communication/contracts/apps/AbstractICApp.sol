// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IInterchainApp} from "../interfaces/IInterchainApp.sol";
import {IInterchainClientV1} from "../interfaces/IInterchainClientV1.sol";

import {InterchainTxDescriptor} from "../libs/InterchainTransaction.sol";
import {TypeCasts} from "../libs/TypeCasts.sol";

abstract contract AbstractICApp is IInterchainApp {
    using TypeCasts for address;

    error InterchainApp__BalanceTooLow(uint256 actual, uint256 required);
    error InterchainApp__InterchainClientZeroAddress();
    error InterchainApp__NotInterchainClient(address caller);
    error InterchainApp__ReceiverNotSet(uint256 chainId);
    error InterchainApp__SameChainId(uint256 chainId);
    error InterchainApp__SenderNotAllowed(uint256 srcChainId, bytes32 sender);

    /// @inheritdoc IInterchainApp
    function appReceive(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        external
        payable
    {
        if (!_isInterchainClient(msg.sender)) {
            revert InterchainApp__NotInterchainClient(msg.sender);
        }
        if (srcChainId == block.chainid) {
            revert InterchainApp__SameChainId(srcChainId);
        }
        if (!_isAllowedSender(srcChainId, sender)) {
            revert InterchainApp__SenderNotAllowed(srcChainId, sender);
        }
        _receiveMessage(srcChainId, sender, dbNonce, entryIndex, message);
    }

    /// @inheritdoc IInterchainApp
    function getReceivingConfig() external view returns (bytes memory appConfig, address[] memory modules) {
        appConfig = _getAppConfig();
        modules = _getModules();
    }

    // ══════════════════════════════════════════════ INTERNAL LOGIC ═══════════════════════════════════════════════════

    /// @dev Thin wrapper around _sendInterchainMessage to accept EVM address as a parameter.
    function _sendInterchainMessageEVM(
        uint256 dstChainId,
        address receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        internal
        returns (InterchainTxDescriptor memory desc)
    {
        return _sendInterchainMessage(dstChainId, receiver.addressToBytes32(), messageFee, options, message);
    }

    /// @dev Performs necessary checks and sends an interchain message.
    function _sendInterchainMessage(
        uint256 dstChainId,
        bytes32 receiver,
        uint256 messageFee,
        bytes memory options,
        bytes memory message
    )
        internal
        returns (InterchainTxDescriptor memory desc)
    {
        address client = _getLatestClient();
        if (client == address(0)) {
            revert InterchainApp__InterchainClientZeroAddress();
        }
        if (dstChainId == block.chainid) {
            revert InterchainApp__SameChainId(dstChainId);
        }
        if (receiver == 0) {
            revert InterchainApp__ReceiverNotSet(dstChainId);
        }
        if (address(this).balance < messageFee) {
            revert InterchainApp__BalanceTooLow({actual: address(this).balance, required: messageFee});
        }
        return IInterchainClientV1(client).interchainSend{value: messageFee}(
            dstChainId, receiver, _getExecutionService(), _getModules(), options, message
        );
    }

    /// @dev Internal logic for receiving messages. At this point the validity of the message is already checked.
    function _receiveMessage(
        uint256 srcChainId,
        bytes32 sender,
        uint256 dbNonce,
        uint64 entryIndex,
        bytes calldata message
    )
        internal
        virtual;

    // ══════════════════════════════════════════════ INTERNAL VIEWS ═══════════════════════════════════════════════════

    /// @dev Returns the fee for sending an Interchain message.
    function _getInterchainFee(
        uint256 dstChainId,
        bytes memory options,
        bytes memory message
    )
        internal
        view
        returns (uint256)
    {
        address client = _getLatestClient();
        if (client == address(0)) {
            revert InterchainApp__InterchainClientZeroAddress();
        }
        return IInterchainClientV1(client).getInterchainFee(
            dstChainId, _getExecutionService(), _getModules(), options, message
        );
    }

    /// @dev Returns the configuration of the app for validating the received messages.
    function _getAppConfig() internal view virtual returns (bytes memory);

    /// @dev Returns the address of the Execution Service to use for sending messages.
    function _getExecutionService() internal view virtual returns (address);

    /// @dev Returns the latest Interchain Client. This is the Client that is used for sending messages.
    function _getLatestClient() internal view virtual returns (address);

    /// @dev Returns the list of modules to use for sending messages, as well as validating the received messages.
    function _getModules() internal view virtual returns (address[] memory);

    /// @dev Checks if the sender is allowed to send messages to this app.
    function _isAllowedSender(uint256 srcChainId, bytes32 sender) internal view virtual returns (bool);

    /// @dev Checks if the caller is an Interchain Client.
    /// Both latest and legacy Interchain Clients are allowed to call `appReceive`.
    function _isInterchainClient(address caller) internal view virtual returns (bool);
}
