import 'package:flutter/material.dart';
import 'package:front/app/components/responsive.dart';
import 'package:front/app/modules/home/components/profile_card.dart';
import 'package:front/colors.dart';

import 'package:get/get.dart';

import '../components/agent_field.dart';
import '../components/controller_field.dart';
import '../components/header.dart';
import '../components/services_field.dart';
import '../controllers/home_controller.dart';

class HomeView extends GetView<HomeController> {
  const HomeView({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        appBar: Responsive.isMobile(context)
            ? AppBar(
                elevation: 0,
                backgroundColor: bgColor,
                leading: IconButton(
                  onPressed: () {
                    controller.scaffoldKey.currentState!.openDrawer();
                  },
                  icon: const Icon(Icons.menu, color: Colors.white),
                ),
                title: Text(
                  "ETRI Smart Farm",
                  overflow: TextOverflow.ellipsis,
                  style: TextStyle(
                    fontSize: 18,
                    color: Colors.grey[100],
                    fontWeight: FontWeight.w800,
                  ),
                ),
                actions: const [
                  Padding(
                    padding: EdgeInsets.all(7.0),
                    child: ProfileCard(),
                  )
                ],
              )
            : const PreferredSize(
                preferredSize: Size.zero,
                child: SizedBox(),
              ),
        body: SafeArea(
          child: SingleChildScrollView(
            padding: const EdgeInsets.all(defaultPadding),
            child: Column(
              children: [
                if (!Responsive.isMobile(context)) const Header(),
                const SizedBox(height: defaultPadding),
                Row(
                  children: [
                    Expanded(
                        child: Column(
                      children: const [
                        ServicesField(),
                        SizedBox(height: defaultPadding),
                        AgentField(),
                        SizedBox(height: defaultPadding),
                        ControllerField(),
                        SizedBox(height: defaultPadding),
                      ],
                    ))
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
