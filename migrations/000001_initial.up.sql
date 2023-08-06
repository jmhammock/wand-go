CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at DATETIME DEFAULT current_timestamp,
    CHECK(length(password) >= 8)
);

CREATE TABLE IF NOT EXISTS lessons (
    id TEXT PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT current_timestamp
);

CREATE TABLE IF NOT EXISTS questions (
    id TEXT PRIMARY KEY,
    lesson_id string NOT NULL,
    text TEXT NOT NULL,
    FOREIGN KEY (lesson_id)
        REFERENCES lessons (id)
); 

CREATE TABLE IF NOT EXISTS options (
    id TEXT PRIMARY KEY,
    question_id TEXT,
    option_type TEXT NOT NULL,
    text TEXT NOT NULL,
    FOREIGN KEY (question_id)
        REFERENCES questions (id),
    CHECK (option_type = 'distractor' OR option_type = 'correct')
); 

CREATE TABLE IF NOT EXISTS answers (
    user_id TEXT NOT NULL,
    question_id TEXT NOT NULL,
    option_id TEXT NOT NULL,
    PRIMARY KEY (user_id, question_id, option_id)
    FOREIGN KEY (user_id)
        REFERENCES users (id),
    FOREIGN KEY (question_id)
        REFERENCES questions (id),
    FOREIGN KEY (option_id)
        REFERENCES options (id)
);

CREATE TABLE IF NOT EXISTS lesson_questions (
    lesson_id TEXT NOT NULL,
    question_id TEXT NOT NULL,
    PRIMARY KEY (lesson_id, question_id)
    FOREIGN KEY (lesson_id)
        REFERENCES lessons (id),
    FOREIGN KEY (question_id)
        REFERENCES questions (id)
);

CREATE TABLE IF NOT EXISTS question_options (
    question_id TEXT NOT NULL,
    option_id TEXT NOT NULL,
    PRIMARY KEY (question_id, option_id)
    FOREIGN KEY (question_id)
        REFERENCES questions (id),
    FOREIGN KEY (option_id)
        REFERENCES options (id)
);