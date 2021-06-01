package main
import(
        "fmt"
        "log"
        "os"
        "io" // use in copyFuncation
        "time"
       "io/ioutil"
       "bufio"
)
var (
newFile *os.File
fileInfo os.FileInfo
err error
)


func createEmptyFile(){
        newFile,err =os.Create("test.txt")
        if err!=nil{
                log.Fatal(err)
        }
        log.Println(newFile)
        newFile.Close()
}

func truncateFile(){
         // Truncate a file to 100 bytes. If file
    // is less than 100 bytes the original contents will remain
    // at the beginning, and the rest of the space is
    // filled will null bytes. If it is over 100 bytes,
    // Everything past 100 bytes will be lost. Either way
    // we will end up with exactly 100 bytes.
    // Pass in 0 to truncate to a completely empty file

    err := os.Truncate("test.txt", 100)
    if err != nil {
        log.Fatal(err)
    }
}

func getFileInfo(){
        //Stat method returns file info. It will
        // return an error if there is no file
        fileInfo , err = os.Stat("test.txt")
        if err!=nil{
                log.Fatal(err)
        }
        fmt.Println("File Name:-",fileInfo.Name())
        fmt.Println("File Size in bytes:-",fileInfo.Size())
        fmt.Println("Permissions ",fileInfo.Mode())
        fmt.Println("Last Modifed",fileInfo.ModTime())
        fmt.Println("Is Directory",fileInfo.IsDir())
        fmt.Printf("System interface type:-%T\n",fileInfo.Sys())
        fmt.Printf("system info : %+v\n\n",fileInfo.Sys())
}

func renameFile(){
        originalpath:="test.txt"
        newpath:="test2.txt"
        err:=os.Rename(originalpath,newpath)
        if err != nil{
                log.Fatal(err)
        }

}
func deleteFile() {
        err := os.Remove("test2.txt")
        if err != nil {
            log.Fatal(err)
        }
    }


    func openAndClose() {
        // Simple read only open. We will cover actually reading
        // and writing to files in examples further down the page
        file, err := os.Open("test.txt")
        if err != nil {
            log.Fatal(err)
        }
        file.Close()
    
        // OpenFile with more options. Last param is the permission mode
        // Second param is the attributes when opening
        file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
        if err != nil {
            log.Fatal(err)
        }
        file.Close()
    
        // Use these attributes individually or combined
        // with an OR for second arg of OpenFile()
        // e.g. os.O_CREATE|os.O_APPEND
        // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
    
        // os.O_RDONLY // Read only
        // os.O_WRONLY // Write only
        // os.O_RDWR // Read and write
        // os.O_APPEND // Append to end of file
        // os.O_CREATE // Create is none exist
        // os.O_TRUNC // Truncate file when opening
    }


    func CheckFileExists() {
        // Stat returns file info. It will return
        // an error if there is no file.
        fileInfo, err := os.Stat("test.txt")
        if err != nil {
            if os.IsNotExist(err) {
                log.Fatal("File does not exist.")
            }
        }
        log.Println("File does exist. File information:")
        log.Println(fileInfo)
    }


    func checkRWPermission() {
        // Test write permissions. It is possible the file
        // does not exist and that will return a different
        // error that can be checked with os.IsNotExist(err)
        file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
        if err != nil {
            if os.IsPermission(err) {
                log.Println("Error: Write permission denied.")
            }
        }
        file.Close()
    
        // Test read permissions
        file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
        if err != nil {
            if os.IsPermission(err) {
                log.Println("Error: Read permission denied.")
            }
        }
        file.Close()
    }


    func changeOwnerShip() {
        // Change perrmissions using Linux style
        err := os.Chmod("test.txt", 0777)
        if err != nil {
            log.Println(err)
        }
    
        // Change ownership
        err = os.Chown("test.txt", os.Getuid(), os.Getgid())
        if err != nil {
            log.Println(err)
        }
    
        // Change timestamps
        twoDaysFromNow := time.Now().Add(48 * time.Hour)
        lastAccessTime := twoDaysFromNow
        lastModifyTime := twoDaysFromNow
        err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
        if err != nil {
            log.Println(err)
        }
    }


    func copyOneFileToAnother() {
        // Open original file
        originalFile, err := os.Open("test.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer originalFile.Close()
    
        // Create new file
        newFile, err := os.Create("test_copy.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer newFile.Close()
    
        // Copy the bytes to destination from source
        bytesWritten, err := io.Copy(newFile, originalFile)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Copied %d bytes.", bytesWritten)
        
        // Commit the file contents
        // Flushes memory to disk
        err = newFile.Sync()
        if err != nil {
            log.Fatal(err)
        }
    }


    func writeBytes() {
        // Open a new file for writing only
        file, err := os.OpenFile(
            "test.txt",
            os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
            0666,
        )
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    
        // Write bytes to file
        byteSlice := []byte("Welcome to IO package in Go !\n")
        bytesWritten, err := file.Write(byteSlice)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Wrote %d bytes.\n", bytesWritten)
    }

// using ioutil package

    func ioutilWritter() {
        err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
        if err != nil {
            log.Fatal(err)
        }
    }

// Using Bufio package
    func bufIoWrting() {
        // Open file for writing
        file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    
        // Create a buffered writer from the file
        bufferedWriter := bufio.NewWriter(file)
    
        // Write bytes to buffer
        bytesWritten, err := bufferedWriter.Write(
            []byte{65, 66, 67},
        )
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Bytes written: %d\n", bytesWritten)
    
        // Write string to buffer
        // Also available are WriteRune() and WriteByte()   
        bytesWritten, err = bufferedWriter.WriteString(
            "\tBuffered string\n",
        )
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Bytes written: %d\n", bytesWritten)
    
        // Check how much is stored in buffer waiting
        unflushedBufferSize := bufferedWriter.Buffered()
        log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
    
        // See how much buffer is available
        bytesAvailable := bufferedWriter.Available()
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Available buffer: %d\n", bytesAvailable)
    
        // Write memory buffer to disk
        bufferedWriter.Flush()
    
        // Revert any changes done to buffer that have
        // not yet been written to file with Flush()
        // We just flushed, so there are no changes to revert
        // The writer that you pass as an argument
        // is where the buffer will output to, if you want
        // to change to a new writer
        bufferedWriter.Reset(bufferedWriter) 
    
        // See how much buffer is available
        bytesAvailable = bufferedWriter.Available()
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Available buffer: %d\n", bytesAvailable)
    
        // Resize buffer. The first argument is a writer
        // where the buffer should output to. In this case
        // we are using the same buffer. If we chose a number
        // that was smaller than the existing buffer, like 10
        // we would not get back a buffer of size 10, we will
        // get back a buffer the size of the original since
        // it was already large enough (default 4096)
        bufferedWriter = bufio.NewWriterSize(
            bufferedWriter,
            8000,
        )
    
        // Check available buffer size after resizing
        bytesAvailable = bufferedWriter.Available()
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Available buffer: %d\n", bytesAvailable)
    }

 func readInBytes() {
        // Open file for reading
        file, err := os.Open("test.txt")
        if err != nil {
            log.Fatal(err)
        }
        defer file.Close()
    
        // Read up to len(b) bytes from the File
        // Zero bytes written means end of file
        // End of file returns error type io.EOF
        byteSlice := make([]byte, 16)
        bytesRead, err := file.Read(byteSlice)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Number of bytes read: %d\n", bytesRead)
        log.Printf("Data read: %s\n", byteSlice)
    }
 func readNBytesFile(){  
        file,err:=os.Open("test.txt")
               if err!=nil{
                   log.Println(err)
               }
               defer file.Close()
               //make a N byte slice 
              // Read up to len(b) bytes from the File
              // Zero bytes written means end of file
              // End of file returns error type io.EOF
               bytesSlice:=make([]byte,16)
               bytesRead,err:=file.Read(bytesSlice)
               if err!=nil{
                   log.Println(err,"Error occur")
               }
               log.Println("No. of bytes read",bytesRead)
               log.Println("Data Read:-",bytesSlice)
    }
    
  func read_ExctalyNBytes_File(){
        file,err:=os.Open("test.txt")
        if err!=nil{
            log.Println("Error:- at Open file ",err)
        }
        // The file.Read() function will happily read a tiny file in to a large
        // byte slice, but io.ReadFull() will return an
        // error if the file is smaller than the byte slice.
        bytesSlice:=make([]byte,2)
        numberBytesRead,err:=io.ReadFull(file,bytesSlice)
        if err!=nil{
            log.Println("File content less than byteBuffer",err)
        }
        log.Println("No of bytes read:-",numberBytesRead)
        log.Println("Data read:-",bytesSlice)
    
    }
 func atLeastNBytesFileRead(){
        file,err:=os.Open("test.txt")
        if err!=nil{
            log.Println(err)
        }
        byteSlice :=make([]byte,512)
        minBytes:=8
        numberBytesRead,err:=io.ReadAtLeast(file,byteSlice,minBytes)
        if err!=nil{
            log.Println(err)
        }
        log.Println("Number of bytes read:-",numberBytesRead)
        log.Println("DATA---",byteSlice)
    }
func quickRead(){
        // Read file to byte slice
        data, err := ioutil.ReadFile("test.txt")
        if err != nil {
            log.Fatal(err)
        }
    
        log.Printf("Data read: %s\n", data)
    
}
 func main(){
	createEmptyFile()
	truncateFile()
	getFileInfo()
	renameFile()
	deleteFile() 
	createEmptyFile()
	openAndClose()
	CheckFileExists()
	checkRWPermission()
	changeOwnerShip()
	writeBytes()
	copyOneFileToAnother()
	ioutilWritter()
	bufIoWrting()
	readNBytesFile()
	read_ExctalyNBytes_File()
	atLeastNBytesFileRead()
    quickRead()
 } 