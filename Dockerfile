# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

FROM golang:1.19-alpine

RUN addgroup -S alts && adduser -S alts -G alts

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /ALTS

##
## STEP 2 - DEPLOY
##
#FROM scratch

#WORKDIR /

#COPY --from=build /ALTS /ALTS
#COPY --from=build /etc/passwd /etc/passwd

USER alts

EXPOSE 8080
ENTRYPOINT [ "/ALTS" ]