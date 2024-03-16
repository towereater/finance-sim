namespace FinanceLib.Models
{
    public class UserAccount
    {
        // User info of the account
        private string name;
        public string Name
        {
            get { return name; }
            set { name = value; }
        }
        
        private string surname;
        public string Surname
        {
            get { return surname; }
            set { surname = value; }
        }
        
        private string birth;
        public string Birth
        {
            get { return birth; }
            set { birth = value; }
        }
        
        // Reference to the bank accounts of the user
        // private Wallet[] wallets;
        // public Wallet[] Wallets
        // {
        //     get { return wallets; }
        //     set { wallets = value; }
        // }
    }
}