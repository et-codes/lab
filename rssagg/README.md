# RSS Feed Aggregator

Server project from [Boot dev's YouTube video](https://www.youtube.com/watch?v=dpXhDzgUSe4).

Allows users to add RSS Feeds to the database. Server will check every minute for new posts. Users can retreive a list of posts from the feeds that they want to follow.

This project makes use of some interesting tools:
- [`sqlc`](https://github.com/sqlc-dev/sqlc) to generate DB access functions from SQL queries
- [`goose`](https://github.com/pressly/goose) DB migration tool
