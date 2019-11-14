module github.com/oraora-io/oraora-be

require github.com/oraora-io/oraora-be/seasons v0.0.0-20191114182100-5486f39b7694

replace github.com/oraora-io/oraora-be/seasons v0.0.0 => ./seasons

go 1.12

require (
	github.com/darenliang/jikan-go v0.0.0-20191111221916-d0d13e1ff382
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
)
