import 'package:flutter/material.dart';

import 'package:bank_app/api/register_user.dart';

class RegisterScreen extends StatelessWidget {
  final nameController = TextEditingController();
  final surnameController = TextEditingController();
  final birthController = TextEditingController();
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  RegisterScreen({super.key});

  Future<void> registerButtonPressed(BuildContext context) async {
    String name = nameController.text;
    String surname = surnameController.text;
    String birth = birthController.text;
    String username = usernameController.text;
    String password = passwordController.text;

    await registerUser(username, password, name, surname, birth).then((value) {
      usernameController.clear();
      passwordController.clear();
      nameController.clear();
      surnameController.clear();
      birthController.clear();

      Navigator.pop(context);

      ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Registration complete!')));
    }).onError((error, stackTrace) {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(error.toString())));
    });
  }

  void backButtonPressed(BuildContext context) {
    Navigator.pop(context);
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
              controller: nameController,
              decoration: const InputDecoration(
                labelText: 'Name',
              ),
            ),
            const SizedBox(height: 20.0),
            TextField(
              controller: surnameController,
              decoration: const InputDecoration(
                labelText: 'Surname',
              ),
            ),
            const SizedBox(height: 20.0),
            TextField(
              controller: birthController,
              decoration: const InputDecoration(
                labelText: 'Birth',
              ),
            ),
            const SizedBox(height: 20.0),
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
              onPressed: () => registerButtonPressed(context),
              child: const Text('Register'),
            ),
            const SizedBox(height: 10.0),
            MaterialButton(
              color: Theme.of(context).primaryColor,
              onPressed: () => backButtonPressed(context),
              child: const Text('Back'),
            ),
          ],
        ),
      ),
    );
  }
}
