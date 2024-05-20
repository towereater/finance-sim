import 'package:flutter/material.dart';

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  String name = 'Andnic';
  List accounts = [
    ['100001', 'Andnic', 12],
    ['100002', 'Andnic', 0],
    ['100005', 'Andnic', 39],
  ];

  int selectedIndex = 0;

  accountTileOnTap() {
    //setState(() {
    //  selectedIndex = index;
    //});
    print('lol');
  }

  Widget accountListBuilder(BuildContext context, int index) {
    return ListTile(
      leading: const Icon(Icons.credit_card),
      title: Text(accounts[index][0]),
      onTap: () => accountTileOnTap(),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('My Home Banking'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          children: [
            Text('Welcome $name'),
            Row(
              children: [
                SizedBox(
                  width: 500,
                  height: 500,
                  child: ListView.builder(
                    itemCount: accounts.length,
                    itemBuilder: accountListBuilder,
                  ),
                ),
                const SizedBox(width: 50.0),
                Builder(
                  builder: (context) {
                    if (selectedIndex == 0) {
                      return Container(
                        child: const Text('Select an account to examine'),
                      );
                    } else {
                      return Row(
                        children: [
                          Text('Account ${accounts[selectedIndex][0]}'),
                          Text('Owner: ${accounts[selectedIndex][1]}'),
                          Text('Cash: ${accounts[selectedIndex][2]}\$'),
                        ],
                      );
                    }
                  },
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
