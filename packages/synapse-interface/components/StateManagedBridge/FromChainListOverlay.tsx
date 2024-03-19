import _ from 'lodash'
import { useCallback, useEffect, useRef, useState } from 'react'
import Fuse from 'fuse.js'
import { useKeyPress } from '@hooks/useKeyPress'
import * as ALL_CHAINS from '@constants/chains/master'
import SlideSearchBox from '@pages/bridge/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { useDispatch } from 'react-redux'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import useCloseOnOutsideClick from '@/utils/hooks/useCloseOnOutsideClick'
import { CloseButton } from './components/CloseButton'
import { SearchResults } from './components/SearchResults'
import { PAUSED_FROM_CHAIN_IDS } from '@constants/chains'

/*
export const FromChainListOverlay = () => {
  const { fromChainIds, fromChainId } = useBridgeState()
  const [currentIdx, setCurrentIdx] = useState(-1)
  const [searchStr, setSearchStr] = useState('')
  const [open, setOpen] = useState(true)

  const dispatch = useDispatch()
  const dataId = 'bridge-origin-chain-list'
  const overlayRef = useRef(null)

  let possibleChains = _(ALL_CHAINS)
    .pickBy((value) => _.includes(fromChainIds, value.id))
    .values()
    .value()
    .filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
      fromChainIds.map((id) => CHAINS_BY_ID[id])
    )
  ).filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  const possibleChainsWithSource = possibleChains.map((chain) => ({
    ...chain,
    source: 'possibleChains',
  }))

  const remainingChainsWithSource = remainingChains.map((chain) => ({
    ...chain,
    source: 'remainingChains',
  }))

  const masterList = [...possibleChainsWithSource, ...remainingChainsWithSource]

  const fuseOptions = {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'id',
      'nativeCurrency.symbol',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter((item) => item.source === 'possibleChains')
    remainingChains = results.filter(
      (item) => item.source === 'remainingChains'
    )
  }

  const escPressed = useKeyPress('Escape')
  const arrowUp = useKeyPress('ArrowUp')
  const arrowDown = useKeyPress('ArrowDown')

  const onClose = useCallback(() => {
    setCurrentIdx(-1)
    setSearchStr('')
    setOpen(false)
  }, [setShowFromChainListOverlay])

  const escFunc = () => {
    if (escPressed) {
      onClose()
    }
  }
  const arrowDownFunc = () => {
    const nextIdx = currentIdx + 1
    if (arrowDown && nextIdx < masterList.length) {
      setCurrentIdx(nextIdx)
    }
  }

  const arrowUpFunc = () => {
    const nextIdx = currentIdx - 1
    if (arrowUp && -1 < nextIdx) {
      setCurrentIdx(nextIdx)
    }
  }

  const onSearch = (str: string) => {
    setSearchStr(str)
    setCurrentIdx(-1)
  }

  useEffect(arrowDownFunc, [arrowDown])
  useEffect(escFunc, [escPressed])
  useEffect(arrowUpFunc, [arrowUp])
  useCloseOnOutsideClick(overlayRef, onClose)

  const handleSetFromChainId = (chainId) => {
    const eventTitle = `[Bridge User Action] Sets new fromChainId`
    const eventData = {
      previousFromChainId: fromChainId,
      newFromChainId: chainId,
    }

    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setFromChainId(chainId))
    onClose()
  }

  // useEffect(() => {
  //   const ref = overlayRef.current
  //   const { y, height } = ref.getBoundingClientRect()
  //   const screen = window.innerHeight
  //   if (y + height > screen) {
  //     ref.style.position = 'fixed'
  //     ref.style.bottom = '4px'
  //   }
  //   if (y < 0) {
  //     ref.style.position = 'fixed'
  //     ref.style.top = '4px'
  //   }
  // }, [])

  return (
    // <div
    //   ref={overlayRef}
    //   data-test-id="fromChain-list-overlay"
    //   className={`${
    //     open ? 'block' : 'hidden'
    //   } z-20 absolute bg-bgLight border border-separator rounded overflow-y-auto max-h-96 animate-slide-down origin-top shadow-md`}
    // >
    //   <div className="p-1 flex items-center font-medium">
    //     <SlideSearchBox
    //       placeholder="Find"
    //       searchStr={searchStr}
    //       onSearch={onSearch}
    //     />
    //     <CloseButton onClick={onClose} />
    //   </div>
    <div data-test-id={dataId}>
      {possibleChains && possibleChains.length > 0 && (
        <>
          <div className="p-2 text-sm text-secondary sticky top-0 bg-bgLight">
            From…
          </div>
          {possibleChains.map(({ id: mapChainId }, idx) => {
            return (
              <SelectSpecificNetworkButton
                key={idx}
                itemChainId={mapChainId}
                isCurrentChain={fromChainId === mapChainId}
                isOrigin={true}
                active={idx === currentIdx}
                onClick={() =>
                  fromChainId === mapChainId
                    ? onClose()
                    : handleSetFromChainId(mapChainId)
                }
                dataId={dataId}
              />
            )
          })}
        </>
      )}
      {remainingChains && remainingChains.length > 0 && (
        <div className="bg-bgBase rounded">
          <div className="px-2 py-2 text-sm text-secondary sticky top-0 bg-bgBase">
            All chains
          </div>
          {remainingChains.map(({ id: mapChainId }, idx) => {
            return (
              <SelectSpecificNetworkButton
                key={mapChainId}
                itemChainId={mapChainId}
                isCurrentChain={fromChainId === mapChainId}
                isOrigin={true}
                active={idx + possibleChains.length === currentIdx}
                onClick={() => handleSetFromChainId(mapChainId)}
                dataId={dataId}
                alternateBackground={true}
              />
            )
          })}
        </div>
      )}
      <SearchResults searchStr={searchStr} type="chain" />
    </div>
    // </div>
  )
}
*/

export function FromChainListArray(searchStr: string = '') {
  const { fromChainIds } = useBridgeState()

  let possibleChains = _(ALL_CHAINS)
    .pickBy((value) => _.includes(fromChainIds, value.id))
    .values()
    .value()
    .filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  possibleChains = sortChains(possibleChains)

  let remainingChains = sortChains(
    _.difference(
      Object.keys(CHAINS_BY_ID).map((id) => CHAINS_BY_ID[id]),
      fromChainIds.map((id) => CHAINS_BY_ID[id])
    )
  ).filter((chain) => !PAUSED_FROM_CHAIN_IDS.includes(chain.id))

  const possibleChainsWithSource = possibleChains.map((chain) => ({
    ...chain,
    source: 'possibleChains',
  }))

  const remainingChainsWithSource = remainingChains.map((chain) => ({
    ...chain,
    source: 'remainingChains',
  }))

  const masterList = [...possibleChainsWithSource, ...remainingChainsWithSource]

  const fuseOptions = {
    includeScore: true,
    threshold: 0.0,
    keys: [
      {
        name: 'name',
        weight: 2,
      },
      'id',
      'nativeCurrency.symbol',
    ],
  }

  const fuse = new Fuse(masterList, fuseOptions)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter((item) => item.source === 'possibleChains')
    remainingChains = results.filter(
      (item) => item.source === 'remainingChains'
    )
  }

  return { possibleChains, remainingChains }
}
