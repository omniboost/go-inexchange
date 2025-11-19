module github.com/omniboost/go-inexchange

go 1.24.0

toolchain go1.24.5

require (
	github.com/gorilla/schema v1.1.0
	github.com/joefitzgerald/passwordcredentials v0.3.1
	github.com/joho/godotenv v1.3.0
	golang.org/x/oauth2 v0.30.0
	gopkg.in/guregu/null.v3 v3.5.0
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
