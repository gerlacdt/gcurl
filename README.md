# gCurl

gCurl is a sample implementation of the original [curl](https://github.com/curl/curl).

## Example calls

```bash
# get help
./gcurl --help

# GET request
./gcurl -X GET https://reqres.in/api/user | jq

# GET is default
./gcurl https://reqres.in/api/user | jq

# set custom http header
./gcurl -X GET -H "X-Api-Key: foobar" https://reqres.in/api/user | jq

# verbose mode outputs request and response headers to STDERR
./gcurl -v -X GET https://reqres.in/api/user | jq

# POST request via STDIN
cat files/post_request.json |  ./gcurl -X POST https://reqres.in/api/user

# POST request via redirect
./gcurl -X POST https://reqres.in/api/user < files/post_request.json

# POST via --data argument
./gcurl -v -X POST -d '{"foo": "bar"}' http://localhost:8080/post  | jq

# PUT works the same as POST
cat files/post_request.json|  ./gcurl -X PUT https://reqres.in/api/user

# DELETE is also supported
./gcurl -v -X DELETE https://reqres.in/api/user/2
```

## Development

```bash
# build linux binary
make

# run tests
make httpbin  # prerequisite: some tests depend on a running container
make test

# run tests with verbose output
make testv

# some tests depend on golden files, if they change the following command will update the golden files
make golden
```
