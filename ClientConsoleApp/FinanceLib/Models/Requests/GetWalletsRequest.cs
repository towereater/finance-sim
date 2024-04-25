
namespace FinanceLib.Models.BankRequests;

public class GetWalletsRequest : BankRequest {
    public GetWalletsRequest() : base()
    {
        APIUrl = "users/accounts";

        Payload = new Dictionary<string, object>();
    }
}