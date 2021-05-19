FROM python:3.8-windowsservercore-1809

RUN New-Item -Path 'C:\app' -Type Directory
WORKDIR C:/app
COPY main.py C:/app
