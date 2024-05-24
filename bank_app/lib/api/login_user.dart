import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:bank_app/models/user.dart';

import 'package:bank_app/config/config.dart' as config;

Future<List<dynamic>> loginUser(String username, String password) async {
  final response = await http.post(
      Uri.http(config.serverHost, config.usersLogin),
      body: jsonEncode(<String, String>{
        'username': username,
        'password': password,
      }
    ),
  );

  if (response.statusCode == 200) {
    return [
      User.fromJson(jsonDecode(response.body) as Map<String, dynamic>),
      response.headers['jwt'],
    ];
  } else {
    return Future.error(Exception('Invalid credentials'));
  }
}