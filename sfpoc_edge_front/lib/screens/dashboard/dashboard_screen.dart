import 'package:sfpoc_edge_front/controllers/DashboardController.dart';
import 'package:sfpoc_edge_front/responsive.dart';
import 'package:sfpoc_edge_front/screens/dashboard/components/controller_field.dart';
import 'package:sfpoc_edge_front/screens/dashboard/components/discover_device_fields.dart';
import 'package:sfpoc_edge_front/screens/dashboard/components/my_services_field.dart';
import 'package:sfpoc_edge_front/screens/dashboard/components/registered_device_field.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

import '../../constants.dart';
import 'components/header.dart';

class DashboardScreen extends StatefulWidget {
  const DashboardScreen({Key? key}) : super(key: key);

  @override
  State<DashboardScreen> createState() => _DashboardScreenState();
}

class _DashboardScreenState extends State<DashboardScreen> {
  @override
  Widget build(BuildContext context) {
    // controller.loadData();
    return SafeArea(
      child: SingleChildScrollView(
        padding: const EdgeInsets.all(defaultPadding),
        child: Column(
          children: [
            const Header(),
            const SizedBox(height: defaultPadding),
            Row(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Expanded(
                  flex: 5,
                  child: Column(
                    children: const [
                      MyServicesField(),
                      SizedBox(height: defaultPadding),
                      ControllerField(),
                      SizedBox(height: defaultPadding),
                      RegisteredDeviceField(),
                      SizedBox(height: defaultPadding),
                      DiscoveredDeviceField(),
                    ],
                  ),
                ),
              ],
            )
          ],
        ),
      ),
    );
  }
}
