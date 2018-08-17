FROM anaxexp/alpine:3.8-2.1.0

COPY ./bin/linux-amd64/anaxexp /usr/local/bin/anaxexp

RUN apk add --update bash docker git

CMD [ "anaxexp" ]
