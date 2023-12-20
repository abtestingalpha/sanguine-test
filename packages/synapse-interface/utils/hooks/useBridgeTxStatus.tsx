import { useState, useEffect } from 'react'
import { useSynapseContext } from '@/utils/providers/SynapseProvider'

interface UseBridgeTxStatusProps {
  originChainId: number
  destinationChainId: number
  transactionHash: string
  bridgeModuleName?: string
  kappa?: string
}

export const useBridgeTxStatus = ({
  originChainId,
  destinationChainId,
  transactionHash,
  bridgeModuleName,
  kappa,
}: UseBridgeTxStatusProps) => {
  const [isComplete, setIsComplete] = useState<boolean>(false)
  const { synapseSDK } = useSynapseContext()

  const getKappa = async (): Promise<string> => {
    if (!bridgeModuleName || !originChainId || !transactionHash) return
    return await synapseSDK.getSynapseTxId(
      originChainId,
      bridgeModuleName,
      transactionHash
    )
  }

  const getBridgeTxStatus = async (
    destinationChainId: number,
    bridgeModuleName: string,
    kappa: string
  ) => {
    if (!destinationChainId || !bridgeModuleName || !kappa) return null
    return await synapseSDK.getBridgeTxStatus(
      destinationChainId,
      bridgeModuleName,
      kappa
    )
  }

  useEffect(() => {
    console.log('running')
    ;(async () => {
      let _kappa

      if (!kappa) {
        _kappa = await getKappa()
      } else {
        _kappa = kappa
      }

      const txStatus = await getBridgeTxStatus(
        destinationChainId,
        bridgeModuleName,
        _kappa
      )

      if (txStatus !== null) {
        setIsComplete(txStatus)
      }
    })()
  }, [])

  return isComplete
}
