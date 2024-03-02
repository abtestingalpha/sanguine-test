import _ from 'lodash'

import { PAUSED_TO_CHAIN_IDS } from '@/constants/chains'

import type { Token } from '@/utils/types'
import { flattenPausedTokens } from '@/utils/tokens/flattenPausedTokens'
import { findTokenByRouteSymbol } from '@/utils/tokens/findTokenByRouteSymbol'
import { getSymbol } from '@/utils/route/getSymbol'

import { getToChainIds } from './getToChainIds'
import { getFromChainIds } from './getFromChainIds'
import { getFromTokens } from './getFromTokens'
import { getToTokens } from './getToTokens'


export const getRoutePossibilities = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}) => {
  const fromTokenRouteSymbol = fromToken && fromToken.routeSymbol
  const toTokenRouteSymbol = toToken && toToken.routeSymbol

  const fromChainIds: number[] = getFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })

  const fromTokens: Token[] = _(
    getFromTokens({
      fromChainId,
      fromTokenRouteSymbol,
      toChainId,
      toTokenRouteSymbol,
    })
  )
    .difference(flattenPausedTokens())
    .map(getSymbol)
    .uniq()
    .map((symbol) => findTokenByRouteSymbol(symbol))
    .compact()
    .value()

  const toChainIds: number[] = getToChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })
    ?.filter((chainId) => !PAUSED_TO_CHAIN_IDS.includes(chainId))
    .filter((chainId) => chainId !== fromChainId)

  const toTokens: Token[] = _(
    getToTokens({
      fromChainId,
      fromTokenRouteSymbol,
      toChainId,
      toTokenRouteSymbol,
    })
  )
    .difference(flattenPausedTokens())
    .filter((token) => {
      return !PAUSED_TO_CHAIN_IDS.some((value) => token.endsWith(`-${value}`))
    })
    .map(getSymbol)
    .uniq()
    .map((symbol) => findTokenByRouteSymbol(symbol))
    .compact()
    .value()

  return {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toChainIds,
    toTokens,
  }
}

