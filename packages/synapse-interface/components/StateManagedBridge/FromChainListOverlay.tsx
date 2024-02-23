import _ from 'lodash'
import { useDispatch } from 'react-redux'
import Fuse from 'fuse.js'
import * as ALL_CHAINS from '@constants/chains/master'
import { SlideSearchBox } from '@components/bridgeSwap/SlideSearchBox'
import { CHAINS_BY_ID, sortChains } from '@constants/chains'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'
import { useBridgeState } from '@/slices/bridge/hooks'
import { setFromChainId } from '@/slices/bridge/reducer'
import { setShowFromChainListOverlay } from '@/slices/bridgeDisplaySlice'
import { SelectSpecificNetworkButton } from './components/SelectSpecificNetworkButton'
import { CloseButton } from '@/components/buttons/CloseButton'
import { NoSearchResultsFound } from '@/components/bridgeSwap/NoSearchResultsFound'
import { PAUSED_FROM_CHAIN_IDS } from '@constants/chains'
import { useOverlaySearch } from '@/utils/hooks/useOverlaySearch'
import { CHAIN_FUSE_OPTIONS } from '@/constants/fuseOptions'
import { SearchResultsContainer } from '@/components/bridgeSwap/SearchResultsContainer'
import { SearchOverlayContent } from '@/components/bridgeSwap/SearchOverlayContent'


export const FromChainListOverlay = () => {
  const { fromChainIds, fromChainId } = useBridgeState()
  const dispatch = useDispatch()
  const dataId = 'bridge-origin-chain-list'

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

  function onCloseOverlay() {
    dispatch(setShowFromChainListOverlay(false))
  }

  const {
    overlayRef,
    onSearch,
    currentIdx,
    searchStr,
    onClose,
  } = useOverlaySearch(masterList.length, onCloseOverlay)



  const fuse = new Fuse(masterList, CHAIN_FUSE_OPTIONS)

  if (searchStr?.length > 0) {
    const results = fuse.search(searchStr).map((i) => i.item)

    possibleChains = results.filter(item => item.source === 'possibleChains')
    remainingChains = results.filter(item => item.source === 'remainingChains')
  }


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

  return (
    <SearchOverlayContent
      overlayRef={overlayRef}
      searchStr={searchStr}
      onSearch={onSearch}
      onClose={onClose}
      type="chain"
    >
        {possibleChains?.length > 0 && (
          <SearchResultsContainer label="From…">
            {possibleChains.map(({ id: mapChainId }, idx) =>
                <SelectSpecificNetworkButton
                  key={idx}
                  itemChainId={mapChainId}
                  isCurrentChain={fromChainId === mapChainId}
                  isOrigin={true}
                  active={idx === currentIdx}
                  onClick={() => {
                    if (fromChainId === mapChainId) {
                      onClose()
                    } else {
                      handleSetFromChainId(mapChainId)
                    }
                  }}
                  dataId={dataId}
                />
            )}
          </SearchResultsContainer>
        )}
        {remainingChains?.length > 0 && (
          <SearchResultsContainer label="All Chains">
            {remainingChains.map(({ id: mapChainId }, idx) =>
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
            )}
          </SearchResultsContainer>
        )}
    </SearchOverlayContent>
  )
}
