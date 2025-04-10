import dotenv from 'dotenv';
dotenv.config();

import pkg from 'pg';
const { Pool } = pkg;

const pool = new Pool({
  connectionString: process.env.API_CONECT,
  ssl: {
    rejectUnauthorized: false, 
  },
});

pool.on('error', (err) => {
  console.error('❌ PostgreSQL Pool Error:', err);
  process.exit(-1); 
});

export default pool;
