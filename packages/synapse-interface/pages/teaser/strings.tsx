const Chains = [
  'Ethereum',
  'Arbitrum',
  'Avalanche',
  'Base',
  'BNB Chain',
  'Optimism',
  'Polygon',
  'DFK Chain',
  'Canto',
  'Fantom',
  'Klaytn',
  'Aurora',
  'Boba Chain',
  'Cronos',
  'Metis',
  'Moonbeam',
  'Moonriver',
]

const Tokens = [
  'USDC',
  'USDT',
  'DAI',
  'crvUSD',
  'FRAX',
  'LUSD',
  'JEWEL',
  'AVAX',
  'BTC.b',
  'AVAX',
  'LINK',
  'ETH',
  'SYN',
  'NOTE',
  'sUSD',
  'GMX',
  'MATIC',
  'MOVR',
  'gOHM',
  'H2O',
  'L2DAO',
  'NEWO',
  'PLS',
  'SDT',
  'SFI',
  'UNIDX',
  'veSOLAR',
  'VSTA',
]

const formatAmount = (amount) => {
  let [, left, right] = amount.toString().match(/(\d+)\.?(\d*)/) ?? ['', '', '']

  for (let i = 3; i < left.length; i += 4)
    left = `${left.slice(0, left.length - i)},${left.slice(-i)}`

  right += '0000'

  return left.length < 4
    ? left + '.' + right.slice(0, left === '0' ? 4 : 4 - left.length)
    : left
}

const randHex = () => {
  const x = () => Math.round(Math.random() * 16).toString(16)
  return `#${x() + x()}…${x() + x() + x() + x()}`
}

export const generateTx = () => {
  let originAmount =
    Math.random() < 0.5 ? Math.round(Math.random() * 10000) : Math.random()
  let destinationAmount = (originAmount * (100 - Math.random() * 5)) / 100

  const origin = {
    payload: Tokens[Math.round(Math.random() * (Tokens.length - 1))],
    chain: Chains[Math.round(Math.random() * (Chains.length - 1))],
    originAmount,
    formattedAmount: formatAmount(originAmount),
    timestamp: Date.now(),
    hash: randHex(),
  }
  const destination = {
    payload: origin.payload,
    chain: Chains[Math.round(Math.random() * (Chains.length - 1))],
    destinationAmount,
    formattedAmount: formatAmount(destinationAmount),
    timestamp: origin.timestamp + Math.round(Math.random() * 600000),
    hash: randHex(),
  }

  return { origin, destination }
}

// TODO: Should be removed and entire file moved to non-pages directory since this is helper functions
export default function Strings() {
  return <div>Strings</div>
}
