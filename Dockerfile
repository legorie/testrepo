FROM golang
COPY . /opt/app
WORKDIR /opt/app
RUN go mod tidy
RUN go build -o first first.go
CMD /opt/app/first.exe
