# 📰 Blog Aggregator
** in progress **

A command-line application built in Go that allows users to register, log in, and aggregate blog content from various sources. Designed as a backend-focused project to demonstrate database interaction, CLI command routing, and type-safe SQL with [`sqlc`](https://github.com/sqlc-dev/sqlc).

---

## 📌 Features

- 🧑‍💻 User registration and login system
- 💾 PostgreSQL database integration
- 📥 Add and follow RSS feeds
- 🐦 CLI-based command interface
- 🗂 Store posts in a local database
- 🔁 Goose migrations for schema management

---

## 🚀 Getting Started

#### 1. Clone the repo

```bash
<pre>
git clone https://github.com/kavancamp/blogAggregator
cd blogAggregator
</pre

#### 2. Run Migrations
<pre>
goose -dir sql/migrations postgres "<YOUR_DB_URL>" up
</pre>

#### 3. Generate SQL code
<pre>
sqlc generate
</pre>

#### 4. Build the binary
<pre>
go build -o gator .
</pre>

## 🛠 Commands
Register a new user
<pre>./gator register <username></pre>
Add a new Feed - auto follows feed
<pre>./gator addfeed "<feed name>" "<feed url>"</pre>
Follow an existing feed
<pre>./gator follow "<feed url>"</pre>
Start aggregating (fetching posts)
<pre>./gator agg 30s</pre>
Browse recent posts - Shows recent posts from feeds you follow. Defaults to 2 posts if no limit is given.
<pre>./gator browse [limit]</pre>
example: 
<pre>./gator browse 10</pre>

### 🧪 Sample RSS Feeds
Here are some feeds to test with:
TechCrunch
Hacker News
Boot.dev Blog

### 🗃 Tech Stack
🐹 Go
🐘 PostgreSQL
🔧 sqlc — SQL to Go type-safe codegen

🧱 goose — DB migrations


🧼 TODO / Improvements
 Store authors & categories

 Support filtering by feed

 Export posts to markdown or HTML

 Configurable concurrency or batching

 📄 License
MIT License. See LICENSE for details.