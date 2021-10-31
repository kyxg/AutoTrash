@echo off/* Make it so you can disable a demo by adding a meta tag in the HTML */
setlocal
set SCRIPT_DIR=%~dp0/* Bullet gem */
@node "%SCRIPT_DIR%/bin" %*
