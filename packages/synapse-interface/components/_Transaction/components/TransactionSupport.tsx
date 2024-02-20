// TODO: Add FAQ link
export const TransactionSupport = ({ isReverted }: { isReverted: boolean }) => {
  return (
    <div
      id="transaction-support"
      className="flex items-center justify-between w-full py-1 pl-3 pr-1 text-sm"
    >
      {isReverted ? (
        <div>Transaction reverted, funds returned</div>
      ) : (
        <div>What's taking so long?</div>
      )}
      <div className="flex items-center">
        {/* <a
          href=""
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-[--synapse-select-bg] hover:no-underline
          `}
        >
          F.A.Q
        </a>
        <div>/</div> */}
        <a
          href="https://discord.gg/synapseprotocol"
          target="_blank"
          className={`
            px-2 py-1 underline text-[--synapse-text]
            hover:rounded hover:bg-zinc-700 hover:no-underline
          `}
        >
          Support
        </a>
      </div>
    </div>
  )
}
