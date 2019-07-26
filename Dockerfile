FROM frolvlad/alpine-glibc

COPY log log

COPY environment environment

RUN apk add --no-cache bash

ADD main /

EXPOSE 8080

EXPOSE 5432

CMD ["/main"]
