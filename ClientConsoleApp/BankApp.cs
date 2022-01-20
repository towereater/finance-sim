using FinanceLib.Managers;

namespace ClientConsoleApp
{
    public class BankApp
    {
        public AccountManager AccountManager { get; private set; } = new AccountManager();

        public void Run()
        {
            while (true) {
                int ans;

                PrintGUI();
                if (!AccountManager.IsLogged) {
                    ans = GetMenuOption(2);

                    if (ans == 1) {
                        string user = GetString("Insert username (-q to exit): ");
                        if (user == "-q")
                            continue;

                        string pass = GetString("Insert password (-q to exit): ");
                        if (pass == "-q")
                            continue;

                        string message = AccountManager.LogIn(user, pass);
                        Console.WriteLine($"{message}\n");
                    }
                    else if (ans == 2) {
                        string user = GetString("Insert a new username (-q to exit): ");
                        if (user == "-q")
                            continue;

                        string pass = GetString("Insert a new password (-q to exit): ");
                        if (pass == "-q")
                            continue;

                        string message = AccountManager.CreateAccount(user, pass);
                        Console.WriteLine($"{message}\n");
                    }
                }
                else {
                    ans = GetMenuOption(3);

                    if (ans == 1) {
                        string message = AccountManager.CreateWallet();
                        Console.WriteLine($"{message}\n");
                    }
                    else if (ans == 2) {
                        string message = AccountManager.GetWallets();
                        Console.WriteLine($"{message}\n");
                    }
                    else if (ans == 3) {
                        string message = AccountManager.LogOut();
                        Console.WriteLine($"{message}\n");
                    }
                }

                Console.Write("Press any key");
                Console.ReadKey(true);
            }
        }

        private string GetString(string message)
        {
            string input = string.Empty;

            while (string.IsNullOrEmpty(input)) {
                Console.Write(message);
                input = Console.ReadLine().Trim();

                if (string.IsNullOrEmpty(input))
                    Console.WriteLine("Invalid input, please insert a new one.\n");
            }

            return input;
        }

        private int GetMenuOption(int maxValue)
        {
            int ans;

            do {
                Console.Write("Insert an option: ");
                
                // The answer should be chosen between 1 and maxValue included
                if (!int.TryParse(Console.ReadLine(), out ans) || ans < 1 || ans > maxValue)
                    Console.WriteLine("Invalid option, please insert a new one.\n");
                else
                    break;
            } while (true);

            return ans;
        }

        private void PrintGUI()
        {
            if (!AccountManager.IsLogged) {
                Console.Clear();
                Console.WriteLine("Welcome to the BANK NAME app\n");
                Console.WriteLine("You need to login first:");
                Console.WriteLine("1. Login");
                Console.WriteLine("2. Create new account");
            }
            else {
                Console.Clear();
                Console.WriteLine($"Welcome back, {AccountManager.UserAccount.Username}\n");
                Console.WriteLine("What do you want to do?");
                Console.WriteLine("1. Create a new wallet");
                Console.WriteLine("2. Get all wallets");
                Console.WriteLine("3. Logout");
            }
        }
    }
}