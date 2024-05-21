import 'package:flutter/material.dart';

import 'package:bank_app/models/user.dart';

class LoginScreen extends StatelessWidget {
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  final User user = const User(
      name: 'And',
      surname: 'Nic',
      birth: '2023-04-23',
      accounts: ['100001', '100002', '100005']);

  LoginScreen({super.key});

  void loginButtonPressed(BuildContext context) {
    if (usernameController.text == 'andnic' &&
        passwordController.text == 'password') {
      Navigator.pushNamed(context, '/home', arguments: user);

      usernameController.clear();
      passwordController.clear();
    } else {
      ScaffoldMessenger.of(context)
          .showSnackBar(const SnackBar(content: Text('Wrong credentials')));
    }
  }

  void registerButtonPressed(BuildContext context) {
    usernameController.clear();
    passwordController.clear();

    Navigator.pushNamed(context, '/register');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).primaryColor,
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          children: [
            TextField(
              controller: usernameController,
              decoration: const InputDecoration(
                labelText: 'Username',
              ),
            ),
            const SizedBox(height: 20.0),
            TextField(
              controller: passwordController,
              decoration: const InputDecoration(
                labelText: 'Password',
              ),
              obscureText: true,
              enableSuggestions: false,
              autocorrect: false,
            ),
            const SizedBox(height: 20.0),
            MaterialButton(
              color: Theme.of(context).highlightColor,
              onPressed: () => loginButtonPressed(context),
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
