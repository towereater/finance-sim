namespace FinanceLib
{
    public class BankAccount
    {
        // IBAN code of the account
        private int iban;
        public int IBAN
        {
            get { return iban; }
            set { iban = value; }
        }
        
        // Cash inside the account
        private double cash;
        public double Cash
        {
            get { return cash; }
            set { cash = value; }
        }
    }
}