from random import randint

class Wallet:
    def __init__(self):
        self._iban = str(randint(1, 1000))
        self._cash = randint(0, 100000)
    
    @property
    def iban(self):
        return self._iban
    
    @property
    def cash(self):
        return self._cash
    
    @cash.setter
    def cash(self, value):
        self._cash = value
    
    def to_json(self):
        return {
            'iban': self.iban,
            'cash': self.cash,
        }