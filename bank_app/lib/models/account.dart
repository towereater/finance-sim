class Account {
  final String id;
  final String iban;
  final String owner;
  final double cash;

  Account({
    required this.id,
    required this.iban,
    required this.owner,
    required this.cash,
  });

  factory Account.fromJson(Map<String, dynamic> json) {
    return switch (json) {
      {
        'id': String id,
        'iban': String iban,
        'owner': String owner,
        'cash': double cash,
      } => 
        Account(
          id: id,
          iban: iban,
          owner: owner,
          cash: cash,
        ),
      _ => throw const FormatException('Failed to load account'),
    };
  }
}