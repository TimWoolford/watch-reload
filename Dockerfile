FROM gcr.io/distroless/base

COPY out/watch-reload /watch-reload

ENTRYPOINT ["/watch-reload"]