import bscImg from '@assets/chains/bnb.svg'
import metisImg from '@assets/chains/metis.svg'
import moonbeamImg from '@assets/chains/moonbeam.svg'
import moonriverImg from '@assets/chains/moonriver.svg'
import harmonyImg from '@assets/chains/harmony.svg'
import cantoImg from '@assets/chains/canto.svg'
import { zeroAddress } from 'viem'

import { Token } from '@/utils/types'
import * as CHAINS from '@/constants/chains/master'

export type GasToken = {
  addresses: { [x: number]: string }
  chainId: number
  decimals: number
  symbol: string
  name: string
  icon: any
}

// export const BNB = new Token({
//   addresses: {
//     [CHAINS.BNB.id]: zeroAddress,
//   },
//   decimals: 18,
//   symbol: 'BNB',
//   name: 'Binance Coin',
//   priorityRank: 1,
//   logo: bscImg,
// })

export const BNB_Gas: GasToken = {
  addresses: {
    [CHAINS.BNB.id]: zeroAddress,
  },
  chainId: 56,
  decimals: 18,
  symbol: 'BNB',
  name: 'Binance Coin',
  icon: bscImg,
}

export const METIS: GasToken = {
  addresses: {
    [CHAINS.METIS.id]: zeroAddress,
  },
  chainId: 1088,
  name: 'Metis',
  symbol: 'METIS',
  decimals: 18,
  icon: metisImg,
}

export const NOTE: GasToken = {
  addresses: {
    [CHAINS.CANTO.id]: zeroAddress,
  },
  chainId: 7700,
  name: 'Canto',
  symbol: 'NOTE',
  decimals: 18,
  icon: cantoImg,
}

export const MOVR: GasToken = {
  addresses: {
    [CHAINS.MOONRIVER.id]: zeroAddress,
  },
  chainId: 1285,
  name: 'Moonriver',
  symbol: 'MOVR',
  decimals: 18,
  icon: moonriverImg,
}

export const GLMR: GasToken = {
  addresses: {
    [CHAINS.MOONBEAM.id]: zeroAddress,
  },
  chainId: 1284,
  name: 'Glimmer',
  symbol: 'GLMR',
  decimals: 18,
  icon: moonbeamImg,
}

export const ONE: GasToken = {
  addresses: {
    [CHAINS.HARMONY.id]: zeroAddress,
  },
  chainId: 1666600000,
  name: 'Harmony One',
  symbol: 'ONE',
  decimals: 18,
  icon: harmonyImg,
}