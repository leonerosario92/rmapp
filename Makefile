.PHONY: openapi
openapi:
	oapi-codegen -generate types -o internal-api/rmapp/ports/openapi_types.gen.go -package ports api/openapi/rmapp.yaml
	oapi-codegen -generate chi-server -o internal-api/rmapp/ports/openapi_api.gen.go -package ports api/openapi/rmapp.yaml
