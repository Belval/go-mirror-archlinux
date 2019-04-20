FROM ubuntu:18.04

# Installing dependencies
RUN apt update -y && apt upgrade -y
RUN apt install -y git \
                   rsync

# Cloning repo
RUN git clone https://github.com/Belval/go-mirror-archlinux.git

# The port we will be serving the files at
EXPOSE 8081

# Starting mirror
CMD ["/go-mirror-archlinux/go-mirror-archlinux/go-mirror-archlinux", \
     "-config=/go-mirror-archlinux/go-mirror-archlinux/config.json"]




