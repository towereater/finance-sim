import 'package:flutter/material.dart';

class RegisterScreen extends StatelessWidget {
  final nameController = TextEditingController();
  final surnameController = TextEditingController();
  final birthController = TextEditingController();
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  RegisterScreen({super.key});

  void registerButtonPressed(BuildContext context) {
    if (usernameController.text == 'andnic' &&
        passwordController.text == 'password') {
      ScaffoldMessenger.of(context)
          .showSnackBar(const SnackBar(content: Text('Registration complete')));

      Navigator.pop(context);

      nameController.clear();
      surnameController.clear();
      birthController.clear();
      usernameController.clear();
      passwordController.clear();
    } else {
      ScaffoldMessenger.of(context)
          .showSnackBar(const SnackBar(content: Text('Registration failed')));
    }
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
