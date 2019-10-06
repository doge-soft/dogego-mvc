FROM alpine

WORKDIR /app

ADD ./main /app
ADD ./web/templates /app

EXPOSE 3000

ENTRYPOINT [ "./main" ]