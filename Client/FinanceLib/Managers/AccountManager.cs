using System.Text.Json;

namespace FinanceLib.Managers
{
    public class AccountManager
    {
        public AccountManager() {}

        public bool LogIn(string user, string password)
        {
            var request = new {
                token = NetManager.RequestToken.LogIn,
                payload = new {
                    user = user,
                    password = password
                }
            };
            NetManager.ServerRequest(JsonSerializer.Serialize(request));

            return false;
        }
    }
}