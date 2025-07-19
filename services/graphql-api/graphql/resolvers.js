const assetDb = require('../db/asset');

module.exports = {
  topAssets: ({ limit }) => assetDb.getTopAssetsByVolume(limit || 10),
  assetHistory: ({ assetId, limit }) => assetDb.getAssetHistory(assetId, limit || 30)
};