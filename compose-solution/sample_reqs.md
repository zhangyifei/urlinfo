## Url Lookup

### Get request

URL components 

scheme:[//[userinfo@]host[:port]]path[?query][#fragment]

curl http://localhost:9090/urlinfo/1/linuxize/q?test=1

curl http://localhost:9090/urlinfo/1/www.linuxize.com/tt/q?test=1

curl http://localhost:9090/urlinfo/1/test:test1@linuxize.com/tt/q?test=1

### response

- Allow access

    ```json
    {"message":"the url is valid","allow":true}
    ```
- Not Allow

    ```json
    {"message":"linuxize--q?test=1is invalid","allow":false}
    ```

## Url Update

### Single Update

  ```shell
    curl -X POST -H "Content-Type: application/json" \
    -d '{"hostnameport": "linuxize", "queryparamter": "q?test=1"}' \
    http://localhost:9090/urlinfo/1/update
  ```

### Batch Update

  ```shell
    curl -X POST -H "Content-Type: application/json" \
    -d '{"requests":[{"hostnameport": "linuxize", "queryparamter": "q?test=4"},{"hostnameport": "linuxize", "queryparamter": "q?test=3"}]}' \
    http://localhost:9090/urlinfo/1/batchupdate
  ```



