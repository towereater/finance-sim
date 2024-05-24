import 'package:flutter/material.dart';

import 'package:bank_app/models/user.dart';
import 'package:bank_app/models/account.dart';
import 'package:bank_app/api/get_account.dart';

import 'package:bank_app/config/config.dart' as config;

class HomeScreen extends StatefulWidget {
  final User user;
  final String authorizationToken;

  const HomeScreen({super.key, required this.user, required this.authorizationToken});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  Account? account;
  int selectedIndex = -1;

  Future<void> accountTileOnTap(BuildContext context, int index) async {
    setState(() {
      selectedIndex = index;
    });

    String acccountId = widget.user.accounts[selectedIndex];

    await getAccount(widget.authorizationToken, acccountId).then((value) {
      setState(() {
        account = value;
      });
    }).onError((error, stackTrace) {
      ScaffoldMessenger.of(context)
          .showSnackBar(SnackBar(content: Text(error.toString())));
    });
  }

  Widget accountListBuilder(BuildContext context, int index) {
    return ListTile(
      leading: const Icon(Icons.credit_card),
      title: Text(widget.user.accounts[index]),
      onTap: () => accountTileOnTap(context, index),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Theme.of(context).primaryColor,
      appBar: AppBar(
        title: Align(
          alignment: Alignment.topLeft,
          child: Text(
            'Welcome ${widget.user.name}',
            style: const TextStyle(
              fontSize: 24,
            ),
          ),
        ),
        backgroundColor: Theme.of(context).primaryColor,
        automaticallyImplyLeading: false,
        actions: [
          IconButton(
              onPressed: () => {
                    Navigator.popUntil(
                        context, ModalRoute.withName(config.routeLogin)),
                  },
              icon: const Icon(Icons.logout)),
        ],
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Row(
          children: [
            Flexible(
              flex: 2,
              child: ListView.builder(
                itemCount: widget.user.accounts.length,
                itemBuilder: accountListBuilder,
              ),
            ),
            Flexible(
              flex: 5,
              child: Builder(
                builder: (context) {
                  if (selectedIndex == -1 || account == null) {
                    return const Padding(
                      padding: EdgeInsets.all(20.0),
                      child: Align(
                        alignment: Alignment.topCenter,
                        child: Text(
                          'Select an account to examine',
                          style: TextStyle(
                            fontSize: 18,
                          ),
                        ),
                      ),
                    );
                  } else {
                    return Padding(
                      padding: const EdgeInsets.all(20.0),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            'Account IBAN: ${account!.iban}',
                            style: const TextStyle(
                              fontSize: 18,
                            ),
                          ),
                          const SizedBox(
                            height: 40,
                          ),
                          Text(
                            'Owner: ${account!.owner}',
                            style: const TextStyle(
                              fontSize: 15,
                            ),
                          ),
                          Text(
                            'Cash: ${account!.cash}\$',
                            style: const TextStyle(
                              fontSize: 15,
                            ),
                          ),
                        ],
                      ),
                    );
                  }
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
