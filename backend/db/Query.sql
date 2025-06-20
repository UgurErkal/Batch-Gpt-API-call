CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    ApiKey TEXT NOT NULL 
);

CREATE TABLE IF NOT EXISTS Prompts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER NOT NULL,
    status TEXT DEFAULT 'pending' CHECK(status IN ('pending', 'completed', 'failed')), 
    prompt TEXT NOT NULL,
    createdAtPrompt DATETIME DEFAULT CURRENT_TIMESTAMP,
    reply TEXT DEFAULT NULL,
    createdAtReply DATETIME DEFAULT NULL,
    FOREIGN KEY (userId) REFERENCES Users(id)
);

/* hi comment */
/* use once from here */

INSERT INTO Users (username, ApiKey) VALUES
  ('test2', 'hehe');

INSERT INTO Prompts (userId, status, prompt, reply) VALUES
  (1, 'completed', 'Hello, how are you?', NULL),
  (1, 'pending', 'What is SQLite?', NULL);

/* Drop table if exists */
/*
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Prompts;
*/
