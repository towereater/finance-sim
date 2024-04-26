using System.Runtime.CompilerServices;
using System.Text.Json;
using FinanceLib.Models;
using FinanceLib.Models.BankRequests;
using FinanceLib.Utils;

namespace FinanceLib.Managers;

public class WalletManager
{
    // Net manager in charge of making requests to the server
    private NewNetManager newNetManager = new("http://127.0.0.1:13000/");

    // Verifies whether or not a login was made by checking the NetManager access token
    public bool IsLogged => !string.IsNullOrEmpty(AuthorizationToken);

    public string AuthorizationToken { get; set; }

    // Wallet connected to the manager
    public Wallet[] Wallets { get; protected set; }

    // Create wallet procedure
    public string CreateWallet()
    {
        // Sets up request
        BankRequest request = new CreateWalletRequest
        {
            AuthorizationToken = AuthorizationToken,
        };

        // Sends the request to the server and gets the response
        BankRequest response = newNetManager.Request(HttpMethod.Post, request);

        if (response.ResponseToken == ResponseToken.Success)
        {
            return $"Wallet created";
        }
        else
        {
            return "Wallet creation failed";
        }
    }

    // Get wallets procedure
    public string GetWallets()
    {
        // Sets up request
        BankRequest request = new GetWalletsRequest()
        {
            AuthorizationToken = AuthorizationToken,
        };

        // Sends the request to the server and gets the response
        BankRequest response = newNetManager.Request(HttpMethod.Get, request);

        if (response.ResponseToken == ResponseToken.Success)
        {
            if (response.Payload == null || !response.Payload.TryGetValue("accounts", out _))
            {
                return "User has no wallets";
            }

            string output = string.Empty;

            Wallets = JsonSerializer.Deserialize<Wallet[]>(response.Payload["accounts"].ToString(),
                new JsonSerializerOptions {
                    AllowTrailingCommas = true,
                    PropertyNameCaseInsensitive = true,
            });

            Console.WriteLine(Wallets.Length);
            
            for (int i = 0; i < Wallets.Length; i++){
                if (Wallets[i] == null){
                    continue;
                }

                Wallet wallet = GetWallet(Wallets[i].Id);
                Wallets[i] = wallet;

                output += string.Format("Wallet IBAN: {0}, Cash: {1}\n", wallet.IBAN, wallet.Cash);
            }

            return output;
        }
        else
        {
            return "Couldn't get user wallets";
        }
    }

    // Get wallet procedure
    public Wallet GetWallet(string walletId)
    {
        // Sets up request
        BankRequest request = new GetWalletRequest(walletId)
        {
            AuthorizationToken = AuthorizationToken,            
        };

        // Sends the request to the server and gets the response
        BankRequest response = newNetManager.Request(HttpMethod.Get, request);

        if (response.ResponseToken == ResponseToken.Success)
        {
            return new Wallet(){
                Id = response.Payload["id"].ToString(),
                IBAN = response.Payload["iban"].ToString(),
                Owner = response.Payload["owner"].ToString(),
                Cash = double.Parse(response.Payload["cash"].ToString()),
            };
        }
        else
        {
            return null;
        }
    }

    // public string CreateWallet()
    // {
    //     // Sets up the request
    //     BankRequest request = new BankRequest() {
    //         RequestToken = RequestToken.CreateWallet,
    //         AuthorizationToken = AuthorizationToken,
    //         Payload = new Dictionary<string, object>(),
    //     };

    //     // Sends the request to the server
    //     BankRequest response;
    //     bool respSuccess = netMan.Request(request, out response);
    //     string message = response.Payload["message"].ToString();

    //     if (respSuccess && response.ResponseToken == ResponseToken.Success) {
    //         Wallet[] wallets = JsonConverter.GetValue<Wallet[]>(response.Payload["wallets"]);
    //         return @$"
    //             {message}\n
    //             Wallet created:\n
    //             IBAN: {wallets[0].IBAN}";
    //     }

    //     return message;
    // }

    // // Retrieves the data of a wallet of the user with the given iban.
    // // If iban equals *, then all user's wallets will be given.
    // //TODO: FIX DESCRIPTION
    // public string GetWallets()
    // {
    //     // Sets up the request
    //     BankRequest request = new BankRequest() {
    //         RequestToken = RequestToken.GetWallet,
    //         AuthorizationToken = AuthorizationToken,
    //         Payload = new Dictionary<string, object>(),
    //     };

    //     // Sends the request to the server
    //     BankRequest response;
    //     bool respSuccess = netMan.Request(request, out response);
    //     string message = response.Payload["message"].ToString();

    //     if (respSuccess && response.ResponseToken == ResponseToken.Success) {
    //         Wallet[] wallets = JsonConverter.GetValue<Wallet[]>(response.Payload["wallets"]);

    //         if (wallets == null)
    //             return $"{message}\nNo wallets found.";
    //         else {
    //             foreach (Wallet wallet in wallets)
    //                 message = @$"{message}\n
    //                     IBAN: {wallet.IBAN}\n
    //                     Cash: {wallet.Cash}";
                        
    //             return message;
    //         }
    //     }

    //     return message;
    // }
}