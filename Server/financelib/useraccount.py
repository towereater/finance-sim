
class UserAccount:
    def __init__(self, user, pwd):
        self._user = user
        self._pwd = pwd
        self._wallets = []
    
    @property
    def user(self):
        return self._user
    
    @property
    def password(self):
        return self._pwd
    
    @property
    def wallets(self):
        return self._wallets
    
    @wallets.setter
    def wallets(self, value):
        self._wallets = value

    @wallets.deleter
    def wallets(self):
        del self._wallets

    def match_credentials(self, user, pwd):
        return self.user == user and self.password == pwd