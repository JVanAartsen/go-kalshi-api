GIT_USER_ID := "jvanaartsen"
GIT_REPO_ID := "go-kalshi-api"

source:
	docker run --rm -u $$(id -u):$$(id -g) -v ${PWD}:/local openapitools/openapi-generator-cli generate \
		--generator-name go \
		--http-user-agent "go-kalshi-api" \
		--git-user-id $(GIT_USER_ID) --git-repo-id $(GIT_REPO_ID) \
		--additional-properties=packageName=kalshiAPI,generateInterfaces=true \
		--input-spec '/local/openapi.json' \
		--output '/local/'
	go mod tidy


clean-source:
	rm -f $$(cat .openapi-generator/FILES)