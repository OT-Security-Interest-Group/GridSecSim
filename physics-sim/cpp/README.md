
---

# cpp/README.md — Physics Engine (Stub) 

```markdown
# C++ Physics Engine (Stub)

Authoritative tick loop that publishes `SimStatus` and per-node `NodeDelta` messages to MQTT. This stub emits deterministic voltages near 1.0 pu; swap the stub for real DC/AC solvers later.

## What it does now

- Connects to MQTT with a Last Will (marks sim OFFLINE on crash).
- Publishes retained `grid/v1/status/sim` = `SimStatus{state:"RUNNING"}` once.
- Every `DT_MS`:
  - For each `NODE_ID`, publish `grid/v1/node/{id}` = `NodeDelta{v_pu, tick}`.

## Layout
```
cpp/
├─ include/gridbus/
│ ├─ topics.hpp # topic helpers
│ ├─ clock.hpp # steady, fixed tick
│ └─ mqtt.hpp # tiny paho wrapper
├─ src/
│ ├─ main.cpp # boot, retained status, tick→publish deltas
│ ├─ engine_stub.cpp# synth_voltage(tick,nodeId)
│ ├─ mqtt.cpp # (empty; keeps compilation units tidy)
│ └─ generated/ # protoc outputs (grid.pb.h/cc)
└─ CMakeLists.txt
```


## Dependencies

- Protobuf (libprotobuf, protoc)
- Paho MQTT C & C++ (`libpaho-mqtt3as`, `libpaho-mqttpp3`)
- CMake ≥ 3.16

Debian/Ubuntu:

```bash
apt-get update && apt-get install -y \
  g++ cmake make \
  libprotobuf-dev protobuf-compiler \
  libpaho-mqtt-dev libpaho-mqttpp3-dev
```