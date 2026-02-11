package exchanges

const (
	createExchange = `
		INSERT INTO exchanges (from_ad_id, to_ad_id, status, comment)
		VALUES ($1, $2, $3, $4)
		RETURNING id, from_ad_id, to_ad_id, status, comment, time
	`
)
