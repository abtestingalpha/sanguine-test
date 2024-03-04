import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import type { Address } from 'wagmi'

import {
  fetchPoolUserData,
  resetPoolUserData,
} from '@/slices/poolUserDataSlice'
import { usePoolDataState, usePoolUserDataState } from '@/slices/pool/hooks'

import LoadingDots from '@tw/LoadingDots'
import LiquidityManagementTabs from '../components/LiquidityManagementTabs'
import Deposit from './Deposit'
import Withdraw from './Withdraw'

const PoolManagement = ({
  address,
  chainId,
}: {
  address: Address
  chainId: number
}) => {
  const [cardNav, setCardNav] = useState(getLiquidityMode('#addLiquidity')) // 'addLiquidity'

  const { pool } = usePoolDataState()
  const { poolUserData, isLoading } = usePoolUserDataState()

  const dispatch: any = useDispatch()

  useEffect(() => {
    if (pool && address) {
      dispatch(resetPoolUserData())
      dispatch(fetchPoolUserData({ pool, address }))
    }
  }, [pool, address])

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <LoadingDots />
      </div>
    )
  }

  return (
    <div>
      <LiquidityManagementTabs
        cardNav={cardNav}
        setCardNav={(val) => {
          setCardNav(val)
        }}
      />
      <div className="pb-3 pl-4 pr-4">
        <div className="mt-8">
          {cardNav === 'addLiquidity' && poolUserData.tokens && (
            <Deposit address={address} chainId={chainId} />
          )}
          {cardNav === 'removeLiquidity' && <Withdraw address={address} />}
        </div>
      </div>
    </div>
  )
}

function getLiquidityMode(browserHash) {
  switch (browserHash) {
    case '#addLiquidity':
      return 'addLiquidity'
    case '#removeLiquidity':
      return 'removeLiquidity'
    default:
      return 'addLiquidity'
  }
}
export default PoolManagement
