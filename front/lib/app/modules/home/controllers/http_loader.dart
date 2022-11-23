import 'dart:convert';

import 'package:front/app/model/agent.dart';
import 'package:front/app/model/controller.dart';
import 'package:front/app/model/service.dart';
import 'package:front/constants.dart';
import 'package:http/http.dart' as http;

Future<List<Service>> loadServices() async {
  var resp = await http.get(Uri.http(serverAddr, '/api/v2/svcs'));

  var body = jsonDecode(resp.body);

  var services = List.generate(
    body.length,
    (index) => Service.fromJson((body[index])),
  ).toList();

  return services;
  // return <Service>[];
}

Future<List<Agent>> loadAgents() async {
  var resp = await http.get(Uri.http(serverAddr, '/api/v2/agents'));

  var body = jsonDecode(resp.body);

  var agents = List.generate(body.length, (index) {
    return Agent.fromJson(body[index]);
  }).toList();

  return agents;
}

Future<List<Controller>> loadCtrls() async {
  var resp = await http.get(Uri.http(serverAddr, '/api/v2/ctrls'));

  var body = jsonDecode(resp.body);
  print(body);

  var ctrls = List.generate(body.length, (index) {
    return Controller.fromJson(body[index]);
  }).toList();

  return ctrls;
}

Future<int> installSvc(Service svc) async {
  var resp = await http.post(
    Uri.http(serverAddr, '/api/v2/svcs'),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
      'service_name': svc.name!,
    },
  );

  return resp.statusCode;
}
