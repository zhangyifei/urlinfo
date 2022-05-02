docker-compose up --build -d
docker-compose down

curl http://localhost:18888/urlinfo/1/linuxize/q?test=1

curl -X POST -H "Content-Type: application/json" \
    -d '{"hostnameport": "linuxize", "queryparamter": "q?test=1"}' \
    http://localhost:18888/urlinfo/1/update

curl -X POST -H "Content-Type: application/json" \
    -d '{"requests":[{"hostnameport": "linuxize", "queryparamter": "q?test=4"},{"hostnameport": "linuxize", "queryparamter": "q?test=3"}]}' \
    http://localhost:18888/urlinfo/1/batchupdate


curl http://localhost:9090/urlinfo/1/linuxize/q?test=1

curl -X POST -H "Content-Type: application/json" \
    -d '{"hostnameport": "linuxize", "queryparamter": "q?test=1"}' \
    http://localhost:9090/urlinfo/1/update