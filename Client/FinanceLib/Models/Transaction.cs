namespace FinanceLib
{
    public class Transaction
    {
        // IBAN of the sender account
        private string sender;
        public string Sender
        {
            get { return sender; }
            set { sender = value; }
        }
        
        // IBAN of the receiver account
        private string receiver;
        public string Receiver
        {
            get { return receiver; }
            set { receiver = value; }
        }
        
        // Amount of cash transferred
        private double amount;
        public double Amount
        {
            get { return amount; }
            set { amount = value; }
        }
    }
}