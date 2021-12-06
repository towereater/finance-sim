using System.Net;
using System.Net.Sockets;
using System.Text;
using System.Text.Json;

namespace FinanceLib.Managers
{
    public class NetManager
    {
        public enum RequestToken {
            LogIn = 0
        }

        // Response results of a given request
        public enum ResponseToken {
            Failure = 0,
            Success = 1
        }

        /*
        // IP of the server to contact when making a request
        private string serverIp;
        public string ServerIP
        {
            get { return serverIp; }
            set { serverIp = value; }
        }
        
        // Default constructor connects to localhost when making requests
        public NetManager()
            : this("localhost") {}

        // Costructor with a custom server IP address
        public NetManager(string serverIp)
        {
            ServerIP = serverIp;
        }
        */

        public static string GetUserAccount(string jsonString)
        {
            return null;
        }

        // Default method used to make requests to the server
        public static string ServerRequest(string request)
        {
            // Input stream data buffer and bytes received
            byte[] inBuffer = new byte[1024];
            int inBytes = 0;

            // Response with validity token
            string response = string.Empty;
            ResponseToken token = ResponseToken.Failure;

            try {
                // Sets up the remote end point to match the localhost on port 11000
                IPHostEntry ipHostInfo = Dns.GetHostEntry("localhost");//Dns.GetHostName()
                IPAddress ipAddress = ipHostInfo.AddressList[0];
                IPEndPoint remoteEP = new IPEndPoint(ipAddress, 11000);

                // Creates the TCP/IP socket
                Socket sender = new Socket(ipAddress.AddressFamily,
                    SocketType.Stream, ProtocolType.Tcp);
                
                try {
                    // Enstablishes the connection with the chosen end point
                    sender.Connect(remoteEP);

                    // Sets up the message and sends it to the remote host
                    byte[] outBuffer = Encoding.ASCII.GetBytes(request);
                    int outBytes = sender.Send(outBuffer);

                    // Receives the response
                    inBytes = sender.Receive(inBuffer);
                    
                    // Releases the socket
                    sender.Shutdown(SocketShutdown.Both);
                    sender.Close();

                    // Converts the data and sets up the response
                    response = Encoding.ASCII.GetString(inBuffer, 0, inBytes);
                    token = ResponseToken.Success;
                }
                catch (ArgumentNullException e) {
                    response = "ArgumentNullException: " + e.ToString();
                }
                catch (SocketException e) {
                    response = "SocketException: " + e.ToString();
                }
                catch (Exception e) {
                    response = "Unexpected exception: " + e.ToString();
                }
            }
            catch (Exception e) {
                response = e.ToString();
            }

            // Generates the JSON file and returns it
            var jsonResponse = new {
                token = token,
                payload = response
            };

            return JsonSerializer.Serialize(jsonResponse);
        }  
    }
}