ARG JOBBER_STRING
FROM golang:latest as builder

# Copying and building Ukaska executable

COPY . /app
WORKDIR /app
RUN go build main.go

# Downloading and building Jobber
ENV JOBBER_VER v1.4.4
WORKDIR /go_wkspc/src/github.com/dshearer
RUN apt update && \
    apt install -y make rsync grep ca-certificates openssl wget gcc
RUN wget "https://api.github.com/repos/dshearer/jobber/tarball/${JOBBER_VER}" -O jobber.tar.gz && \
    tar xzf *.tar.gz && rm *.tar.gz && mv dshearer-* jobber && \
    cd jobber && \
    make check && \
    make install DESTDIR=/jobber-dist/

FROM mongo:latest
RUN mkdir /jobber
COPY --from=builder /jobber-dist/usr/local/libexec/jobbermaster /jobber/jobbermaster
COPY --from=builder //jobber-dist/usr/local/libexec/jobberrunner /jobber/jobberrunner
COPY --from=builder /jobber-dist/usr/local/bin/jobber /jobber/jobber
COPY --from=builder /app/main /app/main
COPY run_proccesses.sh run_proccesses.sh
RUN chmod +x run_proccesses.sh
ENV PATH /jobber:${PATH}
RUN apt-get update
RUN mkdir -p /var/jobber/0/
RUN echo "var-dir: /var\nlibexec-dir: /jobber" > /etc/jobber.conf
RUN touch /jobber/daemon-log.log
CMD ["./run_proccesses.sh"]
