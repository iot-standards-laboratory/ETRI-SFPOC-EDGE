import 'package:flutter/gestures.dart';
import 'package:flutter/material.dart';
import 'package:front/app/components/responsive.dart';
import 'package:front/app/model/agent.dart';
import 'package:front/app/modules/home/controllers/home_controller.dart';
import 'package:front/colors.dart';
import 'package:get/get.dart';

class AgentField extends GetView<HomeController> {
  const AgentField({super.key});

  Widget _render(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        Text(
          'Agents',
          style: Theme.of(context).textTheme.caption!.copyWith(
                color: Colors.white,
                fontSize: 20,
              ),
        ),
        const SizedBox(height: defaultPadding),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20),
          child: Container(
            padding: const EdgeInsets.all(defaultPadding),
            decoration: const BoxDecoration(
              color: secondaryColor,
              borderRadius: BorderRadius.all(Radius.circular(10)),
            ),
            child: Obx(
              () {
                return DataTable(
                  columns: [
                    DataColumn(
                      label: Text(
                        'Name',
                        style: Theme.of(context).textTheme.caption!.copyWith(
                              color: Colors.white,
                              fontSize: 18,
                              fontWeight: FontWeight.w700,
                            ),
                      ),
                    ),
                    DataColumn(
                      label: Text(
                        'ID',
                        style: Theme.of(context).textTheme.caption!.copyWith(
                              color: Colors.white,
                              fontSize: 18,
                              fontWeight: FontWeight.w700,
                            ),
                      ),
                    ),
                    DataColumn(
                      label: Text(
                        'Status',
                        style: Theme.of(context).textTheme.caption!.copyWith(
                              color: Colors.white,
                              fontSize: 18,
                              fontWeight: FontWeight.w700,
                            ),
                      ),
                    ),
                  ],
                  rows: List<DataRow>.generate(controller.agents.length,
                      (idx) => _agentRow(controller.agents[idx])),
                );
              },
            ),
          ),
        ),
      ],
    );
  }

  Widget _renderMobile(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.stretch,
      children: [
        Text(
          'Agents',
          style: Theme.of(context).textTheme.caption!.copyWith(
                color: Colors.white,
                fontSize: 20,
              ),
        ),
        const SizedBox(height: defaultPadding),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 10.0),
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
                  children: controller.agents
                      .map(
                        (e) => Padding(
                          padding: const EdgeInsets.only(right: 10),
                          child: _agentComponent(context, e),
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

  @override
  Widget build(BuildContext context) {
    return Responsive(
      desktop: _render(context),
      tablet: _renderMobile(context),
      mobile: _renderMobile(context),
    );
  }
}

// class AgentFieldComponent extends StatelessWidget {
//   final Agent agent;
//   const AgentFieldComponent({super.key, required this.agent});

//   @override
//   Widget build(BuildContext context) {
//     return Container();
//   }
// }

Widget _agentComponent(BuildContext context, Agent agent) {
  return Container(
    width: 360,
    height: 180,
    padding: const EdgeInsets.all(defaultPadding),
    decoration: const BoxDecoration(
      color: secondaryColor,
      borderRadius: BorderRadius.all(Radius.circular(10)),
    ),
    child: Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Text(
          agent.name,
          style: Theme.of(context).textTheme.caption!.copyWith(
                color: Colors.white,
                fontSize: 18,
              ),
        ),
        Text(
          agent.id,
          maxLines: 1,
          overflow: TextOverflow.ellipsis,
          style: const TextStyle(fontSize: 16),
        ),
        Text(
          agent.status,
          maxLines: 1,
          overflow: TextOverflow.ellipsis,
          style: TextStyle(
            fontSize: 16,
            color: agent.status == 'connected' ? Colors.green : Colors.red,
          ),
        ),
      ],
    ),
  );
}

DataRow _agentRow(Agent agent) {
  return DataRow(
    cells: [
      DataCell(Text(agent.name)),
      DataCell(Text(agent.id)),
      DataCell(Text(agent.status)),
    ],
  );
}
