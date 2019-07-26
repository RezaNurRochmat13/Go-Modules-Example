# Research How To Using Go Modules

This project is example how to use go modules in your projects.

## Built with
List technology used in this example project
- [Golang](https://golang.org/) - Golang Programming Language
- [Visual Studio Code](https://code.visualstudio.com/) - Text Editor
- [Echo](https://echo.labstack.com/) - Golang Library

## How to run
List steps to running this project on local machine
### Natively Run
- Clone the repo using SSH or HTTPS
- Import DB to your local machine
- Run your apps

### Using Docker
- Clone the repo using SSH or HTTPS
- Import DB to your local machine
- Run command shells in manifest using command `sh manifest/compile.sh`
- After finish, you can build to docker images using command `docker build -t $IMAGE-NAME:TAGS .`
- For running apps in docker, you can run command `docker run -it --rm -p $PORT:PORT $NAME-IMAGES`
- After success running,you can access from your agent like Insomnia REST or Postman

## Go Modules Cheatsheet
- cd direectory outside GOPATH # ex : ~/.blabla-workspace
- go mod init <module-name> - Init your module
- go get <package-downloads> - Command for download your dependency and add in your projects
- go list -m all — View final versions that will be used in a build for all direct and indirect dependencies
- go list -u -m all — View available minor and patch upgrades for all direct and indirect dependencies
- go get -u or go get -u=patch — Update all direct and indirect dependencies to latest minor or patch upgrades (pre-releases are ignored)
- go build ./... or go test ./... — Build or test all packages in the module when run from the module root directory
- go mod tidy — Prune any no-longer-needed dependencies from go.mod and add any dependencies needed for other combinations of OS, architecture, and build tags

More documentation about go modules, you can visit this [https://github.com/golang/go/wiki/Modules]

## Sources
List of source for references for building this project

## Contributors
List of member to contribute to this repo

