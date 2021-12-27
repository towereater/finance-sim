namespace FinanceLib.Models
{
    public class UserAccount
    {
        // Username of the account
        private string user;
        public string Username
        {
            get { return user; }
            set { user = value; }
        }
        
        // Reference to the bank accounts of the user
        private Wallet[] wallets;
        public Wallet[] Wallets
        {
            get { return wallets; }
            set { wallets = value; }
        }
    }
}