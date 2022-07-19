import 'dart:convert';
import 'dart:math';

import 'package:sfpoc_edge_front/constants.dart';
import 'package:sfpoc_edge_front/models/service.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;
import 'package:url_launcher/url_launcher.dart';

class ServiceInfoCard extends StatelessWidget {
  ServiceInfoCard({Key? key, required this.info}) : super(key: key);
  final color = colors[next(0, 5)];
  final Service info;
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
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Expanded(
                  child: Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 3),
                    child: Text(
                      "${info.numOfDevs} devices",
                      style: Theme.of(context)
                          .textTheme
                          .caption!
                          .copyWith(color: Colors.white),
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
                  color: Colors.white54,
                  onPressed: () async {
                    if (info.id == '') {
                      var response = await http.post(
                        Uri.http(
                          serverAddr,
                          postSvcs,
                        ),
                        headers: <String, String>{
                          'Content-Type': 'application/json; charset=UTF-8',
                        },
                        body: jsonEncode(<String, String>{
                          'name': info.name!,
                        }),
                      );
                    }
                  },
                )
              ],
            ),
            Text(
              info.name!,
              maxLines: 1,
              overflow: TextOverflow.ellipsis,
            ),
            TextButton(
              child: Text(
                info.id!,
                maxLines: 1,
                overflow: TextOverflow.ellipsis,
              ),
              onPressed: () {
                launch('http://${serverAddr}/svc/${info.id}');
              },
            ),
            // ProgressLine(
            //   color: color,
            //   percentage: info.percentage,
            // ),
          ]),
    );
  }
}

var svgs = [
  "assets/icons/Documents.svg",
  "assets/icons/google_drive.svg",
  "assets/icons/drop_box.svg",
  "assets/icons/one_drive.svg"
];

var colors = [
  Colors.blue,
  Colors.green,
  Colors.orange,
  Colors.purple,
  Colors.amber,
  Colors.lightBlue
];
final _random = new Random();
int next(int min, int max) => min + _random.nextInt(max - min);

class ProgressLine extends StatelessWidget {
  const ProgressLine({
    Key? key,
    this.color = primaryColor,
    required this.percentage,
  }) : super(key: key);

  final Color? color;
  final int? percentage;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Container(
          width: double.infinity,
          height: 5,
          decoration: BoxDecoration(
            color: color!.withOpacity(0.1),
            borderRadius: BorderRadius.all(Radius.circular(10)),
          ),
        ),
        LayoutBuilder(
          builder: (context, constraints) => Container(
            width: constraints.maxWidth * (percentage! / 100),
            height: 5,
            decoration: BoxDecoration(
              color: color,
              borderRadius: BorderRadius.all(Radius.circular(10)),
            ),
          ),
        ),
      ],
    );
  }
}
