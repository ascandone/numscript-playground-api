# syntax=docker/dockerfile:1

FROM golang:1.22
WORKDIR /app

# download numscript release
RUN curl -L https://github.com/formancehq/numscript/releases/download/v0.0.3/numscript_.0.0.3_Linux_arm64.tar.gz > numscript.tar.gz
RUN tar -xvzf ./numscript.tar.gz 
ENV PATH="$PATH:/app"

# TODO this can be optimized
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build .

EXPOSE 3000

CMD ["./numscript_playground_api"]
