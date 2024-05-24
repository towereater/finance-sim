class User {
  final String name;
  final String surname;
  final String birth;
  final List<String> accounts;

  const User({
    required this.name,
    required this.surname,
    required this.birth,
    required this.accounts,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return switch (json) {
      {
        'name': String name,
        'surname': String surname,
        'birth': String birth,
        'accounts': List<dynamic>? accounts,
      } =>
        User(
          name: name,
          surname: surname,
          birth: birth,
          accounts: accounts != null ? List<String>.from(accounts) : [],
        ),
      _ => throw const FormatException('Failed to load user'),
    };
  }
}
