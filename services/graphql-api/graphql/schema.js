const { buildSchema } = require('graphql');

/**
 * GraphQL schema for the crypto aggregator API.
 * It defines types for Asset and AssetHistory, and queries to fetch top assets and asset history.
 */
module.exports = buildSchema(`
  type Asset {
    id: String
    symbol: String
    name: String
    priceUsd: String
    volumeUsd24Hr: String
    timestamp: String
  }

  type AssetHistory {
    priceUsd: String
    volumeUsd24Hr: String
    timestamp: String
  }

  type Query {
    topAssets(limit: Int): [Asset]
    assetHistory(assetId: String!, limit: Int): [AssetHistory]
  }
`);