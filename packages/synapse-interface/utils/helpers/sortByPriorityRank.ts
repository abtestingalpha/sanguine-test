import _ from 'lodash'

import type { Token } from '@/utils/types'

export const sortByPriorityRank = (tokens: Token[]) => {
  return _.orderBy(
    tokens,
    [(token) => token.priorityRank, (token) => token.symbol.toLowerCase()],
    ['asc', 'asc']
  )
}
