# Go Take Home Assignment

## Description

The application you will be working on is a REST web api for displaying lesson data.

There are three routes that have been built for you:

- `GET admin/lessons` returns a list of lesson objects
- `GET admin/lessons/{id}` returns a single lesson and question data
- `POST auth/login` returns a JWT for an authenticated user

The web framework for this project is [go-chi](https://github.com/go-chi/chi). The documentation for the framework is not the best, but it is a framework that embraces a style that is idiomatic to Go. Therefore, you shouldn't have to know much about the framework to get the tasks done.

The data is stored in sqlite3 database and there are two repository structs for interacting with the database. You can get an understanding of the data model by reading the migration sql in the `migrations` directory.

## Your Tasks

1. Create a new endpoint that returns a single lesson, questions data and a users answers. The response should be structured in the following way:

```json
{
    "id": 1,
    "title": "JavaScript Lesson",
    "created_at": "2023-08-07T04:45:32Z",
    "questions": [
        {
            "id": 1,
            "lesson_id": 1,
            "text": "In JS what is the result of [] + []?",
            "created_at": "2023-08-07T04:45:32Z",
            "options": [
                {
                    "id": 1,
                    "option_type": "distractor",
                    "text": "[]",
                    "created_at": "2023-08-07T04:45:32Z"
                },
                {
                    "id": 2,
                    "option_type": "distractor",
                    "text": "[[]]",
                    "created_at": "2023-08-07T04:45:32Z"
                },
                {
                    "id": 3,
                    "option_type": "correct",
                    "text": "empty string",
                    "created_at": "2023-08-07T04:45:32Z"
                }
            ],
            "answer": {
                "question_id": 1,
                "user_id": 1,
                "option": {
                    "id": 3,
                    "option_type": "correct",
                    "created_at": "2023-08-07T04:45:32Z"
                }
            }
        }
    ]
}
```
You have complete freedom to craft the structure of the API path, and code as you see fit. Be prepared to talk about your decisions.

2. Authorize access to the at least one route using a JWT. A user of the API should need to provide a JWT in order to gain access to the data returned by the route. You can choose how the user should provide the JWT, but be prepared to talk about yor decision.
3. (Optional) Write a unit test for any part of the code.
4. Have fun and don't stress.

When you are done with the exercise, please push your code to github and add `jmhammock` as a collaborator.

## Notes

Please don't spend more than **four hours** on this exercise. If you don't get it all done in four hours, that is ok. Just come ready to talk through what you have done.

The code is purposefully simple and doesn't meet the standards of a production application. So, if you see things that could be improved and you have time, please feel free to make improvements that we can talk through!
