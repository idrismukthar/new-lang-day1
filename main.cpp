#include <iostream>
#include <string>

int main()
{
    std::string username;
    std::string password;
    std::string secret_data = "The treasure is buried under the old oak tree.";

    std::cout << "--- SYSTEM LOGIN ---" << std::endl;

    std::cout << "Enter Username: ";
    std::cin >> username;

    std::cout << "Enter Password: ";
    std::cin >> password;

    // Simple authentication logic
    if (username == "admin" && password == "1234")
    {
        std::cout << "\nLOGIN SUCCESSFUL!" << std::endl;
        std::cout << "Secret Note: " << secret_data << std::endl;
    }
    else
    {
        std::cout << "\nACCESS DENIED. Incorrect credentials." << std::endl;
    }

    return 0;
}