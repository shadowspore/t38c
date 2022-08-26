TILE38_CONTAINER_NAME=t38c-test

e2e_start_tile38:
	docker run -d -p 9851:9851 --name ${TILE38_CONTAINER_NAME} tile38/tile38

e2e_run:
	export T38C_TEST_E2E=1; \
	export T38C_TEST_ADDR=localhost:9851; \
	go test --count=1 ./e2e/...

e2e_stop_tile38:
	docker rm -f ${TILE38_CONTAINER_NAME}