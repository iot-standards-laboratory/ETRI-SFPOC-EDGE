import 'dart:convert';

import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/controllers/DashboardController.dart';
import 'package:sfpoc_edge_front/controllers/device_a_controller.dart';
import 'package:sfpoc_edge_front/controllers/device_b_controller.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:sfpoc_edge_front/screens/dashboard/dialog/device_a_dialog.dart';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;
import 'package:sfpoc_edge_front/screens/dashboard/dialog/device_b_dialog.dart';
import 'package:data_table_2/data_table_2.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class RegisteredDeviceField extends StatefulWidget {
  const RegisteredDeviceField({Key? key}) : super(key: key);

  @override
  _RegisteredDeviceFieldState createState() => _RegisteredDeviceFieldState();
}

class _RegisteredDeviceFieldState extends State<RegisteredDeviceField> {
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
            'Registered Device',
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
                  controller.devices.length,
                  (index) => _deviceRow(controller.devices[index], context),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

void showSnackBar(BuildContext ctx, String msg) {
  ScaffoldMessenger.of(ctx).showSnackBar(
    SnackBar(
      backgroundColor: Colors.black,
      content: SizedBox(
        height: 40,
        child: Center(
          child: Text(
            msg,
            style: const TextStyle(
              fontSize: 16,
              color: Colors.white,
            ),
          ),
        ),
      ),
    ),
  );
}

dynamic deviceBListener(BuildContext ctx, Device d) async {
  var response = await http.get(Uri.http(serverAddr, "services"),
      headers: <String, String>{"sname": d.sname!});

  if (response.statusCode != 200) {
    showSnackBar(ctx, "Please install service");
    return;
  }
  d.sid = response.body.toString();

  response = await http.get(Uri.http(serverAddr, "services/${d.sid}/${d.id}"));

  if (response.statusCode != 200) {
    showSnackBar(ctx, "Unexpected error");
    return;
  }

  var param = jsonDecode(response.body);

  var controller = DeviceBController(
    d,
    fanV: param["fan"],
    lightV: param["light"],
    servoV: param["servo"],
  );

  return showDialog(
    context: ctx,
    builder: (context) {
      return DeviceBDialog(
        dev: d,
        controller: controller,
      );
    },
  ).then(
    (value) {
      var payload = {
        "fan": controller.fan.value,
        "servo": controller.servo.value,
        "light": controller.light.value,
      };

      http.post(
          Uri.http(
              serverAddr, 'services/${controller.d.sid!}/${controller.d.id!}'),
          body: jsonEncode(payload));
      Get.delete<DeviceBController>();
    },
  );
}

dynamic deviceAListener(BuildContext ctx, Device d) async {
  var response = await http.get(Uri.http(serverAddr, "services"),
      headers: <String, String>{"sname": d.sname!});

  if (response.statusCode != 200) {
    showSnackBar(ctx, "Please install service");
    return;
  }
  d.sid = response.body.toString();

  response = await http.get(Uri.http(serverAddr, "services/${d.sid}/${d.id}"));

  if (response.statusCode != 200) {
    showSnackBar(ctx, "Unexpected error");
    return;
  }

  var param = jsonDecode(response.body);

  var controller = DeviceAController(d, pmessage: param['msg']);

  return showDialog(
    context: ctx,
    builder: (context) {
      return DeviceADialog(
        dev: d,
        controller: controller,
      );
    },
  ).then(
    (value) {
      Get.delete<DeviceAController>();
    },
  );
}

dynamic defaultListener(BuildContext ctx) async {
  ScaffoldMessenger.of(ctx).showSnackBar(
    const SnackBar(
      backgroundColor: Colors.black,
      content: SizedBox(
        height: 40,
        child: Center(
          child: Text(
            'default listener',
            style: TextStyle(
              fontSize: 16,
              color: Colors.white,
            ),
          ),
        ),
      ),
    ),
  );
}

dynamic cellTabListener(BuildContext ctx, Device d) {
  if (d.sname == "devicemanagerb") {
    return deviceBListener(ctx, d);
  } else if (d.sname == "devicemanagera") {
    return deviceAListener(ctx, d);
  } else {
    return defaultListener(ctx);
  }
}

DataRow _deviceRow(Device d, BuildContext ctx) {
  return DataRow(
    cells: [
      DataCell(Text(d.name!)),
      DataCell(Text(d.cid!), onTap: () {
        Clipboard.setData(ClipboardData(text: '${d.sid}/${d.id}'));
        ScaffoldMessenger.of(ctx).showSnackBar(
          const SnackBar(
            backgroundColor: Colors.black,
            content: SizedBox(
              height: 40,
              child: Center(
                child: Text(
                  'Copied service and device id on clipboard',
                  style: TextStyle(
                    fontSize: 16,
                    color: Colors.white,
                  ),
                ),
              ),
            ),
          ),
        );
      }),
      DataCell(Text(d.sname!), onTap: () => cellTabListener(ctx, d)),
    ],
  );
}
