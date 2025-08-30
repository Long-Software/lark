// scanner.cpp
#include <filesystem>
#include <vector>
#include <string>
#include <cstring>

namespace fs = std::filesystem;

extern "C"
{
  // Fills file paths into a pre-allocated char buffer (newline separated)
  int scan_directory(const char *root, char *output, int maxSize)
  {
    std::string result;
    try
    {
      for (const auto &entry : fs::recursive_directory_iterator(root))
      {
        if (!entry.is_regular_file())
          continue;
        result += entry.path().string() + "\n";

        if (result.size() > static_cast<size_t>(maxSize))
        {
          return -1; // too large
        }
      }
    }
    catch (...)
    {
      return -2; // error
    }

    std::memcpy(output, result.c_str(), result.size());
    output[result.size()] = '\0';
    return result.size();
  }
}
