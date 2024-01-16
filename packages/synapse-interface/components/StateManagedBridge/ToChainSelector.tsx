import React from 'react'
import { useDispatch } from 'react-redux'

import { setShowToChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { useBridgeState } from '@/slices/bridge/hooks'
import { CHAINS_BY_ID } from '@/constants/chains'
import { DropDownArrowSvg } from '../icons/DropDownArrowSvg'
import {
  getNetworkButtonBgClassNameActive,
  getNetworkButtonBorderActive,
  getNetworkButtonBorderHover,
  getNetworkHover,
} from '@/styles/chains'

export const ToChainSelector = () => {
  const dispatch = useDispatch()
  const { toChainId } = useBridgeState()
  const toChain = CHAINS_BY_ID[toChainId]

  return (
    <button
      data-test-id="bridge-origin-chain-list-button"
      className={`
        flex items-center gap-1.5
        p-2 rounded
        border border-transparent
        active:opacity-70
        ${getNetworkHover(toChain?.color)}
        ${getNetworkButtonBorderHover(toChain?.color)}
      `}
      onClick={() => dispatch(setShowToChainListOverlay(true))}
    >
      {toChainId && (
        <img
          src={toChain?.chainImg?.src}
          alt={toChain?.name}
          className="w-6 h-6 rounded-sm"
        />
      )}
      <dl className="text-left">
        <dt className="text-sm opacity-50">To</dt>
        <dd>{toChain?.name || 'Network'}</dd>
      </dl>
      <DropDownArrowSvg className="mx-0.5" />
    </button>
  )
}
