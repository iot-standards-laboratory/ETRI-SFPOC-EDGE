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

class HomeController extends GetxController {
  //TODO: Implement HomeController
  var menuIdx = 0.obs;
  final scaffoldKey = GlobalKey<ScaffoldState>();
  var services = <Service>[
    Service(id: '1-1-1', name: 'devicemanagera', numOfCtrls: 7),
    Service(id: '2-2-2', name: 'devicemanagerb', numOfCtrls: 7),
  ].obs;
  var agents = <Agent>[
    Agent(
        id: '22222-222222-22222-22222222222',
        name: 'devicemanagera',
        status: "connected"),
  ].obs;
  var ctrls = <Controller>[
    Controller(
        id: '1-1-1',
        name: 'etri-Zdtwe2^==',
        agentId: '22222-222222-22222-22222222222'),
  ].obs;
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
    mqttClient = newMqttClient(mqttAddr: "localhost");

    final connMess = MqttConnectMessage()
        .withClientIdentifier('etri/etrismartfarm')
        .withWillTopic(
            'public/statuschanged') // If you set this you must set a will message
        .withWillMessage('I am user')
        .startClean() // Non persistent session for testing
        .withWillQos(MqttQos.atLeastOnce);
    print('Mosquitto client connecting....');
    mqttClient!.connectionMessage = connMess;

    try {
      await mqttClient!.connect();
    } on Exception catch (e) {
      print('client exception - $e');
      mqttClient!.disconnect();
      return;
    }

    Subscribe(
      client: mqttClient!,
      topic: "public/statuschanged",
      onUpdate: (topic, payload) {
        print("topic is $topic");
        print("payload is $payload");
      },
    );
  }

  @override
  void onReady() {
    super.onReady();
    mqttInit();
  }

  @override
  void onClose() {
    super.onClose();
  }
}
