class Controller {
  final String name, id, key;
  int numberOfDevices = 0;
  Controller({required this.name, required this.id, required this.key});
  factory Controller.fromJson(dynamic json) {
    return Controller(
      name: json['cname'],
      id: json['cid'],
      key: json['key'],
    );
  }
}
