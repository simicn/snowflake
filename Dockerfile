FROM alpine:latest
MAINTAINER simi <simiwei@gmail.com>
COPY snowflake /usr/bin/
ENTRYPOINT ["/usr/bin/snowflake"]
EXPOSE 10000 6060
