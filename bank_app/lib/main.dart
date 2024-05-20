import 'package:flutter/material.dart';

import 'package:bank_app/screens/user_access/login_screen.dart';
import 'package:bank_app/screens/user_access/register_screen.dart';
import 'package:bank_app/screens/home/home_screen.dart';
import 'package:bank_app/config/config.dart' as config;

void main() {
  // Loading config from local files
  config.readConfig('assets/config.json');

  runApp(const MyApp(appTitle: 'My Home Banking'));
}

class MyApp extends StatelessWidget {
  final String appTitle;

  const MyApp({super.key, required this.appTitle});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: appTitle,
      theme: ThemeData(
        primaryColor: Colors.white70,
        dividerColor: Colors.blue[50],
        highlightColor: Colors.deepPurple,
      ),
      initialRoute: '/login',
      routes: {
        '/login': (context) => LoginScreen(),
        '/register':(context) => RegisterScreen(),
        '/home':(context) => HomeScreen(),
      },
    );
  }
}
