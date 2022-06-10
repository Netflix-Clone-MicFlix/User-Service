docker compose up ../docker-compose.yml -d
sleep 15
go clean -testcache && go test -v .
sleep 15
docker compose down ../docker-compose.yml 