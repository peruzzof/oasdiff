module github.com/tufin/oasdiff/diff

go 1.15

require (
	github.com/apex/log v1.9.0
	github.com/getkin/kin-openapi v0.39.0
	github.com/stretchr/testify v1.6.1
	github.com/tufin/oasdiff/load v0.0.0-00010101000000-000000000000
)

replace github.com/tufin/oasdiff/load => ../load
replace github.com/tufin/oasdiff/diff => ./diff