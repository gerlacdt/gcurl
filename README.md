### gCurl

gCurl is a sample implementation of the original [curl](https://github.com/curl/curl).

#### Development

```bash
# build
make

# run a command
./gcurl -X GET https://reqres.in/api/user | jq

# GET is default
./gcurl https://reqres.in/api/user | jq


# POST request via pipe
cat files/post_request.json|  ./gcurl -X POST https://reqres.in/api/user

# POST request via redirect
./gcurl -X POST https://reqres.in/api/user < files/post_request.json
```
