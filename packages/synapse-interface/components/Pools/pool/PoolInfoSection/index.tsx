import numeral from 'numeral'
import {
  formatBigIntToPercentString,
  formatBigIntToString,
} from '@/utils/bigint/format'
import { usePoolDataState } from '@/slices/pool/hooks'
import LoadingDots from '@tw/LoadingDots'
import Grid from '@tw/Grid'
import AugmentWithUnits from '../components/AugmentWithUnits'
import InfoSectionCard from './InfoSectionCard'
import CurrencyReservesCard from './CurrencyReservesCard'


const PoolInfoSection = () => {
  const { pool, poolData } = usePoolDataState()

  const usdFormat = poolData.totalLockedUSD > 1000000 ? '$0,0.0' : '$0,0'

  return (
    <Grid cols={{xs: 1}} gap={8}>
      <CurrencyReservesCard />
      <InfoSectionCard title="Pool Info">
        <InfoListItem
          labelText="Trading Fee"
          content={
            poolData && poolData.swapFee ? (
              formatBigIntToPercentString(poolData.swapFee, 8, 2, false)
            ) : (
              <LoadingDots />
            )
          }
        />
        <InfoListItem
          labelText="Virtual Price"
          content={
            poolData && poolData?.virtualPrice ? (
              <AugmentWithUnits
                content={formatBigIntToString(poolData.virtualPrice, 18, 5)}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingDots />
            )
          }
        />
        <InfoListItem
          labelText="Total Liquidity"
          content={
            poolData && poolData?.totalLocked ? (
              <AugmentWithUnits
                content={numeral(poolData.totalLocked).format('0,0')}
                label={pool.priceUnits}
              />
            ) : (
              <LoadingDots />
            )
          }
        />
        <InfoListItem
          labelText="Total Liquidity USD"
          content={
            poolData && poolData?.totalLockedUSD ? (
              <AugmentWithUnits
                content={numeral(poolData.totalLockedUSD).format(usdFormat)}
                label="USD"
              />
            ) : (
              <LoadingDots />
            )
          }
        />
      </InfoSectionCard>
    </Grid>
  )
}

const InfoListItem = ({
  labelText,
  content,
}: {
  labelText: string
  content: any
}) => {
  return (
    <li className="flex w-full py-2">
      <div className="text-white/60">{labelText} </div>
      <div className="self-center ml-auto text-white">{content}</div>
    </li>
  )
}

export default PoolInfoSection
