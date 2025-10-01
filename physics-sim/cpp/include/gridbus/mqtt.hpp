#pragma once
#include <mqtt/async_client.h>
#include <string>
#include <vector>

namespace gridbus {

class Mqtt {
public:
  Mqtt(const std::string& url, const std::string& clientId)
  : cli_(url, clientId) {}

  void connect(const std::string& willTopic = "", const std::string& willPayload = "") {
    mqtt::connect_options_builder b;
    b.automatic_reconnect(true);
    b.clean_session(true);
    b.keep_alive_interval(std::chrono::seconds(20));
    if (!willTopic.empty()) {
      auto will = mqtt::message(willTopic, willPayload, 1, true);
      b.will_message(will);
    }
    auto tok = cli_.connect(b.finalize());
    tok->wait();
  }

  void publish(const std::string& topic, const void* data, size_t n, int qos = 1, bool retain = false) {
    auto msg = mqtt::message(topic, data, n, qos, retain);
    cli_.publish(msg)->wait();
  }

  mqtt::async_client& client() { return cli_; }

  void disconnect() {
    try { cli_.disconnect()->wait(); } catch(...) {}
  }

private:
  mqtt::async_client cli_;
};

} // namespace gridbus
