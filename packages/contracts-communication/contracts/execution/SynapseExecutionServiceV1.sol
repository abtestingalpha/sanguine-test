// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {IExecutionService, ISynapseExecutionServiceV1} from "../interfaces/ISynapseExecutionServiceV1.sol";
import {SynapseExecutionServiceEvents} from "../events/SynapseExecutionServiceEvents.sol";
import {IGasOracle} from "../interfaces/IGasOracle.sol";
import {OptionsLib, OptionsV1} from "../libs/Options.sol";

import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";

contract SynapseExecutionServiceV1 is
    AccessControlUpgradeable,
    SynapseExecutionServiceEvents,
    ISynapseExecutionServiceV1
{
    bytes32 public constant GOVERNOR_ROLE = keccak256("GOVERNOR_ROLE");
    bytes32 public constant IC_CLIENT_ROLE = keccak256("IC_CLIENT_ROLE");

    // TODO: Use ERC-7201 to manage storage
    /// @inheritdoc IExecutionService
    address public executorEOA;
    /// @inheritdoc ISynapseExecutionServiceV1
    address public gasOracle;

    constructor() {
        // Ensure that the implementation contract could not be initialized
        _disableInitializers();
    }

    function initialize(address admin) external initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setExecutorEOA(address executorEOA_) external onlyRole(GOVERNOR_ROLE) {
        if (executorEOA_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        executorEOA = executorEOA_;
        emit ExecutorEOASet(executorEOA_);
    }

    /// @inheritdoc ISynapseExecutionServiceV1
    function setGasOracle(address gasOracle_) external onlyRole(GOVERNOR_ROLE) {
        if (gasOracle_ == address(0)) {
            revert SynapseExecutionService__ZeroAddress();
        }
        gasOracle = gasOracle_;
        emit GasOracleSet(gasOracle_);
    }

    /// @inheritdoc IExecutionService
    function requestExecution(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes32 transactionId,
        uint256 executionFee,
        bytes memory options
    )
        external
        onlyRole(IC_CLIENT_ROLE)
    {
        uint256 requiredFee = getExecutionFee(dstChainId, txPayloadSize, options);
        if (executionFee < requiredFee) {
            revert SynapseExecutionService__FeeAmountTooLow({actual: executionFee, required: requiredFee});
        }
        emit ExecutionRequested({transactionId: transactionId, client: msg.sender});
    }

    /// @inheritdoc IExecutionService
    function getExecutionFee(
        uint256 dstChainId,
        uint256 txPayloadSize,
        bytes memory options
    )
        public
        view
        returns (uint256 executionFee)
    {
        address cachedGasOracle = gasOracle;
        if (cachedGasOracle == address(0)) {
            revert SynapseExecutionService__GasOracleNotSet();
        }
        // TODO: the "exact version" check should be generalized
        (uint8 version,) = OptionsLib.decodeVersionedOptions(options);
        if (version > OptionsLib.OPTIONS_V1) {
            revert SynapseExecutionService__OptionsVersionNotSupported(version);
        }
        OptionsV1 memory optionsV1 = OptionsLib.decodeOptionsV1(options);
        // TODO: there should be a markup applied to the execution fee
        executionFee = IGasOracle(cachedGasOracle).estimateTxCostInLocalUnits({
            remoteChainId: dstChainId,
            gasLimit: optionsV1.gasLimit,
            calldataSize: txPayloadSize
        });
        if (optionsV1.gasAirdrop > 0) {
            executionFee += IGasOracle(cachedGasOracle).convertRemoteValueToLocalUnits({
                remoteChainId: dstChainId,
                value: optionsV1.gasAirdrop
            });
        }
    }
}
