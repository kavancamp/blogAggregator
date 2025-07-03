
-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    iff.*,
    f.name AS feed_name,
    u.name AS user_name
FROM inserted_feed_follow iff
JOIN feeds f ON iff.feed_id = f.id
JOIN users u ON iff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name AS feed_name, u.name AS user_name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE ff.user_id = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;
