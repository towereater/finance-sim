import 'dart:convert';

import 'package:flutter/services.dart';

// App routes
String routeStart = '/login';
String routeLogin = '/login';
String routeRegister = '/register';
String routeHome = '/home';

// Server configuration
String serverHost = '';

// API URIs
String usersLogin = '';
String usersRegister = '';

String accountsDetails = '';

// Loads the config file from local assents
Future<void> readConfig(String path) async {
  final String jsonData = await rootBundle.loadString(path);
  final data = json.decode(jsonData);

  // Config setup
  if (data['routes'] != null) {
    routeStart = data['routes']['start'] ? (data['routes']['start']) : routeStart;
    routeLogin = data['routes']['login'] ? (data['routes']['login']) : routeLogin;
    routeRegister = data['routes']['register'] ? (data['routes']['register']) : routeRegister;
    routeHome = data['routes']['home'] ? (data['routes']['home']) : routeHome;
  }

  serverHost = data['server']['host'];

  usersLogin = data['users']['login'];
  usersRegister = data['users']['register'];

  accountsDetails = data['accounts']['details'];
}
