class Service {
  final String? name, id;
  final int? numOfCtrls;
  final String? status;

  Service({this.name, this.id, this.numOfCtrls, this.status});

  factory Service.fromJson(dynamic json) {
    return Service(
      name: json['name'],
      id: json['id'],
      numOfCtrls: json['num_clnts'],
      status: json['status'],
    );
  }
}
