#pragma once
#include <string>

namespace gridbus {

struct Topics {
  std::string prefix;  // e.g., "grid/v1"

  std::string statusSim() const { return prefix + "/status/sim"; }
  std::string topology()  const { return prefix + "/topology"; }

  std::string node(const std::string& nodeId) const { return prefix + "/node/" + nodeId; }
  std::string nodeAll() const { return prefix + "/node/+"; }
};

} // namespace gridbus
