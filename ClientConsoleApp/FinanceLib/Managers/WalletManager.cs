using FinanceLib.Models;
using FinanceLib.Models.BankRequests;
using FinanceLib.Utils;

namespace FinanceLib.Managers;

public class WalletManager
{
    // Net manager in charge of making requests to the server
    private NetManager netMan = new NetManager();

    // Verifies whether or not a login was made by checking the NetManager access token
    public bool IsLogged => !string.IsNullOrEmpty(AuthorizationToken);

    public string AuthorizationToken { get; set; }

    public string CreateWallet()
    {
        // Sets up the request
        BankRequest request = new BankRequest() {
            RequestToken = RequestToken.CreateWallet,
            AuthorizationToken = AuthorizationToken,
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
            RequestToken = RequestToken.GetWallet,
            AuthorizationToken = AuthorizationToken,
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