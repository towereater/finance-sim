using FinanceLib.Models;
using FinanceLib.Utils;

namespace FinanceLib.Managers
{
    public class AccountManager
    {
        // Net manager in charge of making requests to the server
        private NetManager netMan;

        // User account connected to the manager
        public UserAccount UserAccount { get; protected set; }

        // Initialisation of the user account
        public bool LogIn(string user, string pass)
        {
            // Net manager initialisation
            netMan = NetManager.Initialize(user, pass);

            if (netMan != null) {
                UserAccount = new UserAccount() {
                    Username = user,
                };

                // Downloads all wallets data
                Wallet[] wallets = GetWallets();

                if (wallets != null)
                    UserAccount.Wallets = wallets;

                return true;
            }
            
            return false;
        }

        // Retrieves the data of a wallet of the user with the given iban.
        // If iban equals *, then all user's wallets will be given.
        public Wallet[] GetWallets(string iban = "*")
        {
            // Sets up the request
            BankRequest request = new BankRequest() {
                RequestToken = RequestToken.Wallets,
                Payload = new Dictionary<string, object>() {
                    ["iban"] = iban,
                },
            };

            // Sends the request to the server
            BankRequest response;
            bool respSuccess = netMan.Request(request, out response);

            if (respSuccess && response.ResponseToken == ResponseToken.Success)
                //return (Wallet[])response.Payload["wallets"];
                return JsonConverter.GetValue<Wallet[]>(response.Payload["wallets"]);

            return null;
        }
    }
}