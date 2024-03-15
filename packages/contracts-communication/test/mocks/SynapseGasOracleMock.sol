// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {ISynapseGasOracle} from "../../contracts/interfaces/ISynapseGasOracle.sol";

contract SynapseGasOracleMock is ISynapseGasOracle {
    function receiveRemoteGasData(uint256 srcChainId, bytes calldata data) external {}

    function getLocalGasData() external view returns (bytes memory) {}

    function convertRemoteValueToLocalUnits(
        uint256 remoteChainId,
        uint256 value
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInLocalUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}

    function estimateTxCostInRemoteUnits(
        uint256 remoteChainId,
        uint256 gasLimit,
        uint256 calldataSize
    )
        external
        view
        override
        returns (uint256)
    {}
}