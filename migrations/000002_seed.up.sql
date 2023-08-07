INSERT INTO lessons (title)
VALUES
    ('JavaScript Lesson'),
    ('SQL Lesson');

INSERT INTO questions(lesson_id, text)
VALUES
    (1, 'In JS what is the result of [] + []?'),
    (1, 'In JS what is the result of [] + {}?'),
    (1, 'In JS what is the result of {} + []?'),
    (1, 'In JS what is the result of {} + {}?'),
    (2, 'Does SQLite3 support json data types?'),
    (2, 'How can a developer set a default timestamp on a SQLite3 column?'), 
    (2, 'Should SQLite3 be used as a production database?'), 
    (2, 'What tool can be used to backup SQLite3 in production?');

INSERT INTO options(question_id, option_type, text)
VALUES
    (1, 'distractor', '[]'),
    (1, 'distractor', '[[]]'),
    (1, 'correct', 'empty string'),
    (2, 'distractor', '{}'),
    (2, 'distractor', '[object Object][object Object]'),
    (2, 'correct', '[object Object]'),
    (3, 'distractor', '[object Object]'),
    (3, 'distractor', '{}'),
    (3, 'correct', '0'),
    (4, 'distractor', '[object Object]'),
    (4, 'distractor', '{}'),
    (4, 'correct', 'NaN'),
    (5, 'distractor', 'Yes'),
    (5, 'correct', 'No'),
    (6, 'distractor', 'Use the DEFAULT keyword'),
    (6, 'distractor', 'Use the NOW() function'),
    (6, 'distractor', 'Use the CURRENT_TIMESTAMP function'),
    (6, 'correct', 'Use the DEFAULT CURRENT_TIMESTAMP function'),
    (7, 'distractor', 'Yes'),
    (7, 'distractor', 'No'),
    (7, 'correct', 'It depends'),
    (8, 'distractor', 'pg_dump'),
    (8, 'distractor', 'mysqldump'),
    (8, 'distractor', 'sqlite3dump'),
    (8, 'correct', 'sqlite3 .dump');

INSERT INTO users(first_name, last_name, email, password)
VALUES
    ('John', 'Doe', 'jon.d@fakemail.com', '$2a$10$JfIboTnNhj.lkmxwUO49g.0mdOsVjXcJkOjFF0Dl700HHztSG3mSS'),
    ('Jane', 'Doe', 'jane.d@fakemail.com', '$2a$10$JfIboTnNhj.lkmxwUO49g.0mdOsVjXcJkOjFF0Dl700HHztSG3mSS');

INSERT INTO answers(user_id, question_id, option_id)
VALUES
    (1, 1, 1),
    (1, 2, 2),
    (1, 3, 8),
    (1, 4, 11),
    (1, 5, 14),
    (1, 6, 19),
    (1, 7, 21),
    (1, 8, 25),
    (2, 1, 3),
    (2, 2, 6),
    (2, 3, 9),
    (2, 4, 12),
    (2, 5, 14),
    (2, 6, 19),
    (2, 7, 21),
    (2, 8, 25);

