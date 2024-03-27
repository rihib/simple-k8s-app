SELECT COUNT(*) as count
FROM bigquery-public-data.samples.shakespeare
GROUP BY word
ORDER BY count DESC
LIMIT 10;
