class Service {
  final String? name, id;
  final int? numOfDevs;

  Service({this.name, this.id, this.numOfDevs});

  factory Service.fromJson(dynamic json) {
    return Service(
      name: json['sname'],
      id: json['sid'],
      numOfDevs: json['ndevs'],
    );
  }
}
