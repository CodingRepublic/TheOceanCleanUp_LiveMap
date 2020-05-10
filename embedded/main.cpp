#include <iostream>
#include <string>

#include <httplib.h>
#include <json.hpp>

class Data {
public:
  nlohmann::json objJson;

  std::string id;
  double latitude = 0;
  double longitude = 0;

public:
  std::string toString() { return this->objJson.dump(); };

  void updateJson() {
    this->objJson["id"] = this->id;

    this->objJson["position"]["lat"] = this->latitude;
    this->objJson["position"]["long"] = this->longitude;
  };

public:
  Data() = default;
};

int send_update(httplib::Client &client, Data &data) {
  data.updateJson();
  auto res = client.Post("/api/update_fleet_unit", data.toString(),
                         "application/json");
  std::cout << data.toString() << std::endl;
  if (res) {
    std::cout << res->status << std::endl;
    std::cout << res->body << std::endl;
  } else {
    std::cout << "problem occurred" << std::endl;
  }
  return 0;
}

Data &getUpdate() {
  auto data = new (Data);

  data->id = "alpha_01";
  data->latitude = 15.00;
  data->longitude = 12.05;
  return *data;
}

int main() {
  httplib::Client client("localhost", 8001);

  send_update(client, getUpdate());
  return 0;
}
