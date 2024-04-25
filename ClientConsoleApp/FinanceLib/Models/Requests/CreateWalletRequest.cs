
namespace FinanceLib.Models.BankRequests;

public class CreateWalletRequest : BankRequest {
    public CreateWalletRequest() : base()
    {
        APIUrl = "users/accounts";

        Payload = new Dictionary<string, object>();
    }
}