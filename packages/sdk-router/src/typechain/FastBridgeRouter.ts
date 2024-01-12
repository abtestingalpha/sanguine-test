/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PayableOverrides,
  PopulatedTransaction,
  Signer,
  utils,
} from 'ethers'
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from '@ethersproject/abi'
import type { Listener, Provider } from '@ethersproject/providers'
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
} from './common'

export type SwapQueryStruct = {
  routerAdapter: string
  tokenOut: string
  minAmountOut: BigNumberish
  deadline: BigNumberish
  rawParams: BytesLike
}

export type SwapQueryStructOutput = [
  string,
  string,
  BigNumber,
  BigNumber,
  string
] & {
  routerAdapter: string
  tokenOut: string
  minAmountOut: BigNumber
  deadline: BigNumber
  rawParams: string
}

export interface FastBridgeRouterInterface extends utils.Interface {
  functions: {
    'GAS_REBATE_FLAG()': FunctionFragment
    'adapterSwap(address,address,uint256,address,bytes)': FunctionFragment
    'bridge(address,uint256,address,uint256,(address,address,uint256,uint256,bytes),(address,address,uint256,uint256,bytes))': FunctionFragment
    'fastBridge()': FunctionFragment
    'getOriginAmountOut(address,address[],uint256)': FunctionFragment
    'owner()': FunctionFragment
    'renounceOwnership()': FunctionFragment
    'setFastBridge(address)': FunctionFragment
    'setSwapQuoter(address)': FunctionFragment
    'swapQuoter()': FunctionFragment
    'transferOwnership(address)': FunctionFragment
  }

  getFunction(
    nameOrSignatureOrTopic:
      | 'GAS_REBATE_FLAG'
      | 'adapterSwap'
      | 'bridge'
      | 'fastBridge'
      | 'getOriginAmountOut'
      | 'owner'
      | 'renounceOwnership'
      | 'setFastBridge'
      | 'setSwapQuoter'
      | 'swapQuoter'
      | 'transferOwnership'
  ): FunctionFragment

  encodeFunctionData(
    functionFragment: 'GAS_REBATE_FLAG',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'adapterSwap',
    values: [string, string, BigNumberish, string, BytesLike]
  ): string
  encodeFunctionData(
    functionFragment: 'bridge',
    values: [
      string,
      BigNumberish,
      string,
      BigNumberish,
      SwapQueryStruct,
      SwapQueryStruct
    ]
  ): string
  encodeFunctionData(functionFragment: 'fastBridge', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'getOriginAmountOut',
    values: [string, string[], BigNumberish]
  ): string
  encodeFunctionData(functionFragment: 'owner', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'renounceOwnership',
    values?: undefined
  ): string
  encodeFunctionData(
    functionFragment: 'setFastBridge',
    values: [string]
  ): string
  encodeFunctionData(
    functionFragment: 'setSwapQuoter',
    values: [string]
  ): string
  encodeFunctionData(functionFragment: 'swapQuoter', values?: undefined): string
  encodeFunctionData(
    functionFragment: 'transferOwnership',
    values: [string]
  ): string

  decodeFunctionResult(
    functionFragment: 'GAS_REBATE_FLAG',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'adapterSwap', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'bridge', data: BytesLike): Result
  decodeFunctionResult(functionFragment: 'fastBridge', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'getOriginAmountOut',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'owner', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'renounceOwnership',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'setFastBridge',
    data: BytesLike
  ): Result
  decodeFunctionResult(
    functionFragment: 'setSwapQuoter',
    data: BytesLike
  ): Result
  decodeFunctionResult(functionFragment: 'swapQuoter', data: BytesLike): Result
  decodeFunctionResult(
    functionFragment: 'transferOwnership',
    data: BytesLike
  ): Result

  events: {
    'FastBridgeSet(address)': EventFragment
    'OwnershipTransferred(address,address)': EventFragment
    'SwapQuoterSet(address)': EventFragment
  }

  getEvent(nameOrSignatureOrTopic: 'FastBridgeSet'): EventFragment
  getEvent(nameOrSignatureOrTopic: 'OwnershipTransferred'): EventFragment
  getEvent(nameOrSignatureOrTopic: 'SwapQuoterSet'): EventFragment
}

export interface FastBridgeSetEventObject {
  newFastBridge: string
}
export type FastBridgeSetEvent = TypedEvent<[string], FastBridgeSetEventObject>

export type FastBridgeSetEventFilter = TypedEventFilter<FastBridgeSetEvent>

export interface OwnershipTransferredEventObject {
  previousOwner: string
  newOwner: string
}
export type OwnershipTransferredEvent = TypedEvent<
  [string, string],
  OwnershipTransferredEventObject
>

export type OwnershipTransferredEventFilter =
  TypedEventFilter<OwnershipTransferredEvent>

export interface SwapQuoterSetEventObject {
  newSwapQuoter: string
}
export type SwapQuoterSetEvent = TypedEvent<[string], SwapQuoterSetEventObject>

export type SwapQuoterSetEventFilter = TypedEventFilter<SwapQuoterSetEvent>

export interface FastBridgeRouter extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this
  attach(addressOrName: string): this
  deployed(): Promise<this>

  interface: FastBridgeRouterInterface

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>
  listeners(eventName?: string): Array<Listener>
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this
  removeAllListeners(eventName?: string): this
  off: OnEvent<this>
  on: OnEvent<this>
  once: OnEvent<this>
  removeListener: OnEvent<this>

  functions: {
    GAS_REBATE_FLAG(overrides?: CallOverrides): Promise<[string]>

    adapterSwap(
      recipient: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<ContractTransaction>

    bridge(
      recipient: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<ContractTransaction>

    fastBridge(overrides?: CallOverrides): Promise<[string]>

    getOriginAmountOut(
      tokenIn: string,
      rfqTokens: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<
      [SwapQueryStructOutput[]] & { originQueries: SwapQueryStructOutput[] }
    >

    owner(overrides?: CallOverrides): Promise<[string]>

    renounceOwnership(
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    setFastBridge(
      fastBridge_: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    setSwapQuoter(
      swapQuoter_: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>

    swapQuoter(overrides?: CallOverrides): Promise<[string]>

    transferOwnership(
      newOwner: string,
      overrides?: Overrides & { from?: string }
    ): Promise<ContractTransaction>
  }

  GAS_REBATE_FLAG(overrides?: CallOverrides): Promise<string>

  adapterSwap(
    recipient: string,
    tokenIn: string,
    amountIn: BigNumberish,
    tokenOut: string,
    rawParams: BytesLike,
    overrides?: PayableOverrides & { from?: string }
  ): Promise<ContractTransaction>

  bridge(
    recipient: string,
    chainId: BigNumberish,
    token: string,
    amount: BigNumberish,
    originQuery: SwapQueryStruct,
    destQuery: SwapQueryStruct,
    overrides?: PayableOverrides & { from?: string }
  ): Promise<ContractTransaction>

  fastBridge(overrides?: CallOverrides): Promise<string>

  getOriginAmountOut(
    tokenIn: string,
    rfqTokens: string[],
    amountIn: BigNumberish,
    overrides?: CallOverrides
  ): Promise<SwapQueryStructOutput[]>

  owner(overrides?: CallOverrides): Promise<string>

  renounceOwnership(
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  setFastBridge(
    fastBridge_: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  setSwapQuoter(
    swapQuoter_: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  swapQuoter(overrides?: CallOverrides): Promise<string>

  transferOwnership(
    newOwner: string,
    overrides?: Overrides & { from?: string }
  ): Promise<ContractTransaction>

  callStatic: {
    GAS_REBATE_FLAG(overrides?: CallOverrides): Promise<string>

    adapterSwap(
      recipient: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    bridge(
      recipient: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: CallOverrides
    ): Promise<void>

    fastBridge(overrides?: CallOverrides): Promise<string>

    getOriginAmountOut(
      tokenIn: string,
      rfqTokens: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<SwapQueryStructOutput[]>

    owner(overrides?: CallOverrides): Promise<string>

    renounceOwnership(overrides?: CallOverrides): Promise<void>

    setFastBridge(fastBridge_: string, overrides?: CallOverrides): Promise<void>

    setSwapQuoter(swapQuoter_: string, overrides?: CallOverrides): Promise<void>

    swapQuoter(overrides?: CallOverrides): Promise<string>

    transferOwnership(
      newOwner: string,
      overrides?: CallOverrides
    ): Promise<void>
  }

  filters: {
    'FastBridgeSet(address)'(newFastBridge?: null): FastBridgeSetEventFilter
    FastBridgeSet(newFastBridge?: null): FastBridgeSetEventFilter

    'OwnershipTransferred(address,address)'(
      previousOwner?: string | null,
      newOwner?: string | null
    ): OwnershipTransferredEventFilter
    OwnershipTransferred(
      previousOwner?: string | null,
      newOwner?: string | null
    ): OwnershipTransferredEventFilter

    'SwapQuoterSet(address)'(newSwapQuoter?: null): SwapQuoterSetEventFilter
    SwapQuoterSet(newSwapQuoter?: null): SwapQuoterSetEventFilter
  }

  estimateGas: {
    GAS_REBATE_FLAG(overrides?: CallOverrides): Promise<BigNumber>

    adapterSwap(
      recipient: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<BigNumber>

    bridge(
      recipient: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<BigNumber>

    fastBridge(overrides?: CallOverrides): Promise<BigNumber>

    getOriginAmountOut(
      tokenIn: string,
      rfqTokens: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>

    owner(overrides?: CallOverrides): Promise<BigNumber>

    renounceOwnership(
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    setFastBridge(
      fastBridge_: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    setSwapQuoter(
      swapQuoter_: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>

    swapQuoter(overrides?: CallOverrides): Promise<BigNumber>

    transferOwnership(
      newOwner: string,
      overrides?: Overrides & { from?: string }
    ): Promise<BigNumber>
  }

  populateTransaction: {
    GAS_REBATE_FLAG(overrides?: CallOverrides): Promise<PopulatedTransaction>

    adapterSwap(
      recipient: string,
      tokenIn: string,
      amountIn: BigNumberish,
      tokenOut: string,
      rawParams: BytesLike,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<PopulatedTransaction>

    bridge(
      recipient: string,
      chainId: BigNumberish,
      token: string,
      amount: BigNumberish,
      originQuery: SwapQueryStruct,
      destQuery: SwapQueryStruct,
      overrides?: PayableOverrides & { from?: string }
    ): Promise<PopulatedTransaction>

    fastBridge(overrides?: CallOverrides): Promise<PopulatedTransaction>

    getOriginAmountOut(
      tokenIn: string,
      rfqTokens: string[],
      amountIn: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>

    owner(overrides?: CallOverrides): Promise<PopulatedTransaction>

    renounceOwnership(
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    setFastBridge(
      fastBridge_: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    setSwapQuoter(
      swapQuoter_: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>

    swapQuoter(overrides?: CallOverrides): Promise<PopulatedTransaction>

    transferOwnership(
      newOwner: string,
      overrides?: Overrides & { from?: string }
    ): Promise<PopulatedTransaction>
  }
}
