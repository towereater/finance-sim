import 'package:flutter/material.dart';

import 'package:bank_app/screens/user_access/login_screen.dart';
import 'package:bank_app/screens/user_access/register_screen.dart';
import 'package:bank_app/screens/home/home_screen.dart';

import 'package:bank_app/models/user.dart';

import 'package:bank_app/config/config.dart' as config;

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  // Loading config from local files
  await config.readConfig('assets/config.json');

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'My Home Banking',
      theme: ThemeData(
        primaryColor: Colors.white70,
        dividerColor: Colors.blue[50],
        highlightColor: Colors.deepPurple,
      ),
      initialRoute: config.routeStart,
      routes: {
        config.routeLogin: (context) => LoginScreen(),
        config.routeRegister: (context) => RegisterScreen(),
        config.routeHome: (context) {
          List<dynamic> args =
              ModalRoute.of(context)?.settings.arguments as List<dynamic>;
          return HomeScreen(
            user: args[0] as User,
            authorizationToken: args[1],
          );
        },
      },
    );
  }
}
