import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/controllers/DashboardController.dart';
import 'package:sfpoc_edge_front/models/service.dart';
import 'package:sfpoc_edge_front/responsive.dart';
import 'package:sfpoc_edge_front/screens/dashboard/components/service_info_card.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class MyServicesField extends StatelessWidget {
  const MyServicesField({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GetBuilder<DashboardController>(
      builder: (controller) {
        final Size _size = MediaQuery.of(context).size;
        return Column(
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Text(
                  "My Containers",
                  style: Theme.of(context).textTheme.subtitle1,
                ),
              ],
            ),
            SizedBox(height: defaultPadding),
            Responsive(
              mobile: ServiceInfoCardGridView(
                crossAxisCount: _size.width < 650 ? 2 : 4,
                childAspectRatio:
                    _size.width < 650 && _size.width > 350 ? 1.3 : 1,
                services: controller.services,
              ),
              tablet: ServiceInfoCardGridView(
                services: controller.services,
              ),
              desktop: ServiceInfoCardGridView(
                services: controller.services,
                childAspectRatio: _size.width < 1400 ? 1.1 : 1.4,
              ),
            )
          ],
        );
      },
    );
  }
}

class ServiceInfoCardGridView extends StatelessWidget {
  const ServiceInfoCardGridView({
    Key? key,
    this.crossAxisCount = 4,
    this.childAspectRatio = 1,
    required this.services,
  }) : super(key: key);

  final List<Service> services;
  final int crossAxisCount;
  final double childAspectRatio;

  @override
  Widget build(BuildContext context) {
    return GridView.builder(
      physics: NeverScrollableScrollPhysics(),
      shrinkWrap: true,
      itemCount: services.length,
      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: crossAxisCount,
        crossAxisSpacing: defaultPadding,
        mainAxisSpacing: defaultPadding,
        childAspectRatio: childAspectRatio,
      ),
      itemBuilder: (context, index) => ServiceInfoCard(info: services[index]),
    );
  }
}
