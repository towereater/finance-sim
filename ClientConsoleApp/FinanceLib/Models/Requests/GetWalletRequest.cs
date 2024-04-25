
namespace FinanceLib.Models.BankRequests;

public class GetWalletRequest : BankRequest {
    public GetWalletRequest(
        string walletId) : base()
    {
        APIUrl = "users/accounts/" + walletId;

        Payload = new Dictionary<string, object>();
    }
}