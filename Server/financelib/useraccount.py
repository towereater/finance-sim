import financelib.bankaccount as ba

class UserAccount:
    def __init__(self, user, pwd):
        self.user = user
        self.pwd = pwd
        self.bankAcc = ba.BankAccount()
    
    def match_credentials(self, user, pwd):
        return self.user == user and self.pwd == pwd