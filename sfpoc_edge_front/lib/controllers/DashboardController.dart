import 'dart:convert';

import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/controllers/ws.dart';
import 'package:sfpoc_edge_front/models/Controller.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:sfpoc_edge_front/models/service.dart';
import 'package:get/get.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:http/http.dart' as http;

class DashboardController extends GetxController {
  var devices = <Device>[];
  var discoveredDevices = <Device>[];
  var services = <Service>[];
  var controllers = <Controller>[];

  DashboardController() {
    ws.stream.listen((message) async {
      loadData();
    });
    loadData();
  }

  void loadData() async {
    var response = await http.get(Uri.http(serverAddr, getCtrlsList));

    List<dynamic> body = jsonDecode(response.body);
    // print('${body}');
    controllers = List.generate(
      body.length,
      (index) => Controller.fromJson((body[index])),
    ).toList();

    response = await http.get(Uri.http(serverAddr, getDiscoveredList));
    body = jsonDecode(response.body);
    discoveredDevices = List.generate(
      body.length,
      (index) => Device.fromJson((body[index])),
    ).toList();

    response = await http.get(Uri.http(serverAddr, getDevsList));
    body = jsonDecode(response.body);
    devices = List.generate(
      body.length,
      (index) => Device.fromJson((body[index])),
    ).toList();

    response = await http.get(Uri.http(serverAddr, getSvcsList));
    body = jsonDecode(response.body);
    services = List.generate(
      body.length,
      (index) => Service.fromJson((body[index])),
    ).toList();

    for (var c in controllers) {
      c.numberOfDevices = 0;
      for (var d in devices) {
        if (c.id == d.cid) c.numberOfDevices++;
      }
    }

    update();
  }

  @override
  void onClose() {
    super.onClose();
    ws.sink.close();
  }
}
