from random import randint

class BankAccount:
    def __init__(self):
        self.iban = randint(1, 1000)
        self.cash = randint(0, 100000)