import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/controllers/DashboardController.dart';
import 'package:data_table_2/data_table_2.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:get/get.dart';

class ControllerField extends StatefulWidget {
  const ControllerField({Key? key}) : super(key: key);

  @override
  _ControllerFieldState createState() => _ControllerFieldState();
}

class _ControllerFieldState extends State<ControllerField> {
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
            'Registered Controller',
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
                    label: Text('NumOfDev'),
                  ),
                  DataColumn(
                    label: SelectableText('ID'),
                  ),
                ],
                rows: List.generate(
                  controller.controllers.length,
                  (index) => _controllerRow(
                    controller.controllers[index].name,
                    '${controller.controllers[index].numberOfDevices}',
                    controller.controllers[index].key,
                    context,
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }
}

DataRow _controllerRow(
    String name, String numOfDev, String id, BuildContext ctx) {
  return DataRow(
    cells: [
      DataCell(Text(name)),
      DataCell(Text(numOfDev)),
      DataCell(Text(id), onTap: () {
        Clipboard.setData(ClipboardData(text: id));
        ScaffoldMessenger.of(ctx).showSnackBar(const SnackBar(
            backgroundColor: Colors.black,
            content: SizedBox(
              height: 40,
              child: Center(
                child: Text(
                  'Copied',
                  style: TextStyle(
                    fontSize: 16,
                    color: Colors.white,
                  ),
                ),
              ),
            )));
      }),
    ],
  );
}
