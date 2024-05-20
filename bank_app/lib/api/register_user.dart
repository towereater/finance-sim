import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:bank_app/config/config.dart' as config;

Future<String> registerUser(
  String username,
  String password,
  String name,
  String surname,
  String birth) async {
  final response = await http.post(
      Uri.http(config.serverHost + config.serverPort, config.usersRegister),
      headers: <String, String>{
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': '*/*',
      },
      body: jsonEncode(<String, String>{
        'username': username,
        'password': password,
        'name': name,
        'surname': surname,
        'birth': birth,
      }
    ),
  );

  if (response.statusCode == 200) {
    return 'Registration complete!';
  } else {
    return response.body;
  }
}