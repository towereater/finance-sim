import 'package:flutter/material.dart';

import 'package:bank_app/models/user.dart';

class HomeScreen extends StatefulWidget {
  final User user;

  const HomeScreen({super.key, required this.user});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  List accounts = [
    ['100001', 'Andnic', 12],
    ['100002', 'Andnic', 0],
    ['100005', 'Andnic', 39],
  ];

  int selectedIndex = -1;

  changeSelectedItem(int index) {
    setState(() {
      selectedIndex = index;
    });
  }

  Widget accountListBuilder(BuildContext context, int index) {
    return ListTile(
      leading: const Icon(Icons.credit_card),
      title: Text(widget.user.accounts[index]),
      onTap: () => changeSelectedItem(index),
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
                    Navigator.popUntil(context, ModalRoute.withName('/login')),
                  },
              icon: const Icon(Icons.logout)),
        ],
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Row(
          children: [
            Flexible(
              flex: 1,
              child: ListView.builder(
                itemCount: widget.user.accounts.length,
                itemBuilder: accountListBuilder,
              ),
            ),
            Flexible(
              flex: 4,
              child: Builder(
                builder: (context) {
                  if (selectedIndex == -1) {
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
                            'Account ${accounts[selectedIndex][0]}',
                            style: const TextStyle(
                              fontSize: 18,
                            ),
                          ),
                          const SizedBox(
                            height: 40,
                          ),
                          Text(
                            'Owner: ${accounts[selectedIndex][1]}',
                            style: const TextStyle(
                              fontSize: 15,
                            ),
                          ),
                          Text(
                            'Cash: ${accounts[selectedIndex][2]}\$',
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
