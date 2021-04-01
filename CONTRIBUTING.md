# Contributing

This is a work in progress contributors document, you are encouraged to help us make this better!

## Local Development

**Frontend**: Vue.js app located in the `frontend/` folder  
**Backend**: Golang app located in the `source/` folder

The application runs as a single binary (`./statping`). The frontend code is embedded into the golang binary using [rice](https://github.com/GeertJohan/go.rice).

### Worked Example: Build a frontend change

#### Install prerequisites
- yarn > 1.20
- go > 1.13
- ```
  go get github.com/GeertJohan/go.rice
  go get github.com/GeertJohan/go.rice/rice
  ```
  Ensure your gopath (e.g. `~/go/bin`) is on your PATH so that you can run `rice`  

#### Make and build the change

1. Make changes to frontend (FE)
2. Build FE, resolve dependencies and copy to source/dist folders: `make frontend-build`
3. Embed the frontend into the backend's code: `make compile`
4. Build the backend: `make build`
5. Run the `statping` binary now in your project root with `./statping`
6. You now have a statping running locally on http://localhost:8080 - connect, configure the DB and test your changes. 