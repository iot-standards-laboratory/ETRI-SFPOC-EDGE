class Controller {
  final String name;
  final String id;
  final String agentId;

  int numberOfDevices = 0;
  Controller({required this.name, required this.id, required this.agentId});
  factory Controller.fromJson(dynamic json) {
    return Controller(
      name: json['name'],
      id: json['id'],
      agentId: json['agent_id'],
    );
  }
}
