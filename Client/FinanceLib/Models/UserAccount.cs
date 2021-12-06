namespace FinanceLib
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
        
        // Reference to the bank account of the user
        private BankAccount bankacc;
        public BankAccount BankAccount
        {
            get { return bankacc; }
            set { bankacc = value; }
        }
    }
}