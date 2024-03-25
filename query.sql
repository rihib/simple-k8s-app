WITH complete_data AS (SELECT
  REGEXP_REPLACE(REGEXP_REPLACE(REGEXP_REPLACE(httpRequest.requestUrl, r'\?.*$', ''), r'\/\d{1,}\/', '/NUMBER/'),r'\/\d{1,}', '/NUMBER') AS reqUrl,
  httpRequest.requestMethod AS reqMethod,
  COUNT(*) AS cnt
FROM
  `tron-151603.cloud_logging.requests`
WHERE
    timestamp>="2023-12-31 17:00:00"
    AND timestamp<="2024-01-01 01:00:00"
  AND (
    REGEXP_CONTAINS(httpRequest.requestUrl, '/subscriptions')
    OR REGEXP_CONTAINS(httpRequest.requestUrl, '/plans')
    OR REGEXP_CONTAINS(httpRequest.requestUrl, 'api/icons/entry')
    OR (REGEXP_CONTAINS(httpRequest.requestUrl, 'medias/text') AND httpRequest.requestMethod = "POST")
    OR REGEXP_CONTAINS(httpRequest.requestUrl, '/pointProducts')
    OR REGEXP_CONTAINS(httpRequest.requestUrl, '/bulkLike')
  )
GROUP BY
  reqMethod,
  reqUrl
ORDER BY
  cnt desc
)
SELECT
  reqUrl,
  reqMethod,
  SUM(cnt) total_num_requests,
  SUM(cnt) / TIMESTAMP_DIFF("2024-01-01 01:00:00", "2023-12-31 17:00:00", SECOND) avg_rps
FROM complete_data
GROUP BY reqUrl, reqMethod
ORDER BY avg_rps DESC
