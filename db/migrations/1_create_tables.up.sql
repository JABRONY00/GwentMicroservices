-- Players table
CREATE TABLE players (
    id VARCHAR(40) NOT NULL,
    name VARCHAR(20) UNIQUE NOT NULL,
    email VARCHAR(100) NOT NULL,
    password_hash BYTEA NOT NULL,
   -- profile_id SERIAL,

    PRIMARY KEY (id)
   -- FOREIGN KEY (profile_id) REFERENCES profiles(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- CREATE TABLE profiles (

-- );

-- CREATE TABLE presets(
--     user_id VARCHAR(40) NOT NULL,
--     race VARCHAR(4) NOT NULL,
--     stack []smallint,
--     PRIMARY KEY (user_id, race),
--     FOREIGN KEY user_id REFERENCES players(id) ON DELETE CASCADE ON UPDATE CASCADE
-- );