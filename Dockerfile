FROM ubuntu
WORKDIR test
COPY ./main ./main
CMD ["./main"]