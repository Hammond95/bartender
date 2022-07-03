FROM scratch

ENV PORT 8888
ENV GOOS linux
ENV GOARCH amd64

EXPOSE $PORT

COPY bin/bartender-*-${GOOS}-${GOARCH} /bartender
ADD static /static

ENTRYPOINT ["/bartender"]
CMD "--address=:${PORT}"