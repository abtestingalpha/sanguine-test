import { useMemo } from 'react'
import { useAccount, useNetwork } from 'wagmi'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { MetamaskIcon } from '@icons/WalletIcons/Metamask'
import { CoinbaseWalletIcon } from '@icons/WalletIcons/CoinbaseWalletIcon'
import { WalletConnectIcon } from '@icons/WalletIcons/WalletConnectIcon'
import Spinner from '@icons/Spinner'
import { IconProps, WalletId } from '@utils/types'
import { useHasMounted } from '@/utils/hooks/useHasMounted'


const WALLETS = [
  {
    id: WalletId.MetaMask,
    icon: MetamaskIcon,
  },
  {
    id: WalletId.WalletConnect,
    icon: WalletConnectIcon,
  },
  {
    id: WalletId.CoinbaseWallet,
    icon: CoinbaseWalletIcon,
  },
]

export const WalletIcon = ({
  walletId,
  ...rest
}: IconProps): JSX.Element | null => {
  const SelectedIcon = Object.values(WALLETS).find(
    ({ id }) => id === walletId
  )?.icon

  return SelectedIcon ? <SelectedIcon {...rest} /> : null
}

const SHARED_BTN_CLASSNAME = `
  outline-none active:outline-none
  text-white transition-all duration-100
  rounded-md
  py-2 px-2.5 bg-bgBase/10 hover:bg-opacity-70
  hover:bg-bgBase/20 active:bg-bgBase/20 text-sm
  border border-bgBase/10 hover:border-bgBase/20
  whitespace-nowrap
  `

export const Wallet = () => {
  const { connector: activeConnector, address: connectedAddress } = useAccount()
  const { chain: currentChain } = useNetwork()
  const walletId = activeConnector?.id

  const hasMounted = useHasMounted()

  const render = useMemo(() => {
    return (
      <ConnectButton.Custom>
        {({
          account,
          chain,
          openAccountModal,
          openChainModal,
          openConnectModal,
          authenticationStatus,
          mounted,
        }) => {
          const ready = mounted && authenticationStatus !== 'loading'
          return (
            <div
              {...(!ready && {
                'aria-hidden': true,
              })}
            >
              {(() => {
                if (!connectedAddress) {
                  return (
                    <WalletButton
                      onClick={openConnectModal}
                    >
                      Connect Wallet
                    </WalletButton>
                  )
                }
                if (currentChain?.unsupported || chain?.unsupported) {
                  return (
                    <WalletButton
                      onClick={openChainModal}
                      className="bg-red-500/50"
                    >
                      Wrong Network
                    </WalletButton>
                  )
                }
                return (
                  <div className="flex gap-3">
                    <WalletButton
                      onClick={openChainModal}
                      className={`flex items-center gap-2`}
                    >
                      {account?.displayBalance ? (
                        account.displayBalance
                      ) : (
                        <Spinner />
                      )}
                      {chain && chain.hasIcon && (
                        <div
                          style={{
                            backgroundImage: chain.iconBackground,
                            width: 20,
                            height: 20,
                            borderRadius: 999,
                            overflow: 'hidden',
                            backgroundPosition: 'center',
                            color: '#ffffff',
                          }}
                        >
                          {chain.iconUrl && (
                            <img
                              alt={chain.name ?? 'Chain icon'}
                              src={chain.iconUrl}
                              style={{ width: 20, height: 20 }}
                            />
                          )}
                        </div>
                      )}
                    </WalletButton>

                    <WalletButton
                      onClick={openAccountModal}
                      className={`font-bold`}
                    >
                      {account ? account.displayName : <Spinner />}
                    </WalletButton>
                  </div>
                )
              })()}
            </div>
          )
        }}
      </ConnectButton.Custom>
    )
  }, [connectedAddress, currentChain, walletId])

  return hasMounted && render
}




function WalletButton({className="", children, ...props}) {
  return (
    <button
      type="button"
      className={`
        ${SHARED_BTN_CLASSNAME}
        ${className}
      `}
      {...props}
    >
      {children}
    </button>
  )
}