
namespace FinanceLib.Models.BankRequests;

public class RegisterRequest : BankRequest {
    public RegisterRequest(
        string user,
        string pass,
        string name,
        string surname,
        string birth) : base()
    {
        APIUrl = "users/register";

        Payload = new Dictionary<string, object> {
            ["username"] = user,
            ["password"] = pass,
            ["name"] = name,
            ["surname"] = surname,
            ["birth"] = birth,
        };
    }
}