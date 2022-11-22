class Service {
  final String? name, id;
  final int? numOfCtrls;

  Service({this.name, this.id, this.numOfCtrls});

  factory Service.fromJson(dynamic json) {
    return Service(
      name: json['name'],
      id: json['id'],
      numOfCtrls: json['num_clnts'],
    );
  }
}