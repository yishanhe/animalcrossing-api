.PHONY: swagger-validate
swagger-validate:
	swagger validate openapi.yaml



.PHONY: codegen-model
codegen-model: swagger-validate
	-@rm -r models
	swagger generate model --spec=openapi.yaml



.PHONY: codegen-server
codegen-server: codegen-model
	mkdir -p tmp
	cp restapi/configure_acnh.go tmp/configure_acnh.$(shell date +%FT%T).go
	rm -r restapi
	swagger generate server -f openapi.yaml
	@echo
	@echo "------------        CAVEAT      ------------"
	@echo "Please handle the configure_acnh.go manully."
	@echo "--------------------------------------------"
	@echo



.PHONY: clean
clean:
	@rm -r models


.PHONY: install
install:
	@go get -u -f ./...


.PHONY: run-server
run-server:
	go run cmd/acnh-server/main.go --host 0.0.0.0 --port 8080