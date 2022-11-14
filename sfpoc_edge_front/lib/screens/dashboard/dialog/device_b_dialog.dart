import 'package:sfpoc_edge_front/controllers/device_b_controller.dart';
import 'package:sfpoc_edge_front/models/device.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

double _width = 0;
double _fontSize = 0;
double _imageSize = 0;
Color _textColor = Colors.black;
Widget _makeLabeledImage(String asset, String label) {
  return Column(
    children: [
      ClipRRect(
        borderRadius: BorderRadius.circular(20),
        child: Image.asset(
          asset,
          width: _imageSize,
          height: _imageSize,
        ),
      ),
      const SizedBox(height: 10),
      Text(label,
          overflow: TextOverflow.ellipsis,
          style: TextStyle(color: _textColor, fontSize: _fontSize)),
    ],
  );
}

class DeviceBDialog extends StatelessWidget {
  final Device dev;
  late final DeviceBController controller;
  DeviceBDialog({Key? key, required this.dev, required this.controller})
      : super(key: key) {
    Get.put(controller);
  }

  Widget contentBox(DeviceBController controller) {
    var desktopWidget = Padding(
      padding: const EdgeInsets.fromLTRB(30, 30, 30, 10),
      child: SingleChildScrollView(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            _SensorInfoGridView(controller: controller),
            _ControlInfoGridView(controller: controller)
          ],
        ),
      ),
    );

    var mobileWidget = LayoutBuilder(builder: (ctx, ctis) {
      return SingleChildScrollView(
        scrollDirection: Axis.horizontal,
        child: ClipRRect(
          borderRadius: BorderRadius.circular(20),
          child: Padding(
            padding: const EdgeInsets.all(16.0),
            child: SizedBox(
              width: _width - 32,
              child: SingleChildScrollView(
                child: Column(
                  children: [
                    _SensorInfoGridView(controller: controller),
                    const SizedBox(height: 20),
                    _ControlInfoGridView(controller: controller)
                  ],
                ),
              ),
            ),
          ),
        ),
      );
    });
    return SizedBox(
      width: _width,
      child: _width >= 550 ? desktopWidget : mobileWidget,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Dialog(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(20),
      ),
      elevation: 15,
      backgroundColor: const Color(0xfff5f6fa),
      child: LayoutBuilder(builder: (ctx, ctis) {
        _width = ctis.maxWidth;
        _width = _width < 250 ? 250 : _width;
        _width = _width > 700 ? 700 : _width;
        _fontSize = 10 * ((_width - 250) / 450) + 13;
        double ratio = _width >= 550 ? 6 : 4;
        _imageSize = _width / ratio;
        return GetBuilder<DeviceBController>(builder: (controller) {
          return contentBox(controller);
        });
      }),
    );
  }
}

class _SensorInfoGridView extends StatelessWidget {
  final DeviceBController controller;
  _SensorInfoGridView({
    Key? key,
    required this.controller,
  }) : super(key: key);
  final sensorImages = <String>[
    "assets/images/temperature.png",
    "assets/images/humidity.png",
    "assets/images/cdc_value.png",
    "assets/images/soil_water_value.png",
  ];
  final sensorNames = <String>[
    "Temperature",
    "Humidity",
    "Illuminance",
    "Water Value",
  ];

  @override
  Widget build(BuildContext context) {
    var sensingWidget = <Widget>[
      Text('tem: ${controller.temperature}',
          overflow: TextOverflow.ellipsis,
          style: TextStyle(color: _textColor, fontSize: _fontSize)),
      Text('tem: ${controller.humidity}',
          overflow: TextOverflow.ellipsis,
          style: TextStyle(color: _textColor, fontSize: _fontSize)),
      Text('tem: ${controller.cdc}',
          overflow: TextOverflow.ellipsis,
          style: TextStyle(color: _textColor, fontSize: _fontSize)),
      Text('tem: ${controller.waterValue}',
          overflow: TextOverflow.ellipsis,
          style: TextStyle(color: _textColor, fontSize: _fontSize)),
    ];
    return GridView.builder(
      physics: const NeverScrollableScrollPhysics(),
      shrinkWrap: true,
      itemCount: 4,
      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: _width < 550 ? 2 : 4,
        crossAxisSpacing: 16,
        mainAxisSpacing: 16,
        childAspectRatio: _width < 550
            ? 0.8 + 3.5 * ((_width - 250) / 2500.0)
            : 0.4 + 2 * ((_width - 250) / 2500.0),
      ),
      itemBuilder: (context, index) => Column(
        children: [
          _makeLabeledImage(sensorImages[index], sensorNames[index]),
          sensingWidget[index],
        ],
      ),
    );
  }
}

class _ControlInfoGridView extends StatelessWidget {
  final DeviceBController controller;
  const _ControlInfoGridView({Key? key, required this.controller})
      : super(key: key);
  // final controlImages = <String>["fan_off.png", "LED.png", servo]
  Widget _makeLandscapeWidget() {
    return Row(
      children: [
        Expanded(
          flex: 2,
          child: Column(
            children: [
              Row(
                children: [
                  Column(
                    children: [
                      // LED Image
                      Image.asset('assets/images/LED.png',
                          width: _imageSize, height: _imageSize),
                      // LED Text
                      Text(
                        'LED',
                        style:
                            TextStyle(color: _textColor, fontSize: _fontSize),
                      ),
                    ],
                  ),
                  Expanded(
                    child: GetX<DeviceBController>(builder: (controller) {
                      return SliderTheme(
                        data: const SliderThemeData(
                            trackHeight: 8,
                            activeTickMarkColor: Colors.transparent,
                            activeTrackColor: Color(0xff212331),
                            inactiveTickMarkColor: Colors.transparent,
                            inactiveTrackColor: Color(0xff9b9fbb)),
                        child: Slider(
                          value: controller.light.value.toDouble(),
                          min: 0,
                          max: 100,
                          divisions: 10,
                          // activeColor: Colors.black.withOpacity(0.4),
                          // inactiveColor: Colors.black.withOpacity(0.4),
                          thumbColor: const Color(0xff212331),
                          label: controller.light.value.toString(),
                          onChanged: ((value) {
                            controller.light.value = value.toInt();
                          }),
                        ),
                      );
                    }),
                  ),
                ],
              ),
              const SizedBox(height: 20),
              Row(
                children: [
                  Column(
                    children: [
                      // LED Image
                      Image.asset(
                        'assets/images/servo.png',
                        width: _imageSize,
                        height: _imageSize,
                      ),
                      // LED Text
                      Text('SERVO',
                          style: TextStyle(
                              color: _textColor, fontSize: _fontSize)),
                    ],
                  ),
                  Expanded(
                    child: SliderTheme(
                      data: const SliderThemeData(
                          trackHeight: 8,
                          activeTickMarkColor: Colors.transparent,
                          activeTrackColor: Color(0xff212331),
                          inactiveTickMarkColor: Colors.transparent,
                          inactiveTrackColor: Color(0xff9b9fbb)),
                      child: GetX<DeviceBController>(builder: (controller) {
                        return Slider(
                          value: controller.servo.value.toDouble(),
                          min: 0,
                          max: 100,
                          divisions: 10,
                          // activeColor: Colors.black.withOpacity(0.4),
                          // inactiveColor: Colors.black.withOpacity(0.4),
                          thumbColor: const Color(0xff212331),
                          label: controller.servo.value.toString(),
                          onChanged: ((value) {
                            controller.servo.value = value.toInt();
                          }),
                        );
                      }),
                    ),
                  ),
                ],
              )
            ],
          ),
        ),
        Expanded(
          flex: 1,
          child: GetX<DeviceBController>(builder: (controller) {
            return Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                Column(
                  children: [
                    // LED Image
                    Image.asset(
                      controller.fan.value == 1
                          ? 'assets/images/fan_on.png'
                          : 'assets/images/fan_off.png',
                      width: _imageSize,
                      height: _imageSize,
                    ),
                    // LED Text
                    Text('FAN',
                        style:
                            TextStyle(color: _textColor, fontSize: _fontSize)),
                  ],
                ),
                Column(
                  children: [
                    Text(controller.fan.value == 1 ? 'On' : 'Off',
                        style:
                            TextStyle(color: _textColor, fontSize: _fontSize)),
                    Switch(
                      value: controller.fan.value == 1,
                      activeColor: const Color(0xff212331),
                      inactiveThumbColor: const Color(0xff9b9fbb),
                      inactiveTrackColor: const Color(0xff9b9fbb),
                      onChanged: (value) {
                        controller.fan.value = (value == false ? 0 : 1);
                      },
                    ),
                  ],
                ),
              ],
            );
          }),
        ),
      ],
    );
  }

  Widget _makePortraitWidget() {
    return Column(
      children: [
        Padding(
          padding: const EdgeInsets.fromLTRB(8, 0, 8, 10),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.start,
            children: [
              Column(
                children: [
                  // LED Image
                  Image.asset(
                    'assets/images/fan_off.png',
                    width: _imageSize,
                    height: _imageSize,
                  ),
                  // LED Text
                  Text('FAN',
                      style: TextStyle(color: _textColor, fontSize: _fontSize)),
                ],
              ),
              Padding(
                padding: const EdgeInsets.only(left: 20),
                child: Column(
                  children: [
                    GetX<DeviceBController>(builder: (controller) {
                      return Column(
                        children: [
                          Text(controller.fan.value == 1 ? 'On' : 'Off',
                              style: TextStyle(
                                  color: _textColor, fontSize: _fontSize)),
                          Switch(
                            value: controller.fan.value == 1,
                            activeColor: const Color(0xff212331),
                            inactiveThumbColor: const Color(0xff9b9fbb),
                            inactiveTrackColor: const Color(0xff9b9fbb),
                            onChanged: (value) {
                              controller.fan.value = (value == false ? 0 : 1);
                            },
                          ),
                        ],
                      );
                    }),
                  ],
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 10),
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
            children: [
              Column(
                children: [
                  // LED Image
                  Image.asset(
                    'assets/images/LED.png',
                    width: _imageSize,
                    height: _imageSize,
                  ),
                  // LED Text
                  Text('LED',
                      style: TextStyle(color: _textColor, fontSize: _fontSize)),
                ],
              ),
              Expanded(
                child: GetX<DeviceBController>(builder: (controller) {
                  return SliderTheme(
                    data: const SliderThemeData(
                        trackHeight: 8,
                        activeTickMarkColor: Colors.transparent,
                        activeTrackColor: Color(0xff212331),
                        inactiveTickMarkColor: Colors.transparent,
                        inactiveTrackColor: Color(0xff9b9fbb)),
                    child: Slider(
                      value: controller.light.value.toDouble(),
                      min: 0,
                      max: 100,
                      divisions: 10,
                      // activeColor: Colors.black.withOpacity(0.4),
                      // inactiveColor: Colors.black.withOpacity(0.4),
                      thumbColor: const Color(0xff212331),
                      label: controller.light.value.toString(),
                      onChanged: ((value) {
                        controller.light.value = value.toInt();
                      }),
                    ),
                  );
                }),
              ),
            ],
          ),
        ),
        const SizedBox(
          height: 20,
        ),
        Padding(
          padding: const EdgeInsets.all(8.0),
          child: Row(
            children: [
              Column(
                children: [
                  // LED Image
                  Image.asset(
                    'assets/images/servo.png',
                    width: _imageSize,
                    height: _imageSize,
                  ),
                  // LED Text
                  Text('SERVO',
                      style: TextStyle(color: _textColor, fontSize: _fontSize)),
                ],
              ),
              Expanded(
                child: SliderTheme(
                  data: const SliderThemeData(
                      trackHeight: 8,
                      activeTickMarkColor: Colors.transparent,
                      activeTrackColor: Color(0xff212331),
                      inactiveTickMarkColor: Colors.transparent,
                      inactiveTrackColor: Color(0xff9b9fbb)),
                  child: GetX<DeviceBController>(builder: (controller) {
                    return Slider(
                      value: controller.servo.value.toDouble(),
                      min: 0,
                      max: 100,
                      divisions: 10,
                      // activeColor: Colors.black.withOpacity(0.4),
                      // inactiveColor: Colors.black.withOpacity(0.4),
                      thumbColor: const Color(0xff212331),
                      label: controller.servo.value.toString(),
                      onChanged: ((value) {
                        controller.servo.value = value.toInt();
                      }),
                    );
                  }),
                ),
              ),
            ],
          ),
        ),
      ],
    );
  }

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: const BoxConstraints(maxWidth: 700),
      child: LayoutBuilder(builder: (ctx, ctis) {
        return _width >= 550 ? _makeLandscapeWidget() : _makePortraitWidget();
      }),
    );
  }
}
