// Written by: MAB

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const (
	appCreator = "MAHBUB ALAM (mahbub002@outlook.com)"
	appVersion = "2.0.18251.0"

	// Build spec
	buildSpecInvalid                   = "+"
	buildSpecBuildIncrementPerBuild    = "b+b"
	buildSpecBuildIncrementPerHour     = "b+h"
	buildSpecBuildIncrementPerDay      = "b+d"
	buildSpecBuildUseDate              = "b+t"
	buildSpecRevisionIncrementPerBuild = "r+b"
	buildSpecRevisionIncrementPerHour  = "r+h"
	buildSpecRevisionIncrementPerDay   = "r+d"

	dateTimeFormat = "2006-01-02-15:04:05-Mon"
)

/* NOTE: build-nubmer-file format-
app-version =
build-spec =
build-number =
revision-number =
build-count =
revision-count =
base-datetime =
build-datetime =
*/

type _BuildNumberFileInfo struct {
	appVersion     string
	buildSpec      string
	buildNumber    int
	revisionNumber int
	buildCount     int
	revisionCount  int
	baseDateTime   time.Time
	buildDateTime  time.Time
}

func check(err error, errMsg string) {
	if err != nil {
		err2 := fmt.Errorf(errMsg)
		panic(err2)
	}
}

func fclose(f *os.File) {
	if err := f.Close(); err != nil {
		err := fmt.Errorf("failed to close file <%s>", f.Name())
		panic(err)
	}
}

func verifyOption(option string) string {

	switch option {
	case "-r":

	case "-b":

	default:
		if option != "" {
			option = "-"
		}
	}

	return option
}

func verifyBuildSpec(buildSpec string) string {

	switch buildSpec {
	case buildSpecBuildIncrementPerBuild:

	case buildSpecBuildIncrementPerHour:

	case buildSpecBuildIncrementPerDay:

	case buildSpecBuildUseDate:

	case buildSpecRevisionIncrementPerBuild:

	case buildSpecRevisionIncrementPerHour:

	case buildSpecRevisionIncrementPerDay:

	default:
		buildSpec = buildSpecInvalid
	}

	return buildSpec
}

func createNewBuildNumberFile(fileName string) {

	fo, err := os.Create(fileName)
	check(err, fmt.Sprintf("failed to create file <%s>", fileName))

	// Close file on exit
	defer fclose(fo)

	fmt.Fprintf(fo, "app-version = %s\r\n", appVersion)
	fmt.Fprintf(fo, "build-spec = b+b\r\n")
	fmt.Fprintf(fo, "build-number = 0\r\n")
	fmt.Fprintf(fo, "revision-number = 0\r\n")
	fmt.Fprintf(fo, "build-count = 0\r\n")
	fmt.Fprintf(fo, "revision-count = 0\r\n")

	format := time.Now().Format(dateTimeFormat)

	fmt.Fprintf(fo, "base-datetime = %s\r\n", format)
	fmt.Fprintf(fo, "build-datetime = %s", format)
}

func readBuildNumberFile(fileName string) _BuildNumberFileInfo {

	fi, err := os.Open(fileName)
	check(err, fmt.Sprintf("failed to open file <%s>", fileName))
	defer fclose(fi)

	ret := _BuildNumberFileInfo{}

	fmt.Fscanf(fi, "app-version = %s\r\n", &ret.appVersion)
	fmt.Fscanf(fi, "build-spec = %s\r\n", &ret.buildSpec)
	fmt.Fscanf(fi, "build-number = %d\r\n", &ret.buildNumber)
	fmt.Fscanf(fi, "revision-number = %d\r\n", &ret.revisionNumber)
	fmt.Fscanf(fi, "build-count = %d\r\n", &ret.buildCount)
	fmt.Fscanf(fi, "revision-count = %d\r\n", &ret.revisionCount)

	var year, month, day, hour, min, sec int
	var weekDay, format string

	fmt.Fscanf(fi, "base-datetime = %d-%d-%d-%d:%d:%d-%s\r\n",
		&year, &month, &day, &hour, &min, &sec, &weekDay)

	format = fmt.Sprintf("%04d-%02d-%02d-%02d:%02d:%02d-%s",
		year, month, day, hour, min, sec, weekDay)

	baseDateTime, err := time.Parse(dateTimeFormat, format)
	check(err, fmt.Sprintf("failed to parse base-datetime <%s>", format))
	ret.baseDateTime = baseDateTime

	fmt.Fscanf(fi, "build-datetime = %d-%d-%d-%d:%d:%d-%s\r\n",
		&year, &month, &day, &hour, &min, &sec, &weekDay)

	format = fmt.Sprintf("%04d-%02d-%02d-%02d:%02d:%02d-%s",
		year, month, day, hour, min, sec, weekDay)

	buildDateTime, err := time.Parse(dateTimeFormat, format)
	check(err, fmt.Sprintf("failed to parse build-datetime <%s>", format))
	ret.buildDateTime = buildDateTime

	return ret
}

func updateBuildNumberFile(fileName string, info _BuildNumberFileInfo) {

	fo, err := os.Create(fileName)
	check(err, fmt.Sprintf("failed to open file <%s>", fileName))
	defer fclose(fo)

	fmt.Fprintf(fo, "app-version = %s\r\n", appVersion)
	fmt.Fprintf(fo, "build-spec = %s\r\n", info.buildSpec)
	fmt.Fprintf(fo, "build-number = %d\r\n", info.buildNumber)
	fmt.Fprintf(fo, "revision-number = %d\r\n", info.revisionNumber)
	fmt.Fprintf(fo, "build-count = %d\r\n", info.buildCount)
	fmt.Fprintf(fo, "revision-count = %d\r\n", info.revisionCount)
	fmt.Fprintf(fo, "base-datetime = %s\r\n", info.baseDateTime.Format(dateTimeFormat))
	fmt.Fprintf(fo, "build-datetime = %s", info.buildDateTime.Format(dateTimeFormat))
}

func updateBuildNumberFileInfo(option string, newbuildSpec string,
	info _BuildNumberFileInfo) _BuildNumberFileInfo {

	curDateTime := time.Now()
	prevBuildSpec := info.buildSpec
	build, rev := info.buildNumber, info.revisionNumber
	buildCount, revCount := info.buildCount, info.revisionCount
	baseDateTime, buildDateTime := info.baseDateTime, info.buildDateTime

	switch newbuildSpec {
	case buildSpecBuildIncrementPerBuild:
		if newbuildSpec != prevBuildSpec {
			build = buildCount
		}
		build++

	case buildSpecRevisionIncrementPerBuild:
		if newbuildSpec != prevBuildSpec {
			rev = revCount
		}
		rev++

	case buildSpecBuildIncrementPerDay:
		if newbuildSpec != prevBuildSpec {
			build = buildCount
		}

		if curDateTime.Year() > buildDateTime.Year() ||
			curDateTime.Month() > buildDateTime.Month() ||
			curDateTime.Day() > buildDateTime.Day() {
			build++
		}

	case buildSpecRevisionIncrementPerDay:
		if newbuildSpec != prevBuildSpec {
			rev = revCount
		}

		if curDateTime.Year() > buildDateTime.Year() ||
			curDateTime.Month() > buildDateTime.Month() ||
			curDateTime.Day() > buildDateTime.Day() {
			rev++
		}

	case buildSpecBuildIncrementPerHour:
		if newbuildSpec != prevBuildSpec {
			build = buildCount
		}

		if curDateTime.Year() > buildDateTime.Year() ||
			curDateTime.Month() > buildDateTime.Month() ||
			curDateTime.Day() > buildDateTime.Day() ||
			curDateTime.Hour() > buildDateTime.Hour() {
			build++
		}

	case buildSpecRevisionIncrementPerHour:
		if newbuildSpec != prevBuildSpec {
			rev = revCount
		}

		if curDateTime.Year() > buildDateTime.Year() ||
			curDateTime.Month() > buildDateTime.Month() ||
			curDateTime.Day() > buildDateTime.Day() ||
			curDateTime.Hour() > buildDateTime.Hour() {
			rev++
		}

	case buildSpecBuildUseDate:
		// Old MS style date build number like YMMDD where Y is 1 to 9
		// build = (curDateTime.Year() - baseDateTime.Year() + 1) * 10000
		// build += (int(curDateTime.Month()) * 100)
		// build += curDateTime.Day() * 1

		// Date encoded as : YYAAA where AAA is year day number in 365/366 days
		build = (curDateTime.Year() % 100) * 1000
		build += curDateTime.YearDay()
	}

	// Increment build/revision count and also reset if needed
	if strings.HasPrefix(newbuildSpec, "b+") {
		buildCount++

		if option == "-r" {
			rev = 0
			revCount = 0
		}
	} else if strings.HasPrefix(newbuildSpec, "r+") {
		revCount++

		if option == "-b" {
			build = 0
			buildCount = 0
		}
	}

	retInfo := _BuildNumberFileInfo{}

	retInfo.appVersion = info.appVersion
	retInfo.buildSpec = newbuildSpec
	retInfo.buildNumber = build
	retInfo.revisionNumber = rev
	retInfo.buildCount = buildCount
	retInfo.revisionCount = revCount
	retInfo.baseDateTime = baseDateTime
	retInfo.buildDateTime = curDateTime

	return retInfo
}

func updateVersionFile(outFileName string, inFileName string, build int, revision int) {

	fi, err := os.Open(inFileName)
	check(err, fmt.Sprintf("failed to open file <%s>", inFileName))
	fo, err := os.Create(outFileName)
	check(err, fmt.Sprintf("failed to create/open file <%s>", outFileName))

	defer fclose(fo)
	defer fclose(fi)

	reader := bufio.NewReader(fi)
	writer := bufio.NewWriter(fo)

	var state int
	var seg byte

	for {
		b, err := reader.ReadByte()

		if err != nil && err != io.EOF {
			err2 := fmt.Errorf("failed to read file <%s>", inFileName)
			panic(err2)
		}

		if err == io.EOF {
			break
		}

		switch state {
		case 0:
			if b == '$' {
				state++
			} else {
				err := writer.WriteByte(b)
				check(err, fmt.Sprintf("failed to write to file <%s>", outFileName))
			}

		case 1:
			if b == 'b' || b == 'r' {
				state++
				seg = b
			} else {
				state = 0
				_, err := fmt.Fprintf(writer, "$%c", b)
				check(err, fmt.Sprintf("failed to write to file <%s>", outFileName))
			}

		case 2:
			if b == '$' {
				state = 0
				switch seg {
				case 'b':
					_, err := fmt.Fprintf(writer, "%d", build)
					check(err, fmt.Sprintf("failed to write to file <%s>", outFileName))

				case 'r':
					_, err := fmt.Fprintf(writer, "%d", revision)
					check(err, fmt.Sprintf("failed to write to file <%s>", outFileName))
				}
			} else {
				state = 0
				_, err := fmt.Fprintf(writer, "$%c%c", seg, b)
				check(err, fmt.Sprintf("failed to write to file <%s>", outFileName))
			}
		}
	}

	writer.Flush()
}

func printHelp() {
	fmt.Printf("\nCountBuild %s [ Written by: %s ]\n", appVersion, appCreator)
	fmt.Printf("Usage: countbuild [options] [buildspec] [buildfile] [infile] [outfile]\n")
	fmt.Printf("\nNOTE: To create NEW build-number file use: countbuild newbuild [buildfile]\n")
	fmt.Printf("\n")
	fmt.Printf("[options]\n")
	fmt.Printf("  -r              Reset revision counter when buildspec is changed to b+*.\n")
	fmt.Printf("  -b              Reset build counter when buildspec is changed to r+*.\n")
	fmt.Printf("[buildspec]\n")
	fmt.Printf("  b+b             Increment build number on each build.\n")
	fmt.Printf("  b+h             Increment build number on hourly basis.\n")
	fmt.Printf("  b+d             Increment build number on daily basis.\n")
	fmt.Printf("  b+t             Use date as build number.\n")
	fmt.Printf("  r+b             Increment revision number on each build.\n")
	fmt.Printf("  r+h             Increment revision number on hourly basis.\n")
	fmt.Printf("  r+d             Increment revision number on daily basis.\n")
	fmt.Printf("[buildfile]       Holds the build+revision number+counter and date-time.\n")
	fmt.Printf("[infile]          Version template file, which has build & revision number\n")
	fmt.Printf("                  segment markers ($b$ and $r$).\n")
	fmt.Printf("[outfile]         Version file, created from the version template file; but the\n")
	fmt.Printf("                  markers are replaced by the appropriate version segment number.\n")
	fmt.Printf("\nExample:\n")
	fmt.Printf("countbuild -r b+d build.txt versioninfo.in.txt versioninfo.txt\n")
}

func main() {
	//argv := []string{"countbuild", "-r", "b+b", "build.txt", "versioninfo.in.txt", "versioninfo.txt"}
	argv := os.Args
	argc := len(argv)

	var option, newBuildSpec, buildFile, inFile, outFile string

	if argc < 2 {
		printHelp()
		os.Exit(0)
	} else if argc == 3 {
		option = argv[1]
		buildFile = argv[2]

		if option == "newbuild" {
			createNewBuildNumberFile(buildFile)
			os.Exit(0)
		} else {
			err := fmt.Errorf("invalid option <%s>", option)
			panic(err)
		}
	} else if argc == 5 {
		option = ""
		newBuildSpec = verifyBuildSpec(argv[1])
		buildFile = argv[2]
		inFile = argv[3]
		outFile = argv[4]
	} else if argc == 6 {
		option = verifyOption(argv[1])
		newBuildSpec = verifyBuildSpec(argv[2])
		buildFile = argv[3]
		inFile = argv[4]
		outFile = argv[5]
	} else {
		err := fmt.Errorf("invalid parameters")
		panic(err)
	}

	// Now process

	if option == "-" {
		err := fmt.Errorf("invalid option")
		panic(err)
	}

	if newBuildSpec == buildSpecInvalid {
		err := fmt.Errorf("invalid buildpsec")
		panic(err)
	}

	prevInfo := readBuildNumberFile(buildFile)
	newInfo := updateBuildNumberFileInfo(option, newBuildSpec, prevInfo)
	updateVersionFile(outFile, inFile, newInfo.buildNumber, newInfo.revisionNumber)
	updateBuildNumberFile(buildFile, newInfo)
}
