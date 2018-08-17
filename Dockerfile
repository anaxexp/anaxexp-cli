FROM anaxexp/alpine:3.7-2.0.1

COPY ./bin/linux-amd64/anaxexp /usr/local/bin/anaxexp

RUN apk add --update bash docker git

CMD [ "anaxexp" ]
