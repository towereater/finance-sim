using FinanceLib.Models;
using FinanceLib.Utils;

namespace FinanceLib.Managers
{
    public class AccountManager
    {
        // Net manager in charge of making requests to the server
        private NetManager netMan = new NetManager();

        // Verifies whether or not a login was made by checking the NetManager access token
        public bool IsLogged { get { return !string.IsNullOrEmpty(netMan.AccessToken); } }

        // User account connected to the manager
        public UserAccount UserAccount { get; protected set; }

        // Account creation procedure
        public string CreateAccount(string user, string pass)
        {
            // Sets up the request
            BankRequest request = new BankRequest() {
                RequestToken = RequestToken.CreateAccount,
                Payload = new Dictionary<string, object>() {
                    ["user"] = user,
                    ["password"] = pass,
                },
            };

            // Sends the request to the server and gets the response
            BankRequest response;
            NetManager.ServerRequest(request, out response);
            return response.Payload["message"].ToString();
        }

        // Login procedure
        public string LogIn(string user, string pass)
        {
            // Net manager initialisation
            string message = netMan.Initialize(user, pass);

            // If login was made successfully
            if (IsLogged) {
                UserAccount = new UserAccount() {
                    Username = user,
                };
            }
            
            return message;
        }

        // Logout procedure
        public string LogOut()
        {
            // Drops all session data
            netMan = null;
            UserAccount = null;

            return "Logout successful";
        }

        public string CreateWallet()
        {
            // Sets up the request
            BankRequest request = new BankRequest() {
                RequestToken = RequestToken.Wallets,
                Payload = new Dictionary<string, object>(),
            };

            // Sends the request to the server
            BankRequest response;
            bool respSuccess = netMan.Request(request, out response);
            string message = response.Payload["message"].ToString();

            if (respSuccess && response.ResponseToken == ResponseToken.Success) {
                Wallet[] wallets = JsonConverter.GetValue<Wallet[]>(response.Payload["wallets"]);
                return @$"
                    {message}\n
                    Wallet created:\n
                    IBAN: {wallets[0].IBAN}";
            }

            return message;
        }

        // Retrieves the data of a wallet of the user with the given iban.
        // If iban equals *, then all user's wallets will be given.
        //TODO: FIX DESCRIPTION
        public string GetWallets()
        {
            // Sets up the request
            BankRequest request = new BankRequest() {
                RequestToken = RequestToken.Wallets,
                Payload = new Dictionary<string, object>(),
            };

            // Sends the request to the server
            BankRequest response;
            bool respSuccess = netMan.Request(request, out response);
            string message = response.Payload["message"].ToString();

            if (respSuccess && response.ResponseToken == ResponseToken.Success) {
                Wallet[] wallets = JsonConverter.GetValue<Wallet[]>(response.Payload["wallets"]);

                if (wallets == null)
                    return $"{message}\nNo wallets found.";
                else {
                    foreach (Wallet wallet in wallets)
                        message = @$"{message}\n
                            IBAN: {wallet.IBAN}\n
                            Cash: {wallet.Cash}";
                            
                    return message;
                }
            }

            return message;
        }
    }
}