import 'dart:convert';
import 'package:http/http.dart' as http;

import 'package:bank_app/models/account.dart';

import 'package:bank_app/config/config.dart' as config;

Future<Account> getAccount(String authorizationToken, String accountId) async {
  final response = await http.get(
    Uri.http(config.serverHost, '${config.accountsDetails}/$accountId'),
    headers: {
      'Authorization': authorizationToken,
    },
  );

  if (response.statusCode == 200) {
    return Account.fromJson(jsonDecode(response.body) as Map<String, dynamic>);
  } else {
    return Future.error(Exception('Error while fetching account'));
  }
}
