# 📰 Blog Aggregator

A command-line application built in Go that allows users to register, log in, and aggregate blog content from various sources. Designed as a backend-focused project to demonstrate database interaction, CLI command routing, and type-safe SQL with [`sqlc`](https://github.com/sqlc-dev/sqlc).

## 📌 Features

- 🧑‍💻 User registration and login system
- 💾 PostgreSQL database integration
- 📥 Add and follow RSS feeds
- 🐦 CLI-based command interface
- 🗂 Store posts in a local database
- 🔁 Goose migrations for schema management

## 🚀 Getting Started

#### 1. Clone the repo

<pre>
git clone https://github.com/kavancamp/blogAggregator
cd blogAggregator
</pre>

#### 2. Run Migrations
<pre>goose -dir sql/schema postgres YOUR_DB_URL up <sub> Replace `YOUR_DB_URL` with your actual PostgreSQL connection string. </sub> </pre>

#### 3. Generate SQL code
<pre>sqlc generate </pre>

#### 4. Build the binary
<pre>go build -o gator .</pre>

## 🛠 Commands
Register a new user
<pre>./gator register USERNAME <sub> Replace USERNAME with the actual users name</sub> </pre>
Add a new Feed <sub> auto follows feed, doesn't allow duplicate urls</sub> 
<pre>./gator addfeed FEED_NAME FEED_URL </pre>
Follow an existing feed
<pre>./gator follow FEED_URL </pre>
Start aggregating (fetching posts) by a set time 
<pre>./gator agg 30s</pre>
Browse recent posts - Shows recent posts from feeds you follow. Defaults to 2 posts if no limit is given.
<pre>./gator browse LIMIT</pre>
example: 
<pre>./gator browse 10</pre>

### 🧪 Sample RSS Feeds
Here are some feeds to test with:
TechCrunch
Hacker News
Boot.dev Blog

## 🗃 Tech Stack
- 🐹 Go
- 🐘 PostgreSQL
- 🔧 sqlc — SQL to Go type-safe codegen
- 🧱 goose — DB migrations


### 🧼 TODO / Improvements
- Store authors & categories
- Support filtering by feed
- Export posts to markdown or HTML
- Configurable concurrency or batching

 📄 License
MIT License. See LICENSE for details.
