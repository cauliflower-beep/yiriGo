@echo off
echo Compiling C# code...
csc /target:library /out:Calculate.dll Calculate.cs

echo Compiling C wrapper code...
gcc -shared -o CalculateWrapper.dll CalculateWrapper.c -Wl,--subsystem,windows

echo Build complete!