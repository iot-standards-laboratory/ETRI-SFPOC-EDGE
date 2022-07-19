import 'dart:convert';

import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/controllers/DashboardController.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:data_table_2/data_table_2.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:http/http.dart' as http;

class DiscoveredDeviceField extends StatefulWidget {
  const DiscoveredDeviceField({Key? key}) : super(key: key);

  @override
  State<DiscoveredDeviceField> createState() => _DiscoveredDeviceFieldState();
}

class _DiscoveredDeviceFieldState extends State<DiscoveredDeviceField> {
  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.all(defaultPadding),
      decoration: const BoxDecoration(
        color: secondaryColor,
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            'Discovered Device',
            style: Theme.of(context).textTheme.subtitle1,
          ),
          SizedBox(
            width: double.infinity,
            child: GetBuilder<DashboardController>(
              builder: (controller) => DataTable2(
                columns: const [
                  DataColumn(
                    label: Text('Name'),
                  ),
                  DataColumn(
                    label: Text('Controller ID'),
                  ),
                  DataColumn(
                    label: SelectableText('Service Name'),
                  ),
                ],
                rows: List.generate(
                  controller.discoveredDevices.length,
                  (index) => _discoveredDeviceRow(
                      controller.discoveredDevices[index], context),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

Future<bool?> show(BuildContext context) async {
  return await showDialog<bool?>(
    context: context,
    builder: (BuildContext context) => AlertDialog(
      title: const Text('Device Registration'),
      content: const Text('Do you want to registrate this device??'),
      actions: <Widget>[
        TextButton(
          onPressed: () => Navigator.pop(context, false),
          child: const Text('Cancel'),
        ),
        TextButton(
          onPressed: () => Navigator.pop(context, true),
          child: const Text('OK'),
        ),
      ],
    ),
  );
}

DataRow _discoveredDeviceRow(Device d, BuildContext context) {
  return DataRow(
    cells: [
      DataCell(Text(d.name!)),
      DataCell(Text(d.cid!)),
      DataCell(Text(d.sname!), onTap: () async {
        var payload = jsonEncode(<String, String>{'did': d.id!});
        var result = await show(context);
        if (result!) {
          var response = await http.put(
            Uri.http(
              serverAddr,
              putDiscovered,
            ),
            headers: <String, String>{
              'Content-Type': 'application/json; charset=UTF-8',
            },
            body: payload,
          );
        }
      }),
    ],
  );
}
