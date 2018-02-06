FROM scratch
MAINTAINER thomas.darimont@gmail.com

ADD out/app /

EXPOSE 8080

CMD ["/app"]
