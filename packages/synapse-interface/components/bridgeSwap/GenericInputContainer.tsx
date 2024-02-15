import React, { useEffect, useState, useCallback } from 'react'

import { useAccount } from 'wagmi'


import { formatBigIntToString } from '@/utils/bigint/format'
import { cleanNumberInput } from '@/utils/cleanNumberInput'

import { usePortfolioState } from '@/slices/portfolio/hooks'
import { ConnectStatusIndicator } from '@/components/bridgeSwap/ConnectStatusIndicator'

import MiniMaxButton from '@/components/buttons/MiniMaxButton'

export const GenericInputContainer = ({
  inputRef,
  chainId,
  token,
  value,
  initialStateValue,
  dispatchUpdateFunc,
  chainSelector,
  tokenSelector,
}) => {
  const [showValue, setShowValue] = useState('')

  const { balances } = usePortfolioState()

  const { isConnected } = useAccount()

  const tokenData = balances[chainId]?.find(
    (t) => t.tokenAddress === token?.addresses[chainId]
  )

  const { parsedBalance, balance } = tokenData ?? {}

  useEffect(() => {
    if (token && token?.decimals[chainId]) {
      setShowValue(value)
    }

    if (value === initialStateValue) {
      setShowValue(initialStateValue)
    }
  }, [value, inputRef, chainId, token])

  const handleFromValueChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const fromValueString: string = cleanNumberInput(event.target.value)

    try {
      dispatchUpdateFunc(fromValueString)
      setShowValue(fromValueString)
    } catch (error) {
      console.error('Invalid value for conversion to BigInteger')
      const inputValue = event.target.value
      const regex = /^[0-9]*[.,]?[0-9]*$/
      if (regex.test(inputValue) || inputValue === '') {
        dispatchUpdateFunc(inputValue)
        setShowValue(inputValue)
      }
    }
  }

  const onMaxBalance = useCallback(() => {
    dispatchUpdateFunc(
      formatBigIntToString(balance, token?.decimals[chainId])
    )

  }, [balance, chainId, token])


  return (
    <div
      data-test-id="input-container"
      className="text-left rounded-md p-md bg-bgBase/10 ring-1 ring-white/10"
    >
      <div className="flex items-center justify-between mb-3">
        {chainSelector}
        <ConnectStatusIndicator targetChainId={chainId} />
      </div>
      <div className="flex h-16 mb-2 space-x-2">
        <div
          onClick={() => inputRef.current?.focus()}
          className={`
            flex flex-grow items-center justify-between
            pl-md
            w-full h-16
            rounded-md
            hover:bg-bgBase/[0.05] hover:border-white/20
            focus-within:!bg-bgBase/10
            border border-white/10 focus-within:!border-white/40
          `}
        >
          <div className="flex items-center">
            {tokenSelector}
            <div className="flex flex-col justify-between">
              <div className="pl-4 -mt-0.5" >
                <input
                  ref={inputRef}
                  pattern="^[0-9]*[.,]?[0-9]*$"
                  disabled={false}
                  className={`
                    focus:outline-none
                    focus:ring-0
                    focus:border-none
                    border-none
                    bg-transparent
                    p-0 -mb-2
                    placeholder:text-white/40
                    text-white/90 text-xl md:text-2xl font-medium
                  `}
                  placeholder="0.0000"
                  onChange={handleFromValueChange}
                  value={showValue}
                  name="inputRow"
                  autoComplete="off"
                  minLength={1}
                  maxLength={79}
                  style={{ display: 'table-cell', width: '100%' }}
                />
                <div>
                  {isConnected && (
                    <label
                      htmlFor="inputRow"
                      className="group text-xs text-white/90 transition-all transform-gpu hover:text-white/50 hover:cursor-pointer"
                      onClick={onMaxBalance}
                    >
                      {parsedBalance ?? '0.0'}
                      <span className="text-white/40 group-hover:text-white/30">
                        {' '}
                        available
                      </span>
                    </label>
                  )}
                </div>
              </div>

            </div>
          </div>
          <div>
            {isConnected && (
              <div className="ml-1.5">
                <MiniMaxButton
                  disabled={!balance || balance === 0n}
                  onClickBalance={onMaxBalance}
                />
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
