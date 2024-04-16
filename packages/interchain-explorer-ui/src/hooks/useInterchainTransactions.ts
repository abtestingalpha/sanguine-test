import { useQuery } from '@tanstack/react-query'
import { GraphQLClient, gql } from 'graphql-request'

import { type InterchainTransaction } from '@/types'
import { GET_INTERCHAIN_TRANSACTIONS } from '@/graphql/queries'

const client = new GraphQLClient('https://sanguine-production.up.railway.app')

type InterchainTransactionsResponse = {
  interchainTransactions: {
    pageInfo: {
      startCursor: string | null
      endCursor: string | null
      hasPreviousPage: boolean
      hasNextPage: boolean
    }
    items: InterchainTransaction[]
  }
}

export const useInterchainTransactions = ({
  limit,
  after,
  before,
}: {
  limit?: number | null
  after?: string | null
  before?: string | null
}) => {
  return useQuery({
    queryKey: ['interchain-transactions', limit, after, before],
    queryFn: async () => {
      try {
        const variables = { limit, after, before }

        const data = (await client.request(
          GET_INTERCHAIN_TRANSACTIONS,
          variables
        )) as InterchainTransactionsResponse

        const pageInfo = data.interchainTransactions.pageInfo

        const items = data.interchainTransactions.items.map((d) => ({
          id: d.id,
          interchainTransactionSent: d.interchainTransactionSent,
          interchainTransactionReceived: d.interchainTransactionReceived,
        }))

        return { pageInfo, items }
      } catch (error) {
        console.error('Error fetching interchain transactions:', error)
        throw error
      }
    },
    staleTime: 60_000, // 1 minute
    refetchInterval: 1_000, // 1 second
  })
}
