FROM golang:1.19
WORKDIR /tripplanner
COPY . ./
RUN go build -o app
EXPOSE 8080
CMD [ "./app" ]
