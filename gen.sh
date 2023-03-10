rm -rf client
npx @openapitools/openapi-generator-cli generate -c client_config.json -i proto/index_0.0.2.yml -g go -o client
rm -rf client/api client/docs client/test client/go.mod client/go.sum client/git_push.sh

rm -rf server
npx @openapitools/openapi-generator-cli generate -c server_config.json -i proto/index_0.0.2.yml -g go-gin-server -o server
mv server/go/* server/
rm -rf server/api  server/go server/main.go server/Dockerfile