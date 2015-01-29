echo "Testing Code Coverage"

gocov test > coverage/go-tmdb.json
gocov-html coverage/go-tmdb.json > coverage/coverage_report.html