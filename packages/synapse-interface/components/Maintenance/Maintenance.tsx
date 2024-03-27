import { OPTIMISM, BASE, DOGE } from '@/constants/chains/master'
import { MaintenanceBanner } from './components/MaintenanceBanner'
import { MaintenanceWarningMessage } from './components/MaintenanceWarningMessage'
import { useMaintenanceCountdownProgress } from './components/useMaintenanceCountdownProgress'
import { WarningMessage } from '../Warning'

interface ChainPause {
  id: string
  pausedFromChains: number[]
  pausedToChains: number[]
  startTime: Date
  endTime: Date | null // Indefinite if null
  bannerStartTime: Date
  bannerEndTime: Date | null // Indefinite if null
  warningMessage: any
  bannerMessage: any
  progressBarMessage: any
  disableBanner?: boolean
  disableWarning?: boolean
  disableCountdown?: boolean
}

const PAUSED_CHAINS: ChainPause[] = [
  {
    id: 'optimism-chain-pause',
    pausedFromChains: [OPTIMISM.id],
    pausedToChains: [],
    startTime: new Date(Date.UTC(2024, 2, 21, 18, 0, 0)),
    endTime: new Date(Date.UTC(2024, 2, 21, 19, 40, 0)),
    bannerStartTime: new Date(Date.UTC(2024, 2, 21, 18, 0, 0)),
    bannerEndTime: new Date(Date.UTC(2024, 2, 21, 19, 40, 0)),
    warningMessage: (
      <p> Optimism bridging is paused until maintenance is complete. </p>
    ),
    bannerMessage: (
      <p> Optimism bridging is paused until maintenance is complete. </p>
    ),
    progressBarMessage: <p> Optimism maintenance in progress </p>,
  },
  {
    id: 'base-chain-pause',
    pausedFromChains: [BASE.id],
    pausedToChains: [BASE.id],
    startTime: new Date(Date.UTC(2024, 2, 21, 17, 41, 0)),
    endTime: new Date(Date.UTC(2024, 2, 21, 17, 42, 0)),
    bannerStartTime: new Date(Date.UTC(2024, 2, 21, 17, 40, 0)),
    bannerEndTime: new Date(Date.UTC(2024, 2, 21, 17, 43, 0)),
    warningMessage: (
      <p> Base bridging is paused until maintenance is complete. </p>
    ),
    bannerMessage: (
      <p> Base bridging is paused until maintenance is complete. </p>
    ),
    progressBarMessage: <p> Base maintenance in progress </p>,
  },
  {
    id: 'base-chain-pause',
    pausedFromChains: [BASE.id],
    pausedToChains: [BASE.id],
    startTime: new Date(Date.UTC(2024, 2, 21, 17, 41, 0)),
    endTime: null,
    bannerStartTime: new Date(Date.UTC(2024, 2, 27, 4, 40, 0)),
    bannerEndTime: null,
    warningMessage: (
      <p> Base bridging is paused until maintenance is complete. </p>
    ),
    bannerMessage: (
      <p className="m-auto">
        Base bridging is paused until maintenance is complete.
      </p>
    ),
    progressBarMessage: <p> Base maintenance in progress </p>,
    disableBanner: true,
    disableWarning: true,
    disableCountdown: true,
  },
  {
    id: 'doge-chain-pause',
    pausedFromChains: [],
    pausedToChains: [DOGE.id],
    startTime: new Date(Date.UTC(2024, 2, 21, 17, 41, 0)),
    endTime: null,
    bannerStartTime: new Date(Date.UTC(2024, 2, 27, 4, 40, 0)),
    bannerEndTime: null,
    warningMessage: (
      <div className="space-y-2">
        <strong>
          Alert: Transactions to Dogechain are temporarily paused.
        </strong>
        <p>
          You may still bridge funds from Dogechain to any supported destination
          chain.
        </p>
      </div>
    ),
    bannerMessage: (
      <p className="m-auto">Bridging to Doge Chain is currently unavailable.</p>
    ),
    progressBarMessage: (
      <p> Bridging to Doge Chain is currently unavailable. </p>
    ),
    disableBanner: true,
    disableWarning: false,
    disableCountdown: false,
  },
]

export const MaintenanceBanners = () => {
  return (
    <>
      {PAUSED_CHAINS.map((event) => {
        return (
          <MaintenanceBanner
            id={event.id}
            bannerMessage={event.bannerMessage}
            startDate={event.bannerStartTime}
            endDate={event.bannerEndTime}
            disabled={event.disableBanner}
          />
        )
      })}
    </>
  )
}

export const MaintenanceWarningMessages = () => {
  return (
    <>
      {PAUSED_CHAINS.map((event) => {
        return (
          <MaintenanceWarningMessage
            startDate={event.startTime}
            endDate={event.endTime}
            pausedFromChains={event.pausedFromChains}
            pausedToChains={event.pausedToChains}
            warningMessage={event.warningMessage}
            disabled={event.disableWarning}
          />
        )
      })}
    </>
  )
}

/**
 * Hook that maps through PAUSED_CHAINS to apply the single chain countdown progress logic to each.
 * @returns Array of objects containing maintenance status and components for each paused chain.
 */
export const useMaintenanceCountdownProgresses = () => {
  return PAUSED_CHAINS.map((event) => {
    return useMaintenanceCountdownProgress({
      startDate: event.startTime,
      endDate: event.endTime,
      pausedFromChains: event.pausedFromChains,
      pausedToChains: event.pausedToChains,
      progressBarMessage: event.progressBarMessage,
      disabled: event.disableCountdown,
    })
  })
}
