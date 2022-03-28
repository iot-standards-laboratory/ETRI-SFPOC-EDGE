import 'package:sfpoc_edge_front/controllers/device_a_controller.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class DeviceADialog extends StatelessWidget {
  final Device dev;
  late final DeviceAController controller;
  DeviceADialog({Key? key, required this.dev, required this.controller})
      : super(key: key) {
    Get.put(controller);
  }

  @override
  Widget build(BuildContext context) {
    return Dialog(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(20),
      ),
      elevation: 15,
      backgroundColor: const Color(0xfff5f6fa),
      child: GetX<DeviceAController>(
        builder: (controller) => Padding(
          padding: const EdgeInsets.all(20.0),
          child: Text('message: ${controller.message}',
              style: const TextStyle(
                  fontSize: 20,
                  color: Colors.black87,
                  fontWeight: FontWeight.bold)),
        ),
      ),
    );
  }
}
