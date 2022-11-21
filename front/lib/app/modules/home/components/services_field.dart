import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:front/app/components/responsive.dart';
import 'package:front/app/model/service.dart';
import 'package:front/colors.dart';
import 'package:get/get.dart';

import '../controllers/home_controller.dart';

class ServicesField extends GetView<HomeController> {
  final pageController = PageController(viewportFraction: 1, keepPage: true);

  ServicesField({super.key});

  Widget _render(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        Text(
          'Services',
          style: Theme.of(context).textTheme.caption!.copyWith(
                color: Colors.white,
                fontSize: 20,
              ),
        ),
        const SizedBox(height: defaultPadding),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20.0),
          child: ScrollConfiguration(
            behavior: const MaterialScrollBehavior().copyWith(
              dragDevices: {
                PointerDeviceKind.mouse,
                PointerDeviceKind.touch,
                PointerDeviceKind.stylus,
              },
            ),
            child: SingleChildScrollView(
              scrollDirection: Axis.horizontal,
              physics: const BouncingScrollPhysics(),
              child: Obx(() {
                return Row(
                  children: controller.services
                      .map(
                        (e) => Padding(
                          padding: const EdgeInsets.only(right: 10),
                          child: ServiceFieldComponent(info: e),
                        ),
                      )
                      .toList(),
                );
              }),
            ),
          ),
        )
      ],
    );
  }

  Widget _renderMobile(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        Text(
          'Services',
          style: Theme.of(context).textTheme.caption!.copyWith(
                color: Colors.white,
                fontSize: 18,
              ),
        ),
        const SizedBox(
          height: defaultPadding * 1,
        ),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 10.0),
          child: SizedBox(
            height: 200,
            child: ScrollConfiguration(
              behavior: const MaterialScrollBehavior().copyWith(dragDevices: {
                PointerDeviceKind.mouse,
                PointerDeviceKind.touch,
                PointerDeviceKind.stylus
              }),
              child: PageView.builder(
                itemCount: controller.services.length,
                controller: pageController,

                // itemCount: pages.length,
                itemBuilder: (_, idx) {
                  return Padding(
                    padding: const EdgeInsets.only(right: 10),
                    child:
                        ServiceFieldComponent(info: controller.services[idx]),
                  );
                },
              ),
            ),
          ),
        ),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return Responsive(
      desktop: _render(context),
      tablet: _render(context),
      mobile: _renderMobile(context),
    );
  }
}

class ServiceFieldComponent extends StatelessWidget {
  final Service info;
  const ServiceFieldComponent({super.key, required this.info});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 280,
      height: 200,
      padding: const EdgeInsets.all(defaultPadding),
      decoration: const BoxDecoration(
        color: secondaryColor,
        borderRadius: BorderRadius.all(Radius.circular(10)),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 3),
                child: Text(
                  '${info.numOfCtrls} devices',
                  style: Theme.of(context).textTheme.caption!.copyWith(
                        color: Colors.white,
                        fontSize: 18,
                      ),
                ),
              ),
              IconButton(
                icon: info.id == ''
                    ? const Icon(
                        Icons.cancel_outlined,
                        color: Colors.red,
                      )
                    : const Icon(
                        Icons.check_circle_outline,
                        color: Colors.green,
                      ),
                onPressed: () async {
                  // if (info.id == '') {
                  //       var response = await http.post(
                  //         Uri.http(
                  //           serverAddr,
                  //           postSvcs,
                  //         ),
                  //         headers: <String, String>{
                  //           'Content-Type': 'application/json; charset=UTF-8',
                  //         },
                  //         body: jsonEncode(<String, String>{
                  //           'name': info.name!,
                  //         }),
                  //       );
                  //     }
                },
              ),
            ],
          ),
          Text(
            info.name!,
            maxLines: 1,
            overflow: TextOverflow.ellipsis,
            style: const TextStyle(fontSize: 20),
          ),
          TextButton(
            child: Text(
              info.id!,
              maxLines: 1,
              overflow: TextOverflow.ellipsis,
            ),
            onPressed: () {
              // launch('http://${serverAddr}/svc/${info.id}');
            },
          )
        ],
      ),
    );
  }
}
