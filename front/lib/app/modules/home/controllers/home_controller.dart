import 'package:flutter/material.dart';
import 'package:front/app/model/service.dart';
import 'package:get/get.dart';

class HomeController extends GetxController {
  //TODO: Implement HomeController
  var scaffoldKey = GlobalKey<ScaffoldState>();
  var services = <Service>[
    Service(id: '1-1-1', name: 'devicemanagera', numOfDevs: 7),
    Service(id: '1-1-1', name: 'devicemanagera', numOfDevs: 7),
    Service(id: '1-1-1', name: 'devicemanagera', numOfDevs: 7),
    Service(id: '1-1-1', name: 'devicemanagera', numOfDevs: 7),
    Service(id: '1-1-1', name: 'devicemanagera', numOfDevs: 7),
  ].obs;
  @override
  void onInit() {
    super.onInit();
  }

  @override
  void onReady() {
    super.onReady();
  }

  @override
  void onClose() {
    super.onClose();
  }
}
