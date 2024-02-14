import { useBridgeState } from '@/slices/bridge/hooks'

import LoadingDots from '@tw/LoadingDots'

import { ToChainSelector } from './ToChainSelector'
import { ToTokenSelector } from './ToTokenSelector'


export const OutputContainer = () => {
  const { bridgeQuote, isLoading } = useBridgeState()

  return (
    <div className="mt-[1.125rem] p-md text-left rounded-md bg-bgBase/10 ring-1 ring-white/10">
      <div className="flex items-center justify-between mb-3">
        <ToChainSelector />
      </div>

      <div className="flex h-16 mb-2 space-x-2">
        <div
          className={`
            flex flex-grow items-center
            pl-md
            w-full h-16
            rounded-md
            border border-transparent
          `}
        >
          <ToTokenSelector />
          <div className="flex ml-4">
            {isLoading ? (
              <LoadingDots className="opacity-50" />
            ) : (
              // <BridgeSwapOutputNumber
              //   quote={bridgeQuote}
              // />
              <input
                pattern="[0-9.]+"
                disabled={true}
                className={`
                  focus:outline-none
                  focus:ring-0
                  focus:border-none
                  border-none
                  p-0
                  bg-transparent
                  max-w-[200px]
                placeholder:text-[#88818C]
                text-white/90 text-2xl md:text-3xl font-medium
                -mt-1
                `}
                placeholder="0.0000"
                value={
                  bridgeQuote?.outputAmountString === '0'
                    ? ''
                    : bridgeQuote?.outputAmountString
                }
                name="inputRow"
                autoComplete="off"
              />
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

