import toast from 'react-hot-toast'
import React, { useEffect, useState, useRef, useCallback, useMemo } from 'react'
import { isNull } from 'lodash'
import { useAppSelector, useAppDispatch } from '@/store/hooks'
import { zeroAddress } from 'viem'
import { useAccount, useNetwork } from 'wagmi'
import { initialState, updateFromValue } from '@/slices/bridge/reducer'
import MiniMaxButton from '../buttons/MiniMaxButton'
import { formatBigIntToString } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'
import {
  ConnectToNetworkButton,
  ConnectWalletButton,
  ConnectedIndicator,
} from '@/components/ConnectionIndicators'
import { FromChainSelector } from './FromChainSelector'
import { FromTokenSelector } from './FromTokenSelector'
import { useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioState } from '@/slices/portfolio/hooks'
import { calculateGasCost } from '../../utils/calculateGasCost'
import { hasOnlyZeroes } from '@/utils/hasOnlyZeroes'
import { TokenAndBalance } from '@/utils/actions/fetchPortfolioBalances'
import { isEmpty } from 'lodash'

export const inputRef = React.createRef<HTMLInputElement>()

export const InputContainer = () => {
  const dispatch = useAppDispatch()
  const { chain } = useNetwork()
  const { isConnected } = useAccount()
  const { balances } = usePortfolioState()
  const { fromChainId, fromToken, fromValue } = useBridgeState()
  const [showValue, setShowValue] = useState('')
  const [hasMounted, setHasMounted] = useState(false)

  const { gasData } = useAppSelector((state) => state.gasData)
  const { gasPrice, maxFeePerGas } = gasData?.formatted
  const { rawGasCost, parsedGasCost } = calculateGasCost(maxFeePerGas, 200_000)

  const isGasToken: boolean = fromToken?.addresses[fromChainId] === zeroAddress

  const selectedFromToken: TokenAndBalance = balances[fromChainId]?.find(
    (token) => token.tokenAddress === fromToken?.addresses[fromChainId]
  )

  const { balance: rawBalance, parsedBalance: trimmedParsedBalance } =
    selectedFromToken || {}

  const parsedBalance: string = formatBigIntToString(
    rawBalance,
    fromToken?.decimals[fromChainId]
  )

  useEffect(() => {
    setHasMounted(true)
  }, [])

  const connectedStatus = useMemo(() => {
    if (hasMounted && !isConnected) {
      return <ConnectWalletButton />
    } else if (hasMounted && isConnected && fromChainId === chain.id) {
      return <ConnectedIndicator />
    } else if (hasMounted && isConnected && fromChainId !== chain.id) {
      return <ConnectToNetworkButton chainId={fromChainId} />
    }
  }, [chain, fromChainId, isConnected, hasMounted])

  useEffect(() => {
    if (fromToken && fromToken?.decimals[fromChainId]) {
      setShowValue(fromValue)
    }

    if (fromValue === initialState.fromValue) {
      setShowValue(initialState.fromValue)
    }
  }, [fromValue, inputRef, fromChainId, fromToken])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const fromValueString: string = cleanNumberInput(event.target.value)
    try {
      dispatch(updateFromValue(fromValueString))
      setShowValue(fromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/

      if (regex.test(inputValue) || inputValue === '') {
        dispatch(updateFromValue(inputValue))
        setShowValue(inputValue)
      }
    }
  }

  const calculateMaxBridgeableGas = (
    parsedGasBalance: number,
    parsedGasCost: number
  ): number => {
    const maxBridgeable = parsedGasBalance - parsedGasCost
    return maxBridgeable
  }

  const onAvailableBalance = () => {
    dispatch(
      updateFromValue(
        formatBigIntToString(rawBalance, fromToken?.decimals[fromChainId])
      )
    )
  }

  const maxBridgeableGas: number | null =
    isGasToken && parsedGasCost
      ? calculateMaxBridgeableGas(parseFloat(parsedBalance), parsedGasCost)
      : null

  const onMaxBridgeableBalance = useCallback(() => {
    if (maxBridgeableGas) {
      if (maxBridgeableGas < 0) {
        dispatch(
          updateFromValue(
            formatBigIntToString(0n, fromToken?.decimals[fromChainId])
          )
        )
      } else {
        dispatch(updateFromValue(maxBridgeableGas.toString()))
      }
    } else {
      dispatch(
        updateFromValue(
          formatBigIntToString(rawBalance, fromToken?.decimals[fromChainId])
        )
      )
    }
  }, [
    fromChainId,
    fromToken,
    isGasToken,
    parsedGasCost,
    rawBalance,
    trimmedParsedBalance,
  ])

  const showMaxButton = (): boolean => {
    if (!hasMounted || !isConnected) return false
    if (isGasToken && isNull(parsedGasCost)) return false
    return true
  }

  const showGasReserved = (): boolean => {
    if (!hasMounted || !isConnected) return false
    if (!parsedGasCost) return false
    if (isGasToken && !isEmpty(fromValue) && !hasOnlyZeroes(fromValue)) {
      return true
    }
  }

  const isGasInputMoreThanBridgeableMax = (): boolean => {
    if (isGasToken && parsedGasCost && fromValue && parsedBalance) {
      return parseFloat(fromValue) > parseFloat(parsedBalance) - parsedGasCost
    } else {
      return false
    }
  }

  const isGasBalanceLessThanCost = (): boolean => {
    if (isGasToken && parsedGasCost && parsedBalance) {
      return parsedGasCost > parseFloat(parsedBalance)
    } else {
      return false
    }
  }

  const isTraceBalance = (): boolean => {
    if (!rawBalance || !trimmedParsedBalance) return false
    if (rawBalance && hasOnlyZeroes(trimmedParsedBalance)) return true
    return false
  }

  return (
    <div
      data-test-id="input-container"
      className="text-left rounded-md p-md bg-bgLight"
    >
      <div className="flex items-center justify-between mb-3">
        <FromChainSelector />
        {connectedStatus}
      </div>
      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center justify-between pl-md w-full h-16
            rounded-md border border-white border-opacity-20
          `}
        >
          <div className="flex items-center">
            <FromTokenSelector />
            <div className="flex flex-col justify-between ml-4">
              <div style={{ display: 'table' }}>
                <input
                  ref={inputRef}
                  value={showValue}
                  placeholder="0.0000"
                  pattern="^[0-9]*[.,]?[0-9]*$"
                  onChange={handleFromValueChange}
                  className={`
                    border-none bg-transparent max-w-[190px] p-0 font-medium text-opacity-80
                    placeholder:text-[#88818C] text-white text-xl md:text-2xl
                    focus:outline-none focus:ring-0 focus:border-none
                  `}
                  style={{ display: 'table-cell', width: '100%' }}
                  name="inputRow"
                  autoComplete="off"
                  minLength={1}
                  maxLength={79}
                  disabled={false}
                />
              </div>
              {hasMounted &&
                isConnected &&
                (isGasToken &&
                showGasReserved() &&
                isGasInputMoreThanBridgeableMax() ? (
                  <label
                    htmlFor="inputRow"
                    className={`
                      text-xs text-secondaryTextColor transition-all duration-150 transform-gpu
                      ${
                        (isGasBalanceLessThanCost() ||
                          isGasInputMoreThanBridgeableMax()) &&
                        'text-yellow-500'
                      }
                    `}
                  >
                    {parsedGasCost.toFixed(4)}
                    <span className="text-opacity-50"> estimated gas cost</span>
                  </label>
                ) : (
                  <label
                    htmlFor="inputRow"
                    onClick={onAvailableBalance}
                    className={`
                      text-xs text-white transition-all duration-150 transform-gpu
                      hover:text-opacity-70 hover:cursor-pointer
                    `}
                  >
                    {isTraceBalance()
                      ? '< 0.0001'
                      : trimmedParsedBalance ?? '0.0'}
                    <span className="text-opacity-50 text-secondaryTextColor">
                      {' '}
                      available
                    </span>
                  </label>
                ))}
            </div>
          </div>
          <div>
            {showMaxButton() && (
              <div className="m">
                <MiniMaxButton
                  disabled={
                    !rawBalance || rawBalance === 0n
                      ? true
                      : false || isGasBalanceLessThanCost()
                  }
                  onClickBalance={onMaxBridgeableBalance}
                />
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
