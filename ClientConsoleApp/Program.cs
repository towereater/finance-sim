
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Text.Json;

namespace ClientConsoleApp
{
    public class Program
    {
        static HttpClient client;

        static void Main()
        {
            BankApp app = new();
            app.Run();

            //TestHttp();
        }

        static void TestHttp()
        {
            client = new()
            {
                BaseAddress = new Uri("http://127.0.0.1:13000/"),
                Timeout = TimeSpan.FromSeconds(5),
            };

            client.DefaultRequestHeaders.Clear();
            client.DefaultRequestHeaders.Accept.Add(
                new MediaTypeWithQualityHeaderValue("application/json"));

            var payload = new Dictionary<string, object> {
                {"username", "andnic"},
                {"password", "password"},
                {"name", "And"},
                {"surname", "Nic"},
                {"birth", "1996-03-21"},
            };

            HttpResponseMessage response = client.PostAsJsonAsync("users/register", payload, new JsonSerializerOptions{
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
            }).Result;

            Console.WriteLine(response.StatusCode);
            Console.WriteLine(response.Content.ReadAsStringAsync().Result);
        }
    }
}