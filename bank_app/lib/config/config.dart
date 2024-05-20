import 'dart:convert';

import 'package:flutter/services.dart';

// Server configuration
String serverHost = '';
String serverPort = '';

// API URIs
String usersLogin = '';
String usersRegister = '';

// Loads the config file from local assents
Future<void> readConfig(String path) async {
  final String jsonData = await rootBundle.loadString(path);
  final data = await json.decode(jsonData);

  // Config setup
  serverHost = data['server']['host'];
  serverPort = data['server']['port'];

  usersLogin = data['users']['login'];
  usersRegister = data['users']['register'];
}