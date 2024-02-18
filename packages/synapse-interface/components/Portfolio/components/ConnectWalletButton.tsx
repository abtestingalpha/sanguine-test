
import { useAccount } from 'wagmi'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useHasMounted } from '@/utils/hooks/useHasMounted'

export function ConnectWalletButton() {
  const clientReady = useHasMounted()

  const { address } = useAccount()


  const buttonClassName = `
    h-10 border-[#CA5CFF] border-[1.5px] flex items-center border
    text-base text-white px-6 py-5 hover:opacity-75 rounded-lg
    text-center transform-gpu transition-all duration-75
  `

  const buttonStyle = {
    borderRadius: '30px',
  }

  return (
    <div id="connect-wallet-button">
      {clientReady && (
        <ConnectButton.Custom>
          {({ account, chain, openConnectModal, mounted, openChainModal }) => {
            return (
              <>
                {(() => {
                  if (!mounted || !account || !chain || !address) {
                    return (
                      <button
                        className={buttonClassName}
                        style={buttonStyle}
                        onClick={openConnectModal}
                      >
                        Connect Wallet
                      </button>
                    )
                  }
                })()}
              </>
            )
          }}
        </ConnectButton.Custom>
      )}
    </div>
  )
}
