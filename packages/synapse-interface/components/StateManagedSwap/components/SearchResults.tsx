export const SearchResults = ({
  searchStr,
  type,
}: {
  searchStr: string
  type: string
}) => {
  return (
    <div>
      {searchStr ? (
        <div className="px-12 py-4 text-center text-md">
          No other results found for{' '}
          <i className="text-opacity-60">{searchStr}</i>.
          <div className="pt-2 align-bottom text-md">
            Want to see a {type} supported on Synapse? Let us know!
          </div>
        </div>
      ) : null}
    </div>
  )
}
