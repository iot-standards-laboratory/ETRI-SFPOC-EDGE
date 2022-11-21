import 'dart:async';
import 'package:mqtt_client/mqtt_browser_client.dart';
import 'package:mqtt_client/mqtt_client.dart';

MqttBrowserClient newMqttClient({required String mqttAddr}) {
  var client =
      MqttBrowserClient('ws://$mqttAddr', '', maxConnectionAttempts: 5);
  client.logging(on: false);
  client.setProtocolV311();
  client.connectTimeoutPeriod = 10000;
  client.port = 10001;
  client.onDisconnected = onDisconnected;
  client.onConnected = onConnected;
  client.onSubscribed = onSubscribed;
  client.pongCallback = pong;
  client.websocketProtocols = MqttClientConstants.protocolsSingleDefault;
  client.autoReconnect = true;
  client.resubscribeOnAutoReconnect = true;
  return client;
}

void Subscribe(
    {required MqttBrowserClient client,
    required String topic,
    Function(String topic, String payload)? onUpdate}) {
  if (client.connectionStatus!.state != MqttConnectionState.connected) {
    print("client is invalid error");
    return;
  }

  print('Subscribing to the $topic');
  client.subscribe(topic, MqttQos.atMostOnce);

  /// The client has a change notifier object(see the Observable class) which we then listen to to get
  /// notifications of published updates to each subscribed topic.
  client.updates!.listen((List<MqttReceivedMessage<MqttMessage?>>? c) {
    final recMess = c![0].payload as MqttPublishMessage;
    final pt =
        MqttPublishPayload.bytesToStringAsString(recMess.payload.message);
    print(
        'EXAMPLE::Change notification:: topic is <${c[0].topic}>, payload is <-- $pt -->');
    print('');

    if (onUpdate != null) {
      onUpdate(c[0].topic, pt);
    }
  });
}

/// The subscribed callback
void onSubscribed(String topic) {
  print('EXAMPLE::Subscription confirmed for topic $topic');
}

/// The unsolicited disconnect callback
void onDisconnected() {
  print('EXAMPLE::OnDisconnected client callback - Client disconnection');

  // if (client!.connectionStatus!.disconnectionOrigin ==
  //     MqttDisconnectionOrigin.solicited) {
  //   print('EXAMPLE::OnDisconnected callback is solicited, this is correct');
  // }
}

/// The successful connect callback
void onConnected() {
  print(
      'EXAMPLE::OnConnected client callback - Client connection was sucessful');
}

/// Pong callback
void pong() {
  print('EXAMPLE::Ping response client callback invoked');
}
