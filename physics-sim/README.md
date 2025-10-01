# physics_sim

Event-driven power-grid simulation fabric.

- **Transport:** MQTT (pub/sub)
- **Contract:** Protobuf payloads
- **Physics engine:** C++ service (authoritative tick, publishes Node/Edge deltas)
- **Operator/API gateway:** Go service (HTTP + WebSocket for UIs, mirrors MQTT)

The goal: **fast, selective, bidirectional** data flow. Components send events when things change; the physics engine advances time, computes new state, and publishes deltas only when values change.

---

## Architecture

         +-------------------+                +------------------+
         |   Go API / WS     |  HTTP/WS       |  Browser / HMI   |
         |   (api_go)        |<-------------->|  (clients)       |
         +---------^---------+                +------------------+
                   |
            subscribes/publishes
                   |
             MQTT Broker (mosquitto)
                   |


        +-------------------+-------------------+--------------------------+
        | | | |
        sub grid/v1/node/* sub comp//... pub grid/v1/node/ ... (other comps)
        pub grid/v1/status pub comp//meas sub comp//cmd
        | | |
        +--v-------------------+--+ +---v--------------------------+
        | C++ Physics Engine | | Components (sensors/actors) |
        | (physics_cpp) | | (future, examples under /components)|
        +-------------------------+ +------------------------------+


- **C++ physics** is the single writer of authoritative state (fixed tick).
- **Go API** mirrors MQTT into HTTP/WS for dashboards and exposes simple control endpoints.
- **Components** (future) talk MQTT directly (measurements, acks, commands).

---

## Topics (prefix = `grid/v1`)

- **Physics → all**
  - `status/sim` (retained): `SimStatus`
  - `node/{nodeId}`: `NodeDelta` (event-only; deltas)
  - `edge/{edgeId}`: `EdgeDelta` (optional, later)

- **(Future) Components → physics**
  - `measure/{compId}`: `Measurement`
  - `cmdack/{compId}`: `CommandAck`
  - `status/comp/{compId}` (retained + LWT): `CompStatus`

- **(Future) Physics → component**
  - `cmd/{compId}`: `Command`
  - `ack/{compId}`: `Ack` (hello response)
  - `intent/confirm/{compId}`: `IntentConfirm`

---

## Protobuf contract

Source of truth: `proto/grid/v1/grid.proto`.

Minimal messages used by the current MVP:
- `Clock`, `Envelope`
- `SimStatus` (state, tick)
- `NodeDelta` (node_id, optional v_pu)
- `EdgeDelta` (edge_id, optional breaker_closed)
- (Future) `Command`, `Measurement`, `CommandAck`, etc.

Generate bindings:

```bash
# Go
protoc -I proto \
  --go_out=go --go_opt=paths=source_relative \
  proto/grid/v1/grid.proto

# C++
protoc -I proto \
  --cpp_out=cpp/src/generated \
  proto/grid/v1/grid.proto
