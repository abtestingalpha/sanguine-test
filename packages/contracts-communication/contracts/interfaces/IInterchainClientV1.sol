// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IInterchainClientV1 {
    /**
     * @notice Sets the address of the ExecutionFees contract.
     * @dev Only callable by the contract owner or an authorized account.
     * @param executionFees_ The address of the ExecutionFees contract.
     */
    function setExecutionFees(address executionFees_) external;

    /**
     * @notice Sets the linked client for a specific chain ID.
     * @dev Stores the address of the linked client in a mapping with the chain ID as the key.
     * @param chainId The chain ID for which the client is being set.
     * @param client The address of the client being linked.
     */
    function setLinkedClient(uint256 chainId, bytes32 client) external;

    /**
     * @notice Sets the address of the InterchainDB contract.
     * @dev Only callable by the contract owner or an authorized account.
     * @param _interchainDB The address of the InterchainDB contract.
     */
    function setInterchainDB(address _interchainDB) external;

    /**
     * @notice Sends a message to another chain via the Interchain Communication Protocol.
     * @dev Charges a fee for the message, which is payable upon calling this function:
     * - Verification fees: paid to every module that verifies the message.
     * - Execution fee: paid to the executor that executes the message.
     * Note: while a specific execution service is specified to request the execution of the message,
     * any executor is able to execute the message on destination chain, earning the execution fee.
     * @param dstChainId The chain ID of the destination chain.
     * @param receiver The address of the receiver on the destination chain.
     * @param srcExecutionService The address of the execution service to use for the message.
     * @param message The message being sent.
     * @param options Execution options for the message sent, encoded as bytes, currently primarily gas limit + native gas drop.
     * @param srcModules The source modules involved in the message sending.
     */
    function interchainSend(
        uint256 dstChainId,
        bytes32 receiver,
        address srcExecutionService,
        bytes calldata message,
        bytes calldata options,
        address[] calldata srcModules
    )
        external
        payable;

    /**
     * @notice Executes a transaction that has been sent via the Interchain.
     * @dev The transaction must have been previously sent and recorded.
     * @param transaction The transaction data.
     */
    function interchainExecute(bytes calldata transaction) external;

    /**
     * @notice Checks if a transaction is executable.
     * @dev Determines if a transaction meets the criteria to be executed based on:
     * - If approved modules have written to the InterchainDB
     * - If the threshold of approved modules have been met
     * - If the optimistic window has passed for all modules
     * @param transaction The InterchainTransaction struct to be checked.
     * @return bool Returns true if the transaction is executable, false otherwise.
     */
    function isExecutable(bytes calldata transaction) external view returns (bool);
}