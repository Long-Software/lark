#include <iostream>
#include <Windows.h>
#include <chrono>
#include <Psapi.h>

void printUsage()
{
  std::cout << "Usage: time.exe <program> [args...]" << std::endl;
}

std::wstring toWString(const std::string &str);

int main(int argc, char *argv[])
{
  if (argc < 2)
  {
    printUsage();
    return 1;
  }

  std::string cmd;
  for (int i = 1; i < argc; i++)
  {
    cmd += "\"";
    cmd += argv[i];
    cmd += "\"";
  }

  STARTUPINFOW inf = {sizeof(inf)};
  PROCESS_INFORMATION pi;

  std::wstring wCmd = toWString(cmd);
  auto start = std::chrono::high_resolution_clock::now();

  if (!CreateProcessW(
          nullptr,
          &wCmd[0],
          nullptr,
          nullptr,
          FALSE,
          CREATE_NO_WINDOW,
          nullptr,
          nullptr,
          &inf,
          &pi))
  {
    std::cerr << "Faild to start process." << std::endl;
  }

  SIZE_T maxMem = 0;
  SIZE_T minMem = SIZE_MAX;

  PROCESS_MEMORY_COUNTERS pmc;
  if (GetProcessMemoryInfo(pi.hProcess, &pmc, sizeof(pmc)))
  {
    SIZE_T size = pmc.WorkingSetSize;
    if (size > maxMem)
      maxMem = size;
    if (size < minMem)
      minMem = size;
  }

  while (WaitForSingleObject(pi.hProcess, 1) == WAIT_TIMEOUT)
  {
    if (GetProcessMemoryInfo(pi.hProcess, &pmc, sizeof(pmc)))
    {
      SIZE_T size = pmc.WorkingSetSize;
      if (size > maxMem)
        maxMem = size;
      if (size < minMem)
        minMem = size;
    }
  }

  if (GetProcessMemoryInfo(pi.hProcess, &pmc, sizeof(pmc)))
  {
    SIZE_T size = pmc.WorkingSetSize;
    if (size > maxMem)
      maxMem = size;
    if (size < minMem)
      minMem = size;
  }

  auto end = std::chrono::high_resolution_clock::now();
  auto nanos = std::chrono::duration_cast<std::chrono::nanoseconds>(end - start).count();
  double ms = nanos / 1'000'000.0;

  CloseHandle(pi.hProcess);
  CloseHandle(pi.hThread);

  std::cout << "Execution time: " << ms << "ms" << std::endl;
  std::cout << "Memory usage:" << std::endl;
  std::cout << "\tMin: " << minMem / 1024 << "KB" << std::endl;
  std::cout << "\tMax: " << maxMem / 1024 << "KB" << std::endl;

  return 0;
}

std::wstring toWString(const std::string &str)
{
  int len = MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, nullptr, 0);
  std::wstring wstr(len, L'\0');
  MultiByteToWideChar(CP_UTF8, 0, str.c_str(), -1, &wstr[0], len);
  return wstr;
}