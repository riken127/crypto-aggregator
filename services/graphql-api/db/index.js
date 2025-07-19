const { Pool } = require('pg');
const pool = new Pool({ connectionString: process.env.POSTGRES_DSN });

module.exports = pool;