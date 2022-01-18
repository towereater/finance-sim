namespace FinanceLib.Models
{
    public class Wallet
    {
        // IBAN code of the wallet
        private string iban;
        public string IBAN
        {
            get { return iban; }
            set { iban = value; }
        }
        
        // Cash inside the wallet
        private double cash;
        public double Cash
        {
            get { return cash; }
            set { cash = value; }
        }
    }
}