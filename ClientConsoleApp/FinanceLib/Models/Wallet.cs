namespace FinanceLib.Models
{
    public class Wallet
    {
        // Id of the wallet
        private string id;
        public string Id
        {
            get { return id; }
            set { id = value; }
        }

        // IBAN code of the wallet
        private string iban;
        public string IBAN
        {
            get { return iban; }
            set { iban = value; }
        }

        // Owner of the wallet
        private string owner;
        public string Owner
        {
            get { return owner; }
            set { owner = value; }
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