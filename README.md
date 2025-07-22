#  Important Notice: Rebuild Required for Windows 64-bit

This repository includes a prebuilt `veko-grid.exe` binary.  
However, this binary is built in **16-bit format**, which is **not compatible** with most modern 64-bit Windows systems.

---

##  Why Won’t It Run?

- 16-bit executables are no longer supported by Windows 7/10/11 64-bit editions.
- You will likely see errors such as:
  > "This app can’t run on your PC"  
  > or "The subsystem needed to support the image type is not present"

---

##  How to Rebuild the Tool for 64-bit Systems

To use the tool on a modern Windows PC, you **must rebuild the executable** in 64-bit mode.

###  Requirements:
- Go Programming Language: [https://go.dev/dl/](https://go.dev/dl/)
- Terminal / PowerShell
- Windows 64-bit OS

###  Run this command:

```bash
go build -o veko-grid.exe main.go


Or if you're building from Linux/macOS targeting Windows:
GOOS=windows GOARCH=amd64 go build -o veko-grid.exe main.go

This will generate a 64-bit version that works properly.



 Notes
Only the source code and 16-bit .exe are included in this repository.

The prebuilt .exe is for archival/testing purposes only and not production-ready.

To avoid confusion or execution errors, we recommend you always build from source.

 License & Ethical Usage
This tool is intended for security research and educational use only.
Please do not use it on systems or networks you do not own or have explicit permission to test.
