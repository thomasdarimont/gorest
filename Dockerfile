FROM scratch
LABEL maintainer ="thomas.darimont@gmail.com"
USER 20001
ADD out/app /
EXPOSE 8080
CMD ["/app"]
