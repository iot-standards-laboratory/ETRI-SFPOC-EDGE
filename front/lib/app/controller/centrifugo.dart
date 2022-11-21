import 'package:centrifuge/centrifuge.dart' as centrifuge;

const centrifugoTkn =
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImV0cmkiLCJuYW1lIjoiZXRyaSIsInJvbGUiOlsidXNlciJdLCJiYXNlIjoiLyIsInN1YiI6ImV0cmkiLCJjaGFubmVscyI6WyJwdWJsaWM6c3RhdHVzY2hhbmdlZCJdfQ.LybgxWw40SDeqY3NOREa2_zI2ZwrsSaSoN-S_4clY_0';

const channel = 'public:statuschanged';

void onEvent(event) async {
  print('client> $event');
}

Future<centrifuge.Client?> newCentrifugeClient() async {
  try {
    final client = centrifuge.createClient(
      'ws://localhost:8000/connection/websocket?format=protobuf',
      centrifuge.ClientConfig(
        // Uncomment to use example token based on secret key `secret` for user `testsuite_jwt`.
        token: centrifugoTkn,
        headers: <String, dynamic>{'X-Example-Header': 'example'},
      ),
    );
    client.connecting.listen(onEvent);
    client.connected.listen(onEvent);
    client.disconnected.listen(onEvent);

    // Handle async errors.
    client.error.listen(onEvent);

    // Server-side subscriptions.
    client.subscribing.listen(onEvent);
    client.subscribed.listen(onEvent);
    client.unsubscribed.listen(onEvent);
    client.publication.listen(onEvent);
    client.join.listen(onEvent);
    client.leave.listen(onEvent);

    final subscription = client.newSubscription(
      channel,
      centrifuge.SubscriptionConfig(
        token: centrifugoTkn,
      ),
    );

    onSubscriptionEvent(dynamic event) async {
      print(event);
    }

    // State changes.
    subscription.subscribing.listen(onSubscriptionEvent);
    subscription.subscribed.listen(onSubscriptionEvent);
    subscription.unsubscribed.listen(onSubscriptionEvent);

    // Messages.
    subscription.publication.listen(onSubscriptionEvent);
    subscription.join.listen(onSubscriptionEvent);
    subscription.leave.listen(onSubscriptionEvent);

    // Handle subscription async errors.
    subscription.error.listen(onSubscriptionEvent);

    await subscription.subscribe();
    await client.connect();

    return client;
  } catch (ex) {
    print(ex);
    return null;
  }

  // State changes.
}
