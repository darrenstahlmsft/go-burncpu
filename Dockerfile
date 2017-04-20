FROM microsoft/nanoserver

ADD go-burncpu.exe .
CMD go-burncpu.exe
