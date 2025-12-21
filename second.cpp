#include <iostream>
#include <string>

using namespace std;

int main()
{
    int choice;
    cout << "--- THE MYSTIC CAVE ---" << endl;
    cout << "You stand at the entrance of a dark cave. Do you:" << endl;
    cout << "1. Enter the cave" << endl;
    cout << "2. Run away like a coward" << endl;
    cout << "Choice: ";
    cin >> choice;

    if (choice == 1)
    {
        cout << "\nIt's dark. You find a torch! Do you light it? (1: Yes / 2: No): ";
        cin >> choice;
        if (choice == 1)
        {
            cout << "The light reveals a chest of gold! YOU WIN!" << endl;
        }
        else
        {
            cout << "You tripped over a rock in the dark. GAME OVER." << endl;
        }
    }
    else
    {
        cout << "\nYou went home and took a nap. Boring, but safe." << endl;
    }

    return 0;
}