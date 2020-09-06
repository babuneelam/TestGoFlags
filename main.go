package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

const (
	sec = 1000 * 1000 * 1000
)

// myGetUrl -traceID <traceID> -url <url> -timeout <timeoutInSec>
func main() {
	traceID := flag.String("traceID", "", "Trace/Log Id")
	url := flag.String("url", "", "Url to get")
	timeout := flag.Duration("timeout", 5*sec, "Maximum time expected for the GetUrl")
	//Other types suppprted: flag.Int, flag.Uint, flag.Bool, flag.Float64, flag.Int64, flag.Uint64

	flag.Parse()

	// Verifying whether all mandatory parameters are provided
	mandatoryArgsList := []string{"traceID", "url"}
	providedArgsList := make(map[string]bool)
	// Notes:
	// VisitAll visits the command-line flags in lexicographical order, calling fn for each. It visits all flags, even those not set.
	// Visit visits the command-line flags in lexicographical order, calling fn for each
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != "" {
			providedArgsList[f.Name] = true
		}
	})
	for _, arg := range mandatoryArgsList {
		if !providedArgsList[arg] {
			fmt.Println("Missing mandatory argument: ", arg)
			fmt.Println("usage: ", path.Base(os.Args[0]))
			fmt.Println("options:")
			flag.PrintDefaults() // prints, to standard error unless configured otherwise
			os.Exit(2)           // the same exit code flag.Parse uses
		}
	}
	fmt.Println("Numner of Args with flags:", flag.NFlag())

	// Tried the following to check whether mandatory parameters are provided or not,
	// but "flag.Lookup" succeeded even if it's not provided. The value is empty though !!
	// CODE THAT DOESN't WORK:
	// for _, arg := range mandatoryArgsList {
	// 	if flag.Lookup(arg) == nil {
	// 		fmt.Println("Missing mandatory argument: ", arg)
	// 		fmt.Println("usage: ", path.Base(os.Args[0]))
	// 		fmt.Println("options:")
	// 		flag.PrintDefaults() // prints, to standard error unless configured otherwise
	// 		os.Exit(2)           // the same exit code flag.Parse uses
	// 	} else {
	// 		fmt.Println("Mandatory argument Value: ", flag.Lookup(arg).Value)
	// 	}
	// }

	// The non-flagged arguments are expected to follow flags. I saw errors if we specicy before flag options. 
	// The non-flagged arguments are available as the slice flag.Args() or individually as flag.Arg(i). 
	// The non-flagged arguments are indexed from 0 through flag.NArg()-1.
	fmt.Println("Args without flags:", flag.Args()) 
	// Another way of printig all non-flagged arguments
	for i := 0; i < flag.NArg(); i++ { 
		fmt.Println("args[", i, "]", flag.Arg(i))
	}
	fmt.Println("TraceID", *traceID, "url", *url, "timeout:", *timeout)
}

// TBDs
//    - flagSet
//        - building sub commands: https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
// 
