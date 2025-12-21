#include <iostream>

using namespace std;

int main()
{
    // This header tells the browser to expect an HTML page
    cout << "Content-type:text/html\r\n\r\n";

    cout << "<html>\n";
    cout << "<head><title>C++ Powered Page</title></head>\n";
    cout << "<body style='background-color: #2c3e50; color: white; font-family: sans-serif; text-align: center;'>\n";

    cout << "<h1>Hello from C++!</h1>\n";
    cout << "<p>This page was generated dynamically using C++ logic.</p>\n";

    // Simple logic mix
    int day_goal = 10;
    int current_day = 1;

    cout << "<div style='border: 2px solid #27ae60; padding: 20px; display: inline-block;'>";
    cout << "<h3>Challenge Progress</h3>";
    cout << "<p>Languages Finished: " << current_day << " / " << day_goal << "</p>";
    cout << "</div>\n";

    cout << "</body>\n";
    cout << "</html>\n";

    return 0;
}