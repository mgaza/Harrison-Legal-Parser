# Harrison Legal Parser

## Description

Harrison Legal Parser is a GoLang program for analyzing/engineering
data received from Harrison Texas County.

> Note: index flag is still in development and currently not operable

## Prerequisites

    - GoLang
    
**Make sure all files are in the _same_ directory, not subdirectory**
    
    - Preformatted County Csv File
        - This has to be acquired from an export of the county's data
          living on the TexasFile database.
        - May sometimes already include the data needing to be parsed.
      
    - Optional: Original County Index Files
        - When needing to parse county legal info, you'll most likely needing
          the original data files to read from unless the Preformatted County Csv File
          contains info readable enough to parse fully.
        
## Usage

```golang

// Always assume that the program executable hasn't been built yet.
// Especially on your first use.
// Also important to do this anytime you change code and need to test it.
// This will create/recreate an executable file that will run the program
C:\Users\user\sourceCodePath> go build

// Call ".\Harrison-Legal-Parser.exe" with at least the "source" flag in order to run program.
// The flags "index" and "remarkRead" are only needed if the county's parsing data needs to be read from their original files.
// The end result will create a new folder called "output" in your "source" directory and store the newly written files with parsed legal info there.
// 
// Flags and usage:
//  - "source": 
//      Provide the file path to wherever you're holding the TexasFile Exported/Preformatted Csv files.
//  - "index":
//      Provide the file path to wherever you're holding the original index files given to us by the counties.
//      These files will only be read and processed if the "remarkRead" flag is set to false.
//  - "remarkRead":
//      This tells the program whether or not the legals needing to be parsed are in the "source" path or the "index" path
//      When set to "true" it will parse from "source" files.
//      If no flag is given in command line, then "remarkRead" will automatically be set to "true"
C:\Users\user\sourceCodePath> .\Harrison-Legal-Parser.exe -source="C:\Path\To\TexasFileExportCsvFiles" -index="C:\Path\To\OriginalCountyIndexFiles" -remarkRead=true
```
