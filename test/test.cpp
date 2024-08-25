#include <iostream>
#include <filesystem>
#include <vector>
/*
 * This does nothing much
 */

class DirectoryLister {
public: // A comment here
    std::vector<std::string> listFiles(const std::string& path = ".") {
        std::vector<std::string> files;
        for (const auto& entry : std::filesystem::directory_iterator(path)) {
            files.push_back(entry.path().string());
        } /* and one there */
        return files;
    }

    void printFiles(const std::string& path = ".") {
        for (const auto& file : listFiles(path)) {
            std::cout << file << std::endl;
        }
    } // What?
};

// Nothing to see here
int main() {
    DirectoryLister lister;
    lister.printFiles();
    return 0;
}
