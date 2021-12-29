using FinanceLib.Managers;
using FinanceLib.Models;

namespace Client
{
    class Program
    {
        static void Main()
        {
            string user, pass;
            bool logSuccess = false;
            AccountManager account = new AccountManager();
            
            Console.WriteLine("Login to BANK NAME app\n");
            while (!logSuccess) {
                // Log in data input
                Console.Write("Insert the username: ");
                user = Console.ReadLine();
                Console.Write("Insert the password: ");
                pass = Console.ReadLine();
                
                // Tries to login to the server
                Console.WriteLine("\nLogging in...");
                logSuccess = account.LogIn(user, pass);

                // Manages the login response
                if (logSuccess) {
                    Console.WriteLine("Login complete!");

                    Wallet[] wallets = account.UserAccount.Wallets;
                    if (wallets != null) {
                        Console.WriteLine("\nWallet data retrieved");

                        foreach (Wallet w in wallets)
                            Console.WriteLine($"IBAN: {w.IBAN}\nCash: {w.Cash}");
                        
                        Console.WriteLine("Insert a IBAN code to retrieve info about");
                        wallets = account.GetWallets(Console.ReadLine());
                        
                        foreach (Wallet w in wallets)
                            Console.WriteLine($"IBAN: {w.IBAN}\nCash: {w.Cash}");
                    }
                    else {
                        Console.WriteLine("Error while downloading wallet data!");
                    }
                }
                else {
                    Console.WriteLine("Username and password combination is wrong!\n");
                }
            }
        }
    }
}