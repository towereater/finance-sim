using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Text.Json;

using FinanceLib.Models;

namespace FinanceLib.Managers
{
    public class NetManager
    {
        // IP of the server to contact when making a request
        public static string ServerIP = "127.0.0.1";

        // Validation token used to get informations from server
        public string AccessToken { get; private set; }

        // Costructor with a custom server IP address
        private NetManager(string accessToken)
        {
            AccessToken = accessToken;
        }

        public static NetManager Initialize(string user, string pass)
        {
            // Sets up the request
            BankRequest request = new BankRequest() {
                RequestToken = RequestToken.LogIn,
                Payload = new Dictionary<string, object>() {
                    ["user"] = user,
                    ["password"] = pass,
                },
            };

            // Sends the request to the server and gets the response
            BankRequest response;
            bool respSuccess = NetManager.ServerRequest(request, out response);

            // If both the request was correctly delivered and the login successful,
            // then it creates a new NetManager with the given token
            if (respSuccess && response.ResponseToken == ResponseToken.Success)
                return new NetManager(response.Payload["accessToken"].ToString());
            
            return null;
        }

        public bool Request(BankRequest request, out BankRequest response)
        {
            // Adds the access token to the request
            request.AccessToken = AccessToken;

            // Sends the request using the static version of the function
            return NetManager.ServerRequest(request, out response);
        }

        // Default method used to send requests to the server
        private static bool ServerRequest(BankRequest request, out BankRequest response)
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
                IPEndPoint remoteEP = new IPEndPoint(ipAddress, 11000);

                // Creates the TCP/IP socket
                Socket sender = new Socket(ipAddress.AddressFamily,
                    SocketType.Stream, ProtocolType.Tcp);
                
                try {
                    // Enstablishes the connection with the chosen end point
                    sender.Connect(remoteEP);

                    // Sets up the message and sends it to the remote host
                    string jsonRequest = JsonSerializer.Serialize(request,
                        new JsonSerializerOptions() {
                            AllowTrailingCommas = true,
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

                    Console.WriteLine(jsonResponse);

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
                    errorMessage = $"ArgumentNullException: {e.ToString()}";
                }
                catch (SocketException e) {
                    errorCatched = true;
                    errorMessage = $"SocketException: {e.ToString()}";
                }
                catch (Exception e) {
                    errorCatched = true;
                    errorMessage = $"Unexpected exception: {e.ToString()}";
                }
            }
            catch (Exception e) {
                errorCatched = true;
                errorMessage = $"Unexpected exception: {e.ToString()}";
            }

            // In case of error returns its message
            if (errorCatched) {
                Console.WriteLine($"ERROR CATCHED: {errorMessage}");

                response = new BankRequest() {
                    RequestToken = request.RequestToken,
                    ResponseToken = ResponseToken.Failure,
                    Payload = new Dictionary<string, object>() {
                        ["errorMessage"] = errorMessage,
                    },
                };

                return false;
            }

            // Default return
            response = new BankRequest();
            return false;
        }
    }
}