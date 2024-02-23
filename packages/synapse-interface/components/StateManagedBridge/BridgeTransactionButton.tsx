import { useMemo, useEffect, useState } from 'react'
import { useAccount, useNetwork, useSwitchNetwork } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'

import { isAddress } from '@ethersproject/address'

import { stringToBigInt } from '@/utils/bigint/format'
import { useBridgeDisplayState, useBridgeState } from '@/slices/bridge/hooks'
import { usePortfolioBalances } from '@/slices/portfolio/hooks'
import { PAUSED_FROM_CHAIN_IDS, PAUSED_TO_CHAIN_IDS } from '@/constants/chains'
import { EMPTY_BRIDGE_QUOTE, EMPTY_BRIDGE_QUOTE_ZERO } from '@/constants/bridge'

import { TransactionButton } from '@/components/buttons/TransactionButton'

export const BridgeTransactionButton = ({
  approveTxn,
  executeBridge,
  isApproved,
}) => {
  const [isConnected, setIsConnected] = useState(false)
  const { openConnectModal } = useConnectModal()

  const { chain } = useNetwork()
  const { chains, error, pendingChainId, switchNetwork } = useSwitchNetwork()

  const { address, isConnected: isConnectedInit } = useAccount({
    onDisconnect() {
      setIsConnected(false)
    },
  })

  useEffect(() => {
    setIsConnected(isConnectedInit)
  }, [isConnectedInit])

  const {
    destinationAddress,
    fromToken,
    fromValue,
    toToken,
    fromChainId,
    toChainId,
    isLoading,
    bridgeQuote,
  } = useBridgeState()

  const { showDestinationAddress } = useBridgeDisplayState()

  const balances = usePortfolioBalances()
  const balancesForChain = balances[fromChainId]
  const balanceForToken = balancesForChain?.find(
    (t) => t.tokenAddress === fromToken?.addresses[fromChainId]
  )?.balance

  const sufficientBalance = useMemo(() => {
    if (!fromChainId || !fromToken || !toChainId || !toToken) return false
    return (
      stringToBigInt(fromValue, fromToken?.decimals[fromChainId]) <=
      balanceForToken
    )
  }, [balanceForToken, fromValue, fromChainId, toChainId, toToken])

  const isButtonDisabled =
    isLoading ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE_ZERO ||
    bridgeQuote === EMPTY_BRIDGE_QUOTE ||
    (destinationAddress && !isAddress(destinationAddress)) ||
    (showDestinationAddress && !destinationAddress) ||
    (isConnected && !sufficientBalance) ||
    PAUSED_FROM_CHAIN_IDS.includes(fromChainId) ||
    PAUSED_TO_CHAIN_IDS.includes(toChainId)

  const fromTokenDecimals: number | undefined =
    fromToken && fromToken?.decimals[fromChainId]

  const fromValueBigInt = useMemo(() => {
    return fromTokenDecimals ? stringToBigInt(fromValue, fromTokenDecimals) : 0
  }, [fromValue, fromTokenDecimals])

  let buttonProperties

  if (!fromChainId) {
    buttonProperties = {
      label: 'Please select Origin network',
    }
  } else if (!toChainId) {
    buttonProperties = {
      label: 'Please select Destination network',
    }
  } else if (
    PAUSED_FROM_CHAIN_IDS.includes(fromChainId) ||
    PAUSED_TO_CHAIN_IDS.includes(toChainId)
  ) {
    buttonProperties = {
      label: `Bridge unavailable`,
    }
  } else if (!fromToken) {
    buttonProperties = {
      label: `Unsupported Network`,
    }
  } else if (
    !isLoading &&
    bridgeQuote?.feeAmount === 0n &&
    fromValueBigInt > 0
  ) {
    buttonProperties = {
      label: `Amount must be greater than fee`,
    }
  } else if (!isConnected && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Connect Wallet to Bridge`,
      onClick: openConnectModal,
    }
  } else if (isConnected && !sufficientBalance) {
    buttonProperties = {
      label: 'Insufficient balance',
    }
  } else if (showDestinationAddress && !destinationAddress) {
    buttonProperties = {
      label: 'Please add valid destination address',
    }
  } else if (destinationAddress && !isAddress(destinationAddress)) {
    buttonProperties = {
      label: 'Invalid destination address',
    }
  } else if (chain?.id != fromChainId && fromValueBigInt > 0) {
    buttonProperties = {
      label: `Switch to ${chains.find((c) => c.id === fromChainId)?.name}`,
      onClick: () => switchNetwork(fromChainId),
      pendingLabel: 'Switching chains',
    }
  } else if (!isApproved) {
    buttonProperties = {
      onClick: approveTxn,
      label: `Approve ${fromToken?.symbol}`,
      pendingLabel: 'Approving',
    }
  } else {
    buttonProperties = {
      onClick: executeBridge,
      label: `Bridge ${fromToken?.symbol}`,
      pendingLabel: 'Bridging',
    }
  }

  return (
    <TransactionButton
      {...buttonProperties}
      disabled={isButtonDisabled}
      chainId={fromChainId}
    />
  )
}
