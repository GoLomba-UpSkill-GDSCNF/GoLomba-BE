FROM golang:1.19.3 AS build

WORKDIR /dist

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o appstart

# Create Image for production environment
FROM busybox:stable

# defenie varible and env home directory
ARG USER=golang
ENV HOME /home/$USER/app

# add new user and set default user directory
RUN addgroup -S $USER && adduser -S $USER -G $USER
WORKDIR $HOME

COPY --from=build /dist/appstart $HOME/appstart
RUN chown $USER:$USER -R $HOME/appstart
RUN chmod +x $HOME/appstart

USER $USER
EXPOSE 3000
CMD ["./appstart"]