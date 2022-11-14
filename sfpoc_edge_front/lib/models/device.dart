class Device {
  final String? name, id, type, cid, sname;
  String? sid;
  Device({this.name, this.id, this.type, this.cid, this.sid, this.sname});

  factory Device.fromJson(dynamic json) {
    print(json);
    return Device(
      name: json['dname'],
      id: json['did'],
      type: json['type'],
      cid: json['cid'],
      sid: json['sid'],
      sname: json['sname'],
    );
  }
}
