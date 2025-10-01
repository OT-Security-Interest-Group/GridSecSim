# syntax=docker/dockerfile:1
FROM debian:12 AS build
RUN apt-get update && apt-get install -y --no-install-recommends \
    g++ cmake make git ca-certificates \
    libprotobuf-dev protobuf-compiler \
    libpaho-mqtt-dev libpaho-mqttpp3-dev \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY ./src ./src
COPY ./include ./include
COPY ./CMakeLists.txt ./

# If you generate protos at build time, copy proto and run protoc here instead.
# (We assume you generated cpp/src/generated already.)

RUN cmake -S . -B build -DCMAKE_BUILD_TYPE=Release \
 && cmake --build build -j
# runtime
FROM gcr.io/distroless/cc-debian12
WORKDIR /app
COPY --from=build /app/build/physics_cpp /app/physics_cpp
ENV BROKER_URL=tcp://mqtt:1883
ENV TOPIC_PREFIX=grid/v1
ENV DT_MS=100
ENV NODE_IDS=bus-1,bus-2,bus-3
ENTRYPOINT ["/app/physics_cpp"]