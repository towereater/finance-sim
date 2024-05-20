import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:bank_app/config/config.dart' as config;
import 'package:bank_app/models/user.dart';

Future<User> loginUser(String username, String password) async {
  final response = await http.post(
      Uri.http(config.serverHost + config.serverPort, config.usersLogin),
      headers: <String, String>{
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': '*/*',
      },
      body: jsonEncode(<String, String>{
        'username': username,
        'password': password,
      }
    ),
  );

  if (response.statusCode == 200) {
    return User.fromJson(jsonDecode(response.body) as Map<String, dynamic>);
  } else {
    throw Exception('Invalid credentials');
  }
}