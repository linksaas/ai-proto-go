rm *.go
rm -rf tmp
npx @openapitools/openapi-generator-cli generate -c config.json -i proto/index_0.0.3.yml -g go -o tmp
mv tmp/*.go .
rm -rf tmp
