CREATE TABLE IF NOT EXISTS sharecode(
  id VARCHAR(64) PRIMARY KEY,
  code INTEGER UNIQUE NOT NULL,
  open_id VARCHAR(64) NOT NULL,
  session_id VARCHAR(64) NOT NULL,
  expiration_time TIMESTAMP WITH TIME ZONE NOT NULL
);
