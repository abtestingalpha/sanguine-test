import { createSchema } from '@ponder/core'

export default createSchema((p) => ({
  AppConfigV1: p.createTable({
    id: p.hex(),
    requiredResponses: p.int(),
    optimisticPeriod: p.int(),
    modules: p.hex().list(),
  }),

  InterchainBatch: p.createTable({
    id: p.hex(),
    batchRoot: p.hex(),
    srcDbNonce: p.bigint(),
    srcChainId: p.int().optional(),
    dstDbNonce: p.bigint().optional(),
    dstChainId: p.int().optional(),
    interchainTransactions: p.many('InterchainTransaction.interchainBatchId'),
    status: p.string(),
    verifiedAt: p.bigint().optional(),
    appConfigId: p.hex().references('AppConfigV1.id').optional(),
    appConfig: p.one('appConfigId'),
  }),

  InterchainTransactionSent: p.createTable({
    id: p.string(),
    srcChainId: p.int(),
    name: p.string(),
    transactionId: p.string(),
    dbNonce: p.bigint(),
    entryIndex: p.bigint(),
    dstChainId: p.int(),
    srcSender: p.string(),
    dstReceiver: p.string(),
    verificationFee: p.bigint(),
    executionFee: p.bigint(),
    options: p.string(),
    message: p.string(),
    address: p.string(),
    blockNumber: p.bigint(),
    transactionHash: p.string(),
    timestamp: p.bigint(),
    count: p.int(),
  }),

  InterchainTransactionReceived: p.createTable({
    id: p.string(),
    dstChainId: p.int(),
    name: p.string(),
    transactionId: p.string(),
    dbNonce: p.bigint(),
    entryIndex: p.bigint(),
    srcChainId: p.int(),
    srcSender: p.string(),
    dstReceiver: p.string(),
    address: p.string(),
    blockNumber: p.bigint(),
    transactionHash: p.string(),
    timestamp: p.bigint(),
    count: p.int(),
  }),

  InterchainTransaction: p.createTable({
    id: p.string(),
    srcChainId: p.int(),
    dstChainId: p.int(),
    srcSender: p.string(),
    dstReceiver: p.string(),
    sentAt: p.bigint().optional(),
    receivedAt: p.bigint().optional(),
    createdAt: p.bigint().optional(),
    updatedAt: p.bigint().optional(),
    interchainTransactionSentId: p
      .string()
      .references('InterchainTransactionSent.id')
      .optional(),
    interchainTransactionSent: p.one('interchainTransactionSentId'),
    interchainTransactionReceivedId: p
      .string()
      .references('InterchainTransactionReceived.id')
      .optional(),
    interchainTransactionReceived: p.one('interchainTransactionReceivedId'),
    status: p.string().optional(),
    interchainBatchId: p.hex().references('InterchainBatch.id').optional(),
    interchainBatch: p.one('interchainBatchId'),
  }),
}))
