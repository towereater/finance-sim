using System.Net;
using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Net.Sockets;
using System.Text;
using System.Text.Json;

using FinanceLib.Models.BankRequests;

namespace FinanceLib.Managers;

public class NetManager
{
    public bool Request(BankRequest request, out BankRequest response)
    {
        // Input stream data buffer and bytes received
        byte[] inBuffer = new byte[1024];
        int inBytes = 0;

        // Useful in case of error
        bool errorCatched = false;
        string errorMessage = string.Empty;

        try {
            // Sets up the remote end point to match the localhost on port 11000
            //IPHostEntry ipHostInfo = Dns.GetHostEntry(ServerIP);
            //IPAddress ipAddress = ipHostInfo.AddressList[0];
            IPAddress ipAddress = IPAddress.Parse("127.0.0.1");
            IPEndPoint remoteEP = new(ipAddress, 11000);

            // Creates the TCP/IP socket
            Socket sender = new(ipAddress.AddressFamily,
                SocketType.Stream, ProtocolType.Tcp);
            
            try {
                // Enstablishes the connection with the chosen end point
                sender.Connect(remoteEP);

                // Sets up the message and sends it to the remote host
                string jsonRequest = JsonSerializer.Serialize(request,
                    new JsonSerializerOptions() {
                        PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
                });
                byte[] outBuffer = Encoding.ASCII.GetBytes(jsonRequest);
                int outBytes = sender.Send(outBuffer);

                // Receives the response
                inBytes = sender.Receive(inBuffer);
                
                // Releases the socket
                sender.Shutdown(SocketShutdown.Both);
                sender.Close();

                // Converts the data and sets up the response
                string jsonResponse = Encoding.ASCII.GetString(inBuffer, 0, inBytes);

                response = JsonSerializer.Deserialize<BankRequest>(jsonResponse,
                    new JsonSerializerOptions() {
                        AllowTrailingCommas = true,
                        PropertyNameCaseInsensitive = true,
                });

                // Successful response
                return true;
            }
            catch (ArgumentNullException e) {
                errorCatched = true;
                errorMessage = $"ArgumentNullException: {e}";
            }
            catch (SocketException e) {
                errorCatched = true;
                errorMessage = $"SocketException: {e}";
            }
            catch (Exception e) {
                errorCatched = true;
                errorMessage = $"Unexpected exception: {e}";
            }
        }
        catch (Exception e) {
            errorCatched = true;
            errorMessage = $"Unexpected exception: {e}";
        }

        // In case of error returns its message
        if (errorCatched) {
            //Console.WriteLine($"ERROR CATCHED: {errorMessage}");

            response = new BankRequest() {
                RequestToken = request.RequestToken,
                ResponseToken = ResponseToken.Failure,
                Payload = new Dictionary<string, object>() {
                    ["message"] = errorMessage,
                },
            };

            return false;
        }

        // Default return
        response = new BankRequest();
        return false;
    }
}

public class NewNetManager
{
    private HttpClient client;

    public NewNetManager(string baseAddress)
    {
        client = new(){
            BaseAddress = new Uri(baseAddress),
            Timeout = TimeSpan.FromSeconds(5),
        };

        client.DefaultRequestHeaders.Clear();
        client.DefaultRequestHeaders.Accept.Add(
            new MediaTypeWithQualityHeaderValue("application/json"));
    }

    public BankRequest Request(HttpMethod httpMethod, BankRequest request)
    {
        // Inclusion of the authentication token
        if (request.AuthorizationToken != null)
        {
            client.DefaultRequestHeaders.Add("Authorization", request.AuthorizationToken);
        }

        // Execution of the request
        HttpResponseMessage httpResponse;
        if (httpMethod == HttpMethod.Post)
        {
            httpResponse = client.PostAsJsonAsync(
                request.APIUrl,
                request.Payload,
                new JsonSerializerOptions{
                    PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
                }
            ).Result;
        }
        else
        {
            httpResponse = null;
        }

        // Parsing of the response
        string jsonResponse = httpResponse.Content.ReadAsStringAsync().Result;

        // Creation of the response object
        BankRequest response = new();
        
        if (httpResponse.IsSuccessStatusCode)
        {
            response.ResponseToken = ResponseToken.Success;

            response.Payload = JsonSerializer.Deserialize<Dictionary<string, object>>(jsonResponse,
                new JsonSerializerOptions {
                    AllowTrailingCommas = true,
                    PropertyNameCaseInsensitive = true,
            });

            if (httpResponse.Headers.Contains("Jwt"))
            {
                response.AuthorizationToken = httpResponse.Headers.GetValues("Jwt").ElementAt(0);
            }
        }
        else
        {
            response.ResponseToken = ResponseToken.Failure;
        }

        return response;
    }
}