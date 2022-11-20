import 'package:flutter/gestures.dart';
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
            ? PreferredSize(
                preferredSize: const Size(60, 60),
                child: AppBar(
                  toolbarHeight: 60,
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
                  actions: [
                    Padding(
                      padding: const EdgeInsets.all(7.0),
                      child: ProfileCard(
                        onPressed: () {},
                      ),
                    )
                  ],
                ),
              )
            : const PreferredSize(
                preferredSize: Size.zero,
                child: SizedBox(),
              ),
        body: SafeArea(
          child: Row(
            children: [
              if (!Responsive.isMobile(context))
                Container(
                  color: secondaryColor,
                  width: 200,
                  height: double.infinity,
                ),
              Expanded(
                child: SizedBox.expand(
                  child: ScrollConfiguration(
                    behavior: const MaterialScrollBehavior().copyWith(
                      dragDevices: {
                        PointerDeviceKind.mouse,
                        PointerDeviceKind.touch,
                        PointerDeviceKind.stylus,
                      },
                    ),
                    child: SingleChildScrollView(
                      padding: const EdgeInsets.symmetric(
                        vertical: defaultPadding,
                        horizontal: defaultPadding * 2,
                      ),
                      child: Column(
                        children: [
                          if (!Responsive.isMobile(context)) const Header(),
                          ServicesField(),
                          const SizedBox(height: defaultPadding),
                          const AgentField(),
                          const SizedBox(height: defaultPadding),
                          const ControllerField(),
                          const SizedBox(height: defaultPadding)
                        ],
                      ),
                    ),
                  ),
                ),
              ),
              const SizedBox(height: defaultPadding),
            ],
          ),
        ),
      ),
    );
  }
}
