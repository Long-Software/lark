@echo off
call "C:\Program Files\Microsoft Visual Studio\2022\Community\Common7\Tools\VsDevCmd.bat" -arch=x64

REM Make sure the output folder exists
if not exist ..\..\bin mkdir ..\..\bin

REM Compile with correct output paths
cl main.cpp /O2 /MD /DNDEBUG /GL /Gy /Fo:..\..\bin\ /Fe:..\..\bin\timer.exe /link /LTCG /OPT:REF /OPT:ICF
