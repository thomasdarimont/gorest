FROM phusion/baseimage:0.10.0
RUN useradd -u 10001 app

FROM scratch
LABEL maintainer ="thomas.darimont@gmail.com"
COPY --from=0 /etc/passwd /etc/passwd
USER app
ADD out/app /
EXPOSE 8080
CMD ["/app"]
