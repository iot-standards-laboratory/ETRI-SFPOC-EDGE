import 'package:flutter/material.dart';

const primaryColor = Color(0xFF2697FF);
const secondaryColor = Color(0xFF2A2D3E);
const bgColor = Color(0xFF212332);

const defaultPadding = 16.0;
const defaultRound = 10.0;
// var serverAddr = '${Uri.base.host}:${Uri.base.port}';
late String serverAddr;
const getSvcsList = '/api/v1/svcs/list';
const getCtrlsList = '/api/v1/ctrls/list';
const getDevsList = '/api/v1/devs/list';
const getDiscoveredList = '/api/v1/devs/discover/list';
const putDiscovered = '/api/v1/devs/discover';
const publishUrl = '/api/v1/pub';
