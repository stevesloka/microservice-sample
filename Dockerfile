FROM scratch
ADD microservice /
ENTRYPOINT ["/microservice"]