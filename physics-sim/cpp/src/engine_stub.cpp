#include <cmath>
#include <string>
#include <vector>
#include <functional>
#include "generated/grid/v1/grid.pb.h"

namespace engine {

// Deterministic “fake” voltage around 1.0 pu based on tick + nodeId hash.
static inline uint64_t hash_id(const std::string& s) {
  uint64_t h = 1469598103934665603ull;
  for (auto c : s) { h ^= static_cast<unsigned char>(c); h *= 1099511628211ull; }
  return h;
}

double synth_voltage(uint64_t tick, const std::string& nodeId) {
  // 1.0 ± 0.02 with a smooth variation; deterministic
  const double amp = 0.02;
  double phase = double((hash_id(nodeId) ^ (tick * 911382323ull)) & 0xffff) / 65535.0;
  return 1.0 + amp * std::sin(2.0 * M_PI * phase);
}

} // namespace engine