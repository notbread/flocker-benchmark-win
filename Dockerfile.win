FROM golang:1.16.4

RUN New-Item -Path 'C:\app' -Type Directory
WORKDIR C:/app
COPY ./app.exe C:/app/

ENTRYPOINT ["C:/app/app.exe"]
