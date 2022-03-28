import 'dart:convert';

import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:get/get.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class DeviceAController extends GetxController {
  final Device d;
  late final WebSocketChannel _channel;
  late final message = "".obs;

  DeviceAController(this.d, {String pmessage = ""}) {
    message.value = pmessage;
    _channel = WebSocketChannel.connect(
      Uri.parse('ws://$serverAddr/device/${d.id}'),
    );
    _channel.stream.listen((payload) async {
      var param = jsonDecode(payload);
      message.value = param["msg"];
    });
  }

  @override
  void onClose() {
    super.onClose();
    _channel.sink.close();
  }
}
