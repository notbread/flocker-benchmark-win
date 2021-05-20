FROM python:3.8-windowsservercore-1809

RUN New-Item -Path 'C:\app' -Type Directory
WORKDIR C:/app
COPY . C:/app
RUN pip install -r requirements.txt
