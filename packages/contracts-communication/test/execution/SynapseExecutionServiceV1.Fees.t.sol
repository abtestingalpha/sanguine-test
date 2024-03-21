// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import {OptionsLib, OptionsV1} from "../../contracts/libs/Options.sol";

import {SynapseExecutionServiceV1Test} from "./SynapseExecutionServiceV1.t.sol";

import {SynapseGasOracleMock} from "../mocks/SynapseGasOracleMock.sol";

// solhint-disable func-name-mixedcase
// solhint-disable ordering
contract SynapseExecutionServiceV1ExecutionTest is SynapseExecutionServiceV1Test {
    bytes32 public constant MOCK_TX_ID = keccak256("mock-tx-id");
    uint256 public constant LOCAL_CHAIN_ID = 1337;
    uint256 public constant REMOTE_CHAIN_ID = 7331;

    uint256 public constant MOCK_GAS_LIMIT = 100_000;
    uint256 public constant MOCK_GAS_AIRDROP = 0.01 ether;
    uint256 public constant MOCK_CALLDATA_SIZE = 1024;

    // Execution of tx costs 0.1 remote ETH
    uint256 public constant MOCK_REMOTE_EXEC_COST = 0.1 ether;
    // 1 remote ETH = 2 local ETH
    // Execution cost in local units
    uint256 public constant MOCK_LOCAL_EXEC_COST = MOCK_REMOTE_EXEC_COST * 2;
    uint256 public constant MOCK_LOCAL_AIRDROP_COST = MOCK_GAS_AIRDROP * 2;

    uint256 public constant MOCK_FEE_NO_AIRDROP = MOCK_LOCAL_EXEC_COST;
    uint256 public constant MOCK_FEE_WITH_AIRDROP = MOCK_LOCAL_EXEC_COST + MOCK_LOCAL_AIRDROP_COST;

    OptionsV1 public optionsNoAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: 0});
    OptionsV1 public optionsWithAirdrop = OptionsV1({gasLimit: MOCK_GAS_LIMIT, gasAirdrop: MOCK_GAS_AIRDROP});
    bytes public encodedOptionsNoAirdrop = optionsNoAirdrop.encodeOptionsV1();
    bytes public encodedOptionsWithAirdrop = optionsWithAirdrop.encodeOptionsV1();

    address public icClientA = makeAddr("ICClientA");
    address public icClientB = makeAddr("ICClientB");
    address public executorEOA = makeAddr("ExecutorEOA");
    address public gasOracle;

    function setUp() public override {
        super.setUp();
        gasOracle = address(new SynapseGasOracleMock());
        configureService();
        mockGasPricing();
    }

    function configureService() internal {
        service.initialize(address(this));
        service.grantRole(GOVERNOR_ROLE, address(this));
        service.grantRole(IC_CLIENT_ROLE, icClientA);
        service.grantRole(IC_CLIENT_ROLE, icClientB);
        service.setExecutorEOA(executorEOA);
        service.setGasOracle(gasOracle);
    }

    function mockGasPricing() internal {
        // Mock: execution cost in local units
        vm.mockCall(
            gasOracle,
            abi.encodeCall(
                SynapseGasOracleMock.estimateTxCostInLocalUnits, (REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE)
            ),
            abi.encode(MOCK_LOCAL_EXEC_COST)
        );
        // Mock: execution cost in remote units + its conversion to local units
        vm.mockCall(
            gasOracle,
            abi.encodeCall(
                SynapseGasOracleMock.estimateTxCostInRemoteUnits, (REMOTE_CHAIN_ID, MOCK_GAS_LIMIT, MOCK_CALLDATA_SIZE)
            ),
            abi.encode(MOCK_REMOTE_EXEC_COST)
        );
        vm.mockCall(
            gasOracle,
            abi.encodeCall(
                SynapseGasOracleMock.convertRemoteValueToLocalUnits, (REMOTE_CHAIN_ID, MOCK_REMOTE_EXEC_COST)
            ),
            abi.encode(MOCK_LOCAL_EXEC_COST)
        );
        // Mock: airdrop conversion to local units
        vm.mockCall(
            gasOracle,
            abi.encodeCall(SynapseGasOracleMock.convertRemoteValueToLocalUnits, (REMOTE_CHAIN_ID, MOCK_GAS_AIRDROP)),
            abi.encode(MOCK_LOCAL_AIRDROP_COST)
        );
    }

    function getMockOptionsV2() internal view returns (bytes memory) {
        bytes memory optionsV1 = optionsNoAirdrop.encodeOptionsV1();
        (, bytes memory data) = OptionsLib.decodeVersionedOptions(optionsV1);
        return OptionsLib.encodeVersionedOptions(2, data);
    }

    function test_getExecutionFee_noAirdrop() public {
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsNoAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_NO_AIRDROP);
    }

    function test_getExecutionFee_withAirdrop() public {
        uint256 fee = service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: optionsWithAirdrop.encodeOptionsV1()
        });
        assertEq(fee, MOCK_FEE_WITH_AIRDROP);
    }

    function test_getExecutionFee_revert_unsupportedOptionsVersion() public {
        expectRevertOptionsVersionNotSupported(2);
        service.getExecutionFee({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            options: getMockOptionsV2()
        });
    }

    function test_requestExecution_clientA_noAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_clientA_noAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP + 1,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_clientA_noAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_NO_AIRDROP - 1, MOCK_FEE_NO_AIRDROP);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP - 1,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_clientA_noAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: optionsV2
        });
    }

    function test_requestExecution_clientA_withAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: encodedOptionsWithAirdrop
        });
    }

    function test_requestExecution_clientA_withAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientA);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP + 1,
            options: encodedOptionsWithAirdrop
        });
    }

    function test_requestExecution_clientA_withAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_WITH_AIRDROP - 1, MOCK_FEE_WITH_AIRDROP);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP - 1,
            options: encodedOptionsWithAirdrop
        });
    }

    function test_requestExecution_clientA_withAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        vm.prank(icClientA);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: optionsV2
        });
    }

    function test_requestExecution_clientB_noAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_clientB_noAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP + 1,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_clientB_noAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_NO_AIRDROP - 1, MOCK_FEE_NO_AIRDROP);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP - 1,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_clientB_noAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: optionsV2
        });
    }

    function test_requestExecution_clientB_withAirdrop_exactFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: encodedOptionsWithAirdrop
        });
    }

    function test_requestExecution_clientB_withAirdrop_higherFee() public {
        expectEventExecutionRequested(MOCK_TX_ID, icClientB);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP + 1,
            options: encodedOptionsWithAirdrop
        });
    }

    function test_requestExecution_clientB_withAirdrop_revert_lowerFee() public {
        expectRevertFeeAmountTooLow(MOCK_FEE_WITH_AIRDROP - 1, MOCK_FEE_WITH_AIRDROP);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP - 1,
            options: encodedOptionsWithAirdrop
        });
    }

    function test_requestExecution_clientB_withAirdrop_revert_unsupportedOptionsVersion() public {
        bytes memory optionsV2 = getMockOptionsV2();
        expectRevertOptionsVersionNotSupported(2);
        vm.prank(icClientB);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: optionsV2
        });
    }

    function test_requestExecution_noAirdrop_revert_notInterchainClient(address caller) public {
        vm.assume(caller != icClientA && caller != icClientB);
        expectRevertNotInterchainClient(caller);
        vm.prank(caller);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_NO_AIRDROP,
            options: encodedOptionsNoAirdrop
        });
    }

    function test_requestExecution_withAirdrop_revert_notInterchainClient(address caller) public {
        vm.assume(caller != icClientA && caller != icClientB);
        expectRevertNotInterchainClient(caller);
        vm.prank(caller);
        service.requestExecution({
            dstChainId: REMOTE_CHAIN_ID,
            txPayloadSize: MOCK_CALLDATA_SIZE,
            transactionId: MOCK_TX_ID,
            executionFee: MOCK_FEE_WITH_AIRDROP,
            options: encodedOptionsWithAirdrop
        });
    }
}
