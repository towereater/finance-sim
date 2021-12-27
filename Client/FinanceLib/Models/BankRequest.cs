
namespace FinanceLib.Models
{
    // Possible requests made to the server
    public enum RequestToken {
        LogIn = 0,
        Wallets = 1
    }

    // Response results of a given request
    public enum ResponseToken {
        Failure = 0,
        Success = 1
    }

    // Generic model for a given request (or response) for the bank server
    public class BankRequest
    {
        public RequestToken RequestToken { get; set; }
        public ResponseToken ResponseToken { get; set; }
        public string AccessToken { get; set; }
        public Dictionary<string, object> Payload { get; set; }
    }
}