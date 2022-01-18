using System.Text.Json;

namespace FinanceLib.Utils
{
    public static class JsonConverter
    {
        public static T GetValue<T>(object node)
        {
            return JsonSerializer.Deserialize<T>((JsonElement) node,
                new JsonSerializerOptions() {
                    AllowTrailingCommas = true,
                    PropertyNameCaseInsensitive = true,
                });
        }
    }
}