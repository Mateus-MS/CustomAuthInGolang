-- Use the database
\c users;

-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the tb_users table
CREATE TABLE tb_users (
  id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  
  username      VARCHAR(50),
  password_hash VARCHAR(50),

  session_token VARCHAR(44),
  csrf_token    VARCHAR(44)
);