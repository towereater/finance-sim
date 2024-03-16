using FinanceLib.Models;
using FinanceLib.Models.BankRequests;

namespace FinanceLib.Managers;

public class AccountManager
{
    // Net manager in charge of making requests to the server
    private NewNetManager newNetManager = new("http://127.0.0.1:13000/");

    // Verifies whether or not a login was made by checking the NetManager access token
    public bool IsLogged { get; protected set; }

    // User account connected to the manager
    public UserAccount UserAccount { get; protected set; }

    // Account creation procedure
    public string CreateAccount(string user, string pass, string name, string surname, string birth)
    {
        // Sets up request
        BankRequest request = new RegisterRequest(user, pass, name, surname, birth);

        // Sends the request to the server and gets the response
        BankRequest response = newNetManager.Request(HttpMethod.Post, request);
        return response.ResponseToken == ResponseToken.Success ?
            "Registration complete" : "Registration failed";
    }

    // Login procedure
    public string LogIn(string user, string pass)
    {
        // Sets up request
        BankRequest request = new LoginRequest(user, pass);

        // Sends the request to the server and gets the response
        BankRequest response = newNetManager.Request(HttpMethod.Post, request);

        if (response.ResponseToken == ResponseToken.Success)
        {
            UserAccount = new() {
                Name = response.Payload["name"].ToString(),
                Surname = response.Payload["surname"].ToString(),
                Birth = response.Payload["birth"].ToString(),
            };

            IsLogged = true;

            return $"Login successful";
        }
        else
        {
            return "Login failed";
        }
    }

    // Logout procedure
    public string LogOut()
    {
        // Drops all session data
        IsLogged = false;
        UserAccount = null;

        return "Logout successful";
    }
}