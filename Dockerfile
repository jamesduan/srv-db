FROM alpine:3.2
ADD ./ /srv-db
ENTRYPOINT [ "/srv-db" ]
