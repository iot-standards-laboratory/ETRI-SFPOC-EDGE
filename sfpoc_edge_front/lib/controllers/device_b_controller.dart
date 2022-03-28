import 'dart:convert';

import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:get/get.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class DeviceBController extends GetxController {
  var _cdc = 0;
  var _humidity = 0;
  var _waterValue = 0;
  var _temperature = 0;

  int get cdc => _cdc;
  int get humidity => _humidity;
  int get waterValue => _waterValue;
  int get temperature => _temperature;

  RxInt fan;
  RxInt light;
  RxInt servo;

  final Device d;
  late final WebSocketChannel _channel;

  DeviceBController(this.d, {int fanV = 0, int lightV = 0, int servoV = 0})
      : fan = fanV.obs,
        light = lightV.obs,
        servo = servoV.obs {
    print("sid: ${d.sid}");

    _channel = WebSocketChannel.connect(
      Uri.parse('ws://$serverAddr/device/${d.id}'),
    );
    _channel.stream.listen((message) async {
      message = jsonDecode(message);
      _cdc = message["cdc_value"];
      _humidity = message["humidity"];
      _waterValue = message["water_value"];
      _temperature = message["temperature"];
      update();
      print(message);
    });

    // loadData();
  }

  // dynamic get controlCmd => () {
  //       return {
  //         "fan": this.fan,
  //         "light": this.light,
  //         "servo": this.servo,
  //       };
  //     };

  @override
  void onClose() {
    super.onClose();
    _channel.sink.close();
    print("deleted!!");
  }
}
