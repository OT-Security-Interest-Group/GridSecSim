#pragma once
#include <chrono>
#include <cstdint>

namespace gridbus {

struct TickClock {
  using Steady = std::chrono::steady_clock;

  uint64_t tick = 0;
  uint64_t dt_ms = 100;
  Steady::time_point start = Steady::now();

  uint64_t now_tick() const {
    auto elapsed = std::chrono::duration_cast<std::chrono::milliseconds>(Steady::now() - start).count();
    return static_cast<uint64_t>(elapsed / dt_ms);
  }

  void sleep_until_next() {
    ++tick;
    auto next = start + std::chrono::milliseconds(tick * dt_ms);
    std::this_thread::sleep_until(next);
  }
};

} // namespace gridbus
