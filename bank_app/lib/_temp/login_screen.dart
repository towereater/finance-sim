// import 'package:bank_app/api/login_user.dart';
// import 'package:bank_app/screens/home/home_screen.dart';
// import 'package:bank_app/screens/user_access/register_screen.dart';
// import 'package:flutter/material.dart';

// import 'package:bank_app/models/user.dart';

// class LoginScreen extends StatefulWidget {
//   const LoginScreen({super.key});

//   @override
//   State<LoginScreen> createState() => _LoginScreenState();
// }

// class _LoginScreenState extends State<LoginScreen> {
//   final formKey = GlobalKey<FormState>();

//   String username = '';
//   String password = '';

//   Future<User>? futureUser;

//   @override
//   Widget build(BuildContext context) {
//     return Form(
//       key: formKey,
//       child: Column(
//         crossAxisAlignment: CrossAxisAlignment.center,
//         children: [
//           TextFormField(
//             validator: (value) {
//               if (value == null || value.isEmpty) {
//                 return 'Please enter username';
//               }
//               return null;
//             },
//             onSaved: (value) {
//               username = value!;
//             },
//           ),
//           TextFormField(
//             validator: (value) {
//               if (value == null || value.isEmpty) {
//                 return 'Please enter password';
//               }
//               return null;
//             },
//             onSaved: (value) {
//               password = value!;
//             },
//           ),
//           Padding(
//             padding: const EdgeInsets.symmetric(vertical: 16),
//             child: ElevatedButton(
//               onPressed: () {
//                 if (formKey.currentState!.validate()) {
//                   formKey.currentState!.save();

//                   setState(() {
//                     futureUser = loginUser(username, password);
//                   });
//                 }
//               },
//               child: const Text('Login'),
//             ),
//           ),
//           Padding(
//             padding: const EdgeInsets.symmetric(vertical: 16),
//             child: ElevatedButton(
//               onPressed: () {
//                 Navigator.push(
//                   context,
//                   MaterialPageRoute(
//                     builder: (context) => const RegisterScreen(),
//                   ),
//                 );
//               },
//               child: const Text('Register'),
//             ),
//           ),
//           FutureBuilder(
//             future: futureUser,
//             builder: (context, snapshot) {
//               if (snapshot.hasData) {
//                 Navigator.push(
//                   context,
//                   MaterialPageRoute(
//                     builder: (context) => HomeScreen(
//                       user: snapshot.requireData,
//                     ),
//                   ),
//                 );
//               } else if (snapshot.hasError) {
//                 return Text('${snapshot.error}');
//               }

//               return Container();
//             }
//           ),
//         ],
//       ),
//     );
//   }
// }
