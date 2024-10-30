#include <iostream>
#include <string>
#include <vector>
#include <fstream>
#include <filesystem>

using namespace std;
namespace fs = std::filesystem;

void logAction(const string& action) {
    ofstream logFile("log.txt", ios::app);
    if (logFile.is_open()) {
        logFile << action << endl;
        logFile.close();
    }
    else {
        cout << "Unable to open log file." << endl;
    }
}

void listDirectory() {
    for (const auto& entry : fs::directory_iterator(fs::current_path())) {
        cout << entry.path().filename().string() << endl;
    }
}

void changeDirectory(const string& path) {
    if (path == "..") {
        // Переход на уровень выше
        fs::current_path(fs::current_path().parent_path());
        cout << "Changed directory to: " << fs::current_path() << endl;
    }
    else if (fs::exists(path) && fs::is_directory(path)) {
        fs::current_path(path);
        cout << "Changed directory to: " << path << endl;
    }
    else {
        cout << "Directory does not exist." << endl;
    }
}

void diskUsage() {
    auto space = fs::space(fs::current_path());
    cout << "Available: " << space.available / (1024 * 1024) << " MB" << endl;
    cout << "Capacity: " << space.capacity / (1024 * 1024) << " MB" << endl;
    cout << "Free: " << space.free / (1024 * 1024) << " MB" << endl;
}

void reverseString(const string& str) {
    string reversed(str.rbegin(), str.rend());
    cout << "Reversed: " << reversed << endl;
}

void findFile(const string& filename) {
    bool found = false;
    for (const auto& entry : fs::recursive_directory_iterator(fs::current_path())) {
        if (entry.path().filename() == filename) {
            cout << "Found: " << entry.path() << endl; // Печатаем полный путь
            found = true;
        }
    }
    if (!found) {
        cout << "Not found: " << filename << endl;
    }
}

int main() {
    string command;
    cout << "Welcome to the Linux Console Emulator!" << endl;

    while (true) {
        // Выводим текущую директорию
        cout << fs::current_path().string() << "/:> ";
        getline(cin, command);
        logAction(command); // Логируем действие

        if (command == "exit") {
            break;
        }
        else if (command == "ls") {
            listDirectory();
        }
        else if (command.rfind("cd ", 0) == 0) {
            changeDirectory(command.substr(3));
        }
        else if (command == "du") {
            diskUsage();
        }
        else if (command.rfind("rev ", 0) == 0) {
            reverseString(command.substr(4));
        }
        else if (command.rfind("find ", 0) == 0) {
            findFile(command.substr(5));
        }
        else {
            cout << "Command not recognized." << endl;
        }
    }

    return 0;
}
