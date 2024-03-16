
namespace FinanceLib.Models.BankRequests;

public class LoginRequest : BankRequest {
    public LoginRequest(
        string user,
        string pass) : base()
    {
        APIUrl = "users/login";

        Payload = new Dictionary<string, object> {
            ["username"] = user,
            ["password"] = pass,
        };
    }
}