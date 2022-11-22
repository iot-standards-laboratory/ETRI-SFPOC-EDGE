import 'package:flutter/material.dart';
import 'package:front/app/controller/centrifugo.dart';
import 'package:front/app/controller/mqttclient.dart';
import 'package:front/app/model/agent.dart';
import 'package:front/app/model/controller.dart';
import 'package:front/app/model/service.dart';
import 'package:get/get.dart';
import 'package:centrifuge/centrifuge.dart' as centrifuge;
import 'package:mqtt_client/mqtt_browser_client.dart';
import 'package:mqtt_client/mqtt_client.dart';

import 'http_loader.dart';

class HomeController extends GetxController {
  //TODO: Implement HomeController
  var menuIdx = 0.obs;
  final scaffoldKey = GlobalKey<ScaffoldState>();
  var services = <Service>[].obs;
  var agents = <Agent>[].obs;
  var ctrls = <Controller>[].obs;
  @override
  void onInit() {
    super.onInit();
  }

  // centrifuge.Client? client = null;
  // void initCentrifuge() async {
  //   client = await newCentrifugeClient();
  // }

  MqttBrowserClient? mqttClient;
  void mqttInit() async {
    mqttClient = newMqttClient(mqttAddr: "mqtt.godopu.com");

    final connMess = MqttConnectMessage()
        // .withClientIdentifier('etri/etrismartfarm')
        .withWillTopic(
            'public/statuschanged') // If you set this you must set a will message
        .withWillMessage('I am user')
        .startClean() // Non persistent session for testing
        .withWillQos(MqttQos.atLeastOnce);
    mqttClient!.connectionMessage = connMess;

    try {
      await mqttClient!.connect('etri', 'etrismartfarm');
    } on Exception catch (e) {
      print('client exception - $e');
      mqttClient!.disconnect();
      return;
    }

    Subscribe(
      client: mqttClient!,
      topic: "public/statuschanged",
      onUpdate: (topic, payload) {
        loadData();
      },
    );
  }

  void loadData() async {
    agents.value = await loadAgents();
    services.value = await loadServices();
    ctrls.value = await loadCtrls();
  }

  @override
  void onReady() {
    super.onReady();
    loadData();
    mqttInit();
  }

  @override
  void onClose() {
    super.onClose();
    mqttClient!.disconnect();
  }
}
