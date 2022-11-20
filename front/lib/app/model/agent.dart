class Agent {
  final String name;
  final String id;
  final String status;

  int numberOfDevices = 0;
  Agent({required this.name, required this.id, required this.status});
  factory Agent.fromJson(dynamic json) {
    return Agent(
      name: json['name'],
      id: json['id'],
      status: json['status'],
    );
  }
}
