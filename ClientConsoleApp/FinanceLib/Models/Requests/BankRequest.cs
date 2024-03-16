
namespace FinanceLib.Models.BankRequests;

// Possible requests made to the server
public enum RequestToken {
    CreateAccount = 0,
    LogIn = 1,
    CreateWallet = 21,
    GetWallet = 22,
    DeleteWallet = 23,
}

// Response results of a given request
public enum ResponseToken {
    Failure = 0,
    Success = 1
}

// Generic model for a given request (or response) for the bank server
public class BankRequest
{
    public string APIUrl { get; set; }
    public RequestToken RequestToken { get; set; }
    public ResponseToken ResponseToken { get; set; }
    public string AuthorizationToken { get; set; }
    public Dictionary<string, object> Payload { get; set; }
}