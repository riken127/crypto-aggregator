const pool = require('./index');
const SQL = require('sql-template-strings');
const redis = require('../redis');

/**
 * Fetches the top assets by trading volume in the last 24 hours.
 * @param {*} limit - The maximum number of assets to return. Default is 10.
 * @returns {Promise<Array>} - A promise that resolves to an array of asset objects.
 */
async function getTopAssetsByVolume(limit = 10) {
  const cacheKey = `topAssets:${limit}`;
  const cached = await redis.redisClient.get(cacheKey);

  if (cached) return JSON.parse(cached);

  const { rows } = await pool.query(SQL`
    SELECT a.id, a.symbol, a.name, r.price_usd, r.volume_usd24_hr, r.timestamp
    FROM assets a
    JOIN LATERAL (
      SELECT * FROM asset_records r2
      WHERE r2.asset_id = a.id
      ORDER BY r2.timestamp DESC
      LIMIT 1
    ) r ON true
    ORDER BY r.volume_usd24_hr::float DESC
    LIMIT ${limit}
  `);

  const result = rows.map(row => ({
    id: row.id,
    symbol: row.symbol,
    name: row.name,
    priceUsd: row.price_usd,
    volumeUsd24Hr: row.volume_usd24_hr,
    timestamp: row.timestamp.toString()
  }));

  await redis.redisClient.set(cacheKey, JSON.stringify(result), 'EX', 60);

  return result;
}

/**
 * Fetches the historical data of an asset by its ID.
 * @param {*} assetId  - The ID of the asset to fetch history for.
 * @param {*} limit  - The maximum number of records to return. Default is 30.
 * @returns  {Promise<Array>} - A promise that resolves to an array of asset history records.
 */
async function getAssetHistory(assetId, limit = 30) {
  const { rows } = await pool.query(SQL`
    SELECT price_usd, volume_usd24_hr, timestamp
    FROM asset_records
    WHERE asset_id = ${assetId}
    ORDER BY timestamp DESC
    LIMIT ${limit}
  `);

  return rows.map(row => ({
    priceUsd: row.price_usd,
    volumeUsd24Hr: row.volume_usd24_hr,
    timestamp: row.timestamp.toString()
  }));
}


module.exports = { getTopAssetsByVolume, getAssetHistory };