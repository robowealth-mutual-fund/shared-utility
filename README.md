## Shared Utility
The modules is contains the golang utilities for internal services. 

[[_TOC_]]

## Utilities
* [x] HTTP Client
* [x] Errors
* [x] Validator
    - Basic Validator Rule
    - GRPC Server Interceptor
    - GRPC Client Interceptor (Coming Soon)
* [ ] Logger

## Installation
 1. Use the below Go command to install Shared Utility
 ```bash
 go get github.com/robowealth-mutual-fund/shared-utility
 ```
 2. Import it in your code
 ```go
 import "github.com/robowealth-mutual-fund/shared-utility/httpclient"
 import "github.com/robowealth-mutual-fund/shared-utility/errors"
 import "github.com/robowealth-mutual-fund/shared-utility/grpcerrors"
 import "github.com/robowealth-mutual-fund/shared-utility/validator"
 ```

## Set up your project to support private modules
Mostly setup command in this instructions base on `git.robodev.co (GitLab)`

### Go
Go version >= 1.13 (RECOMMENDED)
```bash
go version # To know your go version

go env -w GO111MODULE=on
go env -w GOPROXY="https://goproxy.io,direct"

# Set go environment to find private module for specified url
go env -w GOPRIVATE="https://git.robodev.co/*"

# Disable checksum the private modules for specified url
go env -w GONOSUMDB="git.robodev.co/*"
```

### Git
Under the `go get` command is using Git to pull the specified versions of your dependencies. So, the git configuration for wherever Go is running ***has to have*** the appropriate credentials to access any private repositories (In this case is GitLab).

>>>
How to get personal access token on GitLab [Here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html)
>>>

```bash
git config \
--global \
url."https://oauth2:${personal_access_token}@git.robodev.co".insteadOf \
"https://git.robodev.co"

#or
git config \
--global \
url."https://${user}:${personal_access_token}@git.robodev.co".insteadOf \
"https://git.robodev.co"
```

>>>
This is great for local development, but what about my CI/CD pipeline?
>>>

### Dockerfile
Here is an example of a Dockerfile that allows for the injection of credentials at build time.
```dockerfile
# ------------------------------------------
#  Base on ROA Microservice's dockerfile
# ------------------------------------------
FROM golang:1.14.1-alpine as dependencies

ENV GO11MODULE=on

# Setup ENV (FOR PRIVATE MODULE)
ENV GOPROXY="https://goproxy.io,direct"
ENV GOPRIVATE="https://git.robodev.co/imp/*"
ENV GONOSUMDB="https://git.robodev.co/*"
ENV GITLAB_ACCESS_TOKEN="apgUaNUFr-EeYFqoruvf"

# Install git.
# Git is required for fetching the dependencies.
RUN apk  update  &&  apk  add  --no-cache  git  make  gcc  libc-dev

# Setup Git URL mapping (FOR PRIVATE MODULE)
RUN git config \
--global \
url."https://oauth2:${GITLAB_ACCESS_TOKEN}@git.robodev.co".insteadOf \
"https://git.robodev.co"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN  go  mod  download

# Copy the source from the current directory to the working Directory inside the container
COPY  .  .

# Build the Go app
RUN  make  build
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server cmd/**/*.go

# Start a new stage from scratch
# FROM scratch
FROM  alpine

RUN  GRPC_HEALTH_PROBE_VERSION=v0.3.1  && \
wget  -qO/bin/grpc_health_probe  https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64  && \
chmod  +x  /bin/grpc_health_probe

# # Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY  --from=dependencies  /app/bin/server  /app/bin/server
COPY  --from=dependencies  /app/entrypoint.sh  /

RUN  chmod  +x  /entrypoint.sh

ENTRYPOINT  ["/entrypoint.sh"]
```