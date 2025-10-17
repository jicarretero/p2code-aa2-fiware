FROM golang:1.24-alpine3.21 AS build

# update certificates to trust github
RUN mkdir /tmp/p2code-aa2-adaptor
COPY . /tmp/p2code-aa2-adaptor

WORKDIR /tmp/p2code-aa2-adaptor
RUN go build -o p2code-aa2-adaptor .

FROM alpine:3.21

ENV KEY_TYPE_TO_GENERATE="EC"


ENV COUNTRY="DE"
ENV STATE="Saxony"
ENV LOCALITY="Dresden"
ENV ORGANIZATION="M&P Operations Inc."
ENV COMMON_NAME="www.mp-operations.org"
ENV STORE_PASS="myPassword"
ENV KEY_ALIAS="myAlias"
ENV KEY_TYPE="P-256"
ENV OUTPUT_FORMAT="json"
ENV DID_TYPE="key"
ENV OUTPUT_FILE="/cert/did.json"

COPY --from=build /tmp/p2code-aa2-adaptor/p2code-aa2-adaptor /usr/bin/p2code-aa2-adaptor
COPY ./help-script/helper.sh /temp/helper.sh
RUN  apk add wget curl bind-tools net-tools gcompat openssl wget && \
  mkdir /config && \
  mkdir /cert && \
  chmod 770 /cert && \
  wget -P /usr/bin https://github.com/wistefan/did-helper/releases/download/0.1.1/did-helper && \
  chmod +x /usr/bin/did-helper && \
  chmod +x /temp/helper.sh
COPY --from=build /tmp/p2code-aa2-adaptor/config/config.toml /config
ENTRYPOINT [ "/usr/bin/p2code-aa2-adaptor" ]
