FROM ubuntu:16.04

COPY out/watch-reload /watch-reload

ENTRYPOINT ["/watch-reload"]