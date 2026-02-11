package exchanges

const (
	createExchange = `
		INSERT INTO exchanges (from_ad_id, to_ad_id, status, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING id, from_ad_id, to_ad_id, status, comment, created_at, updated_at
	`

	testSelect = `
		SELECT * FROM test
	`

	testInsert = `
		INSERT INTO test(comment)
		VALUES ($1)
	`
)
