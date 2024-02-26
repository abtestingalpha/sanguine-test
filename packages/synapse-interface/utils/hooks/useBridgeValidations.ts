import { useNetwork } from 'wagmi'
import { zeroAddress } from 'viem'

import { useBridgeState } from '@/slices/bridge/hooks'
import { BridgeState, initialState } from '@/slices/bridge/reducer'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { stringToBigInt } from '@/utils/bigint/format'

export const useBridgeValidations = (): {
  hasValidRoute: boolean
  hasEnoughBalance: boolean
  hasInputAmount: boolean
  onSelectedChain: boolean
  hasEnoughApproved: boolean
} => {
  const { chain } = useNetwork()
  const {
    debouncedFromValue,
    fromChainId,
    fromToken,
    isLoading,
    bridgeQuote,
  }: BridgeState = useBridgeState()
  const balances = usePortfolioBalances()

  const currentBalance = balances[fromChainId]?.find(
    (token) => token.token === fromToken
  )?.balance

  const preciseDebouncedFromValue = stringToBigInt(
    debouncedFromValue,
    fromToken?.decimals[fromChainId]
  )

  const isEmpty = debouncedFromValue === initialState.debouncedFromValue
  const isZero = parseFloat(debouncedFromValue) === 0

  const hasValidRoute = !isLoading && bridgeQuote.outputAmount > 0n
  const hasEnoughBalance = currentBalance >= preciseDebouncedFromValue
  const hasInputAmount = !isEmpty && !isZero
  const onSelectedChain = chain?.id === fromChainId
  const hasEnoughApproved =
    bridgeQuote.allowance >= preciseDebouncedFromValue ||
    fromToken?.addresses[fromChainId] === zeroAddress

  return {
    hasValidRoute,
    hasEnoughBalance,
    hasInputAmount,
    onSelectedChain,
    hasEnoughApproved,
  }
}
