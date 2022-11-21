import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:get/get.dart';

import '../controllers/home_controller.dart';

class SideMenu extends GetView<HomeController> {
  const SideMenu({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Drawer(
      width: 200,
      child: ListView(
        children: [
          DrawerHeader(
            child: Image.asset("assets/images/logo.png"),
          ),
          Obx(() {
            return DrawerListTile(
              title: "Dashboard",
              svgSrc: "assets/icons/menu_dashbord.svg",
              press: () {
                controller.menuIdx.value = 0;
                if (controller.scaffoldKey.currentState!.isDrawerOpen) {
                  Future.delayed(const Duration(milliseconds: 100), () {
                    Get.back();
                  });
                }
              },
              color:
                  controller.menuIdx.value == 0 ? Colors.white : Colors.white54,
            );
          }),
          GetX<HomeController>(
            builder: (ctrl) => DrawerListTile(
              title: "Logs",
              svgSrc: "assets/icons/menu_tran.svg",
              press: () {
                controller.menuIdx.value = 1;
                if (controller.scaffoldKey.currentState!.isDrawerOpen) {
                  Future.delayed(const Duration(milliseconds: 100), () {
                    Get.back();
                  });
                }
              },
              color:
                  controller.menuIdx.value == 1 ? Colors.white : Colors.white54,
            ),
          ),
        ],
      ),
    );
  }
}

class DrawerListTile extends StatelessWidget {
  const DrawerListTile({
    Key? key,
    // For selecting those three line once press "Command+D"
    required this.title,
    required this.svgSrc,
    required this.press,
    required this.color,
  }) : super(key: key);

  final String title, svgSrc;
  final VoidCallback press;
  final Color color;
  @override
  Widget build(BuildContext context) {
    return ListTile(
      onTap: press,
      horizontalTitleGap: 0.0,
      leading: SvgPicture.asset(
        svgSrc,
        color: color,
        height: 16,
      ),
      title: Text(
        title,
        style: TextStyle(color: color),
      ),
    );
  }
}
