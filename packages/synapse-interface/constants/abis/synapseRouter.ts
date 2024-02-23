export const SYNAPSE_ROUTER_ABI = [
  {
    inputs: [
      {
        internalType: 'address',
        name: '_synapseBridge',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'previousOwner',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'newOwner',
        type: 'address',
      },
    ],
    name: 'OwnershipTransferred',
    type: 'event',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'tokenIn',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amountIn',
        type: 'uint256',
      },
      {
        internalType: 'address',
        name: 'tokenOut',
        type: 'address',
      },
      {
        internalType: 'bytes',
        name: 'rawParams',
        type: 'bytes',
      },
    ],
    name: 'adapterSwap',
    outputs: [
      {
        internalType: 'uint256',
        name: 'amountOut',
        type: 'uint256',
      },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'string',
        name: 'symbol',
        type: 'string',
      },
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'enum LocalBridgeConfig.TokenType',
        name: 'tokenType',
        type: 'uint8',
      },
      {
        internalType: 'address',
        name: 'bridgeToken',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'bridgeFee',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: 'minFee',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: 'maxFee',
        type: 'uint256',
      },
    ],
    name: 'addToken',
    outputs: [
      {
        internalType: 'bool',
        name: 'wasAdded',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: 'string',
            name: 'id',
            type: 'string',
          },
          {
            internalType: 'address',
            name: 'token',
            type: 'address',
          },
          {
            internalType: 'enum LocalBridgeConfig.TokenType',
            name: 'tokenType',
            type: 'uint8',
          },
          {
            internalType: 'address',
            name: 'bridgeToken',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'bridgeFee',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'minFee',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'maxFee',
            type: 'uint256',
          },
        ],
        internalType: 'struct LocalBridgeConfig.BridgeTokenConfig[]',
        name: 'tokens',
        type: 'tuple[]',
      },
    ],
    name: 'addTokens',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'allPools',
    outputs: [
      {
        components: [
          {
            internalType: 'address',
            name: 'pool',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'lpToken',
            type: 'address',
          },
          {
            components: [
              {
                internalType: 'bool',
                name: 'isWeth',
                type: 'bool',
              },
              {
                internalType: 'address',
                name: 'token',
                type: 'address',
              },
            ],
            internalType: 'struct PoolToken[]',
            name: 'tokens',
            type: 'tuple[]',
          },
        ],
        internalType: 'struct Pool[]',
        name: 'pools',
        type: 'tuple[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'chainId',
        type: 'uint256',
      },
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        components: [
          {
            internalType: 'address',
            name: 'swapAdapter',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'tokenOut',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'minAmountOut',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'deadline',
            type: 'uint256',
          },
          {
            internalType: 'bytes',
            name: 'rawParams',
            type: 'bytes',
          },
        ],
        internalType: 'struct SwapQuery',
        name: 'originQuery',
        type: 'tuple',
      },
      {
        components: [
          {
            internalType: 'address',
            name: 'swapAdapter',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'tokenOut',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'minAmountOut',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'deadline',
            type: 'uint256',
          },
          {
            internalType: 'bytes',
            name: 'rawParams',
            type: 'bytes',
          },
        ],
        internalType: 'struct SwapQuery',
        name: 'destQuery',
        type: 'tuple',
      },
    ],
    name: 'bridge',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'bridgeTokens',
    outputs: [
      {
        internalType: 'address[]',
        name: 'tokens',
        type: 'address[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'bridgeTokensAmount',
    outputs: [
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'pool',
        type: 'address',
      },
      {
        internalType: 'uint256[]',
        name: 'amounts',
        type: 'uint256[]',
      },
    ],
    name: 'calculateAddLiquidity',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'calculateBridgeFee',
    outputs: [
      {
        internalType: 'uint256',
        name: 'feeAmount',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'pool',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'calculateRemoveLiquidity',
    outputs: [
      {
        internalType: 'uint256[]',
        name: 'amountsOut',
        type: 'uint256[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'pool',
        type: 'address',
      },
      {
        internalType: 'uint8',
        name: 'tokenIndexFrom',
        type: 'uint8',
      },
      {
        internalType: 'uint8',
        name: 'tokenIndexTo',
        type: 'uint8',
      },
      {
        internalType: 'uint256',
        name: 'dx',
        type: 'uint256',
      },
    ],
    name: 'calculateSwap',
    outputs: [
      {
        internalType: 'uint256',
        name: 'amountOut',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'pool',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'tokenAmount',
        type: 'uint256',
      },
      {
        internalType: 'uint8',
        name: 'tokenIndex',
        type: 'uint8',
      },
    ],
    name: 'calculateWithdrawOneToken',
    outputs: [
      {
        internalType: 'uint256',
        name: 'amountOut',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    name: 'config',
    outputs: [
      {
        internalType: 'enum LocalBridgeConfig.TokenType',
        name: 'tokenType',
        type: 'uint8',
      },
      {
        internalType: 'address',
        name: 'bridgeToken',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    name: 'fee',
    outputs: [
      {
        internalType: 'uint40',
        name: 'bridgeFee',
        type: 'uint40',
      },
      {
        internalType: 'uint104',
        name: 'minFee',
        type: 'uint104',
      },
      {
        internalType: 'uint112',
        name: 'maxFee',
        type: 'uint112',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'tokenIn',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'tokenOut',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amountIn',
        type: 'uint256',
      },
    ],
    name: 'getAmountOut',
    outputs: [
      {
        components: [
          {
            internalType: 'address',
            name: 'swapAdapter',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'tokenOut',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'minAmountOut',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'deadline',
            type: 'uint256',
          },
          {
            internalType: 'bytes',
            name: 'rawParams',
            type: 'bytes',
          },
        ],
        internalType: 'struct SwapQuery',
        name: '',
        type: 'tuple',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'tokenOut',
        type: 'address',
      },
    ],
    name: 'getConnectedBridgeTokens',
    outputs: [
      {
        components: [
          {
            internalType: 'string',
            name: 'symbol',
            type: 'string',
          },
          {
            internalType: 'address',
            name: 'token',
            type: 'address',
          },
        ],
        internalType: 'struct BridgeToken[]',
        name: 'tokens',
        type: 'tuple[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        components: [
          {
            internalType: 'string',
            name: 'symbol',
            type: 'string',
          },
          {
            internalType: 'uint256',
            name: 'amountIn',
            type: 'uint256',
          },
        ],
        internalType: 'struct DestRequest[]',
        name: 'requests',
        type: 'tuple[]',
      },
      {
        internalType: 'address',
        name: 'tokenOut',
        type: 'address',
      },
    ],
    name: 'getDestinationAmountOut',
    outputs: [
      {
        components: [
          {
            internalType: 'address',
            name: 'swapAdapter',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'tokenOut',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'minAmountOut',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'deadline',
            type: 'uint256',
          },
          {
            internalType: 'bytes',
            name: 'rawParams',
            type: 'bytes',
          },
        ],
        internalType: 'struct SwapQuery[]',
        name: 'destQueries',
        type: 'tuple[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'tokenIn',
        type: 'address',
      },
      {
        internalType: 'string[]',
        name: 'tokenSymbols',
        type: 'string[]',
      },
      {
        internalType: 'uint256',
        name: 'amountIn',
        type: 'uint256',
      },
    ],
    name: 'getOriginAmountOut',
    outputs: [
      {
        components: [
          {
            internalType: 'address',
            name: 'swapAdapter',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'tokenOut',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'minAmountOut',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'deadline',
            type: 'uint256',
          },
          {
            internalType: 'bytes',
            name: 'rawParams',
            type: 'bytes',
          },
        ],
        internalType: 'struct SwapQuery[]',
        name: 'originQueries',
        type: 'tuple[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'owner',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'pool',
        type: 'address',
      },
    ],
    name: 'poolInfo',
    outputs: [
      {
        internalType: 'uint256',
        name: '',
        type: 'uint256',
      },
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'pool',
        type: 'address',
      },
    ],
    name: 'poolTokens',
    outputs: [
      {
        components: [
          {
            internalType: 'bool',
            name: 'isWeth',
            type: 'bool',
          },
          {
            internalType: 'address',
            name: 'token',
            type: 'address',
          },
        ],
        internalType: 'struct PoolToken[]',
        name: 'tokens',
        type: 'tuple[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'poolsAmount',
    outputs: [
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
    ],
    name: 'removeToken',
    outputs: [
      {
        internalType: 'bool',
        name: 'wasRemoved',
        type: 'bool',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address[]',
        name: 'tokens',
        type: 'address[]',
      },
    ],
    name: 'removeTokens',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'contract IERC20',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'spender',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'setAllowance',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'contract ISwapQuoter',
        name: '_swapQuoter',
        type: 'address',
      },
    ],
    name: 'setSwapQuoter',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'enum LocalBridgeConfig.TokenType',
        name: 'tokenType',
        type: 'uint8',
      },
      {
        internalType: 'address',
        name: 'bridgeToken',
        type: 'address',
      },
    ],
    name: 'setTokenConfig',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'bridgeFee',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: 'minFee',
        type: 'uint256',
      },
      {
        internalType: 'uint256',
        name: 'maxFee',
        type: 'uint256',
      },
    ],
    name: 'setTokenFee',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        components: [
          {
            internalType: 'address',
            name: 'swapAdapter',
            type: 'address',
          },
          {
            internalType: 'address',
            name: 'tokenOut',
            type: 'address',
          },
          {
            internalType: 'uint256',
            name: 'minAmountOut',
            type: 'uint256',
          },
          {
            internalType: 'uint256',
            name: 'deadline',
            type: 'uint256',
          },
          {
            internalType: 'bytes',
            name: 'rawParams',
            type: 'bytes',
          },
        ],
        internalType: 'struct SwapQuery',
        name: 'query',
        type: 'tuple',
      },
    ],
    name: 'swap',
    outputs: [
      {
        internalType: 'uint256',
        name: 'amountOut',
        type: 'uint256',
      },
    ],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'swapQuoter',
    outputs: [
      {
        internalType: 'contract ISwapQuoter',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'string',
        name: '',
        type: 'string',
      },
    ],
    name: 'symbolToToken',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'synapseBridge',
    outputs: [
      {
        internalType: 'contract ISynapseBridge',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    name: 'tokenToSymbol',
    outputs: [
      {
        internalType: 'string',
        name: '',
        type: 'string',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'newOwner',
        type: 'address',
      },
    ],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    stateMutability: 'payable',
    type: 'receive',
  },
] as const
