class User {
  final String name;
  final String surname;
  final String birth;

  const User({
    required this.name,
    required this.surname,
    required this.birth
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return switch (json) {
      {
        'name': String name,
        'surname': String surname,
        'birth': String birth,
      } => 
        User(
          name: name,
          surname: surname,
          birth: birth
        ),
      _ => throw const FormatException('Failed to load user'),
    };
  }
}