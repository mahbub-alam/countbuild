@echo off

%~dp0CheckBuildConf.exe %~dp3build-conf.txt

if %errorlevel% == 1 goto _DEBUG
if %errorlevel% == 2 goto _RELEASE
if %errorlevel% == 0 goto _COUNTBUILD
goto _END

:_DEBUG
if "%1" == "Debug" goto _COUNTBUILD
if "%1" == "Release" goto _END
goto _END

:_RELEASE
if "%1" == "Release" goto _COUNTBUILD
if "%1" == "Debug" goto _END
goto _END

:_COUNTBUILD
%~dp0CountBuild.exe %2 %3 %4 %5 %6

:_END
REM %~dp0strwrite -overwrite %~dp3build-conf.txt %1
echo %1 > build-conf.txt
echo Current Build Configuration: ^<%1^>
echo ...Done
