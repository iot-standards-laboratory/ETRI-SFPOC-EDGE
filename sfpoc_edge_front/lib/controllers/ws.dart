import 'package:sfpoc_edge_front/constants.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

final ws = WebSocketChannel.connect(
  Uri.parse('ws://$serverAddr$publishUrl'),
);
