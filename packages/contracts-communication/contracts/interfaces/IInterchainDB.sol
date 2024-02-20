// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {InterchainEntry} from "../libs/InterchainEntry.sol";

interface IInterchainDB {
    /// @notice Struct representing an entry from the local Interchain DataBase.
    /// @param writer       The address of the writer on the local chain
    /// @param dataHash     The hash of the data written on the local chain
    struct LocalEntry {
        address writer;
        bytes32 dataHash;
    }

    /// @notice Struct representing an entry from the remote Interchain DataBase verified by the Interchain Module
    /// @param verifiedAt   The block timestamp at which the entry was verified by the module
    /// @param entryValue   The value of the entry: writer + dataHash hashed together
    struct RemoteEntry {
        uint256 verifiedAt;
        bytes32 entryValue;
    }

    error InterchainDB__ConflictingEntries(bytes32 existingEntryValue, InterchainEntry newEntry);
    error InterchainDB__EntryDoesNotExist(uint256 dbNonce);
    error InterchainDB__IncorrectFeeAmount(uint256 actualFee, uint256 expectedFee);
    error InterchainDB__NoModulesSpecified();
    error InterchainDB__SameChainId();

    /// @notice Write data to the Interchain DataBase as a new entry.
    /// Note: there are no guarantees that this entry will be available for reading on any of the remote chains.
    /// Use `verifyEntry` to ensure that the entry is available for reading on the destination chain.
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @return dbNonce     The database nonce of the written entry on this chain
    function writeEntry(bytes32 dataHash) external returns (uint256 dbNonce);

    /// @notice Request the given Interchain Modules to verify the already written entry on the destination chain.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// Note: this method is permissionless, and anyone can request verification for any entry.
    /// @dev Will revert if the entry with the given nonce does not exist.
    /// @param destChainId   The chain id of the destination chain
    /// @param dbNonce       The database nonce of the written entry on this chain
    /// @param srcModules    The source chain addresses of the Interchain Modules to use for verification
    function requestVerification(uint256 destChainId, uint256 dbNonce, address[] memory srcModules) external payable;

    /// @notice Write data to the Interchain DataBase,
    /// and request the given Interchain Modules to verify it on the destination chain.
    /// Note: every module has a separate fee paid in the native gas token of the source chain,
    /// and `msg.value` must be equal to the sum of all fees.
    /// Note: additional verification for the same entry could be later done using `requestVerification`
    /// by providing the returned `dbNonce`.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param destChainId  The chain id of the destination chain
    /// @param dataHash     The hash of the data to be written to the Interchain DataBase as a new entry
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    /// @return dbNonce     The database nonce of the written entry on this chain
    function writeEntryWithVerification(
        uint256 destChainId,
        bytes32 dataHash,
        address[] memory srcModules
    )
        external
        payable
        returns (uint256 dbNonce);

    /// @notice Allows the Interchain Module to verify the entry coming from a remote source chain.
    /// @param entry        The Interchain Entry to confirm
    function verifyEntry(InterchainEntry memory entry) external;

    /// @notice Get the fee for writing data to the Interchain DataBase, and verifying it on the destination chain
    /// using the provided Interchain Modules.
    /// @dev Will revert if the empty array of modules is provided.
    /// @param destChainId  The chain id of the destination chain
    /// @param srcModules   The source chain addresses of the Interchain Modules to use for verification
    function getInterchainFee(uint256 destChainId, address[] memory srcModules) external view returns (uint256);

    /// @notice Get the Interchain Entry by the writer and the writer nonce.
    /// @dev Will revert if the entry with the given nonce does not exist.
    /// @param dbNonce      The database nonce of the written entry on this chain
    function getEntry(uint256 dbNonce) external view returns (InterchainEntry memory);

    /// @notice Get the nonce of the database.
    function getDBNonce() external view returns (uint256);

    /// @notice Read the data written on specific source chain by a specific writer,
    /// and verify it on the destination chain using the provided Interchain Module.
    /// Note: returned zero value indicates that the module has not verified the entry.
    /// @param entry        The Interchain Entry to read
    /// @param dstModule    The destination chain addresses of the Interchain Modules to use for verification
    /// @return moduleVerifiedAt   The block timestamp at which the entry was verified by the module,
    ///                             or zero if the module has not verified the entry.
    function readEntry(
        address dstModule,
        InterchainEntry memory entry
    )
        external
        view
        returns (uint256 moduleVerifiedAt);
}
