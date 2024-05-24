import 'package:flutter/material.dart';

import 'package:bank_app/api/login_user.dart';

import 'package:bank_app/config/config.dart' as config;

class LoginScreen extends StatelessWidget {
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  LoginScreen({super.key});

  Future<void> loginButtonPressed(BuildContext context) async {
    String username = usernameController.text;
    String password = passwordController.text;

    await loginUser(username, password).then((value) {
      usernameController.clear();
      passwordController.clear();

      Navigator.pushNamed(context, config.routeHome, arguments: value);
    }).onError((error, stackTrace) {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(error.toString())));
    });
  }

  void registerButtonPressed(BuildContext context) {
    usernameController.clear();
    passwordController.clear();

    Navigator.pushNamed(context, config.routeRegister);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).primaryColor,
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          children: [
            Expanded(
              child: TextField(
                controller: usernameController,
                decoration: const InputDecoration(
                  labelText: 'Username',
                ),
              ),
            ),
            const SizedBox(height: 20.0),
            Expanded(
              child: TextField(
                controller: passwordController,
                decoration: const InputDecoration(
                  labelText: 'Password',
                ),
                obscureText: true,
                enableSuggestions: false,
                autocorrect: false,
              ),
            ),
            const SizedBox(height: 20.0),
            MaterialButton(
              color: Theme.of(context).highlightColor,
              onPressed: () async {
                await loginButtonPressed(context);
              },
              child: const Text('Login'),
            ),
            const SizedBox(height: 10.0),
            MaterialButton(
              color: Theme.of(context).primaryColor,
              onPressed: () => registerButtonPressed(context),
              child: const Text('Register'),
            ),
          ],
        ),
      ),
    );
  }
}
