CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    CHECK(length(password) >= 8)
);

CREATE TABLE IF NOT EXISTS lessons (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS questions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    lesson_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    FOREIGN KEY (lesson_id)
        REFERENCES lessons (id)
); 

CREATE TABLE IF NOT EXISTS options (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    question_id INTEGER NOT NULL,
    option_type INTEGER NOT NULL,
    text TEXT NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    FOREIGN KEY (question_id)
        REFERENCES questions (id),
    CHECK (option_type = 'distractor' OR option_type = 'correct')
); 

CREATE TABLE IF NOT EXISTS answers (
    user_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    option_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    PRIMARY KEY (user_id, question_id, option_id)
    FOREIGN KEY (user_id)
        REFERENCES users (id),
    FOREIGN KEY (question_id)
        REFERENCES questions (id),
    FOREIGN KEY (option_id)
        REFERENCES options (id)
);