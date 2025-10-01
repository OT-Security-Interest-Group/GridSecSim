cmake_minimum_required(VERSION 3.16)
project(physics_cpp CXX)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(Protobuf REQUIRED)
find_package(PahoMqttCpp REQUIRED)
# On some distros you also need PahoMqttC; the C++ target usually links it.

include_directories(
  ${Protobuf_INCLUDE_DIRS}
  ${PAHO_MQTT_CPP_INCLUDE_DIRS}
  ${CMAKE_CURRENT_SOURCE_DIR}/include
  ${CMAKE_CURRENT_SOURCE_DIR}/src/generated
)

add_executable(physics_cpp
  src/main.cpp
  src/engine_stub.cpp
  src/mqtt.cpp
  src/generated/grid/v1/grid.pb.cc
)

target_link_libraries(physics_cpp
  ${Protobuf_LIBRARIES}
  PahoMqttCpp::paho-mqttpp3
  PahoMqttC::paho-mqtt3as
)