module github.com/omniboost/go-lightspeed-r

go 1.22

require (
	github.com/gorilla/schema v1.3.0
	github.com/pkg/errors v0.9.1
	gopkg.in/guregu/null.v3 v3.5.0
)

require golang.org/x/oauth2 v0.21.0 // indirect

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
