// import 'package:flutter/material.dart';

// import 'package:bank_app/api/register_user.dart';

// class RegisterScreen extends StatefulWidget {
//   const RegisterScreen({super.key});

//   @override
//   State<RegisterScreen> createState() => _RegisterScreenState();
// }

// class _RegisterScreenState extends State<RegisterScreen> {
//   final formKey = GlobalKey<FormState>();

//   String username = '';
//   String password = '';
//   String name = '';
//   String surname = '';
//   String birth = '';

//   Future<String>? registrationMessage;

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
//           TextFormField(
//             validator: (value) {
//               if (value == null || value.isEmpty) {
//                 return 'Please enter name';
//               }
//               return null;
//             },
//             onSaved: (value) {
//               name = value!;
//             },
//           ),
//           TextFormField(
//             validator: (value) {
//               if (value == null || value.isEmpty) {
//                 return 'Please enter surname';
//               }
//               return null;
//             },
//             onSaved: (value) {
//               surname = value!;
//             },
//           ),
//           TextFormField(
//             validator: (value) {
//               if (value == null || value.isEmpty) {
//                 return 'Please enter birth date';
//               }
//               return null;
//             },
//             onSaved: (value) {
//               birth = value!;
//             },
//           ),
//           Padding(
//             padding: const EdgeInsets.symmetric(vertical: 16),
//             child: ElevatedButton(
//               onPressed: () {
//                 if (formKey.currentState!.validate()) {
//                   formKey.currentState!.save();

//                   setState(() {
//                     registrationMessage = registerUser(username, password, name, surname, birth);
//                   });
//                 }
//               },
//               child: const Text('Register'),
//             ),
//           ),
//           FutureBuilder(
//             future: registrationMessage,
//             builder: (context, snapshot) {
//               if (snapshot.hasData) {
//                 Navigator.pop(context);
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
