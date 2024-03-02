import _ from 'lodash'

import type { RouteQueryFields } from '@/utils/route/types'
import { getTokenAndChainId } from '@utils/route/getTokenAndChainId'

export const getFromChainIdsFunc = (routes: _.Dictionary<any>) => {
  return ({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  }: RouteQueryFields) => {
    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .keys()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .keys()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .keys()
        .filter((key) => {
          const { symbol } = getTokenAndChainId(key)
          return symbol === fromTokenRouteSymbol
        })
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .keys()
        .filter((key) => {
          const { symbol } = getTokenAndChainId(key)
          return symbol === fromTokenRouteSymbol
        })
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .entries()
        .filter(([_key, values]) =>
          values.some((v) => v.endsWith(`-${toChainId}`))
        )
        .map(([key]) => getTokenAndChainId(key).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .entries()
        .filter(([_key, values]) =>
          values.some((v) => v.endsWith(`-${toChainId}`))
        )
        .map(([key]) => getTokenAndChainId(key).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .entries()
        .filter(([key, _values]) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .filter(([_key, values]) =>
          values.some((v) => getTokenAndChainId(v).chainId === toChainId)
        )
        .map(([key, _values]) => key)
        .filter((token) => token.endsWith(`-${toChainId}`))
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol &&
      toChainId &&
      toTokenRouteSymbol === null
    ) {
      return _(routes)
        .pickBy((_values, key) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .keys()
        .map((token) => getTokenAndChainId(token).chainId)
        .filter((chainId) => chainId !== toChainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .chain()
        .filter((values, _key) => {
          return values.some((v) => {
            const { symbol } = getTokenAndChainId(v)
            return symbol === toTokenRouteSymbol
          })
        })
        .flatten()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .chain()
        .filter((values, _key) => {
          return values.some((v) => {
            const { symbol } = getTokenAndChainId(v)
            return symbol === toTokenRouteSymbol
          })
        })
        .flatten()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .entries()
        .filter(([key, _values]) => key.startsWith(`${fromTokenRouteSymbol}-`))
        .map(([key, _values]) => key)
        .flatten()
        .filter((token) => token.startsWith(`${toTokenRouteSymbol}-`))
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol &&
      toChainId === null &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .chain()
        .filter((values, key) => {
          return (
            values.some((v) => {
              const { symbol } = getTokenAndChainId(v)
              return symbol === toTokenRouteSymbol
            }) && key.startsWith(`${fromTokenRouteSymbol}-`)
          )
        })
        .flatten()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .pickBy((values, _key) => {
          return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
        })
        .keys()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId &&
      fromTokenRouteSymbol === null &&
      toChainId &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .pickBy((values, _key) => {
          return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
        })
        .keys()
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (
      fromChainId === null &&
      fromTokenRouteSymbol &&
      toChainId &&
      toTokenRouteSymbol
    ) {
      return _(routes)
        .pickBy((values, _key) => {
          return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
        })
        .keys()
        .filter((k) => k.startsWith(`${fromTokenRouteSymbol}-`))
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }

    if (fromChainId && fromTokenRouteSymbol && toChainId && toTokenRouteSymbol) {
      return _(routes)
        .pickBy((values, _key) => {
          return _.includes(values, `${toTokenRouteSymbol}-${toChainId}`)
        })
        .keys()
        .filter((k) => k.startsWith(`${fromTokenRouteSymbol}-`))
        .map((token) => getTokenAndChainId(token).chainId)
        .uniq()
        .value()
    }
  }

}

