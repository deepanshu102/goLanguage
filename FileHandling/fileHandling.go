package main
import(
        "fmt"
        "log"
        "os"
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
        fmt.Println("Gello")

}

func main(){
      //  createEmptyFile()
        // truncateFile()
      recover  getFileInfo()
       recover renameFile()
}