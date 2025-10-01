#include <iostream>
#include <cstdlib>
#include <thread>
#include <vector>
#include <string>
#include <sstream>

#include "gridbus/topics.hpp"
#include "gridbus/clock.hpp"
#include "gridbus/mqtt.hpp"

#include "generated/grid/v1/grid.pb.h"

using grid::v1::Clock;
using grid::v1::Envelope;
using grid::v1::SimStatus;
using grid::v1::NodeDelta;

namespace {

std::vector<std::string> split_csv(const std::string& s) {
  std::vector<std::string> out;
  std::stringstream ss(s);
  std::string item;
  while (std::getline(ss, item, ',')) {
    if (!item.empty()) out.push_back(item);
  }
  return out;
}

} // anonymous

// forward decl
namespace engine { double synth_voltage(uint64_t tick, const std::string& nodeId); }

int main() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  const std::string broker = std::getenv("BROKER_URL") ? std::getenv("BROKER_URL") : "tcp://localhost:1883";
  const std::string prefix = std::getenv("TOPIC_PREFIX") ? std::getenv("TOPIC_PREFIX") : "grid/v1";
  const std::string nodesCsv = std::getenv("NODE_IDS") ? std::getenv("NODE_IDS") : "bus-1,bus-2,bus-3";
  const uint64_t dt_ms = std::getenv("DT_MS") ? std::stoull(std::getenv("DT_MS")) : 100;

  std::vector<std::string> nodeIds = split_csv(nodesCsv);
  gridbus::Topics topics{prefix};
  gridbus::TickClock clk; clk.dt_ms = dt_ms;

  // MQTT connect with a LAST WILL that marks sim offline
  gridbus::Mqtt bus(broker, "physics_cpp");
  bus.connect(topics.statusSim(), R"({"state":"OFFLINE"})");

  // Publish retained "RUNNING" status once connected
  {
    SimStatus st;
    auto env = st.mutable_env();
    env->mutable_clock()->set_tick(0);
    env->mutable_clock()->set_dt_ms(dt_ms);
    st.set_state("RUNNING");

    std::string bytes;
    st.SerializeToString(&bytes);
    bus.publish(topics.statusSim(), bytes.data(), bytes.size(), /*qos*/1, /*retain*/true);
  }

  // Main tick loop: publish NodeDelta for each node
  while (true) {
    uint64_t t = clk.tick; // authoritative tick inside this process

    for (const auto& id : nodeIds) {
      NodeDelta nd;
      auto env = nd.mutable_env();
      env->mutable_clock()->set_tick(t);
      env->mutable_clock()->set_dt_ms(dt_ms);

      nd.set_node_id(id);
      nd.set_v_pu(engine::synth_voltage(t, id));

      std::string bytes;
      nd.SerializeToString(&bytes);
      bus.publish(topics.node(id), bytes.data(), bytes.size(), /*qos*/1, /*retain*/false);
    }

    clk.sleep_until_next();
  }

  bus.disconnect();
  google::protobuf::ShutdownProtobufLibrary();
  return 0;
}
